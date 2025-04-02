package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/oTeeLeko/mystore/db/sqlc"
	"github.com/oTeeLeko/mystore/model"
)

func (server *Server) createCustomer(ctx *gin.Context) {
	var req model.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.AddCustomerParams{
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Gender:       req.Gender,
		Tel:          req.Tel,
		EmailAddress: req.EmailAddress,
	}

	if err := server.store.AddCustomer(ctx, arg); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Create"))

}

func (server *Server) getCustomerByID(ctx *gin.Context) {
	var req model.GetIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	customer, err := server.store.GetCustomer(ctx, req.ID)
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

func (server *Server) getListCustomers(ctx *gin.Context) {
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

	arg := db.GetListCustomersParams{
		Limit:  limit,
		Offset: offset,
	}

	customers, err := server.store.GetListCustomers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, customers)
}

func (server *Server) deleteCustomer(ctx *gin.Context) {
	var req model.GetIDRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := server.store.DeleteCustomer(ctx, req.ID); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Delete"))
}

func (server *Server) updateCustomer(ctx *gin.Context) {
	var req model.UpdateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCustomerParams{
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		Gender:       req.Gender,
		Tel:          req.Tel,
		EmailAddress: req.EmailAddress,
		ID:           req.ID,
	}

	if err := server.store.UpdateCustomer(ctx, arg); err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, successResponse("Update"))
}
