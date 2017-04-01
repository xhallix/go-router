package router

import (
	"html/template"
	"net/http"
)

// gobal variables
// routePaths Lists all url paths for routing
var routePaths = new([]string)

// defines the 404 template
var notFoundTemplate = "notfound.html"

// Router router
type Router struct {
	NotFoundTemplate string
}

// SetRoutes Define all routes used in this router
func (r *Router) SetRoutes(routes []string) {
	routePaths = &routes
}

// SetNotFoundTemplate Allows to customize a NotFoundTemplate
func (r *Router) SetNotFoundTemplate(template string) {
	notFoundTemplate = template
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

	routeFound := false
	for _, v := range *routePaths {
		if v == req.URL.String() {
			routeFound = true
		}
	}

	if routeFound == false {
		handleNotFoundPage()
	}

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
	})
}

func handleNotFoundPage() {
	template.ParseFiles(notFoundTemplate)
}
