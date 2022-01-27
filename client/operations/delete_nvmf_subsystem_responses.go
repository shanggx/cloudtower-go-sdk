// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// DeleteNvmfSubsystemReader is a Reader for the DeleteNvmfSubsystem structure.
type DeleteNvmfSubsystemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNvmfSubsystemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNvmfSubsystemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteNvmfSubsystemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteNvmfSubsystemOK creates a DeleteNvmfSubsystemOK with default headers values
func NewDeleteNvmfSubsystemOK() *DeleteNvmfSubsystemOK {
	return &DeleteNvmfSubsystemOK{}
}

/* DeleteNvmfSubsystemOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteNvmfSubsystemOK struct {
	Payload []*models.WithTaskDeleteNvmfSubsystem
}

func (o *DeleteNvmfSubsystemOK) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-subsystem][%d] deleteNvmfSubsystemOK  %+v", 200, o.Payload)
}
func (o *DeleteNvmfSubsystemOK) GetPayload() []*models.WithTaskDeleteNvmfSubsystem {
	return o.Payload
}

func (o *DeleteNvmfSubsystemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNvmfSubsystemBadRequest creates a DeleteNvmfSubsystemBadRequest with default headers values
func NewDeleteNvmfSubsystemBadRequest() *DeleteNvmfSubsystemBadRequest {
	return &DeleteNvmfSubsystemBadRequest{}
}

/* DeleteNvmfSubsystemBadRequest describes a response with status code 400, with default header values.

DeleteNvmfSubsystemBadRequest delete nvmf subsystem bad request
*/
type DeleteNvmfSubsystemBadRequest struct {
	Payload string
}

func (o *DeleteNvmfSubsystemBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-nvmf-subsystem][%d] deleteNvmfSubsystemBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteNvmfSubsystemBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteNvmfSubsystemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}