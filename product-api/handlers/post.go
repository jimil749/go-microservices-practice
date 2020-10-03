package handlers

import (
	"net/http"

	"github.com/jimil749/go-microservices-practice/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//		200: productResponse
//  422: errorValidation
//  501: errorResponse

//Create handles POST requests to add new products
func (p *Products) Create(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Println("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(prod)
}
