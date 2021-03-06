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

// UpdateUserByIDHandlerFunc turns a function with the right signature into a update user by Id handler
type UpdateUserByIDHandlerFunc func(UpdateUserByIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserByIDHandlerFunc) Handle(params UpdateUserByIDParams) middleware.Responder {
	return fn(params)
}

// UpdateUserByIDHandler interface for that can handle valid update user by Id params
type UpdateUserByIDHandler interface {
	Handle(UpdateUserByIDParams) middleware.Responder
}

// NewUpdateUserByID creates a new http.Handler for the update user by Id operation
func NewUpdateUserByID(ctx *middleware.Context, handler UpdateUserByIDHandler) *UpdateUserByID {
	return &UpdateUserByID{Context: ctx, Handler: handler}
}

/*UpdateUserByID swagger:route PUT /users/{userId} users updateUserById

Updates user by id

Updates user by id

*/
type UpdateUserByID struct {
	Context *middleware.Context
	Handler UpdateUserByIDHandler
}

func (o *UpdateUserByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateUserByIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UpdateUserByIDBadRequestBody update user by ID bad request body
// swagger:model UpdateUserByIDBadRequestBody
type UpdateUserByIDBadRequestBody struct {

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this update user by ID bad request body
func (o *UpdateUserByIDBadRequestBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserByIDBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserByIDBadRequestBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserByIDBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateUserByIDInternalServerErrorBody update user by ID internal server error body
// swagger:model UpdateUserByIDInternalServerErrorBody
type UpdateUserByIDInternalServerErrorBody struct {

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this update user by ID internal server error body
func (o *UpdateUserByIDInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserByIDInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserByIDInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserByIDInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// UpdateUserByIDNotFoundBody update user by ID not found body
// swagger:model UpdateUserByIDNotFoundBody
type UpdateUserByIDNotFoundBody struct {

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// Validate validates this update user by ID not found body
func (o *UpdateUserByIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateUserByIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateUserByIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res UpdateUserByIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
