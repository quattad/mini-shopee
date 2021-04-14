package routes

import (
	"github/quattad/mini-shopee/products-service/src/controllers"
	"net/http"
)

var productsRoutes = []Route{
	Route{
		Uri:          "/products",
		Method:       http.MethodGet,
		Handler:      controllers.GetProducts,
		AuthRequired: false,
	},
	Route{
		Uri:          "/products/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetProduct,
		AuthRequired: false,
	},
	Route{
		Uri:          "/products",
		Method:       http.MethodPost,
		Handler:      controllers.CreateProduct,
		AuthRequired: true,
	},
	Route{
		Uri:          "/products/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateProduct,
		AuthRequired: true,
	},
	Route{
		Uri:          "/products/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeleteProduct,
		AuthRequired: true,
	},
}
