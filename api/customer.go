package api

import (
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
