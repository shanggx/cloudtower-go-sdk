// Code generated by go-swagger; DO NOT EDIT.

package cluster_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// GetClusterImagesReader is a Reader for the GetClusterImages structure.
type GetClusterImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClusterImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetClusterImagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetClusterImagesOK creates a GetClusterImagesOK with default headers values
func NewGetClusterImagesOK() *GetClusterImagesOK {
	return &GetClusterImagesOK{}
}

/* GetClusterImagesOK describes a response with status code 200, with default header values.

Ok
*/
type GetClusterImagesOK struct {
	Payload []*models.ClusterImage
}

func (o *GetClusterImagesOK) Error() string {
	return fmt.Sprintf("[POST /get-cluster-images][%d] getClusterImagesOK  %+v", 200, o.Payload)
}
func (o *GetClusterImagesOK) GetPayload() []*models.ClusterImage {
	return o.Payload
}

func (o *GetClusterImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterImagesBadRequest creates a GetClusterImagesBadRequest with default headers values
func NewGetClusterImagesBadRequest() *GetClusterImagesBadRequest {
	return &GetClusterImagesBadRequest{}
}

/* GetClusterImagesBadRequest describes a response with status code 400, with default header values.

GetClusterImagesBadRequest get cluster images bad request
*/
type GetClusterImagesBadRequest struct {
	Payload string
}

func (o *GetClusterImagesBadRequest) Error() string {
	return fmt.Sprintf("[POST /get-cluster-images][%d] getClusterImagesBadRequest  %+v", 400, o.Payload)
}
func (o *GetClusterImagesBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetClusterImagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
