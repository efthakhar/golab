package main

import (
	"fmt"
	//"net/http"
	// "strings"

	"goapp/reusables/application"
)

func main() {

	// http.HandleFunc("/user/{create}/{id}", userHandler)
	// http.ListenAndServe(":8080", nil)

	application := application.NewApplication("Testing Application", "development")

	fmt.Println(application.Name)
}


// func userHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, r.URL.Path)
// }
