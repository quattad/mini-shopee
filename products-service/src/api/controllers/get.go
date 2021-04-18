package controllers

import (
	"github/quattad/mini-shopee/products-service/src/api/config"
	"github/quattad/mini-shopee/products-service/src/api/db"
	"github/quattad/mini-shopee/products-service/src/api/repository"
	"github/quattad/mini-shopee/products-service/src/api/repository/crud"
	"github/quattad/mini-shopee/products-service/src/api/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	db, err := db.DBService.Connect(config.DBDRIVER, config.DBURL)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	defer db.Close()

	repo := crud.NewRepositoryProductsCrud(db)

	func(pr repository.ProductRepository) {
		products, err := pr.FindAll()

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		responses.JSON(w, http.StatusOK, products)
	}(repo)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// uid is uint64
	uid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
	}

	db, err := db.DBService.Connect(config.DBDRIVER, config.DBURL)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	defer db.Close()

	repo := crud.NewRepositoryProductsCrud(db)

	func(pr repository.ProductRepository) {
		products, err := pr.FindById(uint32(uid))

		if err != nil {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}

		responses.JSON(w, http.StatusOK, products)
	}(repo)
}
