package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/api/controllers"
)

func OrderRoutes(router *gin.Engine, h *controllers.OrderHandler) {
	orderGroup := router.Group("api/orders")
	orderGroup.POST("/create", h.CreateOrder)
	orderGroup.GET("/:id", h.GetOrder)
	orderGroup.GET("/list", h.GetListOrders)
	orderGroup.DELETE("/:id", h.DeleteOrder)
	orderGroup.PUT("/:id", h.UpdateOrder)
}
