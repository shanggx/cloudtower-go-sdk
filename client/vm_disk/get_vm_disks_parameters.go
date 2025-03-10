// Code generated by go-swagger; DO NOT EDIT.

package vm_disk

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

// NewGetVMDisksParams creates a new GetVMDisksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMDisksParams() *GetVMDisksParams {
	return &GetVMDisksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMDisksParamsWithTimeout creates a new GetVMDisksParams object
// with the ability to set a timeout on a request.
func NewGetVMDisksParamsWithTimeout(timeout time.Duration) *GetVMDisksParams {
	return &GetVMDisksParams{
		timeout: timeout,
	}
}

// NewGetVMDisksParamsWithContext creates a new GetVMDisksParams object
// with the ability to set a context for a request.
func NewGetVMDisksParamsWithContext(ctx context.Context) *GetVMDisksParams {
	return &GetVMDisksParams{
		Context: ctx,
	}
}

// NewGetVMDisksParamsWithHTTPClient creates a new GetVMDisksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMDisksParamsWithHTTPClient(client *http.Client) *GetVMDisksParams {
	return &GetVMDisksParams{
		HTTPClient: client,
	}
}

/* GetVMDisksParams contains all the parameters to send to the API endpoint
   for the get Vm disks operation.

   Typically these are written to a http.Request.
*/
type GetVMDisksParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMDisksRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm disks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMDisksParams) WithDefaults() *GetVMDisksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm disks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMDisksParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMDisksParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm disks params
func (o *GetVMDisksParams) WithTimeout(timeout time.Duration) *GetVMDisksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm disks params
func (o *GetVMDisksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm disks params
func (o *GetVMDisksParams) WithContext(ctx context.Context) *GetVMDisksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm disks params
func (o *GetVMDisksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm disks params
func (o *GetVMDisksParams) WithHTTPClient(client *http.Client) *GetVMDisksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm disks params
func (o *GetVMDisksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm disks params
func (o *GetVMDisksParams) WithContentLanguage(contentLanguage *string) *GetVMDisksParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm disks params
func (o *GetVMDisksParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm disks params
func (o *GetVMDisksParams) WithRequestBody(requestBody *models.GetVMDisksRequestBody) *GetVMDisksParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm disks params
func (o *GetVMDisksParams) SetRequestBody(requestBody *models.GetVMDisksRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMDisksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
