/*
Copyright © 2021 Nikita Ivanovski info@slnt-opp.xyz

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

	"github.com/arangodb/go-driver"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/access"
	"github.com/slntopp/nocloud/pkg/nocloud/roles"
	"go.uber.org/zap"
)

const (
	NAMESPACES_COL = "Namespaces"
	NS2ACC = NAMESPACES_COL + "2" + ACCOUNTS_COL
)

type Namespace struct {
	Title string `json:"title"`
	driver.DocumentMeta
}

type NamespacesController struct {
	col driver.Collection
	log *zap.Logger
}

func NewNamespacesController(log *zap.Logger, col driver.Collection) NamespacesController {
	return NamespacesController{log: log, col: col}
}

func (ctrl *NamespacesController) Get(ctx context.Context, id string) (Namespace, error) {
	var r Namespace
	_, err := ctrl.col.ReadDocument(nil, id, &r)
	return r, err
}

func (ctrl *NamespacesController) List(ctx context.Context, requestor Account, req_depth *int32) ([]Namespace, error) {
	var depth int32
	if req_depth == nil || *req_depth < 2 {
		depth = 2
	} else {
		depth = *req_depth
	}

	query := `FOR node IN 0..@depth OUTBOUND @account GRAPH @permissions_graph OPTIONS {order: "bfs", uniqueVertices: "global"} FILTER IS_SAME_COLLECTION(@@namespaces, node) RETURN node`
	bindVars := map[string]interface{}{
		"depth": depth,
		"account": requestor.ID.String(),
		"permissions_graph": PERMISSIONS_GRAPH.Name,
		"@namespaces": NAMESPACES_COL,
	}
	ctrl.log.Debug("Ready to build query", zap.Any("bindVars", bindVars))

	c, err := ctrl.col.Database().Query(ctx, query, bindVars)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	var r []Namespace
	for {
		var ns Namespace 
		_, err := c.ReadDocument(ctx, &ns)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		ctrl.log.Debug("Got document", zap.Any("namespace", ns))
		r = append(r, ns)
	}

	return r, nil
}

func (ctrl *NamespacesController) Create(ctx context.Context, title string) (Namespace, error) {
	ns := Namespace{
		Title: title,
	}
	meta, err := ctrl.col.CreateDocument(ctx, ns)
	if err != nil {
		return Namespace{}, err
	}
	
	ns.DocumentMeta = meta
	key := ctx.Value(nocloud.NoCloudAccount).(string)
	acc := Account{
		DocumentMeta: driver.DocumentMeta {
			Key: key,
			ID: driver.NewDocumentID(ACCOUNTS_COL, key),
		},
	}

	return ns, ctrl.Link(ctx, acc, ns, access.ADMIN, roles.OWNER)
}

func (ctrl *NamespacesController) Link(ctx context.Context, acc Account, ns Namespace, access int32, role string) (error) {
	edge, _ := ctrl.col.Database().Collection(ctx, ACC2NS)
	return acc.LinkNamespace(ctx, edge, ns, access, role)
}

func (ctrl *NamespacesController) Join(ctx context.Context, acc Account, ns Namespace, access int32, role string) (error) {
	edge, _ := ctrl.col.Database().Collection(ctx, NS2ACC)
	return acc.JoinNamespace(ctx, edge, ns, access, role)
}

func (ns *Namespace) Delete(ctx context.Context, db driver.Database) (error) {
	err := DeleteNodeChildren(ctx, db, ns.ID.String())
	if err != nil {
		return err
	}

	graph, _ := db.Graph(ctx, PERMISSIONS_GRAPH.Name)
	col, _ := graph.VertexCollection(ctx, NAMESPACES_COL)
	_, err = col.RemoveDocument(ctx, ns.Key)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *NamespacesController) Delete(ctx context.Context, id string) (error) {
	ns, err := ctrl.Get(ctx, id)
	if err != nil {
		return err
	}
	return ns.Delete(ctx, ctrl.col.Database())
}
