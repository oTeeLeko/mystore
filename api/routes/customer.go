package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/api/controllers"
)

func CustomerRoutes(router *gin.Engine, h *controllers.CustomerHandler) {
	customerGroup := router.Group("api/customers")
	customerGroup.POST("/create", h.CreateCustomer)
	customerGroup.GET("/:id", h.GetCustomer)
	customerGroup.GET("/list", h.GetListCustomers)
	customerGroup.DELETE("/:id", h.DeleteCustomer)
	customerGroup.PUT("/:id", h.UpdateCustomer)
}
