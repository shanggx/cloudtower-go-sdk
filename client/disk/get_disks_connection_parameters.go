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

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// NewGetDisksConnectionParams creates a new GetDisksConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDisksConnectionParams() *GetDisksConnectionParams {
	return &GetDisksConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDisksConnectionParamsWithTimeout creates a new GetDisksConnectionParams object
// with the ability to set a timeout on a request.
func NewGetDisksConnectionParamsWithTimeout(timeout time.Duration) *GetDisksConnectionParams {
	return &GetDisksConnectionParams{
		timeout: timeout,
	}
}

// NewGetDisksConnectionParamsWithContext creates a new GetDisksConnectionParams object
// with the ability to set a context for a request.
func NewGetDisksConnectionParamsWithContext(ctx context.Context) *GetDisksConnectionParams {
	return &GetDisksConnectionParams{
		Context: ctx,
	}
}

// NewGetDisksConnectionParamsWithHTTPClient creates a new GetDisksConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDisksConnectionParamsWithHTTPClient(client *http.Client) *GetDisksConnectionParams {
	return &GetDisksConnectionParams{
		HTTPClient: client,
	}
}

/* GetDisksConnectionParams contains all the parameters to send to the API endpoint
   for the get disks connection operation.

   Typically these are written to a http.Request.
*/
type GetDisksConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetDisksConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get disks connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDisksConnectionParams) WithDefaults() *GetDisksConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get disks connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDisksConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetDisksConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get disks connection params
func (o *GetDisksConnectionParams) WithTimeout(timeout time.Duration) *GetDisksConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get disks connection params
func (o *GetDisksConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get disks connection params
func (o *GetDisksConnectionParams) WithContext(ctx context.Context) *GetDisksConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get disks connection params
func (o *GetDisksConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get disks connection params
func (o *GetDisksConnectionParams) WithHTTPClient(client *http.Client) *GetDisksConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get disks connection params
func (o *GetDisksConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get disks connection params
func (o *GetDisksConnectionParams) WithContentLanguage(contentLanguage *string) *GetDisksConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get disks connection params
func (o *GetDisksConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get disks connection params
func (o *GetDisksConnectionParams) WithRequestBody(requestBody *models.GetDisksConnectionRequestBody) *GetDisksConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get disks connection params
func (o *GetDisksConnectionParams) SetRequestBody(requestBody *models.GetDisksConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetDisksConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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