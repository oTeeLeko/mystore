package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/core/contracts"
	dto "github.com/oTeeLeko/mystore/core/model"
	"github.com/oTeeLeko/mystore/util"
)

type CustomerHandler struct {
	customerRepo contracts.CustomerRepository
}

func NewCustomerHandler(customerRepo contracts.CustomerRepository) *CustomerHandler {
	return &CustomerHandler{customerRepo: customerRepo}
}

// CreateCustomer godoc
// @Summary Create a new customer
// @Description Create a new customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Param customer body dto.CreateCustomerRequest true "Customer"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/customers/create [post]
func (h *CustomerHandler) CreateCustomer(ctx *gin.Context) {
	var req dto.CreateCustomerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.customerRepo.CreateCustomer(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Customer created successfully", nil))
}

// UpdateCustomer godoc
// @Summary Update a customer
// @Description Update a customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Param customer body dto.UpdateCustomerRequest true "Customer"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/customers/{id} [put]
func (h *CustomerHandler) UpdateCustomer(ctx *gin.Context) {
	var req dto.UpdateCustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.customerRepo.UpdateCustomer(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Customer updated successfully", nil))
}

// DeleteCustomer godoc
// @Summary Delete a customer
// @Description Delete a customer
// @Tags customers
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/customers/{id} [delete]
func (h *CustomerHandler) DeleteCustomer(ctx *gin.Context) {
	var req dto.CustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.customerRepo.DeleteCustomer(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Customer deleted successfully", nil))
}

// GetCustomer godoc
// @Summary Get a customer
// @Description Get a customer by ID
// @Tags customers
// @Accept  json
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/customers/{id} [get]
func (h *CustomerHandler) GetCustomer(ctx *gin.Context) {
	var req dto.CustomerRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	customer, err := h.customerRepo.GetCustomer(ctx, &req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Customer retrieved successfully", customer))
}

// GetListCustomers godoc
// @Summary Get list of customers
// @Description Get all customers
// @Tags customers
// @Accept  json
// @Produce  json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/customers/list [get]
func (h *CustomerHandler) GetListCustomers(ctx *gin.Context) {
	customers, err := h.customerRepo.GetListCustomers(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Customers retrieved successfully", customers))
}
