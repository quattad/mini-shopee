package routes

import (
	"github/quattad/mini-shopee/products-service/src/api/controllers"
	"net/http"
)

var productsRoutes = []Route{
	{
		Uri:          "/products",
		Method:       http.MethodGet,
		Handler:      controllers.GetProducts,
		AuthRequired: false,
	},
	{
		Uri:          "/products/{id}",
		Method:       http.MethodGet,
		Handler:      controllers.GetProduct,
		AuthRequired: false,
	},
	{
		Uri:          "/products",
		Method:       http.MethodPost,
		Handler:      controllers.CreateProduct,
		AuthRequired: false,
	},
	{
		Uri:          "/products/{id}",
		Method:       http.MethodPut,
		Handler:      controllers.UpdateProduct,
		AuthRequired: false,
	},
	{
		Uri:          "/products/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeleteProduct,
		AuthRequired: false,
	},
}
