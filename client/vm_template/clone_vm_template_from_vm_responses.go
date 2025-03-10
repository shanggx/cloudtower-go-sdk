// Code generated by go-swagger; DO NOT EDIT.

package vm_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CloneVMTemplateFromVMReader is a Reader for the CloneVMTemplateFromVM structure.
type CloneVMTemplateFromVMReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneVMTemplateFromVMReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneVMTemplateFromVMOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloneVMTemplateFromVMBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneVMTemplateFromVMOK creates a CloneVMTemplateFromVMOK with default headers values
func NewCloneVMTemplateFromVMOK() *CloneVMTemplateFromVMOK {
	return &CloneVMTemplateFromVMOK{}
}

/* CloneVMTemplateFromVMOK describes a response with status code 200, with default header values.

Ok
*/
type CloneVMTemplateFromVMOK struct {
	Payload []*models.WithTaskVMTemplate
}

func (o *CloneVMTemplateFromVMOK) Error() string {
	return fmt.Sprintf("[POST /clone-vm-template-from-vm][%d] cloneVmTemplateFromVmOK  %+v", 200, o.Payload)
}
func (o *CloneVMTemplateFromVMOK) GetPayload() []*models.WithTaskVMTemplate {
	return o.Payload
}

func (o *CloneVMTemplateFromVMOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneVMTemplateFromVMBadRequest creates a CloneVMTemplateFromVMBadRequest with default headers values
func NewCloneVMTemplateFromVMBadRequest() *CloneVMTemplateFromVMBadRequest {
	return &CloneVMTemplateFromVMBadRequest{}
}

/* CloneVMTemplateFromVMBadRequest describes a response with status code 400, with default header values.

CloneVMTemplateFromVMBadRequest clone Vm template from Vm bad request
*/
type CloneVMTemplateFromVMBadRequest struct {
	Payload string
}

func (o *CloneVMTemplateFromVMBadRequest) Error() string {
	return fmt.Sprintf("[POST /clone-vm-template-from-vm][%d] cloneVmTemplateFromVmBadRequest  %+v", 400, o.Payload)
}
func (o *CloneVMTemplateFromVMBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CloneVMTemplateFromVMBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
