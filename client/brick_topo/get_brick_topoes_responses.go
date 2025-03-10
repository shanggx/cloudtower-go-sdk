// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetBrickTopoesReader is a Reader for the GetBrickTopoes structure.
type GetBrickTopoesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBrickTopoesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBrickTopoesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBrickTopoesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBrickTopoesOK creates a GetBrickTopoesOK with default headers values
func NewGetBrickTopoesOK() *GetBrickTopoesOK {
	return &GetBrickTopoesOK{}
}

/* GetBrickTopoesOK describes a response with status code 200, with default header values.

Ok
*/
type GetBrickTopoesOK struct {
	Payload []*models.BrickTopo
}

func (o *GetBrickTopoesOK) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes][%d] getBrickTopoesOK  %+v", 200, o.Payload)
}
func (o *GetBrickTopoesOK) GetPayload() []*models.BrickTopo {
	return o.Payload
}

func (o *GetBrickTopoesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBrickTopoesBadRequest creates a GetBrickTopoesBadRequest with default headers values
func NewGetBrickTopoesBadRequest() *GetBrickTopoesBadRequest {
	return &GetBrickTopoesBadRequest{}
}

/* GetBrickTopoesBadRequest describes a response with status code 400, with default header values.

GetBrickTopoesBadRequest get brick topoes bad request
*/
type GetBrickTopoesBadRequest struct {
	Payload string
}

func (o *GetBrickTopoesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-brick-topoes][%d] getBrickTopoesBadRequest  %+v", 400, o.Payload)
}
func (o *GetBrickTopoesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetBrickTopoesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
