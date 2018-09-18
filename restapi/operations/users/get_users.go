// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
)

// GetUsersHandlerFunc turns a function with the right signature into a get users handler
type GetUsersHandlerFunc func(GetUsersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUsersHandlerFunc) Handle(params GetUsersParams) middleware.Responder {
	return fn(params)
}

// GetUsersHandler interface for that can handle valid get users params
type GetUsersHandler interface {
	Handle(GetUsersParams) middleware.Responder
}

// NewGetUsers creates a new http.Handler for the get users operation
func NewGetUsers(ctx *middleware.Context, handler GetUsersHandler) *GetUsers {
	return &GetUsers{Context: ctx, Handler: handler}
}

/*GetUsers swagger:route GET /users users getUsers

Returns all users list

Returns all users list for test

*/
type GetUsers struct {
	Context *middleware.Context
	Handler GetUsersHandler
}

func (o *GetUsers) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetUsersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetUsersInternalServerErrorBody get users internal server error body
// swagger:model GetUsersInternalServerErrorBody
type GetUsersInternalServerErrorBody struct {

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this get users internal server error body
func (o *GetUsersInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetUsersInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetUsersInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res GetUsersInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
