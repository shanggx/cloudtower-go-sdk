// Code generated by go-swagger; DO NOT EDIT.

package migrate_transmitter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetMigrateTransmittersConnectionReader is a Reader for the GetMigrateTransmittersConnection structure.
type GetMigrateTransmittersConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMigrateTransmittersConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMigrateTransmittersConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMigrateTransmittersConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMigrateTransmittersConnectionOK creates a GetMigrateTransmittersConnectionOK with default headers values
func NewGetMigrateTransmittersConnectionOK() *GetMigrateTransmittersConnectionOK {
	return &GetMigrateTransmittersConnectionOK{}
}

/* GetMigrateTransmittersConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetMigrateTransmittersConnectionOK struct {
	Payload *models.MigrateTransmitterConnection
}

func (o *GetMigrateTransmittersConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-migrate-transmitters-connection][%d] getMigrateTransmittersConnectionOK  %+v", 200, o.Payload)
}
func (o *GetMigrateTransmittersConnectionOK) GetPayload() *models.MigrateTransmitterConnection {
	return o.Payload
}

func (o *GetMigrateTransmittersConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MigrateTransmitterConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrateTransmittersConnectionBadRequest creates a GetMigrateTransmittersConnectionBadRequest with default headers values
func NewGetMigrateTransmittersConnectionBadRequest() *GetMigrateTransmittersConnectionBadRequest {
	return &GetMigrateTransmittersConnectionBadRequest{}
}

/* GetMigrateTransmittersConnectionBadRequest describes a response with status code 400, with default header values.

GetMigrateTransmittersConnectionBadRequest get migrate transmitters connection bad request
*/
type GetMigrateTransmittersConnectionBadRequest struct {
	Payload string
}

func (o *GetMigrateTransmittersConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-migrate-transmitters-connection][%d] getMigrateTransmittersConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetMigrateTransmittersConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetMigrateTransmittersConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
