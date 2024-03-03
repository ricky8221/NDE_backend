package api

import (
	"NDE_backend/token"
	"NDE_backend/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Server servers HTTP requests
type Server struct {
	config     util.Config
	store      Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{config: config, store: store, tokenMaker: tokenMaker}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.POST("/users/create", server.createUser)
	router.POST("/users/login", server.loginUser)

	// authRoutes Validating token before reaching auth needed path
	groupAuthMiddleware := router.Group("/").Use(groupAuthMiddleware(server.tokenMaker, server.config.AllRights))

	groupAuthMiddleware.POST("/createCompany", server.createCompany)
	groupAuthMiddleware.POST("/getCompany", server.getCompany)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
