// Code generated by go-swagger; DO NOT EDIT.

package host

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

// NewGetHostsParams creates a new GetHostsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostsParams() *GetHostsParams {
	return &GetHostsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostsParamsWithTimeout creates a new GetHostsParams object
// with the ability to set a timeout on a request.
func NewGetHostsParamsWithTimeout(timeout time.Duration) *GetHostsParams {
	return &GetHostsParams{
		timeout: timeout,
	}
}

// NewGetHostsParamsWithContext creates a new GetHostsParams object
// with the ability to set a context for a request.
func NewGetHostsParamsWithContext(ctx context.Context) *GetHostsParams {
	return &GetHostsParams{
		Context: ctx,
	}
}

// NewGetHostsParamsWithHTTPClient creates a new GetHostsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostsParamsWithHTTPClient(client *http.Client) *GetHostsParams {
	return &GetHostsParams{
		HTTPClient: client,
	}
}

/* GetHostsParams contains all the parameters to send to the API endpoint
   for the get hosts operation.

   Typically these are written to a http.Request.
*/
type GetHostsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetHostsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostsParams) WithDefaults() *GetHostsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get hosts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetHostsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get hosts params
func (o *GetHostsParams) WithTimeout(timeout time.Duration) *GetHostsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hosts params
func (o *GetHostsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hosts params
func (o *GetHostsParams) WithContext(ctx context.Context) *GetHostsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hosts params
func (o *GetHostsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hosts params
func (o *GetHostsParams) WithHTTPClient(client *http.Client) *GetHostsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hosts params
func (o *GetHostsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get hosts params
func (o *GetHostsParams) WithContentLanguage(contentLanguage *string) *GetHostsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get hosts params
func (o *GetHostsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get hosts params
func (o *GetHostsParams) WithRequestBody(requestBody *models.GetHostsRequestBody) *GetHostsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get hosts params
func (o *GetHostsParams) SetRequestBody(requestBody *models.GetHostsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
