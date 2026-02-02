package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/core/contracts"
	dto "github.com/oTeeLeko/mystore/core/model"
	"github.com/oTeeLeko/mystore/util"
)

type OrderHandler struct {
	orderRepo contracts.OrderRepository
}

func NewOrderHandler(orderRepo contracts.OrderRepository) *OrderHandler {
	return &OrderHandler{orderRepo: orderRepo}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param order body dto.CreateOrderRequest true "Order"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/orders/create [post]
func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req dto.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.orderRepo.CreateOrder(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Order created successfully", nil))
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update an order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Param order body dto.UpdateOrderRequest true "Order"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/orders/{id} [put]
func (h *OrderHandler) UpdateOrder(ctx *gin.Context) {
	var req dto.UpdateOrderRequest
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

	if err := h.orderRepo.UpdateOrder(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Order updated successfully", nil))
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(ctx *gin.Context) {
	var req dto.OrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.orderRepo.DeleteOrder(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Order deleted successfully", nil))
}

// GetOrder godoc
// @Summary Get an order
// @Description Get an order by ID
// @Tags orders
// @Accept  json
// @Produce  json
// @Param id path int true "Order ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/orders/{id} [get]
func (h *OrderHandler) GetOrder(ctx *gin.Context) {
	var req dto.OrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	order, err := h.orderRepo.GetOrder(ctx, &req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Order retrieved successfully", order))
}

// GetListOrders godoc
// @Summary Get list of orders
// @Description Get all orders
// @Tags orders
// @Accept  json
// @Produce  json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/orders/list [get]
func (h *OrderHandler) GetListOrders(ctx *gin.Context) {
	orders, err := h.orderRepo.GetListOrders(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Orders retrieved successfully", orders))
}
