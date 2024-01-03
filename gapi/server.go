package gapi

import (
	"fmt"
	db "simple-bank/db/sqlc"
	pb "simple-bank/pb"
	"simple-bank/token"
	"simple-bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store  *db.SQLStore
	// jwtMaker    token.JwtMaker  // can use this method
	pasetoMaker token.PMaker
}

func NewServer(config util.Config, store *db.SQLStore) (*Server, error) {
	pasetoMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:      config,
		store:       store,
		pasetoMaker: pasetoMaker,
	}

	return server, nil
}