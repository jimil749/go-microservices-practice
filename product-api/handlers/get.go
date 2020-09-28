package handlers

import (
	"net/http"

	"github.com/jimil749/go-microservices-practice/product-api/data"
)

//	swagger:route GET /products products listProducts
//	Returns a list of products

//ListAll gets the products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal to json", http.StatusInternalServerError)
	}
}
