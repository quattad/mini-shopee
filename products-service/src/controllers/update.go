package controllers

import (
	"encoding/json"
	"github/quattad/mini-shopee/products-service/models"
	"github/quattad/mini-shopee/products-service/src/api/responses"
	"github/quattad/mini-shopee/products-service/src/config"
	"github/quattad/mini-shopee/products-service/src/db"
	"github/quattad/mini-shopee/products-service/src/repository"
	"github/quattad/mini-shopee/products-service/src/repository/crud"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	product := models.Product{}

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	err = json.Unmarshal(body, &product)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := db.DBService.Connect(config.DBDRIVER, config.DBURL)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	repo := crud.NewRepositoryProductsCrud(db)

	go func(pr repository.ProductRepository) {
		rows, err := pr.Update(uint32(pid), product)

		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}

		responses.JSON(w, http.StatusOK, rows)
	}(repo)
}
