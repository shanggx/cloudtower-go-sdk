// Code generated by go-swagger; DO NOT EDIT.

package vm_nic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetVMNicsReader is a Reader for the GetVMNics structure.
type GetVMNicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVMNicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVMNicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVMNicsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVMNicsOK creates a GetVMNicsOK with default headers values
func NewGetVMNicsOK() *GetVMNicsOK {
	return &GetVMNicsOK{}
}

/* GetVMNicsOK describes a response with status code 200, with default header values.

Ok
*/
type GetVMNicsOK struct {
	Payload []*models.VMNic
}

func (o *GetVMNicsOK) Error() string {
	return fmt.Sprintf("[POST /get-vm-nics][%d] getVmNicsOK  %+v", 200, o.Payload)
}
func (o *GetVMNicsOK) GetPayload() []*models.VMNic {
	return o.Payload
}

func (o *GetVMNicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVMNicsBadRequest creates a GetVMNicsBadRequest with default headers values
func NewGetVMNicsBadRequest() *GetVMNicsBadRequest {
	return &GetVMNicsBadRequest{}
}

/* GetVMNicsBadRequest describes a response with status code 400, with default header values.

GetVMNicsBadRequest get Vm nics bad request
*/
type GetVMNicsBadRequest struct {
	Payload string
}

func (o *GetVMNicsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vm-nics][%d] getVmNicsBadRequest  %+v", 400, o.Payload)
}
func (o *GetVMNicsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVMNicsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
