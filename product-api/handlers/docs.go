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

import "github.com/jimil749/go-microservices-practice/product-api/data"

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

//	A list of products returns in the response
//	swagger:response productsResponse
type productsResponse struct {
	//	All products in the system
	//	in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

//	swagger:parameters updateProduct
type productIDParameterWrapper struct {
	//	The id of the product to delete from database
	//	in: path
	//	required: true
	ID int `json:"id"`
}
