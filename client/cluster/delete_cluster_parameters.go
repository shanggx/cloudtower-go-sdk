// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewDeleteClusterParams creates a new DeleteClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteClusterParams() *DeleteClusterParams {
	return &DeleteClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterParamsWithTimeout creates a new DeleteClusterParams object
// with the ability to set a timeout on a request.
func NewDeleteClusterParamsWithTimeout(timeout time.Duration) *DeleteClusterParams {
	return &DeleteClusterParams{
		timeout: timeout,
	}
}

// NewDeleteClusterParamsWithContext creates a new DeleteClusterParams object
// with the ability to set a context for a request.
func NewDeleteClusterParamsWithContext(ctx context.Context) *DeleteClusterParams {
	return &DeleteClusterParams{
		Context: ctx,
	}
}

// NewDeleteClusterParamsWithHTTPClient creates a new DeleteClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteClusterParamsWithHTTPClient(client *http.Client) *DeleteClusterParams {
	return &DeleteClusterParams{
		HTTPClient: client,
	}
}

/* DeleteClusterParams contains all the parameters to send to the API endpoint
   for the delete cluster operation.

   Typically these are written to a http.Request.
*/
type DeleteClusterParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ClusterDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterParams) WithDefaults() *DeleteClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteClusterParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteClusterParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) WithTimeout(timeout time.Duration) *DeleteClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster params
func (o *DeleteClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster params
func (o *DeleteClusterParams) WithContext(ctx context.Context) *DeleteClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster params
func (o *DeleteClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) WithHTTPClient(client *http.Client) *DeleteClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster params
func (o *DeleteClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete cluster params
func (o *DeleteClusterParams) WithContentLanguage(contentLanguage *string) *DeleteClusterParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete cluster params
func (o *DeleteClusterParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete cluster params
func (o *DeleteClusterParams) WithRequestBody(requestBody *models.ClusterDeletionParams) *DeleteClusterParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete cluster params
func (o *DeleteClusterParams) SetRequestBody(requestBody *models.ClusterDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
