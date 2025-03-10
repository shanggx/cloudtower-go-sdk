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

// NewCreateVMParams creates a new CreateVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVMParams() *CreateVMParams {
	return &CreateVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVMParamsWithTimeout creates a new CreateVMParams object
// with the ability to set a timeout on a request.
func NewCreateVMParamsWithTimeout(timeout time.Duration) *CreateVMParams {
	return &CreateVMParams{
		timeout: timeout,
	}
}

// NewCreateVMParamsWithContext creates a new CreateVMParams object
// with the ability to set a context for a request.
func NewCreateVMParamsWithContext(ctx context.Context) *CreateVMParams {
	return &CreateVMParams{
		Context: ctx,
	}
}

// NewCreateVMParamsWithHTTPClient creates a new CreateVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVMParamsWithHTTPClient(client *http.Client) *CreateVMParams {
	return &CreateVMParams{
		HTTPClient: client,
	}
}

/* CreateVMParams contains all the parameters to send to the API endpoint
   for the create Vm operation.

   Typically these are written to a http.Request.
*/
type CreateVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VMCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMParams) WithDefaults() *CreateVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create Vm params
func (o *CreateVMParams) WithTimeout(timeout time.Duration) *CreateVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create Vm params
func (o *CreateVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create Vm params
func (o *CreateVMParams) WithContext(ctx context.Context) *CreateVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create Vm params
func (o *CreateVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create Vm params
func (o *CreateVMParams) WithHTTPClient(client *http.Client) *CreateVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create Vm params
func (o *CreateVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create Vm params
func (o *CreateVMParams) WithContentLanguage(contentLanguage *string) *CreateVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create Vm params
func (o *CreateVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create Vm params
func (o *CreateVMParams) WithRequestBody(requestBody []*models.VMCreationParams) *CreateVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create Vm params
func (o *CreateVMParams) SetRequestBody(requestBody []*models.VMCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
