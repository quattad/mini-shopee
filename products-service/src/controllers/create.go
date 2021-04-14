package controllers

import (
	"encoding/json"
	"fmt"
	"github/quattad/mini-shopee/products-service/models"
	"github/quattad/mini-shopee/products-service/src/api/responses"
	"github/quattad/mini-shopee/products-service/src/config"
	"github/quattad/mini-shopee/products-service/src/db"
	"github/quattad/mini-shopee/products-service/src/repository"
	"github/quattad/mini-shopee/products-service/src/repository/crud"
	"io/ioutil"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	body, err := ioutil.ReadAll(r.Body) // read from req body which is []byte

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	err = json.Unmarshal(body, &product)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	product.Prepare()
	product.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := db.DBService.Connect(config.DBDRIVER, config.DBURL)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	repo := crud.NewRepositoryProductsCrud(db)

	func(pr repository.ProductRepository) {
		product, err := pr.Save(product)

		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}

		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, product.ID))
		responses.JSON(w, http.StatusCreated, product)
	}(repo)
}
