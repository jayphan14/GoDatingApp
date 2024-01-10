package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/jayphan14/GoDatingApp/sqlc"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	newRouter := gin.Default()
	newServer := &Server{store: store}

	// DEFINED ROUTES:
	newRouter.POST("/users", newServer.CreateUser)
	newRouter.GET("/users/:email", newServer.GetUserByEmail)

	newServer.router = newRouter
	return newServer
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
