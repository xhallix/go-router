# Custom Router build in Golang

A small router implementation, primary build as a training purpose for looking at golang

## Usage

```go

func main() {
    
	r := router.Router{}

	// setup routes
	r.GET("/foo", func(writer http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("views/templates/index.html")
		pageData := struct {
			Title string
		}{
			"Homepage",
		}
		t.Execute(writer, pageData)
	})

	http.ListenAndServe(":8000", http.HandlerFunc(r.RouteHandler))
}

``` 