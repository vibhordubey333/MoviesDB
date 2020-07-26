// Code generated by go-swagger; DO NOT EDIT.

package add_comment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostcommentsHandlerFunc turns a function with the right signature into a postcomments handler
type PostcommentsHandlerFunc func(PostcommentsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostcommentsHandlerFunc) Handle(params PostcommentsParams) middleware.Responder {
	return fn(params)
}

// PostcommentsHandler interface for that can handle valid postcomments params
type PostcommentsHandler interface {
	Handle(PostcommentsParams) middleware.Responder
}

// NewPostcomments creates a new http.Handler for the postcomments operation
func NewPostcomments(ctx *middleware.Context, handler PostcommentsHandler) *Postcomments {
	return &Postcomments{Context: ctx, Handler: handler}
}

/*Postcomments swagger:route POST /imdbservice/addcomment addComment postcomments

Add movie into DB.

Add movie into DB.

*/
type Postcomments struct {
	Context *middleware.Context
	Handler PostcommentsHandler
}

func (o *Postcomments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostcommentsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostcommentsBody postcomments body
//
// swagger:model PostcommentsBody
type PostcommentsBody struct {

	// Comment for movie.
	MovieComment string `json:"MovieComment,omitempty"`

	// Movie name.
	MovieName string `json:"MovieName,omitempty"`
}

// Validate validates this postcomments body
func (o *PostcommentsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostcommentsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostcommentsBody) UnmarshalBinary(b []byte) error {
	var res PostcommentsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// PostcommentsOKBody postcomments o k body
//
// swagger:model PostcommentsOKBody
type PostcommentsOKBody struct {

	// Response Code
	Code string `json:"code,omitempty"`

	// Comment addded message.
	Message string `json:"message,omitempty"`
}

// Validate validates this postcomments o k body
func (o *PostcommentsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostcommentsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostcommentsOKBody) UnmarshalBinary(b []byte) error {
	var res PostcommentsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}