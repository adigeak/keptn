// Code generated by go-swagger; DO NOT EDIT.

package service_default_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/keptn/keptn/configuration-service/models"
)

// PutProjectProjectNameServiceServiceNameResourceResourceURICreatedCode is the HTTP code returned for type PutProjectProjectNameServiceServiceNameResourceResourceURICreated
const PutProjectProjectNameServiceServiceNameResourceResourceURICreatedCode int = 201

/*PutProjectProjectNameServiceServiceNameResourceResourceURICreated Success. Service default resource has been updated. The version of the new configuration is returned.

swagger:response putProjectProjectNameServiceServiceNameResourceResourceUriCreated
*/
type PutProjectProjectNameServiceServiceNameResourceResourceURICreated struct {

	/*
	  In: Body
	*/
	Payload *models.Version `json:"body,omitempty"`
}

// NewPutProjectProjectNameServiceServiceNameResourceResourceURICreated creates PutProjectProjectNameServiceServiceNameResourceResourceURICreated with default headers values
func NewPutProjectProjectNameServiceServiceNameResourceResourceURICreated() *PutProjectProjectNameServiceServiceNameResourceResourceURICreated {

	return &PutProjectProjectNameServiceServiceNameResourceResourceURICreated{}
}

// WithPayload adds the payload to the put project project name service service name resource resource Uri created response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURICreated) WithPayload(payload *models.Version) *PutProjectProjectNameServiceServiceNameResourceResourceURICreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name service service name resource resource Uri created response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURICreated) SetPayload(payload *models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURICreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequestCode is the HTTP code returned for type PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest
const PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequestCode int = 400

/*PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest Failed. Service default resource could not be updated.

swagger:response putProjectProjectNameServiceServiceNameResourceResourceUriBadRequest
*/
type PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest creates PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest with default headers values
func NewPutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest() *PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest {

	return &PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest{}
}

// WithPayload adds the payload to the put project project name service service name resource resource Uri bad request response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest) WithPayload(payload *models.Error) *PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name service service name resource resource Uri bad request response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PutProjectProjectNameServiceServiceNameResourceResourceURIDefault Error

swagger:response putProjectProjectNameServiceServiceNameResourceResourceUriDefault
*/
type PutProjectProjectNameServiceServiceNameResourceResourceURIDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutProjectProjectNameServiceServiceNameResourceResourceURIDefault creates PutProjectProjectNameServiceServiceNameResourceResourceURIDefault with default headers values
func NewPutProjectProjectNameServiceServiceNameResourceResourceURIDefault(code int) *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault {
	if code <= 0 {
		code = 500
	}

	return &PutProjectProjectNameServiceServiceNameResourceResourceURIDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put project project name service service name resource resource URI default response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault) WithStatusCode(code int) *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put project project name service service name resource resource URI default response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put project project name service service name resource resource URI default response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault) WithPayload(payload *models.Error) *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put project project name service service name resource resource URI default response
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutProjectProjectNameServiceServiceNameResourceResourceURIDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
