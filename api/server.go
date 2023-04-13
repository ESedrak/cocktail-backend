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
	router := gin.Default()

	// recipe
	router.POST("/recipe", server.createRecipe)
	router.GET("/recipe/:id", server.getRecipe)
	router.GET("/recipes", server.listRecipes)
	router.POST("/recipe/:id", server.updateRecipe)
	router.DELETE("/recipe/:id", server.deleteRecipe)

	// ingredient
	router.POST("/ingredient", server.createIngredient)
	router.GET("/ingredient/:id", server.getIngredient)
	router.GET("/ingredients", server.listIngredients)
	router.POST("/ingredient/:id", server.updateIngredient)
	router.DELETE("/ingredient/:id", server.deleteIngredient)

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
