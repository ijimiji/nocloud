/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package graph

import (
	"context"
	"errors"
	"fmt"

	"github.com/arangodb/go-driver"
	hasher "github.com/slntopp/nocloud/pkg/hasher"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	pb "github.com/slntopp/nocloud/pkg/services/proto"
	"go.uber.org/zap"
)

type Service struct {
	*pb.Service
	driver.DocumentMeta
}

type Provision struct {
	From driver.DocumentID `json:"_from"`
	To driver.DocumentID `json:"_to"`
	Group string `json:"group"`

	driver.DocumentMeta
}

type ServicesController struct {
	log *zap.Logger

	col driver.Collection // Services Collection
	ig_ctrl InstancesGroupsController

	db driver.Database
}

func NewServicesController(log *zap.Logger, db driver.Database) ServicesController {
	col, _ := db.Collection(context.TODO(), schema.SERVICES_COL)
	return ServicesController{log: log, col: col, ig_ctrl: NewInstancesGroupsController(log, db), db:db}
}

const createServiceQuery =
`LET service = @body
LET instances = (FOR gid IN ATTRIBUTES(service.instances_groups)
    LET group = service.instances_groups[gid]
    FOR instance IN group.instances
        RETURN instance.uuid)

INSERT MERGE(service, {instances: instances}) INTO @@services RETURN NEW`

// Create Service and underlaying entities and store in DB
func (ctrl *ServicesController) Create(ctx context.Context, service *pb.Service) (*Service, error) {
	ctrl.log.Debug("Creating Service", zap.Any("service", service))
	for _, ig := range service.GetInstancesGroups() {
		err := ctrl.ig_ctrl.Create(ctx, ig)
		if err != nil {
			return nil, err
		}
	}

	err := hasher.SetHash(service.ProtoReflect())
	if err != nil {
		return nil, err
	}

	service.Status = "init"

	c, err := ctrl.db.Query(ctx, createServiceQuery, map[string]interface{}{
		"body": service,
		"@services": schema.SERVICES_COL,
	})
	if err != nil {
		ctrl.log.Debug("Error creating document", zap.Error(err))
		return nil, err
	}
	defer c.Close()

	meta, err := c.ReadDocument(ctx, service)
	if err != nil {
		ctrl.log.Debug("Error reading new document", zap.Error(err))
		return nil, err
	}
	service.Uuid = meta.ID.Key()
	return &Service{service, meta}, nil
}

const updateServiceQuery = 
`LET service = @body
LET instances = (FOR gid IN ATTRIBUTES(service.instances_groups)
    LET group = service.instances_groups[gid]
    FOR instance IN group.instances
        RETURN instance.uuid)

REPLACE @service WITH MERGE(service, {instances: instances}) IN @@services`

// Update Service and underlaying entities and store in DB
func (ctrl *ServicesController) Update(ctx context.Context, service *pb.Service, hash bool) (error) {
	ctrl.log.Debug("Updating Service", zap.Any("service", service))

	if hash {
		hash := service.GetHash()
		err := hasher.SetHash(service.ProtoReflect())
		if err != nil {
			return err
		}
		if service.GetHash() != hash {
			service.Status = "modified"
		}
	}

	c, err := ctrl.db.Query(ctx, updateServiceQuery, map[string]interface{}{
		"service": service.GetUuid(),
		"body": service,
		"@services": schema.SERVICES_COL,
	})
	if err != nil {
		return err
	}
	defer c.Close()

	return nil
}

// Get Service from DB
func (ctrl *ServicesController) Get(ctx context.Context, id string) (*Service, error) {
	ctrl.log.Debug("Getting Service", zap.String("id", id))
	var service pb.Service
	meta, err := ctrl.col.ReadDocument(ctx, id, &service)
	if err != nil {
		ctrl.log.Debug("Error reading document(Service)", zap.Error(err))
		return nil, errors.New("error reading document")
	}
	ctrl.log.Debug("ReadDocument.Result", zap.Any("meta", meta), zap.Any("service", &service))
	service.Uuid = meta.ID.Key()
	return &Service{&service, meta}, nil
}

// List Services in DB
func (ctrl *ServicesController) List(ctx context.Context, requestor string, request *pb.ListRequest) ([]*Service, error) {
	ctrl.log.Debug("Getting Services", zap.String("requestor", requestor))

	depth := request.GetDepth()
	if depth < 2 {
		depth = 5
	}
	showDeleted := request.GetShowDeleted() == "true"
	var query string
	if showDeleted {
		query = `FOR node IN 0..@depth OUTBOUND @account GRAPH @permissions_graph OPTIONS {order: "bfs", uniqueVertices: "global"} FILTER IS_SAME_COLLECTION(@@services, node) RETURN node`
	} else {
		query = `FOR node IN 0..@depth OUTBOUND @account GRAPH @permissions_graph OPTIONS {order: "bfs", uniqueVertices: "global"} FILTER IS_SAME_COLLECTION(@@services, node) FILTER node.status != "del" RETURN node`
	}
	bindVars := map[string]interface{}{
		"depth": depth,
		"account": driver.NewDocumentID(schema.ACCOUNTS_COL, requestor),
		"permissions_graph": schema.PERMISSIONS_GRAPH.Name,
		"@services": schema.SERVICES_COL,
	}
	ctrl.log.Debug("Ready to build query", zap.Any("bindVars", bindVars), zap.Bool("show_deleted", showDeleted))

	c, err := ctrl.db.Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	var r []*Service
	for {
		var s pb.Service
		meta, err := c.ReadDocument(ctx, &s)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		ctrl.log.Debug("Got document", zap.Any("service", &s))
		s.Uuid = meta.ID.Key()
		r = append(r, &Service{&s, meta})
	}

	return r,  nil
}

// Join Service into Namespace
func (ctrl *ServicesController) Join(ctx context.Context, service *Service, namespace *Namespace, access int32, role string) (error) {
	ctrl.log.Debug("Joining service to namespace")
	edge, _ := ctrl.db.Collection(ctx, schema.NS2SERV)
	_, err := edge.CreateDocument(ctx, Access{
		From: namespace.ID,
		To: service.ID,
		Level: access,
		Role: role,
	})
	return err
}

// Create Link between Service/Group and Services Provider group is Provisioned(deployed) to
func (ctrl *ServicesController) Provide(ctx context.Context, sp, service driver.DocumentID, group string) (error) {
	ctrl.log.Debug("Providing group to service provider")
	edge, _ := ctrl.db.Collection(ctx, schema.SP2SERV)
	_, err := edge.CreateDocument(ctx, Provision{
		From: sp,
		To: service,
		Group: group,
		DocumentMeta: driver.DocumentMeta{Key: group},
	})
	return err
}

// Delete Link between Service/Group and Services Provider group have beem Unprovisioned(undeployed) from
func (ctrl *ServicesController) Unprovide(ctx context.Context, group string) (err error) {
	ctrl.log.Debug("Unproviding group from service provider", zap.String("group", group))
	g, _ := ctrl.db.Graph(ctx, schema.SERVICES_GRAPH.Name)
	edge, _, _ := g.EdgeCollection(ctx, schema.SP2SERV)
	_, err = edge.RemoveDocument(ctx, group)
	return err
}

// Get Provisions, map of InstancesGroups to ServicesProviders, those groups are deployed to
func (ctrl *ServicesController) GetProvisions(ctx context.Context, service string) (r map[string]string, err error) {
	ctrl.log.Debug("Getting groups provisions")
	query := `FOR service, provision IN INBOUND @service GRAPH @services RETURN provision`
	bindVars := map[string]interface{}{
		"service": service,
		"services": schema.SERVICES_GRAPH.Name,
	}
	ctrl.log.Debug("Ready to build query", zap.Any("bindVars", bindVars))

	c, err := ctrl.db.Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	r = make(map[string]string)
	for {
		var p Provision
		_, err = c.ReadDocument(ctx, &p)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		ctrl.log.Debug("Got document", zap.Any("provision", p))
		r[p.Group] = p.From.Key()
	}

	return r, nil
}

func (ctrl *ServicesController) Delete(ctx context.Context, s *Service) (err error) {
	log := ctrl.log.Named("Service.Delete")
	log.Debug("Deleting Service")
	if s.GetStatus() != "init" {
		return fmt.Errorf("cannot delete Service, status: %s", s.GetStatus())
	}

	return ctrl.SetStatus(ctx, s, "del")
}

func (ctrl *ServicesController) SetStatus(ctx context.Context, s *Service, status string) (err error) {
	// TODO: check if valid status message
	s.Status = status
	return ctrl.Update(ctx, s.Service, false)
}

const findServiceByInstanceQuery =
`FOR service IN @@services
FILTER @instance IN service.instances[*]
LIMIT 1
    RETURN service`

// Returns Service by one of it's Instances UUID
func (ctrl *ServicesController) FindServiceByInstance(ctx context.Context, instance string) (*Service, error) {
	log := ctrl.log.Named("FindServiceByInstance")

	c, err := ctrl.db.Query(ctx, findServiceByInstanceQuery, map[string]interface{}{
		"instance": instance,
		"@services": schema.SERVICES_COL,
	})
	if err != nil {
		return nil, err
	}
	defer c.Close()

	var service *pb.Service
	meta, err := c.ReadDocument(ctx, service)
	if err != nil {
		log.Error("Error reading new document", zap.String("instance", instance), zap.Error(err))
		return nil, err
	}
	service.Uuid = meta.ID.Key()
	return &Service{service, meta}, nil
}