// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PutProductsNameNoContentCode is the HTTP code returned for type PutProductsNameNoContent
const PutProductsNameNoContentCode int = 204

/*PutProductsNameNoContent Product successfuly withdrawn from the fridge

swagger:response putProductsNameNoContent
*/
type PutProductsNameNoContent struct {
}

// NewPutProductsNameNoContent creates PutProductsNameNoContent with default headers values
func NewPutProductsNameNoContent() *PutProductsNameNoContent {

	return &PutProductsNameNoContent{}
}

// WriteResponse to the client
func (o *PutProductsNameNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PutProductsNameNotFoundCode is the HTTP code returned for type PutProductsNameNotFound
const PutProductsNameNotFoundCode int = 404

/*PutProductsNameNotFound No such product in the fridge

swagger:response putProductsNameNotFound
*/
type PutProductsNameNotFound struct {
}

// NewPutProductsNameNotFound creates PutProductsNameNotFound with default headers values
func NewPutProductsNameNotFound() *PutProductsNameNotFound {

	return &PutProductsNameNotFound{}
}

// WriteResponse to the client
func (o *PutProductsNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
