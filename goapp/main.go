package main

import (
	"fmt"
	"net/http"
	"log"
	// "strings"

	"goapp/reusables/application"
	"goapp/reusables/helpers"
)

func main() {

	// Create application
	application := application.NewApplication("Testing Application", "development")
	fmt.Println(application.Name)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		helpers.NetHttpJsonResponse(w, map[string]string{"msg": "From helper!"}, 404)
	})


	http.ListenAndServe(":8080", mux)

}


func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("Request: %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r) // call the next handler
    })
}

