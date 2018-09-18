// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-openapi/runtime"

	models "users-api/models"
)

// GetUsersOKCode is the HTTP code returned for type GetUsersOK
const GetUsersOKCode int = 200

/*GetUsersOK OK

swagger:response getUsersOK
*/
type GetUsersOK struct {
	/*
	  In: Body
	*/
	Payload []*models.User `json:"body,omitempty"`
}

// NewGetUsersOK creates GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {

	forReturn := make([]*models.User, 0)   //новая переменная slice внутри данные с типом models.User (slice пока имеет в себе 0 записей)
	elementUserName := "hello_user"    // простой стринг
	elementUserName2 := "how_are_you_user"    // простой стринг
	element1 := &models.User{UserID: 12, Username: &elementUserName} // готовим новый объект на основе models.User
	element2 := &models.User{UserID: 13, Username: &elementUserName2} // готовим новый объект на основе models.User

	forReturn = append(forReturn, element1, element2)   //
	//return &GetUsersOK{}
	//forReturn := make([]*GetUsersOK,0)
	fmt.Println(&GetUsersOK{Payload:forReturn})
	return &GetUsersOK{Payload:forReturn} // forReturn
}

// WithPayload adds the payload to the get users o k response
func (o *GetUsersOK) WithPayload(payload []*models.User) *GetUsersOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users o k response
func (o *GetUsersOK) SetPayload(payload []*models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.User, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// GetUsersInternalServerErrorCode is the HTTP code returned for type GetUsersInternalServerError
const GetUsersInternalServerErrorCode int = 500

/*GetUsersInternalServerError Something went wrong

swagger:response getUsersInternalServerError
*/ 

type GetUsersInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *GetUsersInternalServerErrorBody `json:"body,omitempty"`
}

// NewGetUsersInternalServerError creates GetUsersInternalServerError with default headers values
func NewGetUsersInternalServerError() *GetUsersInternalServerError {

	return &GetUsersInternalServerError{}
}

// WithPayload adds the payload to the get users internal server error response
func (o *GetUsersInternalServerError) WithPayload(payload *GetUsersInternalServerErrorBody) *GetUsersInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get users internal server error response
func (o *GetUsersInternalServerError) SetPayload(payload *GetUsersInternalServerErrorBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUsersInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}