// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Balun-courses/microservices_like_in_bigtech/lecture_2/docker/models"
)

// GetUserOKCode is the HTTP code returned for type GetUserOK
const GetUserOKCode int = 200

/*
GetUserOK User successfully created

swagger:response getUserOK
*/
type GetUserOK struct {

	/*
	  In: Body
	*/
	Payload *models.GetUserResponse `json:"body,omitempty"`
}

// NewGetUserOK creates GetUserOK with default headers values
func NewGetUserOK() *GetUserOK {

	return &GetUserOK{}
}

// WithPayload adds the payload to the get user o k response
func (o *GetUserOK) WithPayload(payload *models.GetUserResponse) *GetUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user o k response
func (o *GetUserOK) SetPayload(payload *models.GetUserResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserNotFoundCode is the HTTP code returned for type GetUserNotFound
const GetUserNotFoundCode int = 404

/*
GetUserNotFound Not found

swagger:response getUserNotFound
*/
type GetUserNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetUserNotFound creates GetUserNotFound with default headers values
func NewGetUserNotFound() *GetUserNotFound {

	return &GetUserNotFound{}
}

// WithPayload adds the payload to the get user not found response
func (o *GetUserNotFound) WithPayload(payload *models.ErrorMessage) *GetUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user not found response
func (o *GetUserNotFound) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserInternalServerErrorCode is the HTTP code returned for type GetUserInternalServerError
const GetUserInternalServerErrorCode int = 500

/*
GetUserInternalServerError Internal server error

swagger:response getUserInternalServerError
*/
type GetUserInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorMessage `json:"body,omitempty"`
}

// NewGetUserInternalServerError creates GetUserInternalServerError with default headers values
func NewGetUserInternalServerError() *GetUserInternalServerError {

	return &GetUserInternalServerError{}
}

// WithPayload adds the payload to the get user internal server error response
func (o *GetUserInternalServerError) WithPayload(payload *models.ErrorMessage) *GetUserInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user internal server error response
func (o *GetUserInternalServerError) SetPayload(payload *models.ErrorMessage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
