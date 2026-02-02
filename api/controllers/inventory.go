package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/core/contracts"
	dto "github.com/oTeeLeko/mystore/core/model"
	"github.com/oTeeLeko/mystore/util"
)

type InventoryHandler struct {
	inventoryRepo contracts.InventoryRepository
}

func NewInventoryHandler(inventoryRepo contracts.InventoryRepository) *InventoryHandler {
	return &InventoryHandler{inventoryRepo: inventoryRepo}
}

// CreateInventory godoc
// @Summary Create a new inventory item
// @Description Create a new inventory item
// @Tags inventories
// @Accept  json
// @Produce  json
// @Param inventory body dto.CreateInventoryRequest true "Inventory"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/inventories/create [post]
func (h *InventoryHandler) CreateInventory(ctx *gin.Context) {
	var req dto.CreateInventoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.inventoryRepo.CreateInventory(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Inventory created successfully", nil))
}

// UpdateInventory godoc
// @Summary Update an inventory item
// @Description Update an inventory item
// @Tags inventories
// @Accept  json
// @Produce  json
// @Param id path int true "Inventory ID"
// @Param inventory body dto.UpdateInventoryRequest true "Inventory"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/inventories/{id} [put]
func (h *InventoryHandler) UpdateInventory(ctx *gin.Context) {
	var req dto.UpdateInventoryRequest
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

	if err := h.inventoryRepo.UpdateInventory(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Inventory updated successfully", nil))
}

// DeleteInventory godoc
// @Summary Delete an inventory item
// @Description Delete an inventory item
// @Tags inventories
// @Accept  json
// @Produce  json
// @Param id path int true "Inventory ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/inventories/{id} [delete]
func (h *InventoryHandler) DeleteInventory(ctx *gin.Context) {
	var req dto.InventoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.inventoryRepo.DeleteInventory(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Inventory deleted successfully", nil))
}

// GetInventory godoc
// @Summary Get an inventory item
// @Description Get an inventory item by ID
// @Tags inventories
// @Accept  json
// @Produce  json
// @Param id path int true "Inventory ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/inventories/{id} [get]
func (h *InventoryHandler) GetInventory(ctx *gin.Context) {
	var req dto.InventoryRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	inventory, err := h.inventoryRepo.GetInventory(ctx, &req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Inventory retrieved successfully", inventory))
}

// GetListInventories godoc
// @Summary Get list of inventory items
// @Description Get all inventory items
// @Tags inventories
// @Accept  json
// @Produce  json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/inventories/list [get]
func (h *InventoryHandler) GetListInventories(ctx *gin.Context) {
	inventories, err := h.inventoryRepo.GetListInventories(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Inventories retrieved successfully", inventories))
}
