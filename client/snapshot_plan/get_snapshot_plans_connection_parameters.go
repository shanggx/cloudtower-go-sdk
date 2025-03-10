// Code generated by go-swagger; DO NOT EDIT.

package snapshot_plan

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

// NewGetSnapshotPlansConnectionParams creates a new GetSnapshotPlansConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSnapshotPlansConnectionParams() *GetSnapshotPlansConnectionParams {
	return &GetSnapshotPlansConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotPlansConnectionParamsWithTimeout creates a new GetSnapshotPlansConnectionParams object
// with the ability to set a timeout on a request.
func NewGetSnapshotPlansConnectionParamsWithTimeout(timeout time.Duration) *GetSnapshotPlansConnectionParams {
	return &GetSnapshotPlansConnectionParams{
		timeout: timeout,
	}
}

// NewGetSnapshotPlansConnectionParamsWithContext creates a new GetSnapshotPlansConnectionParams object
// with the ability to set a context for a request.
func NewGetSnapshotPlansConnectionParamsWithContext(ctx context.Context) *GetSnapshotPlansConnectionParams {
	return &GetSnapshotPlansConnectionParams{
		Context: ctx,
	}
}

// NewGetSnapshotPlansConnectionParamsWithHTTPClient creates a new GetSnapshotPlansConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSnapshotPlansConnectionParamsWithHTTPClient(client *http.Client) *GetSnapshotPlansConnectionParams {
	return &GetSnapshotPlansConnectionParams{
		HTTPClient: client,
	}
}

/* GetSnapshotPlansConnectionParams contains all the parameters to send to the API endpoint
   for the get snapshot plans connection operation.

   Typically these are written to a http.Request.
*/
type GetSnapshotPlansConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetSnapshotPlansConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get snapshot plans connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotPlansConnectionParams) WithDefaults() *GetSnapshotPlansConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get snapshot plans connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSnapshotPlansConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetSnapshotPlansConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) WithTimeout(timeout time.Duration) *GetSnapshotPlansConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) WithContext(ctx context.Context) *GetSnapshotPlansConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) WithHTTPClient(client *http.Client) *GetSnapshotPlansConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) WithContentLanguage(contentLanguage *string) *GetSnapshotPlansConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) WithRequestBody(requestBody *models.GetSnapshotPlansConnectionRequestBody) *GetSnapshotPlansConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get snapshot plans connection params
func (o *GetSnapshotPlansConnectionParams) SetRequestBody(requestBody *models.GetSnapshotPlansConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotPlansConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
