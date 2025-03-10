// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// NewMountDiskParams creates a new MountDiskParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMountDiskParams() *MountDiskParams {
	return &MountDiskParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMountDiskParamsWithTimeout creates a new MountDiskParams object
// with the ability to set a timeout on a request.
func NewMountDiskParamsWithTimeout(timeout time.Duration) *MountDiskParams {
	return &MountDiskParams{
		timeout: timeout,
	}
}

// NewMountDiskParamsWithContext creates a new MountDiskParams object
// with the ability to set a context for a request.
func NewMountDiskParamsWithContext(ctx context.Context) *MountDiskParams {
	return &MountDiskParams{
		Context: ctx,
	}
}

// NewMountDiskParamsWithHTTPClient creates a new MountDiskParams object
// with the ability to set a custom HTTPClient for a request.
func NewMountDiskParamsWithHTTPClient(client *http.Client) *MountDiskParams {
	return &MountDiskParams{
		HTTPClient: client,
	}
}

/* MountDiskParams contains all the parameters to send to the API endpoint
   for the mount disk operation.

   Typically these are written to a http.Request.
*/
type MountDiskParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.DiskMountParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mount disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MountDiskParams) WithDefaults() *MountDiskParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mount disk params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MountDiskParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := MountDiskParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the mount disk params
func (o *MountDiskParams) WithTimeout(timeout time.Duration) *MountDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mount disk params
func (o *MountDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mount disk params
func (o *MountDiskParams) WithContext(ctx context.Context) *MountDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mount disk params
func (o *MountDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mount disk params
func (o *MountDiskParams) WithHTTPClient(client *http.Client) *MountDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mount disk params
func (o *MountDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the mount disk params
func (o *MountDiskParams) WithContentLanguage(contentLanguage *string) *MountDiskParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the mount disk params
func (o *MountDiskParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the mount disk params
func (o *MountDiskParams) WithRequestBody(requestBody *models.DiskMountParams) *MountDiskParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the mount disk params
func (o *MountDiskParams) SetRequestBody(requestBody *models.DiskMountParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *MountDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentLanguage != nil {

		// header param content-language
		if err := r.SetHeaderParam("content-language", *o.ContentLanguage); err != nil {
			return err
		}
	}
	if o.RequestBody != nil {
		if err := r.SetBodyParam(o.RequestBody); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
