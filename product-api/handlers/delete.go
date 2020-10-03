package handlers

import (
	"net/http"

	"github.com/jimil749/go-microservices-practice/product-api/data"
)

//	swagger:route DELETE /products/{id} products deleteProducts
//	Update a list of products
//
//	responses:
//		201: noContent
//	404: errorResponse
//	404: errorResponse

//Delete handles DELETE requests and removes items from database
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] deleting record id", id)

	err := data.DeleteProduct(id)

	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] deleting record id does not exist")

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	if err != nil {
		p.l.Println("[ERROR] deleting record", err)

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
