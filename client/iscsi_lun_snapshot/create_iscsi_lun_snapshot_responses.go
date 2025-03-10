// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// CreateIscsiLunSnapshotReader is a Reader for the CreateIscsiLunSnapshot structure.
type CreateIscsiLunSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateIscsiLunSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateIscsiLunSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateIscsiLunSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateIscsiLunSnapshotOK creates a CreateIscsiLunSnapshotOK with default headers values
func NewCreateIscsiLunSnapshotOK() *CreateIscsiLunSnapshotOK {
	return &CreateIscsiLunSnapshotOK{}
}

/* CreateIscsiLunSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type CreateIscsiLunSnapshotOK struct {
	Payload []*models.WithTaskIscsiLunSnapshot
}

func (o *CreateIscsiLunSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /create-iscsi-lun-snapshot][%d] createIscsiLunSnapshotOK  %+v", 200, o.Payload)
}
func (o *CreateIscsiLunSnapshotOK) GetPayload() []*models.WithTaskIscsiLunSnapshot {
	return o.Payload
}

func (o *CreateIscsiLunSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateIscsiLunSnapshotBadRequest creates a CreateIscsiLunSnapshotBadRequest with default headers values
func NewCreateIscsiLunSnapshotBadRequest() *CreateIscsiLunSnapshotBadRequest {
	return &CreateIscsiLunSnapshotBadRequest{}
}

/* CreateIscsiLunSnapshotBadRequest describes a response with status code 400, with default header values.

CreateIscsiLunSnapshotBadRequest create iscsi lun snapshot bad request
*/
type CreateIscsiLunSnapshotBadRequest struct {
	Payload string
}

func (o *CreateIscsiLunSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /create-iscsi-lun-snapshot][%d] createIscsiLunSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *CreateIscsiLunSnapshotBadRequest) GetPayload() string {
	return o.Payload
}

func (o *CreateIscsiLunSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
