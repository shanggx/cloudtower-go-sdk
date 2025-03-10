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

	"github.com/smartxworks/cloudtower-go-sdk/models"
)

// NewRemoveVMNicParams creates a new RemoveVMNicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRemoveVMNicParams() *RemoveVMNicParams {
	return &RemoveVMNicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveVMNicParamsWithTimeout creates a new RemoveVMNicParams object
// with the ability to set a timeout on a request.
func NewRemoveVMNicParamsWithTimeout(timeout time.Duration) *RemoveVMNicParams {
	return &RemoveVMNicParams{
		timeout: timeout,
	}
}

// NewRemoveVMNicParamsWithContext creates a new RemoveVMNicParams object
// with the ability to set a context for a request.
func NewRemoveVMNicParamsWithContext(ctx context.Context) *RemoveVMNicParams {
	return &RemoveVMNicParams{
		Context: ctx,
	}
}

// NewRemoveVMNicParamsWithHTTPClient creates a new RemoveVMNicParams object
// with the ability to set a custom HTTPClient for a request.
func NewRemoveVMNicParamsWithHTTPClient(client *http.Client) *RemoveVMNicParams {
	return &RemoveVMNicParams{
		HTTPClient: client,
	}
}

/* RemoveVMNicParams contains all the parameters to send to the API endpoint
   for the remove Vm nic operation.

   Typically these are written to a http.Request.
*/
type RemoveVMNicParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMRemoveNicParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the remove Vm nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMNicParams) WithDefaults() *RemoveVMNicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the remove Vm nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RemoveVMNicParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := RemoveVMNicParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the remove Vm nic params
func (o *RemoveVMNicParams) WithTimeout(timeout time.Duration) *RemoveVMNicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove Vm nic params
func (o *RemoveVMNicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove Vm nic params
func (o *RemoveVMNicParams) WithContext(ctx context.Context) *RemoveVMNicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove Vm nic params
func (o *RemoveVMNicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove Vm nic params
func (o *RemoveVMNicParams) WithHTTPClient(client *http.Client) *RemoveVMNicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove Vm nic params
func (o *RemoveVMNicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the remove Vm nic params
func (o *RemoveVMNicParams) WithContentLanguage(contentLanguage *string) *RemoveVMNicParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the remove Vm nic params
func (o *RemoveVMNicParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the remove Vm nic params
func (o *RemoveVMNicParams) WithRequestBody(requestBody *models.VMRemoveNicParams) *RemoveVMNicParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the remove Vm nic params
func (o *RemoveVMNicParams) SetRequestBody(requestBody *models.VMRemoveNicParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveVMNicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
