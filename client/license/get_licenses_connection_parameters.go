// Code generated by go-swagger; DO NOT EDIT.

package license

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

// NewGetLicensesConnectionParams creates a new GetLicensesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLicensesConnectionParams() *GetLicensesConnectionParams {
	return &GetLicensesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLicensesConnectionParamsWithTimeout creates a new GetLicensesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetLicensesConnectionParamsWithTimeout(timeout time.Duration) *GetLicensesConnectionParams {
	return &GetLicensesConnectionParams{
		timeout: timeout,
	}
}

// NewGetLicensesConnectionParamsWithContext creates a new GetLicensesConnectionParams object
// with the ability to set a context for a request.
func NewGetLicensesConnectionParamsWithContext(ctx context.Context) *GetLicensesConnectionParams {
	return &GetLicensesConnectionParams{
		Context: ctx,
	}
}

// NewGetLicensesConnectionParamsWithHTTPClient creates a new GetLicensesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLicensesConnectionParamsWithHTTPClient(client *http.Client) *GetLicensesConnectionParams {
	return &GetLicensesConnectionParams{
		HTTPClient: client,
	}
}

/* GetLicensesConnectionParams contains all the parameters to send to the API endpoint
   for the get licenses connection operation.

   Typically these are written to a http.Request.
*/
type GetLicensesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetLicensesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get licenses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicensesConnectionParams) WithDefaults() *GetLicensesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get licenses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLicensesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetLicensesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get licenses connection params
func (o *GetLicensesConnectionParams) WithTimeout(timeout time.Duration) *GetLicensesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get licenses connection params
func (o *GetLicensesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get licenses connection params
func (o *GetLicensesConnectionParams) WithContext(ctx context.Context) *GetLicensesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get licenses connection params
func (o *GetLicensesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get licenses connection params
func (o *GetLicensesConnectionParams) WithHTTPClient(client *http.Client) *GetLicensesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get licenses connection params
func (o *GetLicensesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get licenses connection params
func (o *GetLicensesConnectionParams) WithContentLanguage(contentLanguage *string) *GetLicensesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get licenses connection params
func (o *GetLicensesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get licenses connection params
func (o *GetLicensesConnectionParams) WithRequestBody(requestBody *models.GetLicensesConnectionRequestBody) *GetLicensesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get licenses connection params
func (o *GetLicensesConnectionParams) SetRequestBody(requestBody *models.GetLicensesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetLicensesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
