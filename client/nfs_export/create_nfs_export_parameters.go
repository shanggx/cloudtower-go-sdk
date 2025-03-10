// Code generated by go-swagger; DO NOT EDIT.

package nfs_export

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

// NewCreateNfsExportParams creates a new CreateNfsExportParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateNfsExportParams() *CreateNfsExportParams {
	return &CreateNfsExportParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNfsExportParamsWithTimeout creates a new CreateNfsExportParams object
// with the ability to set a timeout on a request.
func NewCreateNfsExportParamsWithTimeout(timeout time.Duration) *CreateNfsExportParams {
	return &CreateNfsExportParams{
		timeout: timeout,
	}
}

// NewCreateNfsExportParamsWithContext creates a new CreateNfsExportParams object
// with the ability to set a context for a request.
func NewCreateNfsExportParamsWithContext(ctx context.Context) *CreateNfsExportParams {
	return &CreateNfsExportParams{
		Context: ctx,
	}
}

// NewCreateNfsExportParamsWithHTTPClient creates a new CreateNfsExportParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateNfsExportParamsWithHTTPClient(client *http.Client) *CreateNfsExportParams {
	return &CreateNfsExportParams{
		HTTPClient: client,
	}
}

/* CreateNfsExportParams contains all the parameters to send to the API endpoint
   for the create nfs export operation.

   Typically these are written to a http.Request.
*/
type CreateNfsExportParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.NfsExportCreationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create nfs export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNfsExportParams) WithDefaults() *CreateNfsExportParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create nfs export params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateNfsExportParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateNfsExportParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create nfs export params
func (o *CreateNfsExportParams) WithTimeout(timeout time.Duration) *CreateNfsExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create nfs export params
func (o *CreateNfsExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create nfs export params
func (o *CreateNfsExportParams) WithContext(ctx context.Context) *CreateNfsExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create nfs export params
func (o *CreateNfsExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create nfs export params
func (o *CreateNfsExportParams) WithHTTPClient(client *http.Client) *CreateNfsExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create nfs export params
func (o *CreateNfsExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create nfs export params
func (o *CreateNfsExportParams) WithContentLanguage(contentLanguage *string) *CreateNfsExportParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create nfs export params
func (o *CreateNfsExportParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create nfs export params
func (o *CreateNfsExportParams) WithRequestBody(requestBody []*models.NfsExportCreationParams) *CreateNfsExportParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create nfs export params
func (o *CreateNfsExportParams) SetRequestBody(requestBody []*models.NfsExportCreationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNfsExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
