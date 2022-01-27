// Code generated by go-swagger; DO NOT EDIT.

package vds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetVdsesReader is a Reader for the GetVdses structure.
type GetVdsesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVdsesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVdsesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVdsesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVdsesOK creates a GetVdsesOK with default headers values
func NewGetVdsesOK() *GetVdsesOK {
	return &GetVdsesOK{}
}

/* GetVdsesOK describes a response with status code 200, with default header values.

Ok
*/
type GetVdsesOK struct {
	Payload []*models.Vds
}

func (o *GetVdsesOK) Error() string {
	return fmt.Sprintf("[POST /get-vdses][%d] getVdsesOK  %+v", 200, o.Payload)
}
func (o *GetVdsesOK) GetPayload() []*models.Vds {
	return o.Payload
}

func (o *GetVdsesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVdsesBadRequest creates a GetVdsesBadRequest with default headers values
func NewGetVdsesBadRequest() *GetVdsesBadRequest {
	return &GetVdsesBadRequest{}
}

/* GetVdsesBadRequest describes a response with status code 400, with default header values.

GetVdsesBadRequest get vdses bad request
*/
type GetVdsesBadRequest struct {
	Payload string
}

func (o *GetVdsesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vdses][%d] getVdsesBadRequest  %+v", 400, o.Payload)
}
func (o *GetVdsesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVdsesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}