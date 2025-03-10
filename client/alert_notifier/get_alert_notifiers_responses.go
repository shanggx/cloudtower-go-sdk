// Code generated by go-swagger; DO NOT EDIT.

package alert_notifier

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetAlertNotifiersReader is a Reader for the GetAlertNotifiers structure.
type GetAlertNotifiersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertNotifiersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAlertNotifiersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAlertNotifiersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAlertNotifiersOK creates a GetAlertNotifiersOK with default headers values
func NewGetAlertNotifiersOK() *GetAlertNotifiersOK {
	return &GetAlertNotifiersOK{}
}

/* GetAlertNotifiersOK describes a response with status code 200, with default header values.

Ok
*/
type GetAlertNotifiersOK struct {
	Payload []*models.AlertNotifier
}

func (o *GetAlertNotifiersOK) Error() string {
	return fmt.Sprintf("[POST /get-alert-notifiers][%d] getAlertNotifiersOK  %+v", 200, o.Payload)
}
func (o *GetAlertNotifiersOK) GetPayload() []*models.AlertNotifier {
	return o.Payload
}

func (o *GetAlertNotifiersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertNotifiersBadRequest creates a GetAlertNotifiersBadRequest with default headers values
func NewGetAlertNotifiersBadRequest() *GetAlertNotifiersBadRequest {
	return &GetAlertNotifiersBadRequest{}
}

/* GetAlertNotifiersBadRequest describes a response with status code 400, with default header values.

GetAlertNotifiersBadRequest get alert notifiers bad request
*/
type GetAlertNotifiersBadRequest struct {
	Payload string
}

func (o *GetAlertNotifiersBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-alert-notifiers][%d] getAlertNotifiersBadRequest  %+v", 400, o.Payload)
}
func (o *GetAlertNotifiersBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetAlertNotifiersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
