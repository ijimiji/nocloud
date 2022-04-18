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
package services

import (
	"context"
	"encoding/json"
	"time"

	pb "github.com/slntopp/nocloud/pkg/services/proto"
	settingspb "github.com/slntopp/nocloud/pkg/settings/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"go.uber.org/zap"
)

var (
	settingsClient settingspb.SettingsServiceClient
)

// Settings Key storing monitoring routine conf
const monFreqKey string = "services-monitoring-routine"

type MonitoringRoutineConf struct {
	Frequency int `json:"freq"` // Frequency in Seconds
}

var (
	defaultConf = MonitoringRoutineConf{
		Frequency: 60,
	}

	description = "Services Monitoring Routine Configuration"
	public = false
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("SETTINGS_HOST", "settings:8080")
    host := viper.GetString("SETTINGS_HOST")
    
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        panic(err)
    }

    settingsClient = settingspb.NewSettingsServiceClient(conn)
}

func MakeConf(ctx context.Context, log *zap.Logger) MonitoringRoutineConf {

	var conf MonitoringRoutineConf
	var r_str string
	r, err := settingsClient.Get(ctx, &settingspb.GetRequest{Keys: []string{monFreqKey}})
	if err != nil {
		log.Debug("Failed to Get conf", zap.Error(err))
		goto set_default
	}
	if _, ok := r.GetFields()[monFreqKey]; !ok {
		goto set_default
	}
	r_str = r.GetFields()[monFreqKey].GetStringValue()
	err = json.Unmarshal([]byte(r_str), &conf)
	if err != nil {
		goto set_default
	}
	return conf

	set_default:
	log.Info("Setting default conf")
	conf = defaultConf
	payload, err := json.Marshal(conf)
	if err == nil {
		_, err := settingsClient.Put(ctx, &settingspb.PutRequest{
			Key: monFreqKey,
			Value: string(payload),
			Description: &description,
			Public: &public,
		})
		if err != nil {
			log.Error("Error Putting Monitoring Configuration", zap.Error(err))
		}
	}
	return conf
}

func (s *ServicesServer) MonitoringRoutineState() Routine {
	return s.monitoring
}

func (s *ServicesServer) MonitoringRoutine(ctx context.Context) {
	log := s.log.Named("MonitoringRoutine")

	conf := MakeConf(ctx, log)
	log.Info("Got Monitoring Configuration", zap.Any("conf", conf))
	ticker := time.NewTicker(time.Second * time.Duration(conf.Frequency))

	for tick := range ticker.C {
		log.Info("Starting", zap.Time("tick", tick))
		s.monitoring.Running = true

		pool, err := s.ctrl.List(ctx, schema.ROOT_ACCOUNT_KEY, &pb.ListRequest{})
		if err != nil {
			log.Error("Failed to get Services", zap.Error(err))
			continue
		}
		log.Debug("Got Services", zap.Int("length", len(pool)))

		for _, service := range pool {
			go func(service *pb.Service) {
				states, err := s.GetStatesInternal(ctx, service)
				if err != nil {
					log.Error("Failed to get Service Instances states", zap.String("service", service.GetUuid()), zap.Error(err))
					return
				}

				for _, group := range service.GetInstancesGroups() {
					for _, inst := range group.GetInstances() {
						state, ok := states.GetStates()[inst.GetUuid()]
						if !ok {
							log.Debug("State for Instance not found", zap.String("instance", inst.GetUuid()))
							continue
						}
						inst.State = state
					}
				}

				err = s.ctrl.Update(ctx, service, false)
				if err != nil {
					log.Error("Failed to update Service", zap.String("service", service.GetUuid()), zap.Error(err))
				}
			}(service)
		}

		s.monitoring.LastExec = tick.Format("2006-01-02T15:04:05Z07:00")
	}
}