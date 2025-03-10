// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateSnapshotPlanReader is a Reader for the UpdateSnapshotPlan structure.
type UpdateSnapshotPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSnapshotPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSnapshotPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateSnapshotPlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateSnapshotPlanOK creates a UpdateSnapshotPlanOK with default headers values
func NewUpdateSnapshotPlanOK() *UpdateSnapshotPlanOK {
	return &UpdateSnapshotPlanOK{}
}

/* UpdateSnapshotPlanOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateSnapshotPlanOK struct {
	Payload []*models.WithTaskSnapshotPlan
}

func (o *UpdateSnapshotPlanOK) Error() string {
	return fmt.Sprintf("[POST /update-snapshot-plan][%d] updateSnapshotPlanOK  %+v", 200, o.Payload)
}
func (o *UpdateSnapshotPlanOK) GetPayload() []*models.WithTaskSnapshotPlan {
	return o.Payload
}

func (o *UpdateSnapshotPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSnapshotPlanBadRequest creates a UpdateSnapshotPlanBadRequest with default headers values
func NewUpdateSnapshotPlanBadRequest() *UpdateSnapshotPlanBadRequest {
	return &UpdateSnapshotPlanBadRequest{}
}

/* UpdateSnapshotPlanBadRequest describes a response with status code 400, with default header values.

UpdateSnapshotPlanBadRequest update snapshot plan bad request
*/
type UpdateSnapshotPlanBadRequest struct {
	Payload string
}

func (o *UpdateSnapshotPlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-snapshot-plan][%d] updateSnapshotPlanBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateSnapshotPlanBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateSnapshotPlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
