// Code generated by go-swagger; DO NOT EDIT.

package nvmf_subsystem

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetNvmfSubsystemsConnectionReader is a Reader for the GetNvmfSubsystemsConnection structure.
type GetNvmfSubsystemsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNvmfSubsystemsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNvmfSubsystemsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNvmfSubsystemsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNvmfSubsystemsConnectionOK creates a GetNvmfSubsystemsConnectionOK with default headers values
func NewGetNvmfSubsystemsConnectionOK() *GetNvmfSubsystemsConnectionOK {
	return &GetNvmfSubsystemsConnectionOK{}
}

/* GetNvmfSubsystemsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNvmfSubsystemsConnectionOK struct {
	Payload *models.NvmfSubsystemConnection
}

func (o *GetNvmfSubsystemsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-subsystems-connection][%d] getNvmfSubsystemsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNvmfSubsystemsConnectionOK) GetPayload() *models.NvmfSubsystemConnection {
	return o.Payload
}

func (o *GetNvmfSubsystemsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmfSubsystemConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNvmfSubsystemsConnectionBadRequest creates a GetNvmfSubsystemsConnectionBadRequest with default headers values
func NewGetNvmfSubsystemsConnectionBadRequest() *GetNvmfSubsystemsConnectionBadRequest {
	return &GetNvmfSubsystemsConnectionBadRequest{}
}

/* GetNvmfSubsystemsConnectionBadRequest describes a response with status code 400, with default header values.

GetNvmfSubsystemsConnectionBadRequest get nvmf subsystems connection bad request
*/
type GetNvmfSubsystemsConnectionBadRequest struct {
	Payload string
}

func (o *GetNvmfSubsystemsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-subsystems-connection][%d] getNvmfSubsystemsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNvmfSubsystemsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNvmfSubsystemsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
