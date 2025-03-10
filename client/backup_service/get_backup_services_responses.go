// Code generated by go-swagger; DO NOT EDIT.

package backup_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetBackupServicesReader is a Reader for the GetBackupServices structure.
type GetBackupServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupServicesOK creates a GetBackupServicesOK with default headers values
func NewGetBackupServicesOK() *GetBackupServicesOK {
	return &GetBackupServicesOK{}
}

/* GetBackupServicesOK describes a response with status code 200, with default header values.

Ok
*/
type GetBackupServicesOK struct {
	Payload []*models.BackupService
}

func (o *GetBackupServicesOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-services][%d] getBackupServicesOK  %+v", 200, o.Payload)
}
func (o *GetBackupServicesOK) GetPayload() []*models.BackupService {
	return o.Payload
}

func (o *GetBackupServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupServicesBadRequest creates a GetBackupServicesBadRequest with default headers values
func NewGetBackupServicesBadRequest() *GetBackupServicesBadRequest {
	return &GetBackupServicesBadRequest{}
}

/* GetBackupServicesBadRequest describes a response with status code 400, with default header values.

GetBackupServicesBadRequest get backup services bad request
*/
type GetBackupServicesBadRequest struct {
	Payload string
}

func (o *GetBackupServicesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-services][%d] getBackupServicesBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupServicesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetBackupServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
