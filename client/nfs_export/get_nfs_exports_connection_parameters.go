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

// NewGetNfsExportsConnectionParams creates a new GetNfsExportsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNfsExportsConnectionParams() *GetNfsExportsConnectionParams {
	return &GetNfsExportsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNfsExportsConnectionParamsWithTimeout creates a new GetNfsExportsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetNfsExportsConnectionParamsWithTimeout(timeout time.Duration) *GetNfsExportsConnectionParams {
	return &GetNfsExportsConnectionParams{
		timeout: timeout,
	}
}

// NewGetNfsExportsConnectionParamsWithContext creates a new GetNfsExportsConnectionParams object
// with the ability to set a context for a request.
func NewGetNfsExportsConnectionParamsWithContext(ctx context.Context) *GetNfsExportsConnectionParams {
	return &GetNfsExportsConnectionParams{
		Context: ctx,
	}
}

// NewGetNfsExportsConnectionParamsWithHTTPClient creates a new GetNfsExportsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNfsExportsConnectionParamsWithHTTPClient(client *http.Client) *GetNfsExportsConnectionParams {
	return &GetNfsExportsConnectionParams{
		HTTPClient: client,
	}
}

/* GetNfsExportsConnectionParams contains all the parameters to send to the API endpoint
   for the get nfs exports connection operation.

   Typically these are written to a http.Request.
*/
type GetNfsExportsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetNfsExportsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get nfs exports connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsExportsConnectionParams) WithDefaults() *GetNfsExportsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get nfs exports connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNfsExportsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetNfsExportsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) WithTimeout(timeout time.Duration) *GetNfsExportsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) WithContext(ctx context.Context) *GetNfsExportsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) WithHTTPClient(client *http.Client) *GetNfsExportsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) WithContentLanguage(contentLanguage *string) *GetNfsExportsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) WithRequestBody(requestBody *models.GetNfsExportsConnectionRequestBody) *GetNfsExportsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get nfs exports connection params
func (o *GetNfsExportsConnectionParams) SetRequestBody(requestBody *models.GetNfsExportsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetNfsExportsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
