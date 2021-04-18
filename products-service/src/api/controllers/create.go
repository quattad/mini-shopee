package controllers

import (
	"encoding/json"
	"fmt"
	"github/quattad/mini-shopee/products-service/src/api/config"
	"github/quattad/mini-shopee/products-service/src/api/db"
	"github/quattad/mini-shopee/products-service/src/api/models"
	"github/quattad/mini-shopee/products-service/src/api/repository"
	"github/quattad/mini-shopee/products-service/src/api/repository/crud"
	"github/quattad/mini-shopee/products-service/src/api/responses"
	"io/ioutil"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	body, err := ioutil.ReadAll(r.Body) // read from req body which is []byte

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = json.Unmarshal(body, &product)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	product.Prepare()
	err = product.Validate("")

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := db.DBService.Connect(config.DBDRIVER, config.DBURL)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
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
