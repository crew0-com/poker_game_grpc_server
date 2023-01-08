package api

import (
	"github.com/crew_0/poker/internal/config"
	"github.com/crew_0/poker/internal/grpc/pb"

	db "github.com/crew_0/poker/database/sqlc"
)

type Server struct {
	pb.UnimplementedPokerServiceServer
	c     config.Config
	store db.Queries
}

func NewServer(c config.Config, store db.Queries) (*Server, error) {
	server := &Server{
		c:     c,
		store: store,
	}

	return server, nil
}
