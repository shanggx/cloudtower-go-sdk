// Code generated by go-swagger; DO NOT EDIT.

package consistency_group_snapshot

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

// NewDeleteConsistencyGroupSnapshotParams creates a new DeleteConsistencyGroupSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteConsistencyGroupSnapshotParams() *DeleteConsistencyGroupSnapshotParams {
	return &DeleteConsistencyGroupSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteConsistencyGroupSnapshotParamsWithTimeout creates a new DeleteConsistencyGroupSnapshotParams object
// with the ability to set a timeout on a request.
func NewDeleteConsistencyGroupSnapshotParamsWithTimeout(timeout time.Duration) *DeleteConsistencyGroupSnapshotParams {
	return &DeleteConsistencyGroupSnapshotParams{
		timeout: timeout,
	}
}

// NewDeleteConsistencyGroupSnapshotParamsWithContext creates a new DeleteConsistencyGroupSnapshotParams object
// with the ability to set a context for a request.
func NewDeleteConsistencyGroupSnapshotParamsWithContext(ctx context.Context) *DeleteConsistencyGroupSnapshotParams {
	return &DeleteConsistencyGroupSnapshotParams{
		Context: ctx,
	}
}

// NewDeleteConsistencyGroupSnapshotParamsWithHTTPClient creates a new DeleteConsistencyGroupSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteConsistencyGroupSnapshotParamsWithHTTPClient(client *http.Client) *DeleteConsistencyGroupSnapshotParams {
	return &DeleteConsistencyGroupSnapshotParams{
		HTTPClient: client,
	}
}

/* DeleteConsistencyGroupSnapshotParams contains all the parameters to send to the API endpoint
   for the delete consistency group snapshot operation.

   Typically these are written to a http.Request.
*/
type DeleteConsistencyGroupSnapshotParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ConsistencyGroupSnapshotDeletionParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete consistency group snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConsistencyGroupSnapshotParams) WithDefaults() *DeleteConsistencyGroupSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete consistency group snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteConsistencyGroupSnapshotParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := DeleteConsistencyGroupSnapshotParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) WithTimeout(timeout time.Duration) *DeleteConsistencyGroupSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) WithContext(ctx context.Context) *DeleteConsistencyGroupSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) WithHTTPClient(client *http.Client) *DeleteConsistencyGroupSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) WithContentLanguage(contentLanguage *string) *DeleteConsistencyGroupSnapshotParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) WithRequestBody(requestBody *models.ConsistencyGroupSnapshotDeletionParams) *DeleteConsistencyGroupSnapshotParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the delete consistency group snapshot params
func (o *DeleteConsistencyGroupSnapshotParams) SetRequestBody(requestBody *models.ConsistencyGroupSnapshotDeletionParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteConsistencyGroupSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
