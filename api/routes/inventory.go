package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/api/controllers"
)

func InventoryRoutes(router *gin.Engine, h *controllers.InventoryHandler) {
	inventoryGroup := router.Group("api/inventories")
	inventoryGroup.POST("/create", h.CreateInventory)
	inventoryGroup.GET("/:id", h.GetInventory)
	inventoryGroup.GET("/list", h.GetListInventories)
	inventoryGroup.DELETE("/:id", h.DeleteInventory)
	inventoryGroup.PUT("/:id", h.UpdateInventory)
}
