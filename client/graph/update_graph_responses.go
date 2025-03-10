// Code generated by go-swagger; DO NOT EDIT.

package graph

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// UpdateGraphReader is a Reader for the UpdateGraph structure.
type UpdateGraphReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGraphReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateGraphOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateGraphBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateGraphOK creates a UpdateGraphOK with default headers values
func NewUpdateGraphOK() *UpdateGraphOK {
	return &UpdateGraphOK{}
}

/* UpdateGraphOK describes a response with status code 200, with default header values.

Ok
*/
type UpdateGraphOK struct {
	Payload []*models.WithTaskGraph
}

func (o *UpdateGraphOK) Error() string {
	return fmt.Sprintf("[POST /update-graph][%d] updateGraphOK  %+v", 200, o.Payload)
}
func (o *UpdateGraphOK) GetPayload() []*models.WithTaskGraph {
	return o.Payload
}

func (o *UpdateGraphOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGraphBadRequest creates a UpdateGraphBadRequest with default headers values
func NewUpdateGraphBadRequest() *UpdateGraphBadRequest {
	return &UpdateGraphBadRequest{}
}

/* UpdateGraphBadRequest describes a response with status code 400, with default header values.

UpdateGraphBadRequest update graph bad request
*/
type UpdateGraphBadRequest struct {
	Payload string
}

func (o *UpdateGraphBadRequest) Error() string {
	return fmt.Sprintf("[POST /update-graph][%d] updateGraphBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateGraphBadRequest) GetPayload() string {
	return o.Payload
}

func (o *UpdateGraphBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
