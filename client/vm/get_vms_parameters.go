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

// NewGetVmsParams creates a new GetVmsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVmsParams() *GetVmsParams {
	return &GetVmsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVmsParamsWithTimeout creates a new GetVmsParams object
// with the ability to set a timeout on a request.
func NewGetVmsParamsWithTimeout(timeout time.Duration) *GetVmsParams {
	return &GetVmsParams{
		timeout: timeout,
	}
}

// NewGetVmsParamsWithContext creates a new GetVmsParams object
// with the ability to set a context for a request.
func NewGetVmsParamsWithContext(ctx context.Context) *GetVmsParams {
	return &GetVmsParams{
		Context: ctx,
	}
}

// NewGetVmsParamsWithHTTPClient creates a new GetVmsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVmsParamsWithHTTPClient(client *http.Client) *GetVmsParams {
	return &GetVmsParams{
		HTTPClient: client,
	}
}

/* GetVmsParams contains all the parameters to send to the API endpoint
   for the get vms operation.

   Typically these are written to a http.Request.
*/
type GetVmsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVmsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get vms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVmsParams) WithDefaults() *GetVmsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get vms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVmsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVmsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get vms params
func (o *GetVmsParams) WithTimeout(timeout time.Duration) *GetVmsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vms params
func (o *GetVmsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vms params
func (o *GetVmsParams) WithContext(ctx context.Context) *GetVmsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vms params
func (o *GetVmsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vms params
func (o *GetVmsParams) WithHTTPClient(client *http.Client) *GetVmsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vms params
func (o *GetVmsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get vms params
func (o *GetVmsParams) WithContentLanguage(contentLanguage *string) *GetVmsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get vms params
func (o *GetVmsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get vms params
func (o *GetVmsParams) WithRequestBody(requestBody *models.GetVmsRequestBody) *GetVmsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get vms params
func (o *GetVmsParams) SetRequestBody(requestBody *models.GetVmsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVmsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
