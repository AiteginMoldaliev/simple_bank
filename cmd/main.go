package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"simple-bank/api"
	db "simple-bank/db/sqlc"
	"simple-bank/gapi"
	pb "simple-bank/pb"
	"simple-bank/util"
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	conn, err := sql.Open(config.Dbdriver, config.Dbsource)
	if err != nil {
		panic(err)
	}
	store := db.NewStore(conn)
	
	go runGRPCServer(config, store)

	go runHTTPGateawayServer(config)

	runGinServer(config, store)
}

func runGRPCServer(config util.Config, store *db.SQLStore) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		panic(err)
	}
	
	log.Printf("gRPC server started at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}

func runHTTPGateawayServer(config util.Config) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := pb.RegisterSimpleBankHandlerFromEndpoint(context.Background(), mux, config.GRPCServerAddress, opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("GRPC Gateway server started on :50051")
	if err := http.ListenAndServe(":50051", mux); err != nil {
		panic(err)
	}
}

func runGinServer(config util.Config, store *db.SQLStore) {
	server, err := api.NewServer(config, store)
	if err != nil {
		panic(err)
	}

	if err = server.Start(config.GinServerAddress); err != nil {
		panic(err)
	}

	log.Printf("HTTP server started on: %v", config.GinServerAddress)
}