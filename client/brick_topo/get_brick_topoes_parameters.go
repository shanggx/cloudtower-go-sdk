// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

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

// NewGetBrickTopoesParams creates a new GetBrickTopoesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBrickTopoesParams() *GetBrickTopoesParams {
	return &GetBrickTopoesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBrickTopoesParamsWithTimeout creates a new GetBrickTopoesParams object
// with the ability to set a timeout on a request.
func NewGetBrickTopoesParamsWithTimeout(timeout time.Duration) *GetBrickTopoesParams {
	return &GetBrickTopoesParams{
		timeout: timeout,
	}
}

// NewGetBrickTopoesParamsWithContext creates a new GetBrickTopoesParams object
// with the ability to set a context for a request.
func NewGetBrickTopoesParamsWithContext(ctx context.Context) *GetBrickTopoesParams {
	return &GetBrickTopoesParams{
		Context: ctx,
	}
}

// NewGetBrickTopoesParamsWithHTTPClient creates a new GetBrickTopoesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBrickTopoesParamsWithHTTPClient(client *http.Client) *GetBrickTopoesParams {
	return &GetBrickTopoesParams{
		HTTPClient: client,
	}
}

/* GetBrickTopoesParams contains all the parameters to send to the API endpoint
   for the get brick topoes operation.

   Typically these are written to a http.Request.
*/
type GetBrickTopoesParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetBrickTopoesRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get brick topoes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBrickTopoesParams) WithDefaults() *GetBrickTopoesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get brick topoes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBrickTopoesParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetBrickTopoesParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get brick topoes params
func (o *GetBrickTopoesParams) WithTimeout(timeout time.Duration) *GetBrickTopoesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get brick topoes params
func (o *GetBrickTopoesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get brick topoes params
func (o *GetBrickTopoesParams) WithContext(ctx context.Context) *GetBrickTopoesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get brick topoes params
func (o *GetBrickTopoesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get brick topoes params
func (o *GetBrickTopoesParams) WithHTTPClient(client *http.Client) *GetBrickTopoesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get brick topoes params
func (o *GetBrickTopoesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get brick topoes params
func (o *GetBrickTopoesParams) WithContentLanguage(contentLanguage *string) *GetBrickTopoesParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get brick topoes params
func (o *GetBrickTopoesParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get brick topoes params
func (o *GetBrickTopoesParams) WithRequestBody(requestBody *models.GetBrickTopoesRequestBody) *GetBrickTopoesParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get brick topoes params
func (o *GetBrickTopoesParams) SetRequestBody(requestBody *models.GetBrickTopoesRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetBrickTopoesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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