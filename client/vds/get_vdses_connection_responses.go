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

// GetVdsesConnectionReader is a Reader for the GetVdsesConnection structure.
type GetVdsesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVdsesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVdsesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVdsesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVdsesConnectionOK creates a GetVdsesConnectionOK with default headers values
func NewGetVdsesConnectionOK() *GetVdsesConnectionOK {
	return &GetVdsesConnectionOK{}
}

/* GetVdsesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetVdsesConnectionOK struct {
	Payload *models.VdsConnection
}

func (o *GetVdsesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-vdses-connection][%d] getVdsesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetVdsesConnectionOK) GetPayload() *models.VdsConnection {
	return o.Payload
}

func (o *GetVdsesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VdsConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVdsesConnectionBadRequest creates a GetVdsesConnectionBadRequest with default headers values
func NewGetVdsesConnectionBadRequest() *GetVdsesConnectionBadRequest {
	return &GetVdsesConnectionBadRequest{}
}

/* GetVdsesConnectionBadRequest describes a response with status code 400, with default header values.

GetVdsesConnectionBadRequest get vdses connection bad request
*/
type GetVdsesConnectionBadRequest struct {
	Payload string
}

func (o *GetVdsesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vdses-connection][%d] getVdsesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetVdsesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVdsesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
