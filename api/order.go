package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/oTeeLeko/mystore/db/sqlc"
	"github.com/oTeeLeko/mystore/model"
)

func (server *Server) createOrder(ctx *gin.Context) {
	var req model.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddOrderParams{
		Customerid: req.CustomerID,
		Productid:  req.ProductID,
		Quantity:   req.Quantity,
	}

	if err := server.store.AddOrder(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Create"))
}

func (server *Server) getOrderByID(ctx *gin.Context) {
	var req model.GetIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	customer, err := server.store.GetOrder(ctx, req.ID)
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

func (server *Server) getListOrders(ctx *gin.Context) {
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

	arg := db.GetListOrdersParams{
		Limit:  limit,
		Offset: offset,
	}

	customers, err := server.store.GetListOrders(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (server *Server) deleteOrder(ctx *gin.Context) {
	var req model.GetIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteOrder(ctx, req.ID); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Delete"))
}
