package api

import (
	db "cocktail-backend/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createIngredientRequest struct {
	IngredientName string `json:"ingredient_name" binding:"required"`
}

func (server *Server) createIngredient(ctx *gin.Context) {
	var req createIngredientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	Ingredient, err := server.store.CreateIngredient(ctx, req.IngredientName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, Ingredient)
}

type getIngredientRequest struct {
	IngredientID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getIngredient(ctx *gin.Context) {
	var req getIngredientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	Ingredient, err := server.store.GetIngredient(ctx, req.IngredientID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Ingredient)
}

type listIngredientRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listIngredients(ctx *gin.Context) {
	var req listIngredientRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListIngredientParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	Ingredients, err := server.store.ListIngredient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Ingredients)
}

type deleteIngredientRequest struct {
	IngredientID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteIngredient(ctx *gin.Context) {
	var req deleteIngredientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteIngredient(ctx, req.IngredientID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, err)
}

type updateIngredientRequest struct {
	IngredientID   int64   `json:"ingredient_id" binding:"required,min=1"`
	IngredientName *string `json:"ingredient_name"`
}

func (server *Server) updateIngredient(ctx *gin.Context) {
	var req updateIngredientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	existingIngredient, err := server.store.GetIngredient(ctx, req.IngredientID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	updateExistingIngredientDetails(&req, &existingIngredient)

	arg := db.UpdateIngredientParams{
		IngredientID:   existingIngredient.IngredientID,
		IngredientName: existingIngredient.IngredientName,
	}
	Ingredients, err := server.store.UpdateIngredient(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Ingredients)
}

func updateExistingIngredientDetails(in *updateIngredientRequest, existingIngredientDetails *db.Ingredient) {
	if in.IngredientName != nil {
		existingIngredientDetails.IngredientName = *in.IngredientName
	}
}
