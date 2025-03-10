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

// NewGetDisksParams creates a new GetDisksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDisksParams() *GetDisksParams {
	return &GetDisksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDisksParamsWithTimeout creates a new GetDisksParams object
// with the ability to set a timeout on a request.
func NewGetDisksParamsWithTimeout(timeout time.Duration) *GetDisksParams {
	return &GetDisksParams{
		timeout: timeout,
	}
}

// NewGetDisksParamsWithContext creates a new GetDisksParams object
// with the ability to set a context for a request.
func NewGetDisksParamsWithContext(ctx context.Context) *GetDisksParams {
	return &GetDisksParams{
		Context: ctx,
	}
}

// NewGetDisksParamsWithHTTPClient creates a new GetDisksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDisksParamsWithHTTPClient(client *http.Client) *GetDisksParams {
	return &GetDisksParams{
		HTTPClient: client,
	}
}

/* GetDisksParams contains all the parameters to send to the API endpoint
   for the get disks operation.

   Typically these are written to a http.Request.
*/
type GetDisksParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetDisksRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get disks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDisksParams) WithDefaults() *GetDisksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get disks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDisksParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetDisksParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get disks params
func (o *GetDisksParams) WithTimeout(timeout time.Duration) *GetDisksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get disks params
func (o *GetDisksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get disks params
func (o *GetDisksParams) WithContext(ctx context.Context) *GetDisksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get disks params
func (o *GetDisksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get disks params
func (o *GetDisksParams) WithHTTPClient(client *http.Client) *GetDisksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get disks params
func (o *GetDisksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get disks params
func (o *GetDisksParams) WithContentLanguage(contentLanguage *string) *GetDisksParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get disks params
func (o *GetDisksParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get disks params
func (o *GetDisksParams) WithRequestBody(requestBody *models.GetDisksRequestBody) *GetDisksParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get disks params
func (o *GetDisksParams) SetRequestBody(requestBody *models.GetDisksRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetDisksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
