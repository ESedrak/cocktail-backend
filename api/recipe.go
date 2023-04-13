package api

import (
	db "cocktail-backend/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createRecipeRequest struct {
	DrinkName    string  `json:"drink_name" binding:"required"`
	Instructions string  `json:"instructions" binding:"required"`
	ImageUrl     *string `json:"image_url"`
}

func (server *Server) createRecipe(ctx *gin.Context) {
	var req createRecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRecipeParams{
		DrinkName:    req.DrinkName,
		Instructions: req.Instructions,
		ImageUrl:     req.ImageUrl,
	}

	recipe, err := server.store.CreateRecipe(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, recipe)
}

type getRecipeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getRecipe(ctx *gin.Context) {
	var req getRecipeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	recipe, err := server.store.GetRecipe(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, recipe)
}

type listRecipeRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listRecipe(ctx *gin.Context) {
	var req listRecipeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListRecipesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	recipes, err := server.store.ListRecipes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, recipes)
}

type deleteRecipeRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteRecipe(ctx *gin.Context) {
	var req deleteRecipeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteRecipe(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, err)
}

type updateRecipeRequest struct {
	RecipeID     int64   `json:"recipe_id" binding:"required,min=1"`
	DrinkName    *string `json:"drink_name"`
	Instructions *string `json:"instructions"`
	ImageUrl     *string `json:"image_url"`
}

func (server *Server) updateRecipe(ctx *gin.Context) {
	var req updateRecipeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	existingRecipe, err := server.store.GetRecipe(ctx, req.RecipeID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	updateExistingRecipeDetails(&req, &existingRecipe)

	arg := db.UpdateRecipeParams{
		RecipeID:     existingRecipe.RecipeID,
		DrinkName:    existingRecipe.DrinkName,
		Instructions: existingRecipe.Instructions,
		ImageUrl:     existingRecipe.ImageUrl,
	}
	recipes, err := server.store.UpdateRecipe(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, recipes)
}

func updateExistingRecipeDetails(in *updateRecipeRequest, existingRecipeDetails *db.Recipe) {
	if in.DrinkName != nil {
		existingRecipeDetails.DrinkName = *in.DrinkName
	}
	if in.Instructions != nil {
		existingRecipeDetails.Instructions = *in.Instructions
	}
	if in.ImageUrl != nil {
		existingRecipeDetails.ImageUrl = in.ImageUrl
	}
}
