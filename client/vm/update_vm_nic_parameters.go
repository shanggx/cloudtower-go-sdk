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

// NewUpdateVMNicParams creates a new UpdateVMNicParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateVMNicParams() *UpdateVMNicParams {
	return &UpdateVMNicParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVMNicParamsWithTimeout creates a new UpdateVMNicParams object
// with the ability to set a timeout on a request.
func NewUpdateVMNicParamsWithTimeout(timeout time.Duration) *UpdateVMNicParams {
	return &UpdateVMNicParams{
		timeout: timeout,
	}
}

// NewUpdateVMNicParamsWithContext creates a new UpdateVMNicParams object
// with the ability to set a context for a request.
func NewUpdateVMNicParamsWithContext(ctx context.Context) *UpdateVMNicParams {
	return &UpdateVMNicParams{
		Context: ctx,
	}
}

// NewUpdateVMNicParamsWithHTTPClient creates a new UpdateVMNicParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateVMNicParamsWithHTTPClient(client *http.Client) *UpdateVMNicParams {
	return &UpdateVMNicParams{
		HTTPClient: client,
	}
}

/* UpdateVMNicParams contains all the parameters to send to the API endpoint
   for the update Vm nic operation.

   Typically these are written to a http.Request.
*/
type UpdateVMNicParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMUpdateNicParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update Vm nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMNicParams) WithDefaults() *UpdateVMNicParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update Vm nic params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateVMNicParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateVMNicParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update Vm nic params
func (o *UpdateVMNicParams) WithTimeout(timeout time.Duration) *UpdateVMNicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Vm nic params
func (o *UpdateVMNicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Vm nic params
func (o *UpdateVMNicParams) WithContext(ctx context.Context) *UpdateVMNicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Vm nic params
func (o *UpdateVMNicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Vm nic params
func (o *UpdateVMNicParams) WithHTTPClient(client *http.Client) *UpdateVMNicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Vm nic params
func (o *UpdateVMNicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update Vm nic params
func (o *UpdateVMNicParams) WithContentLanguage(contentLanguage *string) *UpdateVMNicParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update Vm nic params
func (o *UpdateVMNicParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update Vm nic params
func (o *UpdateVMNicParams) WithRequestBody(requestBody *models.VMUpdateNicParams) *UpdateVMNicParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update Vm nic params
func (o *UpdateVMNicParams) SetRequestBody(requestBody *models.VMUpdateNicParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVMNicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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