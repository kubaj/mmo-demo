package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"sync"

	"github.com/kubaj/mmo-demo/auth"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/kubaj/mmo-demo/order"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.DebugLevel)
}

func main() {

	viper.SetDefault("db.conn", "host=db-svc user=goo dbname=goo sslmode=disable password=goo")
	viper.SetDefault("server.binds.grpc", ":50051")
	viper.SetDefault("server.binds.gw", ":80")
	viper.SetDefault("server.links.auth", "auth:50051")

	lis, err := net.Listen("tcp", viper.GetString("server.binds.grpc"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create the grpc server
	s := grpc.NewServer()

	authCC, err := grpc.Dial(viper.GetString("server.links.auth"), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to auth service")
	}

	authClient := auth.NewAuthServiceClient(authCC)

	// register the service
	order.RegisterOrderServiceServer(s, order.NewService(authClient))

	// Register reflection service on gRPC server.
	reflection.Register(s)

	// if any service fails, whole app should fail
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		log.Infoln("Starting gRPC server on", viper.GetString("server.binds.grpc"))
		if err := s.Serve(lis); err != nil {
			log.Fatalf("grpc: failed to serve: %v", err)
		}

		wg.Done()
	}()

	go func() {
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux()
		opts := []grpc.DialOption{grpc.WithInsecure()}
		err := order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, viper.GetString("server.binds.grpc"), opts)
		if err != nil {
			log.Fatalf("gw: failed to register: %v", err)
		}

		log.Infoln("Starting gateway server on", viper.GetString("server.binds.gw"))
		log.Fatalf("gw: failed to server: %v", http.ListenAndServe(viper.GetString("server.binds.gw"), mux))
	}()

	wg.Wait()
}
