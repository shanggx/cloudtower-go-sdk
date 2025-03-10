// Code generated by go-swagger; DO NOT EDIT.

package vlan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// DeleteVlanReader is a Reader for the DeleteVlan structure.
type DeleteVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVlanOK creates a DeleteVlanOK with default headers values
func NewDeleteVlanOK() *DeleteVlanOK {
	return &DeleteVlanOK{}
}

/* DeleteVlanOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVlanOK struct {
	Payload []*models.WithTaskDeleteVlan
}

func (o *DeleteVlanOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-vlan][%d] deleteVlanOK  %+v", 200, o.Payload)
}
func (o *DeleteVlanOK) GetPayload() []*models.WithTaskDeleteVlan {
	return o.Payload
}

func (o *DeleteVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVlanBadRequest creates a DeleteVlanBadRequest with default headers values
func NewDeleteVlanBadRequest() *DeleteVlanBadRequest {
	return &DeleteVlanBadRequest{}
}

/* DeleteVlanBadRequest describes a response with status code 400, with default header values.

DeleteVlanBadRequest delete vlan bad request
*/
type DeleteVlanBadRequest struct {
	Payload string
}

func (o *DeleteVlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-vlan][%d] deleteVlanBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVlanBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteVlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
