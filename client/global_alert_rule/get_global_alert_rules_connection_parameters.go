// Code generated by go-swagger; DO NOT EDIT.

package global_alert_rule

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

// NewGetGlobalAlertRulesConnectionParams creates a new GetGlobalAlertRulesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGlobalAlertRulesConnectionParams() *GetGlobalAlertRulesConnectionParams {
	return &GetGlobalAlertRulesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGlobalAlertRulesConnectionParamsWithTimeout creates a new GetGlobalAlertRulesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetGlobalAlertRulesConnectionParamsWithTimeout(timeout time.Duration) *GetGlobalAlertRulesConnectionParams {
	return &GetGlobalAlertRulesConnectionParams{
		timeout: timeout,
	}
}

// NewGetGlobalAlertRulesConnectionParamsWithContext creates a new GetGlobalAlertRulesConnectionParams object
// with the ability to set a context for a request.
func NewGetGlobalAlertRulesConnectionParamsWithContext(ctx context.Context) *GetGlobalAlertRulesConnectionParams {
	return &GetGlobalAlertRulesConnectionParams{
		Context: ctx,
	}
}

// NewGetGlobalAlertRulesConnectionParamsWithHTTPClient creates a new GetGlobalAlertRulesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGlobalAlertRulesConnectionParamsWithHTTPClient(client *http.Client) *GetGlobalAlertRulesConnectionParams {
	return &GetGlobalAlertRulesConnectionParams{
		HTTPClient: client,
	}
}

/* GetGlobalAlertRulesConnectionParams contains all the parameters to send to the API endpoint
   for the get global alert rules connection operation.

   Typically these are written to a http.Request.
*/
type GetGlobalAlertRulesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetGlobalAlertRulesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get global alert rules connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalAlertRulesConnectionParams) WithDefaults() *GetGlobalAlertRulesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get global alert rules connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGlobalAlertRulesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetGlobalAlertRulesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) WithTimeout(timeout time.Duration) *GetGlobalAlertRulesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) WithContext(ctx context.Context) *GetGlobalAlertRulesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) WithHTTPClient(client *http.Client) *GetGlobalAlertRulesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) WithContentLanguage(contentLanguage *string) *GetGlobalAlertRulesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) WithRequestBody(requestBody *models.GetGlobalAlertRulesConnectionRequestBody) *GetGlobalAlertRulesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get global alert rules connection params
func (o *GetGlobalAlertRulesConnectionParams) SetRequestBody(requestBody *models.GetGlobalAlertRulesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetGlobalAlertRulesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
