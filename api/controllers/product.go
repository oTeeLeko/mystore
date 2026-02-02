package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/core/contracts"
	dto "github.com/oTeeLeko/mystore/core/model"
	"github.com/oTeeLeko/mystore/util"
)

type ProductHandler struct {
	productRepo contracts.ProductRepository
}

func NewProductHandler(productRepo contracts.ProductRepository) *ProductHandler {
	return &ProductHandler{productRepo: productRepo}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param product body dto.CreateProductRequest true "Product"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/products/create [post]
func (h *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req dto.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.productRepo.CreateProduct(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Product created successfully", nil))
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Param product body dto.UpdateProductRequest true "Product"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/products/{id} [put]
func (h *ProductHandler) UpdateProduct(ctx *gin.Context) {
	var req dto.UpdateProductRequest
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

	if err := h.productRepo.UpdateProduct(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Product updated successfully", nil))
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/products/{id} [delete]
func (h *ProductHandler) DeleteProduct(ctx *gin.Context) {
	var req dto.ProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := h.productRepo.DeleteProduct(ctx, &req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Product deleted successfully", nil))
}

// GetProduct godoc
// @Summary Get a product
// @Description Get a product by ID
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} util.APIResponse
// @Failure 400 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/products/{id} [get]
func (h *ProductHandler) GetProduct(ctx *gin.Context) {
	var req dto.ProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	product, err := h.productRepo.GetProduct(ctx, &req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Product retrieved successfully", product))
}

// GetListProducts godoc
// @Summary Get list of products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200 {object} util.APIResponse
// @Failure 500 {object} util.APIResponse
// @Router /api/products/list [get]
func (h *ProductHandler) GetListProducts(ctx *gin.Context) {
	products, err := h.productRepo.GetListProducts(ctx)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, util.SuccessResponse("Products retrieved successfully", products))
}
