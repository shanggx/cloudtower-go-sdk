// Code generated by go-swagger; DO NOT EDIT.

package nvmf_namespace_snapshot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetNvmfNamespaceSnapshotsConnectionReader is a Reader for the GetNvmfNamespaceSnapshotsConnection structure.
type GetNvmfNamespaceSnapshotsConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNvmfNamespaceSnapshotsConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNvmfNamespaceSnapshotsConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNvmfNamespaceSnapshotsConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNvmfNamespaceSnapshotsConnectionOK creates a GetNvmfNamespaceSnapshotsConnectionOK with default headers values
func NewGetNvmfNamespaceSnapshotsConnectionOK() *GetNvmfNamespaceSnapshotsConnectionOK {
	return &GetNvmfNamespaceSnapshotsConnectionOK{}
}

/* GetNvmfNamespaceSnapshotsConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetNvmfNamespaceSnapshotsConnectionOK struct {
	Payload *models.NvmfNamespaceSnapshotConnection
}

func (o *GetNvmfNamespaceSnapshotsConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-namespace-snapshots-connection][%d] getNvmfNamespaceSnapshotsConnectionOK  %+v", 200, o.Payload)
}
func (o *GetNvmfNamespaceSnapshotsConnectionOK) GetPayload() *models.NvmfNamespaceSnapshotConnection {
	return o.Payload
}

func (o *GetNvmfNamespaceSnapshotsConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmfNamespaceSnapshotConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNvmfNamespaceSnapshotsConnectionBadRequest creates a GetNvmfNamespaceSnapshotsConnectionBadRequest with default headers values
func NewGetNvmfNamespaceSnapshotsConnectionBadRequest() *GetNvmfNamespaceSnapshotsConnectionBadRequest {
	return &GetNvmfNamespaceSnapshotsConnectionBadRequest{}
}

/* GetNvmfNamespaceSnapshotsConnectionBadRequest describes a response with status code 400, with default header values.

GetNvmfNamespaceSnapshotsConnectionBadRequest get nvmf namespace snapshots connection bad request
*/
type GetNvmfNamespaceSnapshotsConnectionBadRequest struct {
	Payload string
}

func (o *GetNvmfNamespaceSnapshotsConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-nvmf-namespace-snapshots-connection][%d] getNvmfNamespaceSnapshotsConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetNvmfNamespaceSnapshotsConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetNvmfNamespaceSnapshotsConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}