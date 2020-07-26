// Code generated by go-swagger; DO NOT EDIT.

package movieinfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"MoviesDB/models"
)

// GetMovieInfoOKCode is the HTTP code returned for type GetMovieInfoOK
const GetMovieInfoOKCode int = 200

/*GetMovieInfoOK Success.

swagger:response getMovieInfoOK
*/
type GetMovieInfoOK struct {

	/*
	  In: Body
	*/
	Payload *GetMovieInfoOKBody `json:"body,omitempty"`
}

// NewGetMovieInfoOK creates GetMovieInfoOK with default headers values
func NewGetMovieInfoOK() *GetMovieInfoOK {

	return &GetMovieInfoOK{}
}

// WithPayload adds the payload to the get movie info o k response
func (o *GetMovieInfoOK) WithPayload(payload *GetMovieInfoOKBody) *GetMovieInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie info o k response
func (o *GetMovieInfoOK) SetPayload(payload *GetMovieInfoOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetMovieInfoInternalServerErrorCode is the HTTP code returned for type GetMovieInfoInternalServerError
const GetMovieInfoInternalServerErrorCode int = 500

/*GetMovieInfoInternalServerError Server Error

swagger:response getMovieInfoInternalServerError
*/
type GetMovieInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMovieInfoInternalServerError creates GetMovieInfoInternalServerError with default headers values
func NewGetMovieInfoInternalServerError() *GetMovieInfoInternalServerError {

	return &GetMovieInfoInternalServerError{}
}

// WithPayload adds the payload to the get movie info internal server error response
func (o *GetMovieInfoInternalServerError) WithPayload(payload *models.Error) *GetMovieInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie info internal server error response
func (o *GetMovieInfoInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetMovieInfoDefault error

swagger:response getMovieInfoDefault
*/
type GetMovieInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetMovieInfoDefault creates GetMovieInfoDefault with default headers values
func NewGetMovieInfoDefault(code int) *GetMovieInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetMovieInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get movie info default response
func (o *GetMovieInfoDefault) WithStatusCode(code int) *GetMovieInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get movie info default response
func (o *GetMovieInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get movie info default response
func (o *GetMovieInfoDefault) WithPayload(payload *models.Error) *GetMovieInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get movie info default response
func (o *GetMovieInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetMovieInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
