// Code generated by go-swagger; DO NOT EDIT.

package add_movie

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"MoviesDB/models"
)

// AddMovieOKCode is the HTTP code returned for type AddMovieOK
const AddMovieOKCode int = 200

/*AddMovieOK Movie Uploaded Successfully

swagger:response addMovieOK
*/
type AddMovieOK struct {

	/*
	  In: Body
	*/
	Payload *AddMovieOKBody `json:"body,omitempty"`
}

// NewAddMovieOK creates AddMovieOK with default headers values
func NewAddMovieOK() *AddMovieOK {

	return &AddMovieOK{}
}

// WithPayload adds the payload to the add movie o k response
func (o *AddMovieOK) WithPayload(payload *AddMovieOKBody) *AddMovieOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add movie o k response
func (o *AddMovieOK) SetPayload(payload *AddMovieOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddMovieOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddMovieInternalServerErrorCode is the HTTP code returned for type AddMovieInternalServerError
const AddMovieInternalServerErrorCode int = 500

/*AddMovieInternalServerError Server Error

swagger:response addMovieInternalServerError
*/
type AddMovieInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddMovieInternalServerError creates AddMovieInternalServerError with default headers values
func NewAddMovieInternalServerError() *AddMovieInternalServerError {

	return &AddMovieInternalServerError{}
}

// WithPayload adds the payload to the add movie internal server error response
func (o *AddMovieInternalServerError) WithPayload(payload *models.Error) *AddMovieInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add movie internal server error response
func (o *AddMovieInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddMovieInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*AddMovieDefault error

swagger:response addMovieDefault
*/
type AddMovieDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewAddMovieDefault creates AddMovieDefault with default headers values
func NewAddMovieDefault(code int) *AddMovieDefault {
	if code <= 0 {
		code = 500
	}

	return &AddMovieDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add movie default response
func (o *AddMovieDefault) WithStatusCode(code int) *AddMovieDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add movie default response
func (o *AddMovieDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add movie default response
func (o *AddMovieDefault) WithPayload(payload *models.Error) *AddMovieDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add movie default response
func (o *AddMovieDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddMovieDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
