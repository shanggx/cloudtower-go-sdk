// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateNvmfNamespaceReader is a Reader for the UpdateNvmfNamespace structure.
type UpdateNvmfNamespaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNvmfNamespaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNvmfNamespaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNvmfNamespaceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNvmfNamespaceOK creates a UpdateNvmfNamespaceOK with default headers values
func NewUpdateNvmfNamespaceOK() *UpdateNvmfNamespaceOK {
	return &UpdateNvmfNamespaceOK{}
}

/* UpdateNvmfNamespaceOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateNvmfNamespaceOK struct {
	Payload []*models.WithTaskNvmfNamespace
}

func (o *UpdateNvmfNamespaceOK) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-namespace][%d] updateNvmfNamespaceOK  %+v", 200, o.Payload)
}
func (o *UpdateNvmfNamespaceOK) GetPayload() []*models.WithTaskNvmfNamespace {
	return o.Payload
}

func (o *UpdateNvmfNamespaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNvmfNamespaceBadRequest creates a UpdateNvmfNamespaceBadRequest with default headers values
func NewUpdateNvmfNamespaceBadRequest() *UpdateNvmfNamespaceBadRequest {
	return &UpdateNvmfNamespaceBadRequest{}
}

/* UpdateNvmfNamespaceBadRequest describes a response with status code 400, with default header values.

UpdateNvmfNamespaceBadRequest update nvmf namespace bad request
*/
type UpdateNvmfNamespaceBadRequest struct {
	Payload string
}

func (o *UpdateNvmfNamespaceBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-nvmf-namespace][%d] updateNvmfNamespaceBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateNvmfNamespaceBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateNvmfNamespaceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
