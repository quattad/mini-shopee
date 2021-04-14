package controllers

import (
	"fmt"
	"github/quattad/mini-shopee/products-service/src/api/responses"
	"github/quattad/mini-shopee/products-service/src/config"
	"github/quattad/mini-shopee/products-service/src/db"
	"github/quattad/mini-shopee/products-service/src/repository"
	"github/quattad/mini-shopee/products-service/src/repository/crud"
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

	db, err := db.DBService.Connect(config.DBURL, config.DBDRIVER)

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
