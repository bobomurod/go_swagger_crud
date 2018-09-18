// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "users-api/models"
)

// GetUserByIDOKCode is the HTTP code returned for type GetUserByIDOK
const GetUserByIDOKCode int = 200

/*GetUserByIDOK OK

swagger:response getUserByIdOK
*/
type GetUserByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserByIDOK creates GetUserByIDOK with default headers values
func NewGetUserByIDOK() *GetUserByIDOK {

	return &GetUserByIDOK{}
}

// WithPayload adds the payload to the get user by Id o k response
func (o *GetUserByIDOK) WithPayload(payload *models.User) *GetUserByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id o k response
func (o *GetUserByIDOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByIDBadRequestCode is the HTTP code returned for type GetUserByIDBadRequest
const GetUserByIDBadRequestCode int = 400

/*GetUserByIDBadRequest Invalid request

swagger:response getUserByIdBadRequest
*/
type GetUserByIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *GetUserByIDBadRequestBody `json:"body,omitempty"`
}

// NewGetUserByIDBadRequest creates GetUserByIDBadRequest with default headers values
func NewGetUserByIDBadRequest() *GetUserByIDBadRequest {

	return &GetUserByIDBadRequest{}
}

// WithPayload adds the payload to the get user by Id bad request response
func (o *GetUserByIDBadRequest) WithPayload(payload *GetUserByIDBadRequestBody) *GetUserByIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id bad request response
func (o *GetUserByIDBadRequest) SetPayload(payload *GetUserByIDBadRequestBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByIDNotFoundCode is the HTTP code returned for type GetUserByIDNotFound
const GetUserByIDNotFoundCode int = 404

/*GetUserByIDNotFound Not found

swagger:response getUserByIdNotFound
*/
type GetUserByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *GetUserByIDNotFoundBody `json:"body,omitempty"`
}

// NewGetUserByIDNotFound creates GetUserByIDNotFound with default headers values
func NewGetUserByIDNotFound() *GetUserByIDNotFound {

	return &GetUserByIDNotFound{}
}

// WithPayload adds the payload to the get user by Id not found response
func (o *GetUserByIDNotFound) WithPayload(payload *GetUserByIDNotFoundBody) *GetUserByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id not found response
func (o *GetUserByIDNotFound) SetPayload(payload *GetUserByIDNotFoundBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserByIDInternalServerErrorCode is the HTTP code returned for type GetUserByIDInternalServerError
const GetUserByIDInternalServerErrorCode int = 500

/*GetUserByIDInternalServerError Something went wrong

swagger:response getUserByIdInternalServerError
*/
type GetUserByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetUserByIDInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUserByIDInternalServerError creates GetUserByIDInternalServerError with default headers values
func NewGetUserByIDInternalServerError() *GetUserByIDInternalServerError {

	return &GetUserByIDInternalServerError{}
}

// WithPayload adds the payload to the get user by Id internal server error response
func (o *GetUserByIDInternalServerError) WithPayload(payload *GetUserByIDInternalServerErrorBody) *GetUserByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user by Id internal server error response
func (o *GetUserByIDInternalServerError) SetPayload(payload *GetUserByIDInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}