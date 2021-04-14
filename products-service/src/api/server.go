package api

import (
	"fmt"
	"github/quattad/mini-shopee/products-service/src/config"
	"github/quattad/mini-shopee/products-service/src/router"
	"log"
	"net/http"
)

// Loads configurations runs Listen function based on port
func Run() {
	config.Load()
	fmt.Printf("\n \t Listening on port [::]:%d\n", config.PORT)
	Listen(config.PORT)
}

func Listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
