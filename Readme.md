# Custom Router build in Golang

A small router implementation, primary build as a training purpose for looking at golang

## Usage

```go

func main() {
    
	r := router.Router{}

	// declare all routes
	appRoutes := []string{
		"/",
		"/homepage",
		"/about",
	}

	// set the routes
	r.SetRoutes(appRoutes)

	r.SetNotFoundTemplate("views/templates/errors/notfound.html")

	// set actions for routes
	r.GET("/foo", func(writer http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("views/templates/index.html")
		pageData := struct {
			Title string
		}{
			"Homepage",
		}
		t.Execute(writer, pageData)
	})

	// use the route as a handler
	http.ListenAndServe(":8000", http.HandlerFunc(r.RouteHandler))
}

``` 