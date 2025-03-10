// Code generated by go-swagger; DO NOT EDIT.

package rack_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateRackTopoReader is a Reader for the CreateRackTopo structure.
type CreateRackTopoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRackTopoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateRackTopoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRackTopoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRackTopoOK creates a CreateRackTopoOK with default headers values
func NewCreateRackTopoOK() *CreateRackTopoOK {
	return &CreateRackTopoOK{}
}

/* CreateRackTopoOK describes a response with status code 200, with default header values.

Ok
*/
type CreateRackTopoOK struct {
	Payload []*models.WithTaskRackTopo
}

func (o *CreateRackTopoOK) Error() string {
	return fmt.Sprintf("[POST /create-rack-topo][%d] createRackTopoOK  %+v", 200, o.Payload)
}
func (o *CreateRackTopoOK) GetPayload() []*models.WithTaskRackTopo {
	return o.Payload
}

func (o *CreateRackTopoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRackTopoBadRequest creates a CreateRackTopoBadRequest with default headers values
func NewCreateRackTopoBadRequest() *CreateRackTopoBadRequest {
	return &CreateRackTopoBadRequest{}
}

/* CreateRackTopoBadRequest describes a response with status code 400, with default header values.

CreateRackTopoBadRequest create rack topo bad request
*/
type CreateRackTopoBadRequest struct {
	Payload string
}

func (o *CreateRackTopoBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-rack-topo][%d] createRackTopoBadRequest  %+v", 400, o.Payload)
}
func (o *CreateRackTopoBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateRackTopoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
