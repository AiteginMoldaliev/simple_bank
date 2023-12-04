package api

import (
	db "simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	api := router.Group("/api/v1")
	// accounts
	api.POST("/accounts", server.creatAccount)
	api.GET("/accounts", server.listAccount)
	api.GET("/accounts/:id", server.getAccount)
	api.DELETE("/accounts/:id", server.deleteAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
