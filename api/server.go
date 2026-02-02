package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/oTeeLeko/mystore/api/controllers"
	"github.com/oTeeLeko/mystore/api/routes"
	store "github.com/oTeeLeko/mystore/core/sqlstore"
	"github.com/oTeeLeko/mystore/middleware"
	"github.com/oTeeLeko/mystore/util"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	config util.Config
	store  *store.Store
	router *gin.Engine
}

func NewServer(config util.Config, store *store.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(middleware.AccessLogger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	routes.CustomerRoutes(router, controllers.NewCustomerHandler(server.store.CustomerRepo))
	routes.ProductRoutes(router, controllers.NewProductHandler(server.store.ProductRepo))
	routes.InventoryRoutes(router, controllers.NewInventoryHandler(server.store.InventoryRepo))
	routes.OrderRoutes(router, controllers.NewOrderHandler(server.store.OrderRepo))

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
