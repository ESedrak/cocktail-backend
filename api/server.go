package api

import (
	db "cocktail-backend/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// recipe
	router.POST("/recipe", server.createRecipe)
	router.GET("/recipe/:id", server.getRecipe)
	router.GET("/recipe", server.listRecipe)
	router.POST("/recipe/:id", server.updateRecipe)
	router.DELETE("/recipe/:id", server.deleteRecipe)

	server.router = router
	return server
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
