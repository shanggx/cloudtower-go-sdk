// Code generated by go-swagger; DO NOT EDIT.

package cluster_topo

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

// NewGetClusterTopoesParams creates a new GetClusterTopoesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetClusterTopoesParams() *GetClusterTopoesParams {
	return &GetClusterTopoesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterTopoesParamsWithTimeout creates a new GetClusterTopoesParams object
// with the ability to set a timeout on a request.
func NewGetClusterTopoesParamsWithTimeout(timeout time.Duration) *GetClusterTopoesParams {
	return &GetClusterTopoesParams{
		timeout: timeout,
	}
}

// NewGetClusterTopoesParamsWithContext creates a new GetClusterTopoesParams object
// with the ability to set a context for a request.
func NewGetClusterTopoesParamsWithContext(ctx context.Context) *GetClusterTopoesParams {
	return &GetClusterTopoesParams{
		Context: ctx,
	}
}

// NewGetClusterTopoesParamsWithHTTPClient creates a new GetClusterTopoesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetClusterTopoesParamsWithHTTPClient(client *http.Client) *GetClusterTopoesParams {
	return &GetClusterTopoesParams{
		HTTPClient: client,
	}
}

/* GetClusterTopoesParams contains all the parameters to send to the API endpoint
   for the get cluster topoes operation.

   Typically these are written to a http.Request.
*/
type GetClusterTopoesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetClusterTopoesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get cluster topoes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTopoesParams) WithDefaults() *GetClusterTopoesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get cluster topoes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetClusterTopoesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetClusterTopoesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get cluster topoes params
func (o *GetClusterTopoesParams) WithTimeout(timeout time.Duration) *GetClusterTopoesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster topoes params
func (o *GetClusterTopoesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster topoes params
func (o *GetClusterTopoesParams) WithContext(ctx context.Context) *GetClusterTopoesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster topoes params
func (o *GetClusterTopoesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster topoes params
func (o *GetClusterTopoesParams) WithHTTPClient(client *http.Client) *GetClusterTopoesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster topoes params
func (o *GetClusterTopoesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get cluster topoes params
func (o *GetClusterTopoesParams) WithContentLanguage(contentLanguage *string) *GetClusterTopoesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get cluster topoes params
func (o *GetClusterTopoesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get cluster topoes params
func (o *GetClusterTopoesParams) WithRequestBody(requestBody *models.GetClusterTopoesRequestBody) *GetClusterTopoesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get cluster topoes params
func (o *GetClusterTopoesParams) SetRequestBody(requestBody *models.GetClusterTopoesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterTopoesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
