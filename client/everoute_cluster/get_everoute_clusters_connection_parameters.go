// Code generated by go-swagger; DO NOT EDIT.

package everoute_cluster

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

// NewGetEverouteClustersConnectionParams creates a new GetEverouteClustersConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEverouteClustersConnectionParams() *GetEverouteClustersConnectionParams {
	return &GetEverouteClustersConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEverouteClustersConnectionParamsWithTimeout creates a new GetEverouteClustersConnectionParams object
// with the ability to set a timeout on a request.
func NewGetEverouteClustersConnectionParamsWithTimeout(timeout time.Duration) *GetEverouteClustersConnectionParams {
	return &GetEverouteClustersConnectionParams{
		timeout: timeout,
	}
}

// NewGetEverouteClustersConnectionParamsWithContext creates a new GetEverouteClustersConnectionParams object
// with the ability to set a context for a request.
func NewGetEverouteClustersConnectionParamsWithContext(ctx context.Context) *GetEverouteClustersConnectionParams {
	return &GetEverouteClustersConnectionParams{
		Context: ctx,
	}
}

// NewGetEverouteClustersConnectionParamsWithHTTPClient creates a new GetEverouteClustersConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEverouteClustersConnectionParamsWithHTTPClient(client *http.Client) *GetEverouteClustersConnectionParams {
	return &GetEverouteClustersConnectionParams{
		HTTPClient: client,
	}
}

/* GetEverouteClustersConnectionParams contains all the parameters to send to the API endpoint
   for the get everoute clusters connection operation.

   Typically these are written to a http.Request.
*/
type GetEverouteClustersConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetEverouteClustersConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get everoute clusters connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEverouteClustersConnectionParams) WithDefaults() *GetEverouteClustersConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get everoute clusters connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEverouteClustersConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetEverouteClustersConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) WithTimeout(timeout time.Duration) *GetEverouteClustersConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) WithContext(ctx context.Context) *GetEverouteClustersConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) WithHTTPClient(client *http.Client) *GetEverouteClustersConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) WithContentLanguage(contentLanguage *string) *GetEverouteClustersConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) WithRequestBody(requestBody *models.GetEverouteClustersConnectionRequestBody) *GetEverouteClustersConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get everoute clusters connection params
func (o *GetEverouteClustersConnectionParams) SetRequestBody(requestBody *models.GetEverouteClustersConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetEverouteClustersConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
