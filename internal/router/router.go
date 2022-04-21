package router

import (
	"Ys7/internal/router/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routers []route

var rootRoute = routers{
	route{
		"login",
		"POST",
		"/login",
		handlers.Login,
	},
}

var apiRoute = routers{
	route{
		"accessToken",
		"GET",
		"/token",
		handlers.AccessToken,
	},
	route{
		"deviceCapture",
		"POST",
		"/capture",
		handlers.DeviceCapture,
	},
	route{
		"deviceSearch",
		"POST",
		"/search",
		handlers.DeviceSearch,
	},
	route{
		"deviceList",
		"POST",
		"/list",
		handlers.DeviceList,
	},
	route{
		"deviceInfo",
		"POST",
		"/info",
		handlers.DeviceInfo,
	},
	route{
		"changeName",
		"POST",
		"/changename",
		handlers.ChangeName,
	},
	route{
		"encryptDevice",
		"POST",
		"/encrypt",
		handlers.Encrypt,
	},
	route{
		"decryptDevice",
		"POST",
		"/decrypt",
		handlers.Decrypt,
	},
	route{
		"liveAddress",
		"POST",
		"/live",
		handlers.LiveAddress,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.Use(SetCorsHeader)

	for _, route := range rootRoute {
		router.Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.Use(CheackCookie)

	for _, route := range apiRoute {
		apiRouter.Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	return router
}
