// Code generated by go-swagger; DO NOT EDIT.

package user_role_next

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateRoleReader is a Reader for the UpdateRole structure.
type UpdateRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRoleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRoleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRoleOK creates a UpdateRoleOK with default headers values
func NewUpdateRoleOK() *UpdateRoleOK {
	return &UpdateRoleOK{}
}

/* UpdateRoleOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateRoleOK struct {
	Payload []*models.WithTaskUserRoleNext
}

func (o *UpdateRoleOK) Error() string {
	return fmt.Sprintf("[POST /update-role][%d] updateRoleOK  %+v", 200, o.Payload)
}
func (o *UpdateRoleOK) GetPayload() []*models.WithTaskUserRoleNext {
	return o.Payload
}

func (o *UpdateRoleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleBadRequest creates a UpdateRoleBadRequest with default headers values
func NewUpdateRoleBadRequest() *UpdateRoleBadRequest {
	return &UpdateRoleBadRequest{}
}

/* UpdateRoleBadRequest describes a response with status code 400, with default header values.

UpdateRoleBadRequest update role bad request
*/
type UpdateRoleBadRequest struct {
	Payload string
}

func (o *UpdateRoleBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-role][%d] updateRoleBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateRoleBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateRoleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
