package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/jimil749/go-microservices-practice/product-api/data"
)

// Products struct
type Products struct {
	l *log.Logger
	v *data.Validation
}

// NewProducts creates a new products handler with the given logger.
func NewProducts(l *log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

//KeyProduct for context key
type KeyProduct struct{}

//ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

//GenericError is a generic message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

//ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"message"`
}

func getProductID(r *http.Request) int {
	//	parse the product id from the url
	vars := mux.Vars(r)

	//	convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id

}
