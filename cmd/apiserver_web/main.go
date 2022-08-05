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
package main

import (
	"context"
	"crypto/tls"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
	billingpb "github.com/slntopp/nocloud/pkg/billing/proto"
	dnspb "github.com/slntopp/nocloud/pkg/dns/proto"
	healthpb "github.com/slntopp/nocloud/pkg/health/proto"
	instancespb "github.com/slntopp/nocloud/pkg/instances/proto"
	"github.com/slntopp/nocloud/pkg/nocloud"
	registrypb "github.com/slntopp/nocloud/pkg/registry/proto"
	servicespb "github.com/slntopp/nocloud/pkg/services/proto"
	sppb "github.com/slntopp/nocloud/pkg/services_providers/proto"
	settingspb "github.com/slntopp/nocloud/pkg/settings/proto"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/gabriel-vasile/mimetype"
)

var (
	log *zap.Logger

	apiserver   string
	corsAllowed []string
	insecure    bool
	with_block  bool
)

func init() {
	viper.AutomaticEnv()
	log = nocloud.NewLogger()

	viper.SetDefault("CORS_ALLOWED", []string{"*"})
	viper.SetDefault("APISERVER_HOST", "proxy:8000")
	viper.SetDefault("INSECURE", true)
	viper.SetDefault("WITH_BLOCK", false)

	apiserver = viper.GetString("APISERVER_HOST")
	corsAllowed = strings.Split(viper.GetString("CORS_ALLOWED"), ",")
	insecure = viper.GetBool("INSECURE")
	with_block = viper.GetBool("WITH_BLOCK")
}

func getContentType(path string) (mime string, ok bool) {
	chunks := strings.Split(path, ".")
	switch chunks[len(chunks)-1] {
	case "css":
		return "text/css; charset=utf-8", true
	default:
		return "", false
	}
}

func staticHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	log.Debug("Static request", zap.Any("path", pathParams))
	file := "index.html"
	if path, ok := pathParams["path"]; ok {
		file = path
	}
	if path, ok := pathParams["file"]; ok {
		file += "/" + path
	}
	index, err := os.ReadFile("/dist/" + file)
	if err != nil {
		log.Error("Error reading file", zap.Error(err))
		w.WriteHeader(404)
		return
	}

	mime, ok := getContentType(file)
	if !ok {
		mime = mimetype.Detect(index).String()
	}
	log.Debug("Content-Type", zap.String("mime", mime))
	w.Header().Set("Content-Type", mime)
	w.Write(index)
}

func main() {
	defer func() {
		_ = log.Sync()
	}()

	log.Info("Starting REST-API Server")
	log.Info("Registering Endpoints", zap.String("server", apiserver))
	var err error

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{}
	if insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}
	if with_block {
		opts = append(opts, grpc.WithBlock())
	}
	log.Info("Registering HealthService Gateway")
	err = healthpb.RegisterHealthServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register HealthService gateway", zap.Error(err))
	}
	log.Info("Registering AccountsService Gateway")
	err = registrypb.RegisterAccountsServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register AccountsService gateway", zap.Error(err))
	}
	log.Info("Registering NamespacesService Gateway")
	err = registrypb.RegisterNamespacesServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register NamespacesService gateway", zap.Error(err))
	}
	log.Info("Registering ServicesProvidersService Gateway")
	err = sppb.RegisterServicesProvidersServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register ServicesProvidersService gateway", zap.Error(err))
	}
	log.Info("Registering ServicesService Gateway")
	err = servicespb.RegisterServicesServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register ServicesService gateway", zap.Error(err))
	}
	log.Info("Registering InstancesService Gateway")
	err = instancespb.RegisterInstancesServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register InstancesService gateway", zap.Error(err))
	}
	log.Info("Registering DNS Gateway")
	err = dnspb.RegisterDNSHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register DNS gateway", zap.Error(err))
	}
	log.Info("Registering SettingsService Gateway")
	err = settingspb.RegisterSettingsServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register SettingsService gateway", zap.Error(err))
	}
	log.Info("Registering BillingService Gateway")
	err = billingpb.RegisterBillingServiceHandlerFromEndpoint(context.Background(), gwmux, apiserver, opts)
	if err != nil {
		log.Fatal("Failed to register BillingService gateway", zap.Error(err))
	}

	gwmux.HandlePath("GET", "/admin", staticHandler)
	gwmux.HandlePath("GET", "/admin/{path}", staticHandler)
	gwmux.HandlePath("GET", "/admin/{path}/{file}", staticHandler)

	log.Info("Allowed Origins", zap.Strings("hosts", corsAllowed))
	handler := cors.New(cors.Options{
		AllowedOrigins:   corsAllowed,
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "PATCH", "OPTIONS", "HEAD"},
		AllowCredentials: true,
	}).Handler(gwmux)

	log.Info("Serving gRPC-Gateway on http://0.0.0.0:8000")
	log.Fatal("Failed to Listen and Serve Gateway-Server", zap.Error(http.ListenAndServe(":8000", wsproxy.WebsocketProxy(handler))))
}
