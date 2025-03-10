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

// GetDisksReader is a Reader for the GetDisks structure.
type GetDisksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDisksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDisksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDisksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDisksOK creates a GetDisksOK with default headers values
func NewGetDisksOK() *GetDisksOK {
	return &GetDisksOK{}
}

/* GetDisksOK describes a response with status code 200, with default header values.

Ok
*/
type GetDisksOK struct {
	Payload []*models.Disk
}

func (o *GetDisksOK) Error() string {
	return fmt.Sprintf("[POST /get-disks][%d] getDisksOK  %+v", 200, o.Payload)
}
func (o *GetDisksOK) GetPayload() []*models.Disk {
	return o.Payload
}

func (o *GetDisksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDisksBadRequest creates a GetDisksBadRequest with default headers values
func NewGetDisksBadRequest() *GetDisksBadRequest {
	return &GetDisksBadRequest{}
}

/* GetDisksBadRequest describes a response with status code 400, with default header values.

GetDisksBadRequest get disks bad request
*/
type GetDisksBadRequest struct {
	Payload string
}

func (o *GetDisksBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-disks][%d] getDisksBadRequest  %+v", 400, o.Payload)
}
func (o *GetDisksBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetDisksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
