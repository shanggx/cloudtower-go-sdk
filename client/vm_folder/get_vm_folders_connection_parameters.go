// Code generated by go-swagger; DO NOT EDIT.

package vm_folder

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

// NewGetVMFoldersConnectionParams creates a new GetVMFoldersConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMFoldersConnectionParams() *GetVMFoldersConnectionParams {
	return &GetVMFoldersConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMFoldersConnectionParamsWithTimeout creates a new GetVMFoldersConnectionParams object
// with the ability to set a timeout on a request.
func NewGetVMFoldersConnectionParamsWithTimeout(timeout time.Duration) *GetVMFoldersConnectionParams {
	return &GetVMFoldersConnectionParams{
		timeout: timeout,
	}
}

// NewGetVMFoldersConnectionParamsWithContext creates a new GetVMFoldersConnectionParams object
// with the ability to set a context for a request.
func NewGetVMFoldersConnectionParamsWithContext(ctx context.Context) *GetVMFoldersConnectionParams {
	return &GetVMFoldersConnectionParams{
		Context: ctx,
	}
}

// NewGetVMFoldersConnectionParamsWithHTTPClient creates a new GetVMFoldersConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMFoldersConnectionParamsWithHTTPClient(client *http.Client) *GetVMFoldersConnectionParams {
	return &GetVMFoldersConnectionParams{
		HTTPClient: client,
	}
}

/* GetVMFoldersConnectionParams contains all the parameters to send to the API endpoint
   for the get Vm folders connection operation.

   Typically these are written to a http.Request.
*/
type GetVMFoldersConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMFoldersConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm folders connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMFoldersConnectionParams) WithDefaults() *GetVMFoldersConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm folders connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMFoldersConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMFoldersConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) WithTimeout(timeout time.Duration) *GetVMFoldersConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) WithContext(ctx context.Context) *GetVMFoldersConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) WithHTTPClient(client *http.Client) *GetVMFoldersConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) WithContentLanguage(contentLanguage *string) *GetVMFoldersConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) WithRequestBody(requestBody *models.GetVMFoldersConnectionRequestBody) *GetVMFoldersConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm folders connection params
func (o *GetVMFoldersConnectionParams) SetRequestBody(requestBody *models.GetVMFoldersConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMFoldersConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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