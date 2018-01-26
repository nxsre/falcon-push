package main

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = accessLogger(handler, route.Name)
		methods := []string{}
		if strings.ToLower(route.Method) == "any" {
			methods = []string{"POST", "PUT", "HEAD", "GET", "DELETE", "PATCH"}
		} else {
			methods = strings.Split(route.Method, ",")
		}
		router.
			Methods(methods...).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
