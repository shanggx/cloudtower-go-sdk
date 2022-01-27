// Code generated by go-swagger; DO NOT EDIT.

package usb_device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// GetUsbDevicesConnectionReader is a Reader for the GetUsbDevicesConnection structure.
type GetUsbDevicesConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsbDevicesConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsbDevicesConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsbDevicesConnectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsbDevicesConnectionOK creates a GetUsbDevicesConnectionOK with default headers values
func NewGetUsbDevicesConnectionOK() *GetUsbDevicesConnectionOK {
	return &GetUsbDevicesConnectionOK{}
}

/* GetUsbDevicesConnectionOK describes a response with status code 200, with default header values.

Ok
*/
type GetUsbDevicesConnectionOK struct {
	Payload *models.UsbDeviceConnection
}

func (o *GetUsbDevicesConnectionOK) Error() string {
	return fmt.Sprintf("[POST /get-usb-devices-connection][%d] getUsbDevicesConnectionOK  %+v", 200, o.Payload)
}
func (o *GetUsbDevicesConnectionOK) GetPayload() *models.UsbDeviceConnection {
	return o.Payload
}

func (o *GetUsbDevicesConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UsbDeviceConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsbDevicesConnectionBadRequest creates a GetUsbDevicesConnectionBadRequest with default headers values
func NewGetUsbDevicesConnectionBadRequest() *GetUsbDevicesConnectionBadRequest {
	return &GetUsbDevicesConnectionBadRequest{}
}

/* GetUsbDevicesConnectionBadRequest describes a response with status code 400, with default header values.

GetUsbDevicesConnectionBadRequest get usb devices connection bad request
*/
type GetUsbDevicesConnectionBadRequest struct {
	Payload string
}

func (o *GetUsbDevicesConnectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-usb-devices-connection][%d] getUsbDevicesConnectionBadRequest  %+v", 400, o.Payload)
}
func (o *GetUsbDevicesConnectionBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetUsbDevicesConnectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}