// Code generated by go-swagger; DO NOT EDIT.

package elf_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateElfImageReader is a Reader for the CreateElfImage structure.
type CreateElfImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateElfImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateElfImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateElfImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateElfImageOK creates a CreateElfImageOK with default headers values
func NewCreateElfImageOK() *CreateElfImageOK {
	return &CreateElfImageOK{}
}

/* CreateElfImageOK describes a response with status code 200, with default header values.

Ok
*/
type CreateElfImageOK struct {
	Payload []*models.UploadTask
}

func (o *CreateElfImageOK) Error() string {
	return fmt.Sprintf("[POST /upload-elf-image][%d] createElfImageOK  %+v", 200, o.Payload)
}
func (o *CreateElfImageOK) GetPayload() []*models.UploadTask {
	return o.Payload
}

func (o *CreateElfImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateElfImageBadRequest creates a CreateElfImageBadRequest with default headers values
func NewCreateElfImageBadRequest() *CreateElfImageBadRequest {
	return &CreateElfImageBadRequest{}
}

/* CreateElfImageBadRequest describes a response with status code 400, with default header values.

CreateElfImageBadRequest create elf image bad request
*/
type CreateElfImageBadRequest struct {
	Payload string
}

func (o *CreateElfImageBadRequest) Error() string {
	return fmt.Sprintf("[POST /upload-elf-image][%d] createElfImageBadRequest  %+v", 400, o.Payload)
}
func (o *CreateElfImageBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateElfImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
