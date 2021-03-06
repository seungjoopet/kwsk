// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/projectodd/kwsk/models"
)

// GetActionByNameOKCode is the HTTP code returned for type GetActionByNameOK
const GetActionByNameOKCode int = 200

/*GetActionByNameOK Returned action

swagger:response getActionByNameOK
*/
type GetActionByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.Action `json:"body,omitempty"`
}

// NewGetActionByNameOK creates GetActionByNameOK with default headers values
func NewGetActionByNameOK() *GetActionByNameOK {

	return &GetActionByNameOK{}
}

// WithPayload adds the payload to the get action by name o k response
func (o *GetActionByNameOK) WithPayload(payload *models.Action) *GetActionByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action by name o k response
func (o *GetActionByNameOK) SetPayload(payload *models.Action) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionByNameUnauthorizedCode is the HTTP code returned for type GetActionByNameUnauthorized
const GetActionByNameUnauthorizedCode int = 401

/*GetActionByNameUnauthorized Unauthorized request

swagger:response getActionByNameUnauthorized
*/
type GetActionByNameUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionByNameUnauthorized creates GetActionByNameUnauthorized with default headers values
func NewGetActionByNameUnauthorized() *GetActionByNameUnauthorized {

	return &GetActionByNameUnauthorized{}
}

// WithPayload adds the payload to the get action by name unauthorized response
func (o *GetActionByNameUnauthorized) WithPayload(payload *models.ErrorMessage) *GetActionByNameUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action by name unauthorized response
func (o *GetActionByNameUnauthorized) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionByNameUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionByNameForbiddenCode is the HTTP code returned for type GetActionByNameForbidden
const GetActionByNameForbiddenCode int = 403

/*GetActionByNameForbidden Unauthorized request

swagger:response getActionByNameForbidden
*/
type GetActionByNameForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionByNameForbidden creates GetActionByNameForbidden with default headers values
func NewGetActionByNameForbidden() *GetActionByNameForbidden {

	return &GetActionByNameForbidden{}
}

// WithPayload adds the payload to the get action by name forbidden response
func (o *GetActionByNameForbidden) WithPayload(payload *models.ErrorMessage) *GetActionByNameForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action by name forbidden response
func (o *GetActionByNameForbidden) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionByNameForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionByNameNotFoundCode is the HTTP code returned for type GetActionByNameNotFound
const GetActionByNameNotFoundCode int = 404

/*GetActionByNameNotFound Item not found

swagger:response getActionByNameNotFound
*/
type GetActionByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionByNameNotFound creates GetActionByNameNotFound with default headers values
func NewGetActionByNameNotFound() *GetActionByNameNotFound {

	return &GetActionByNameNotFound{}
}

// WithPayload adds the payload to the get action by name not found response
func (o *GetActionByNameNotFound) WithPayload(payload *models.ErrorMessage) *GetActionByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action by name not found response
func (o *GetActionByNameNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetActionByNameInternalServerErrorCode is the HTTP code returned for type GetActionByNameInternalServerError
const GetActionByNameInternalServerErrorCode int = 500

/*GetActionByNameInternalServerError Server error

swagger:response getActionByNameInternalServerError
*/
type GetActionByNameInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetActionByNameInternalServerError creates GetActionByNameInternalServerError with default headers values
func NewGetActionByNameInternalServerError() *GetActionByNameInternalServerError {

	return &GetActionByNameInternalServerError{}
}

// WithPayload adds the payload to the get action by name internal server error response
func (o *GetActionByNameInternalServerError) WithPayload(payload *models.ErrorMessage) *GetActionByNameInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get action by name internal server error response
func (o *GetActionByNameInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetActionByNameInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
