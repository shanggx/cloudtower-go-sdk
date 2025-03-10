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

// GetRackTopoesConnectionReader is a Reader for the GetRackTopoesConnection structure.
type GetRackTopoesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRackTopoesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRackTopoesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRackTopoesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRackTopoesConnectionOK creates a GetRackTopoesConnectionOK with default headers values
func NewGetRackTopoesConnectionOK() *GetRackTopoesConnectionOK {
	return &GetRackTopoesConnectionOK{}
}

/* GetRackTopoesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetRackTopoesConnectionOK struct {
	Payload *models.RackTopoConnection
}

func (o *GetRackTopoesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-rack-topoes-connection][%d] getRackTopoesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetRackTopoesConnectionOK) GetPayload() *models.RackTopoConnection {
	return o.Payload
}

func (o *GetRackTopoesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackTopoConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRackTopoesConnectionBadRequest creates a GetRackTopoesConnectionBadRequest with default headers values
func NewGetRackTopoesConnectionBadRequest() *GetRackTopoesConnectionBadRequest {
	return &GetRackTopoesConnectionBadRequest{}
}

/* GetRackTopoesConnectionBadRequest describes a response with status code 400, with default header values.

GetRackTopoesConnectionBadRequest get rack topoes connection bad request
*/
type GetRackTopoesConnectionBadRequest struct {
	Payload string
}

func (o *GetRackTopoesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-rack-topoes-connection][%d] getRackTopoesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetRackTopoesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetRackTopoesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
