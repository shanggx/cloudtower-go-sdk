// Code generated by go-swagger; DO NOT EDIT.

package vsphere_esxi_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetVsphereEsxiAccountsReader is a Reader for the GetVsphereEsxiAccounts structure.
type GetVsphereEsxiAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVsphereEsxiAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVsphereEsxiAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetVsphereEsxiAccountsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetVsphereEsxiAccountsOK creates a GetVsphereEsxiAccountsOK with default headers values
func NewGetVsphereEsxiAccountsOK() *GetVsphereEsxiAccountsOK {
	return &GetVsphereEsxiAccountsOK{}
}

/* GetVsphereEsxiAccountsOK describes a response with status code 200, with default header values.

Ok
*/
type GetVsphereEsxiAccountsOK struct {
	Payload []*models.VsphereEsxiAccount
}

func (o *GetVsphereEsxiAccountsOK) Error() string {
	return fmt.Sprintf("[POST /get-vsphere-esxi-accounts][%d] getVsphereEsxiAccountsOK  %+v", 200, o.Payload)
}
func (o *GetVsphereEsxiAccountsOK) GetPayload() []*models.VsphereEsxiAccount {
	return o.Payload
}

func (o *GetVsphereEsxiAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVsphereEsxiAccountsBadRequest creates a GetVsphereEsxiAccountsBadRequest with default headers values
func NewGetVsphereEsxiAccountsBadRequest() *GetVsphereEsxiAccountsBadRequest {
	return &GetVsphereEsxiAccountsBadRequest{}
}

/* GetVsphereEsxiAccountsBadRequest describes a response with status code 400, with default header values.

GetVsphereEsxiAccountsBadRequest get vsphere esxi accounts bad request
*/
type GetVsphereEsxiAccountsBadRequest struct {
	Payload string
}

func (o *GetVsphereEsxiAccountsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-vsphere-esxi-accounts][%d] getVsphereEsxiAccountsBadRequest  %+v", 400, o.Payload)
}
func (o *GetVsphereEsxiAccountsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetVsphereEsxiAccountsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
