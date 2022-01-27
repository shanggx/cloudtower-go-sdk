// Code generated by go-swagger; DO NOT EDIT.

package entity_filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetEntityFiltersReader is a Reader for the GetEntityFilters structure.
type GetEntityFiltersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEntityFiltersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEntityFiltersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetEntityFiltersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEntityFiltersOK creates a GetEntityFiltersOK with default headers values
func NewGetEntityFiltersOK() *GetEntityFiltersOK {
	return &GetEntityFiltersOK{}
}

/* GetEntityFiltersOK describes a response with status code 200, with default header values.

Ok
*/
type GetEntityFiltersOK struct {
	Payload []*models.EntityFilter
}

func (o *GetEntityFiltersOK) Error() string {
	return fmt.Sprintf("[POST /get-entity-filters][%d] getEntityFiltersOK  %+v", 200, o.Payload)
}
func (o *GetEntityFiltersOK) GetPayload() []*models.EntityFilter {
	return o.Payload
}

func (o *GetEntityFiltersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEntityFiltersBadRequest creates a GetEntityFiltersBadRequest with default headers values
func NewGetEntityFiltersBadRequest() *GetEntityFiltersBadRequest {
	return &GetEntityFiltersBadRequest{}
}

/* GetEntityFiltersBadRequest describes a response with status code 400, with default header values.

GetEntityFiltersBadRequest get entity filters bad request
*/
type GetEntityFiltersBadRequest struct {
	Payload string
}

func (o *GetEntityFiltersBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-entity-filters][%d] getEntityFiltersBadRequest  %+v", 400, o.Payload)
}
func (o *GetEntityFiltersBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetEntityFiltersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}