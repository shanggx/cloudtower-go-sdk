// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateOrganizationReader is a Reader for the CreateOrganization structure.
type CreateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateOrganizationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationOK creates a CreateOrganizationOK with default headers values
func NewCreateOrganizationOK() *CreateOrganizationOK {
	return &CreateOrganizationOK{}
}

/* CreateOrganizationOK describes a response with status code 200, with default header values.

Ok
*/
type CreateOrganizationOK struct {
	Payload []*models.WithTaskOrganization
}

func (o *CreateOrganizationOK) Error() string {
	return fmt.Sprintf("[POST /create-organization][%d] createOrganizationOK  %+v", 200, o.Payload)
}
func (o *CreateOrganizationOK) GetPayload() []*models.WithTaskOrganization {
	return o.Payload
}

func (o *CreateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationBadRequest creates a CreateOrganizationBadRequest with default headers values
func NewCreateOrganizationBadRequest() *CreateOrganizationBadRequest {
	return &CreateOrganizationBadRequest{}
}

/* CreateOrganizationBadRequest describes a response with status code 400, with default header values.

CreateOrganizationBadRequest create organization bad request
*/
type CreateOrganizationBadRequest struct {
	Payload string
}

func (o *CreateOrganizationBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-organization][%d] createOrganizationBadRequest  %+v", 400, o.Payload)
}
func (o *CreateOrganizationBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateOrganizationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
