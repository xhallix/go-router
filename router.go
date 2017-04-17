package router

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

// gobal variables
// routePaths Lists all url paths for routing
var routePaths = new([]string)

// defines the 404 template
var notFoundTemplate = "notfound.html"

var assetPath = "/assets"

// Router router
type Router struct{}

// Init should be called after everything is setup in the router
func (r *Router) Init() {
	setupStaticPath()
}

// SetRoutes Define all routes used in this router
func (r *Router) SetRoutes(routes []string) {
	routePaths = &routes
}

// SetAssetsPath Define a custom asset path for your application
func (r *Router) SetAssetsPath(newAssetPath string) {
	assetPath = newAssetPath
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

// Redirect redirect to https
// thanks @https://gist.github.com/d-schmidt/587ceec34ce1334a5e60
func (r *Router) Redirect(w http.ResponseWriter, req *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + req.Host + req.URL.Path
	if len(req.URL.RawQuery) > 0 {
		target += "?" + req.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, req, target, http.StatusTemporaryRedirect)
}

// RouteHandler Used in listen and serve to catch all requests
func (r *Router) RouteHandler(writer http.ResponseWriter, req *http.Request) {

	routeFound := false
	for _, v := range *routePaths {
		if v == req.URL.String() {
			routeFound = true
		}
	}

	isAssetsRoute(req.URL.String())
	if routeFound == false && isAssetsRoute(req.URL.String()) == false {
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

func isAssetsRoute(url string) bool {
	return strings.Contains(url, assetPath)
}

func setupStaticPath() {

	assetDir := assetPath
	if assetPath[0] == '/' {
		assetDir = strings.Replace(assetDir, "/", "", -1)
	}
	http.Handle(assetPath+"/", http.StripPrefix(assetPath+"/", http.FileServer(http.Dir(assetDir))))
}
