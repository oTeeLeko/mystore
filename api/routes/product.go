package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/api/controllers"
)

func ProductRoutes(router *gin.Engine, h *controllers.ProductHandler) {
	productGroup := router.Group("api/products")
	productGroup.POST("/create", h.CreateProduct)
	productGroup.GET("/:id", h.GetProduct)
	productGroup.GET("/list", h.GetListProducts)
	productGroup.DELETE("/:id", h.DeleteProduct)
	productGroup.PUT("/:id", h.UpdateProduct)
}
