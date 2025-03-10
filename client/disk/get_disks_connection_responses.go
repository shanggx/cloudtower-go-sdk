// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetDisksConnectionReader is a Reader for the GetDisksConnection structure.
type GetDisksConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDisksConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDisksConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDisksConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDisksConnectionOK creates a GetDisksConnectionOK with default headers values
func NewGetDisksConnectionOK() *GetDisksConnectionOK {
	return &GetDisksConnectionOK{}
}

/* GetDisksConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetDisksConnectionOK struct {
	Payload *models.DiskConnection
}

func (o *GetDisksConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-disks-connection][%d] getDisksConnectionOK  %+v", 200, o.Payload)
}
func (o *GetDisksConnectionOK) GetPayload() *models.DiskConnection {
	return o.Payload
}

func (o *GetDisksConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DiskConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDisksConnectionBadRequest creates a GetDisksConnectionBadRequest with default headers values
func NewGetDisksConnectionBadRequest() *GetDisksConnectionBadRequest {
	return &GetDisksConnectionBadRequest{}
}

/* GetDisksConnectionBadRequest describes a response with status code 400, with default header values.

GetDisksConnectionBadRequest get disks connection bad request
*/
type GetDisksConnectionBadRequest struct {
	Payload string
}

func (o *GetDisksConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-disks-connection][%d] getDisksConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetDisksConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetDisksConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
