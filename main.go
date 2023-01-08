package main

import (
	"database/sql"
	db "github.com/crew_0/poker/database/sqlc"
	"github.com/crew_0/poker/internal/config"
	"github.com/crew_0/poker/internal/grpc/api"
	"github.com/crew_0/poker/internal/grpc/pb"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.New(conn)
	runGrpcServer(c, store)
}

func runGrpcServer(c config.Config, store *db.Queries) {
	server, err := api.NewServer(c, *store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPokerServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", c.GRPCServerAddress)
	if err != nil {
		log.Fatalf("cannot create gRPC server: %v", err)
	}

	log.Printf("starting gRPC server at %s", c.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("cannot start gRPC server: %v", err)
	}
}
