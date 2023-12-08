package api

import (
	"fmt"
	db "simple-bank/db/sqlc"
	middleware "simple-bank/internal/auth"
	"simple-bank/token"
	"simple-bank/util"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config      util.Config
	store       *db.SQLStore
	jwtMaker    token.JwtMaker
	pasetoMaker token.PMaker
	router      *gin.Engine
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

	server.setupRouter()
	
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	api := router.Group("/api/v1")

	// users
	api.POST("/users", server.creatUser)
	api.GET("/users/:username", server.getUser)
	api.POST("/users/login", server.loginUser)

	authRoutes := api.Group("/").Use(middleware.PasetoAuthMiddleware(server.pasetoMaker))

	// accounts
	authRoutes.POST("/accounts", server.creatAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	// transfers
	authRoutes.POST("/transfer", server.createTransfer)
	
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
