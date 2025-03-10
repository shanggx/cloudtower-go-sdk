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

// RemoveVMToFolderReader is a Reader for the RemoveVMToFolder structure.
type RemoveVMToFolderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveVMToFolderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveVMToFolderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveVMToFolderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveVMToFolderOK creates a RemoveVMToFolderOK with default headers values
func NewRemoveVMToFolderOK() *RemoveVMToFolderOK {
	return &RemoveVMToFolderOK{}
}

/* RemoveVMToFolderOK describes a response with status code 200, with default header values.

Ok
*/
type RemoveVMToFolderOK struct {
	Payload []*models.WithTaskVM
}

func (o *RemoveVMToFolderOK) Error() string {
	return fmt.Sprintf("[POST /remove-vm-from-folder][%d] removeVmToFolderOK  %+v", 200, o.Payload)
}
func (o *RemoveVMToFolderOK) GetPayload() []*models.WithTaskVM {
	return o.Payload
}

func (o *RemoveVMToFolderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveVMToFolderBadRequest creates a RemoveVMToFolderBadRequest with default headers values
func NewRemoveVMToFolderBadRequest() *RemoveVMToFolderBadRequest {
	return &RemoveVMToFolderBadRequest{}
}

/* RemoveVMToFolderBadRequest describes a response with status code 400, with default header values.

RemoveVMToFolderBadRequest remove Vm to folder bad request
*/
type RemoveVMToFolderBadRequest struct {
	Payload string
}

func (o *RemoveVMToFolderBadRequest) Error() string {
	return fmt.Sprintf("[POST /remove-vm-from-folder][%d] removeVmToFolderBadRequest  %+v", 400, o.Payload)
}
func (o *RemoveVMToFolderBadRequest) GetPayload() string {
	return o.Payload
}

func (o *RemoveVMToFolderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
