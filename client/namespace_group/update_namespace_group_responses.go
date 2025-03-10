// Code generated by go-swagger; DO NOT EDIT.

package namespace_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateNamespaceGroupReader is a Reader for the UpdateNamespaceGroup structure.
type UpdateNamespaceGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNamespaceGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNamespaceGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateNamespaceGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNamespaceGroupOK creates a UpdateNamespaceGroupOK with default headers values
func NewUpdateNamespaceGroupOK() *UpdateNamespaceGroupOK {
	return &UpdateNamespaceGroupOK{}
}

/* UpdateNamespaceGroupOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateNamespaceGroupOK struct {
	Payload []*models.WithTaskNamespaceGroup
}

func (o *UpdateNamespaceGroupOK) Error() string {
	return fmt.Sprintf("[POST /update-namespace-group][%d] updateNamespaceGroupOK  %+v", 200, o.Payload)
}
func (o *UpdateNamespaceGroupOK) GetPayload() []*models.WithTaskNamespaceGroup {
	return o.Payload
}

func (o *UpdateNamespaceGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateNamespaceGroupBadRequest creates a UpdateNamespaceGroupBadRequest with default headers values
func NewUpdateNamespaceGroupBadRequest() *UpdateNamespaceGroupBadRequest {
	return &UpdateNamespaceGroupBadRequest{}
}

/* UpdateNamespaceGroupBadRequest describes a response with status code 400, with default header values.

UpdateNamespaceGroupBadRequest update namespace group bad request
*/
type UpdateNamespaceGroupBadRequest struct {
	Payload string
}

func (o *UpdateNamespaceGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-namespace-group][%d] updateNamespaceGroupBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateNamespaceGroupBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateNamespaceGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
