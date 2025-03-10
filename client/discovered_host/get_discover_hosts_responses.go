// Code generated by go-swagger; DO NOT EDIT.

package discovered_host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetDiscoverHostsReader is a Reader for the GetDiscoverHosts structure.
type GetDiscoverHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDiscoverHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDiscoverHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDiscoverHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDiscoverHostsOK creates a GetDiscoverHostsOK with default headers values
func NewGetDiscoverHostsOK() *GetDiscoverHostsOK {
	return &GetDiscoverHostsOK{}
}

/* GetDiscoverHostsOK describes a response with status code 200, with default header values.

Ok
*/
type GetDiscoverHostsOK struct {
	Payload []*models.DiscoveredHost
}

func (o *GetDiscoverHostsOK) Error() string {
	return fmt.Sprintf("[POST /get-discover-hosts][%d] getDiscoverHostsOK  %+v", 200, o.Payload)
}
func (o *GetDiscoverHostsOK) GetPayload() []*models.DiscoveredHost {
	return o.Payload
}

func (o *GetDiscoverHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDiscoverHostsBadRequest creates a GetDiscoverHostsBadRequest with default headers values
func NewGetDiscoverHostsBadRequest() *GetDiscoverHostsBadRequest {
	return &GetDiscoverHostsBadRequest{}
}

/* GetDiscoverHostsBadRequest describes a response with status code 400, with default header values.

GetDiscoverHostsBadRequest get discover hosts bad request
*/
type GetDiscoverHostsBadRequest struct {
	Payload string
}

func (o *GetDiscoverHostsBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-discover-hosts][%d] getDiscoverHostsBadRequest  %+v", 400, o.Payload)
}
func (o *GetDiscoverHostsBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetDiscoverHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
