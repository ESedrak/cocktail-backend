package api

import (
	db "cocktail-backend/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createMeasurementUnitRequest struct {
	MeasurementUnit string `json:"unit" binding:"required"`
}

func (server *Server) createMeasurementUnit(ctx *gin.Context) {
	var req createMeasurementUnitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	measurementUnit, err := server.store.CreateMeasurementUnit(ctx, req.MeasurementUnit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, measurementUnit)
}

type getMeasurementUnitRequest struct {
	MeasurementUnitID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getMeasurementUnit(ctx *gin.Context) {
	var req getMeasurementUnitRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	measurementUnit, err := server.store.GetMeasurementUnit(ctx, req.MeasurementUnitID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, measurementUnit)
}

type listMeasurementUnitsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMeasurementUnits(ctx *gin.Context) {
	var req listMeasurementUnitsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListMeasurementUnitsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	measurementUnits, err := server.store.ListMeasurementUnits(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, measurementUnits)
}

type deleteMeasurementUnitRequest struct {
	MeasurementUnitID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteMeasurementUnit(ctx *gin.Context) {
	var req deleteMeasurementUnitRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteMeasurementUnits(ctx, req.MeasurementUnitID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, err)
}

type updateMeasurementUnitRequest struct {
	MeasurementUnitID int64   `json:"measurement_unit_id" binding:"required,min=1"`
	Unit              *string `json:"unit"`
}

func (server *Server) updateMeasurementUnit(ctx *gin.Context) {
	var req updateMeasurementUnitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	existingMeasurementUnit, err := server.store.GetMeasurementUnit(ctx, req.MeasurementUnitID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	updateExistingMeasurementUnitDetails(&req, &existingMeasurementUnit)

	arg := db.UpdateMeasurementUnitsParams{
		MeasurementUnitsID: existingMeasurementUnit.MeasurementUnitsID,
		Unit:               existingMeasurementUnit.Unit,
	}
	measurementUnit, err := server.store.UpdateMeasurementUnits(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, measurementUnit)
}

func updateExistingMeasurementUnitDetails(in *updateMeasurementUnitRequest, existingMeasurementUnitDetails *db.MeasurementUnit) {
	if in.Unit != nil {
		existingMeasurementUnitDetails.Unit = *in.Unit
	}
}
