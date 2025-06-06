package middlewares

import (
	"fmt"
	"net/http"
)

// ==== Middlewares
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Applying logging middleware...")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Applying authentication middleware...")
		// In a real scenario, you'd check auth here
		next.ServeHTTP(w, r)
	})
}

// ==== Group Middlewares
type Middleware func(http.Handler) http.Handler

var GlobalMiddlewares = []Middleware{
	LoggingMiddleware,
	AuthMiddleware,
}

// ==== Function to apply middleares to a handler
func ApplyMiddleware(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}
