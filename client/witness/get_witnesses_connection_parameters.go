// Code generated by go-swagger; DO NOT EDIT.

package witness

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

// NewGetWitnessesConnectionParams creates a new GetWitnessesConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWitnessesConnectionParams() *GetWitnessesConnectionParams {
	return &GetWitnessesConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWitnessesConnectionParamsWithTimeout creates a new GetWitnessesConnectionParams object
// with the ability to set a timeout on a request.
func NewGetWitnessesConnectionParamsWithTimeout(timeout time.Duration) *GetWitnessesConnectionParams {
	return &GetWitnessesConnectionParams{
		timeout: timeout,
	}
}

// NewGetWitnessesConnectionParamsWithContext creates a new GetWitnessesConnectionParams object
// with the ability to set a context for a request.
func NewGetWitnessesConnectionParamsWithContext(ctx context.Context) *GetWitnessesConnectionParams {
	return &GetWitnessesConnectionParams{
		Context: ctx,
	}
}

// NewGetWitnessesConnectionParamsWithHTTPClient creates a new GetWitnessesConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWitnessesConnectionParamsWithHTTPClient(client *http.Client) *GetWitnessesConnectionParams {
	return &GetWitnessesConnectionParams{
		HTTPClient: client,
	}
}

/* GetWitnessesConnectionParams contains all the parameters to send to the API endpoint
   for the get witnesses connection operation.

   Typically these are written to a http.Request.
*/
type GetWitnessesConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetWitnessesConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get witnesses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWitnessesConnectionParams) WithDefaults() *GetWitnessesConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get witnesses connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWitnessesConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetWitnessesConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get witnesses connection params
func (o *GetWitnessesConnectionParams) WithTimeout(timeout time.Duration) *GetWitnessesConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get witnesses connection params
func (o *GetWitnessesConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get witnesses connection params
func (o *GetWitnessesConnectionParams) WithContext(ctx context.Context) *GetWitnessesConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get witnesses connection params
func (o *GetWitnessesConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get witnesses connection params
func (o *GetWitnessesConnectionParams) WithHTTPClient(client *http.Client) *GetWitnessesConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get witnesses connection params
func (o *GetWitnessesConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get witnesses connection params
func (o *GetWitnessesConnectionParams) WithContentLanguage(contentLanguage *string) *GetWitnessesConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get witnesses connection params
func (o *GetWitnessesConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get witnesses connection params
func (o *GetWitnessesConnectionParams) WithRequestBody(requestBody *models.GetWitnessesConnectionRequestBody) *GetWitnessesConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get witnesses connection params
func (o *GetWitnessesConnectionParams) SetRequestBody(requestBody *models.GetWitnessesConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetWitnessesConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
