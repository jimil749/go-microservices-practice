// Package classification of Product API
//
// Documentation for Product API
//
//    Schemes: http
//    BasePath: /
//    Version: 1.0.0
//
//    Consumes:
//    - application/json
//
//    Produces:
//    - application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jimil749/go-microservices-practice/product-api/data"
)

// Products struc
type Products struct {
	l *log.Logger
}

// NewProducts creates a new products handler with the given logger.
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

//KeyProduct for context key
type KeyProduct struct{}

//MiddleWareValidateProduct serializes the list from json to go struct
func (p Products) MiddleWareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)

		if err != nil {
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		//validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
