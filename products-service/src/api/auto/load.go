package auto

import (
	"fmt"
	"github/quattad/mini-shopee/products-service/src/api/config"
	"github/quattad/mini-shopee/products-service/src/api/db"
	"github/quattad/mini-shopee/products-service/src/api/models"
	"log"
)

func Load() {
	db, err := db.DBService.Connect(config.DBDRIVER, config.DBURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.Product{}).Error

	if err != nil {
		log.Fatal(err)
	}

	// Creates tables for models based on schema defined in package 'models'
	err = db.Debug().AutoMigrate(&models.Product{}).Error

	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range products {
		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error

		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Successfully connected to db ... ")
}
