package api

import (
	db "cocktail-backend/db/sqlc"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createMeasurementQtyRequest struct {
	QtyAmount int64 `json:"qty_amount" binding:"required"`
}

func (server *Server) createMeasurementQty(ctx *gin.Context) {
	var req createMeasurementQtyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	measurementQty, err := server.store.CreateMeasurementQty(ctx, req.QtyAmount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, measurementQty)
}

type getMeasurementQtyRequest struct {
	MeasurementQtyID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getMeasurementQty(ctx *gin.Context) {
	var req getMeasurementQtyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	measurementQty, err := server.store.GetMeasurementQty(ctx, req.MeasurementQtyID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, measurementQty)
}

type listMeasurementQtyRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMeasurementQty(ctx *gin.Context) {
	var req listMeasurementQtyRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListMeasurementQtyParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	measurementQty, err := server.store.ListMeasurementQty(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, measurementQty)
}

type deleteMeasurementQtyRequest struct {
	MeasurementQtyID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteMeasurementQty(ctx *gin.Context) {
	var req deleteMeasurementQtyRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteMeasurementQty(ctx, req.MeasurementQtyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, err)
}

type updateMeasurementQtyRequest struct {
	MeasurementQtyID int64  `json:"measurement_qty_id" binding:"required,min=1"`
	QtyAmount        *int64 `json:"qty_amount"`
}

func (server *Server) updateMeasurementQty(ctx *gin.Context) {
	var req updateMeasurementQtyRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	existingMeasurementQty, err := server.store.GetMeasurementQty(ctx, req.MeasurementQtyID)

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	updateExistingMeasurementQtyDetails(&req, &existingMeasurementQty)

	arg := db.UpdateMeasurementQtyParams{
		MeasurementQtyID: existingMeasurementQty.MeasurementQtyID,
		QtyAmount:        existingMeasurementQty.QtyAmount,
	}
	measurementQty, err := server.store.UpdateMeasurementQty(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, measurementQty)
}

func updateExistingMeasurementQtyDetails(in *updateMeasurementQtyRequest, existingMeasurementQtyDetails *db.MeasurementQty) {
	if in.QtyAmount != nil {
		existingMeasurementQtyDetails.QtyAmount = *in.QtyAmount
	}
}
