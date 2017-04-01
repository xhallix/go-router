package router

import (
	"fmt"
	"net/http"
)

// Router router
type Router struct {
}

// NotFoundResponseWriter todo
type NotFoundResponseWriter struct {
	http.ResponseWriter
	status int
}

// HTTPHandler Handles http requests
type HTTPHandler func(http.ResponseWriter, *http.Request)

// RouteHandler Used in listen and serve to catch all requests
func (r *Router) RouteHandler(writer http.ResponseWriter, req *http.Request) {
	ServerMux := http.DefaultServeMux
	ServerMux.ServeHTTP(writer, req)
}

// GET handles get request
func (r *Router) GET(path string, handler HTTPHandler) {
	if path[0] != '/' {
		path = "/" + path
	}
	http.HandleFunc(path, func(writer http.ResponseWriter, req *http.Request) {
		handler(writer, req)
		fmt.Print("TTEST")
	})

}
