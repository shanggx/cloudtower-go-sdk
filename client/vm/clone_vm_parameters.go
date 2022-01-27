// Code generated by go-swagger; DO NOT EDIT.

package vm

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

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// NewCloneVMParams creates a new CloneVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloneVMParams() *CloneVMParams {
	return &CloneVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloneVMParamsWithTimeout creates a new CloneVMParams object
// with the ability to set a timeout on a request.
func NewCloneVMParamsWithTimeout(timeout time.Duration) *CloneVMParams {
	return &CloneVMParams{
		timeout: timeout,
	}
}

// NewCloneVMParamsWithContext creates a new CloneVMParams object
// with the ability to set a context for a request.
func NewCloneVMParamsWithContext(ctx context.Context) *CloneVMParams {
	return &CloneVMParams{
		Context: ctx,
	}
}

// NewCloneVMParamsWithHTTPClient creates a new CloneVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloneVMParamsWithHTTPClient(client *http.Client) *CloneVMParams {
	return &CloneVMParams{
		HTTPClient: client,
	}
}

/* CloneVMParams contains all the parameters to send to the API endpoint
   for the clone Vm operation.

   Typically these are written to a http.Request.
*/
type CloneVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VMCloneParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the clone Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneVMParams) WithDefaults() *CloneVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the clone Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloneVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CloneVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the clone Vm params
func (o *CloneVMParams) WithTimeout(timeout time.Duration) *CloneVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone Vm params
func (o *CloneVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone Vm params
func (o *CloneVMParams) WithContext(ctx context.Context) *CloneVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone Vm params
func (o *CloneVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone Vm params
func (o *CloneVMParams) WithHTTPClient(client *http.Client) *CloneVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone Vm params
func (o *CloneVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the clone Vm params
func (o *CloneVMParams) WithContentLanguage(contentLanguage *string) *CloneVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the clone Vm params
func (o *CloneVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the clone Vm params
func (o *CloneVMParams) WithRequestBody(requestBody []*models.VMCloneParams) *CloneVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the clone Vm params
func (o *CloneVMParams) SetRequestBody(requestBody []*models.VMCloneParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CloneVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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