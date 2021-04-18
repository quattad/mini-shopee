package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// PORT stores port number, loaded from env variable
var PORT int

// DBURL stores database URL
var DBURL string

// DBDRIVER stores db type e.g. MySQL
var DBDRIVER string

// SECRETKEY stores hash key of API used to generate JWT
var SECRETKEY []byte

func Load() {
	var err error

	if os.Getenv("ENV_SOURCE") != "docker" {
		fmt.Println("Loading envs from .env ... ")
		err = godotenv.Load(os.ExpandEnv("$GOPATH/src/github.com/quattad/mini-shopee/products-service/src/api/env/.env"))

		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Docker env detected ... ")
	}

	PORT, err = strconv.Atoi((os.Getenv("PORT")))

	if err != nil {
		log.Println(err)
		PORT = 9000
	}

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
}
