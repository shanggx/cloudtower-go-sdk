// Code generated by go-swagger; DO NOT EDIT.

package vm_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// DeleteVMSnapshotReader is a Reader for the DeleteVMSnapshot structure.
type DeleteVMSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVMSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteVMSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteVMSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteVMSnapshotOK creates a DeleteVMSnapshotOK with default headers values
func NewDeleteVMSnapshotOK() *DeleteVMSnapshotOK {
	return &DeleteVMSnapshotOK{}
}

/* DeleteVMSnapshotOK describes a response with status code 200, with default header values.

Ok
*/
type DeleteVMSnapshotOK struct {
	Payload []*models.WithTaskDeleteVMSnapshot
}

func (o *DeleteVMSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /delete-vm-snapshot][%d] deleteVmSnapshotOK  %+v", 200, o.Payload)
}
func (o *DeleteVMSnapshotOK) GetPayload() []*models.WithTaskDeleteVMSnapshot {
	return o.Payload
}

func (o *DeleteVMSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVMSnapshotBadRequest creates a DeleteVMSnapshotBadRequest with default headers values
func NewDeleteVMSnapshotBadRequest() *DeleteVMSnapshotBadRequest {
	return &DeleteVMSnapshotBadRequest{}
}

/* DeleteVMSnapshotBadRequest describes a response with status code 400, with default header values.

DeleteVMSnapshotBadRequest delete Vm snapshot bad request
*/
type DeleteVMSnapshotBadRequest struct {
	Payload string
}

func (o *DeleteVMSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /delete-vm-snapshot][%d] deleteVmSnapshotBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteVMSnapshotBadRequest) GetPayload() string {
	return o.Payload
}

func (o *DeleteVMSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
