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

// CreateVdsWithAccessVlanReader is a Reader for the CreateVdsWithAccessVlan structure.
type CreateVdsWithAccessVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVdsWithAccessVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVdsWithAccessVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVdsWithAccessVlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVdsWithAccessVlanOK creates a CreateVdsWithAccessVlanOK with default headers values
func NewCreateVdsWithAccessVlanOK() *CreateVdsWithAccessVlanOK {
	return &CreateVdsWithAccessVlanOK{}
}

/* CreateVdsWithAccessVlanOK describes a response with status code 200, with default header values.

Ok
*/
type CreateVdsWithAccessVlanOK struct {
	Payload []*models.WithTaskVds
}

func (o *CreateVdsWithAccessVlanOK) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-access-vlan][%d] createVdsWithAccessVlanOK  %+v", 200, o.Payload)
}
func (o *CreateVdsWithAccessVlanOK) GetPayload() []*models.WithTaskVds {
	return o.Payload
}

func (o *CreateVdsWithAccessVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVdsWithAccessVlanBadRequest creates a CreateVdsWithAccessVlanBadRequest with default headers values
func NewCreateVdsWithAccessVlanBadRequest() *CreateVdsWithAccessVlanBadRequest {
	return &CreateVdsWithAccessVlanBadRequest{}
}

/* CreateVdsWithAccessVlanBadRequest describes a response with status code 400, with default header values.

CreateVdsWithAccessVlanBadRequest create vds with access vlan bad request
*/
type CreateVdsWithAccessVlanBadRequest struct {
	Payload string
}

func (o *CreateVdsWithAccessVlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vds-with-access-vlan][%d] createVdsWithAccessVlanBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVdsWithAccessVlanBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateVdsWithAccessVlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
