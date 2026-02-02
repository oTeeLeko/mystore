package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/oTeeLeko/mystore/api"
	"github.com/oTeeLeko/mystore/configs"
	store "github.com/oTeeLeko/mystore/core/sqlstore"
	_ "github.com/oTeeLeko/mystore/docs"
	"github.com/oTeeLeko/mystore/util"
)

// @title MyStore API
// @version 1.0
// @description MyStore API Documentation
// @contact.name Supakan Sriwichai
// @BasePath  /
func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	configs.ConnectDatabase()
	db := configs.DB

	store := store.NewStore(db)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
