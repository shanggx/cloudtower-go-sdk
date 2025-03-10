// Code generated by go-swagger; DO NOT EDIT.

package backup_plan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetBackupPlansConnectionReader is a Reader for the GetBackupPlansConnection structure.
type GetBackupPlansConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBackupPlansConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBackupPlansConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetBackupPlansConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBackupPlansConnectionOK creates a GetBackupPlansConnectionOK with default headers values
func NewGetBackupPlansConnectionOK() *GetBackupPlansConnectionOK {
	return &GetBackupPlansConnectionOK{}
}

/* GetBackupPlansConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetBackupPlansConnectionOK struct {
	Payload *models.BackupPlanConnection
}

func (o *GetBackupPlansConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-backup-plans-connection][%d] getBackupPlansConnectionOK  %+v", 200, o.Payload)
}
func (o *GetBackupPlansConnectionOK) GetPayload() *models.BackupPlanConnection {
	return o.Payload
}

func (o *GetBackupPlansConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BackupPlanConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBackupPlansConnectionBadRequest creates a GetBackupPlansConnectionBadRequest with default headers values
func NewGetBackupPlansConnectionBadRequest() *GetBackupPlansConnectionBadRequest {
	return &GetBackupPlansConnectionBadRequest{}
}

/* GetBackupPlansConnectionBadRequest describes a response with status code 400, with default header values.

GetBackupPlansConnectionBadRequest get backup plans connection bad request
*/
type GetBackupPlansConnectionBadRequest struct {
	Payload string
}

func (o *GetBackupPlansConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-backup-plans-connection][%d] getBackupPlansConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetBackupPlansConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetBackupPlansConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
