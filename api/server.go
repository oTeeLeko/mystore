package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/oTeeLeko/mystore/db/sqlc"
	"github.com/oTeeLeko/mystore/middleware"
	"github.com/oTeeLeko/mystore/util"
)

type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(middleware.Logger())

	CustomerRoutes(router, server)

	server.router = router
}

func CustomerRoutes(router *gin.Engine, server *Server) {
	customerGroup := router.Group("api/customers")
	customerGroup.POST("/create", server.createCustomer)
	customerGroup.GET("", server.getCustomerByID)
	customerGroup.GET("/list", server.getListCustomers)
	customerGroup.DELETE("/delete", server.deleteCustomer)
	customerGroup.PUT("/update/:id", server.updateCustomer)
}

func ProductRoutes(router *gin.Engine, server *Server) {
	productGroup := router.Group("api/products")
	productGroup.POST("/create", server.createProduct)
	productGroup.GET("", server.getProductByID)
	productGroup.GET("/list", server.getListProducts)
	productGroup.DELETE("/delete", server.deleteProduct)
	productGroup.PUT("/update/:id", server.updateProduct)
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func successResponse(message string) gin.H {
	return gin.H{
		"status":  "success",
		"message": fmt.Sprintf("%s successfully", message),
	}
}
