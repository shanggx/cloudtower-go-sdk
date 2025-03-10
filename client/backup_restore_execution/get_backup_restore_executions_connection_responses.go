// Code generated by go-swagger; DO NOT EDIT.

package backup_restore_execution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetBackupRestoreExecutionsConnectionReader is a Reader for the GetBackupRestoreExecutionsConnection structure.
type GetBackupRestoreExecutionsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupRestoreExecutionsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupRestoreExecutionsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupRestoreExecutionsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupRestoreExecutionsConnectionOK creates a GetBackupRestoreExecutionsConnectionOK with default headers values
func NewGetBackupRestoreExecutionsConnectionOK() *GetBackupRestoreExecutionsConnectionOK {
	return &GetBackupRestoreExecutionsConnectionOK{}
}

/* GetBackupRestoreExecutionsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetBackupRestoreExecutionsConnectionOK struct {
	Payload *models.BackupRestoreExecutionConnection
}

func (o *GetBackupRestoreExecutionsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-executions-connection][%d] getBackupRestoreExecutionsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetBackupRestoreExecutionsConnectionOK) GetPayload() *models.BackupRestoreExecutionConnection {
	return o.Payload
}

func (o *GetBackupRestoreExecutionsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupRestoreExecutionConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupRestoreExecutionsConnectionBadRequest creates a GetBackupRestoreExecutionsConnectionBadRequest with default headers values
func NewGetBackupRestoreExecutionsConnectionBadRequest() *GetBackupRestoreExecutionsConnectionBadRequest {
	return &GetBackupRestoreExecutionsConnectionBadRequest{}
}

/* GetBackupRestoreExecutionsConnectionBadRequest describes a response with status code 400, with default header values.

GetBackupRestoreExecutionsConnectionBadRequest get backup restore executions connection bad request
*/
type GetBackupRestoreExecutionsConnectionBadRequest struct {
	Payload string
}

func (o *GetBackupRestoreExecutionsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-restore-executions-connection][%d] getBackupRestoreExecutionsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupRestoreExecutionsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetBackupRestoreExecutionsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
