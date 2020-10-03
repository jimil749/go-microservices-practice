package handlers

import (
	"net/http"

	"github.com/jimil749/go-microservices-practice/product-api/data"
)

//	swagger:route GET /products products listProducts
//	Returns a list of products
//	responses:
//		200: productsResponse

//ListAll gets the products
func (p *Products) ListAll(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, rw)

	if err != nil {
		p.l.Println("[ERROR] serializing the product", err)
	}
}
