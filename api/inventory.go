package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/oTeeLeko/mystore/db/sqlc"
	"github.com/oTeeLeko/mystore/model"
)

func (server *Server) createInventory(ctx *gin.Context) {
	var req model.CreateInventoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddInventoryParams{
		Productid: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := server.store.AddInventory(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Create"))
}

func (server *Server) updateInventory(ctx *gin.Context) {
	var req model.UpdateInventoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateInventoryQuantityParams{
		Productid: req.ProductID,
		Quantity:  req.Quantity,
	}

	if err := server.store.UpdateInventoryQuantity(ctx, arg); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Update"))
}

func (server *Server) getInventoryByID(ctx *gin.Context) {
	var req model.GetIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	customer, err := server.store.GetInventory(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customer)
}

func (server *Server) getListInventories(ctx *gin.Context) {
	var req model.PaginationRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	limit := int32(10000)
	offset := int32(0)

	if req.PageSize > 0 {
		limit = req.PageSize
	}
	if req.PageID > 0 {
		offset = (req.PageID - 1) * req.PageSize
	}

	arg := db.GetListInventoriesParams{
		Limit:  limit,
		Offset: offset,
	}

	customers, err := server.store.GetListInventories(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (server *Server) deleteInventory(ctx *gin.Context) {
	var req model.GetIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteInventory(ctx, req.ID); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Delete"))
}
