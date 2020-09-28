package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jimil749/go-microservices-practice/product-api/data"
)

//Delete deletes the item from the list
func (p *Products) Delete(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, "Unable to convertID", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
	if err != nil {
		http.Error(rw, "Unale to find the element", http.StatusNotFound)
		return
	}
}
