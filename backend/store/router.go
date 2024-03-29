package store

import (
	"github.com/gorilla/mux"
	"net/http"
)

var controller = &Controller{Repository: Repository{}}

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/get_token",
		controller.GetToken,
	},
	Route{
		"GetCatalog",
		"GET",
		"/products",
		controller.GetCatalog,
	},
	Route{
		"AddProduct",
		"POST",
		"/product",
		AuthMiddleware(controller.AddProduct),
	},
	Route{
		"UpdateProduct",
		"PUT",
		"/product",
		AuthMiddleware(controller.UpdateProduct),
	},
	Route{
		"GetProduct",
		"GET",
		"/products/{id}",
		controller.GetProduct,
	},
	Route{
		"DeleteProduct",
		"DELETE",
		"/product/{id}",
		AuthMiddleware(controller.DeleteProduct),
	},
	Route{
		"RegisterUser",
		"POST",
		"/register",
		controller.RegisterUser,
	}}

func CreateRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
