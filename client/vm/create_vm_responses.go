// Code generated by go-swagger; DO NOT EDIT.

package vm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateVMReader is a Reader for the CreateVM structure.
type CreateVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateVMOK creates a CreateVMOK with default headers values
func NewCreateVMOK() *CreateVMOK {
	return &CreateVMOK{}
}

/* CreateVMOK describes a response with status code 200, with default header values.

Ok
*/
type CreateVMOK struct {
	Payload []*models.WithTaskVM
}

func (o *CreateVMOK) Error() string {
	return fmt.Sprintf("[POST /create-vm][%d] createVmOK  %+v", 200, o.Payload)
}
func (o *CreateVMOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *CreateVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateVMBadRequest creates a CreateVMBadRequest with default headers values
func NewCreateVMBadRequest() *CreateVMBadRequest {
	return &CreateVMBadRequest{}
}

/* CreateVMBadRequest describes a response with status code 400, with default header values.

CreateVMBadRequest create Vm bad request
*/
type CreateVMBadRequest struct {
	Payload string
}

func (o *CreateVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-vm][%d] createVmBadRequest  %+v", 400, o.Payload)
}
func (o *CreateVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
