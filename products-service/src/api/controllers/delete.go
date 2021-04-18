package controllers

import (
	"fmt"
	"github/quattad/mini-shopee/products-service/src/api/config"
	"github/quattad/mini-shopee/products-service/src/api/db"
	"github/quattad/mini-shopee/products-service/src/api/repository"
	"github/quattad/mini-shopee/products-service/src/api/repository/crud"
	"github/quattad/mini-shopee/products-service/src/api/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)

	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	db, err := db.DBService.Connect(config.DBDRIVER, config.DBURL)

	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
	}

	repo := crud.NewRepositoryProductsCrud(db)

	func(pr repository.ProductRepository) {
		_, err := pr.Delete(uint32(pid))

		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
		}

		w.Header().Set("Entity", fmt.Sprintf("%d", pid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
