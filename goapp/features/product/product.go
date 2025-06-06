package product

import (
	"net/http"

	"goapp/reusables/helpers"
)

type Product struct {
}

func getProductList(w http.ResponseWriter, r *http.Request) {
	helpers.NetHttpJsonResponse(w, map[string]string{"msg": "Product List"}, 200)
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	data := "Single Product of ID: " + id
	helpers.NetHttpJsonResponse(w, data, 200)
}

// instead of instance of router pass Application instance pointer here and 
// extract the router *http.ServeMux from the application property
func AddProductFeature(router *http.ServeMux) {
	router.HandleFunc("GET /products", getProductList)
	router.HandleFunc("GET /products/{id}", getProductById)
}
