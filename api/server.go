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
	newRouter.GET("/users/email/:email", newServer.GetUserByEmail)
	newRouter.GET("/users/id/:id", newServer.GetUserById)
	// Must be called with quotation:
	// For exampla: localhost:8080/users/id/"1432bf1f-a448-4f50-a20b-5b7ed4a9ad2b"
	newRouter.PATCH("/users", newServer.UpdateUser)
	newRouter.DELETE("/users/:id", newServer.DeleteUser)
	newRouter.GET("/users", newServer.ListUsers)

	newServer.router = newRouter
	return newServer
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
