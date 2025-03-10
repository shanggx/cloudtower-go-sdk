// Code generated by go-swagger; DO NOT EDIT.

package vds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateVdsReader is a Reader for the CreateVds structure.
type CreateVdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVdsOK creates a CreateVdsOK with default headers values
func NewCreateVdsOK() *CreateVdsOK {
	return &CreateVdsOK{}
}

/* CreateVdsOK describes a response with status code 200, with default header values.

Ok
*/
type CreateVdsOK struct {
	Payload []*models.WithTaskVds
}

func (o *CreateVdsOK) Error() string {
	return fmt.Sprintf("[POST /create-vds][%d] createVdsOK  %+v", 200, o.Payload)
}
func (o *CreateVdsOK) GetPayload() []*models.WithTaskVds {
	return o.Payload
}

func (o *CreateVdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVdsBadRequest creates a CreateVdsBadRequest with default headers values
func NewCreateVdsBadRequest() *CreateVdsBadRequest {
	return &CreateVdsBadRequest{}
}

/* CreateVdsBadRequest describes a response with status code 400, with default header values.

CreateVdsBadRequest create vds bad request
*/
type CreateVdsBadRequest struct {
	Payload string
}

func (o *CreateVdsBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vds][%d] createVdsBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVdsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateVdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
