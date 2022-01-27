// Code generated by go-swagger; DO NOT EDIT.

package nfs_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetNfsExportsReader is a Reader for the GetNfsExports structure.
type GetNfsExportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNfsExportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNfsExportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNfsExportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNfsExportsOK creates a GetNfsExportsOK with default headers values
func NewGetNfsExportsOK() *GetNfsExportsOK {
	return &GetNfsExportsOK{}
}

/* GetNfsExportsOK describes a response with status code 200, with default header values.

Ok
*/
type GetNfsExportsOK struct {
	Payload []*models.NfsExport
}

func (o *GetNfsExportsOK) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports][%d] getNfsExportsOK  %+v", 200, o.Payload)
}
func (o *GetNfsExportsOK) GetPayload() []*models.NfsExport {
	return o.Payload
}

func (o *GetNfsExportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNfsExportsBadRequest creates a GetNfsExportsBadRequest with default headers values
func NewGetNfsExportsBadRequest() *GetNfsExportsBadRequest {
	return &GetNfsExportsBadRequest{}
}

/* GetNfsExportsBadRequest describes a response with status code 400, with default header values.

GetNfsExportsBadRequest get nfs exports bad request
*/
type GetNfsExportsBadRequest struct {
	Payload string
}

func (o *GetNfsExportsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nfs-exports][%d] getNfsExportsBadRequest  %+v", 400, o.Payload)
}
func (o *GetNfsExportsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNfsExportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}