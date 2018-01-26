package main

import (
	"net/http"

	"github.com/dickmanben/rebel-api/adsets"
	"github.com/dickmanben/rebel-api/alert"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PassThru",
		"GET",
		"/Passthru/{AdSetID}",
		adsets.PassThru,
	},
	Route{
		"CreateAdSet",
		"POST",
		"/Adset",
		adsets.CreateAdset,
	},
	Route{
		"UpdateAdset",
		"POST",
		"/Adset/{AdSetID}",
		adsets.UpdateAdset,
	},
	Route{
		"GetAdset",
		"GET",
		"/Adset/{AdSetID}",
		adsets.GetAdset,
	},
	Route{
		"GetAdsets",
		"GET",
		"/Adset",
		adsets.GetAdsets,
	},
	Route{
		"Alert",
		"POST",
		"/alert",
		alert.All,
	},
}
