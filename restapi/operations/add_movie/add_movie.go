// Code generated by go-swagger; DO NOT EDIT.

package add_movie

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddMovieHandlerFunc turns a function with the right signature into a add movie handler
type AddMovieHandlerFunc func(AddMovieParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddMovieHandlerFunc) Handle(params AddMovieParams) middleware.Responder {
	return fn(params)
}

// AddMovieHandler interface for that can handle valid add movie params
type AddMovieHandler interface {
	Handle(AddMovieParams) middleware.Responder
}

// NewAddMovie creates a new http.Handler for the add movie operation
func NewAddMovie(ctx *middleware.Context, handler AddMovieHandler) *AddMovie {
	return &AddMovie{Context: ctx, Handler: handler}
}

/*AddMovie swagger:route POST /imdbservice/addmovie addMovie addMovie

Add movie into DB.

Add movie into DB.

*/
type AddMovie struct {
	Context *middleware.Context
	Handler AddMovieHandler
}

func (o *AddMovie) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddMovieParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddMovieBody add movie body
//
// swagger:model AddMovieBody
type AddMovieBody struct {

	// Movie comments.
	MovieComments string `json:"MovieComments,omitempty"`

	// Movie name.
	MovieName string `json:"MovieName,omitempty"`

	// Movie rating.
	MovieRating string `json:"MovieRating,omitempty"`
}

// Validate validates this add movie body
func (o *AddMovieBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMovieBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMovieBody) UnmarshalBinary(b []byte) error {
	var res AddMovieBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddMovieOKBody add movie o k body
//
// swagger:model AddMovieOKBody
type AddMovieOKBody struct {

	// Response Code
	Code string `json:"code,omitempty"`

	// Success message
	Message string `json:"message,omitempty"`
}

// Validate validates this add movie o k body
func (o *AddMovieOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddMovieOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddMovieOKBody) UnmarshalBinary(b []byte) error {
	var res AddMovieOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}