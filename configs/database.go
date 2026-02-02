package configs

import (
	"log"

	tbl "github.com/oTeeLeko/mystore/models"
	"github.com/oTeeLeko/mystore/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	database, err := gorm.Open(mysql.Open(config.DBSource), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	err = database.AutoMigrate(
		&tbl.Customers{},
		&tbl.Products{},
		&tbl.Orders{},
		&tbl.Inventory{},
	)

	if err != nil {
		log.Fatal("auto migrate failed:", err)
	}

	DB = database
}
