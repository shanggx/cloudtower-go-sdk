// Code generated by go-swagger; DO NOT EDIT.

package snapshot_group

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

// NewRollbackSnapshotGroupParams creates a new RollbackSnapshotGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRollbackSnapshotGroupParams() *RollbackSnapshotGroupParams {
	return &RollbackSnapshotGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackSnapshotGroupParamsWithTimeout creates a new RollbackSnapshotGroupParams object
// with the ability to set a timeout on a request.
func NewRollbackSnapshotGroupParamsWithTimeout(timeout time.Duration) *RollbackSnapshotGroupParams {
	return &RollbackSnapshotGroupParams{
		timeout: timeout,
	}
}

// NewRollbackSnapshotGroupParamsWithContext creates a new RollbackSnapshotGroupParams object
// with the ability to set a context for a request.
func NewRollbackSnapshotGroupParamsWithContext(ctx context.Context) *RollbackSnapshotGroupParams {
	return &RollbackSnapshotGroupParams{
		Context: ctx,
	}
}

// NewRollbackSnapshotGroupParamsWithHTTPClient creates a new RollbackSnapshotGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewRollbackSnapshotGroupParamsWithHTTPClient(client *http.Client) *RollbackSnapshotGroupParams {
	return &RollbackSnapshotGroupParams{
		HTTPClient: client,
	}
}

/* RollbackSnapshotGroupParams contains all the parameters to send to the API endpoint
   for the rollback snapshot group operation.

   Typically these are written to a http.Request.
*/
type RollbackSnapshotGroupParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SnapshotGroupRollbackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rollback snapshot group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackSnapshotGroupParams) WithDefaults() *RollbackSnapshotGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rollback snapshot group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackSnapshotGroupParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := RollbackSnapshotGroupParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) WithTimeout(timeout time.Duration) *RollbackSnapshotGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) WithContext(ctx context.Context) *RollbackSnapshotGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) WithHTTPClient(client *http.Client) *RollbackSnapshotGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) WithContentLanguage(contentLanguage *string) *RollbackSnapshotGroupParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) WithRequestBody(requestBody *models.SnapshotGroupRollbackParams) *RollbackSnapshotGroupParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the rollback snapshot group params
func (o *RollbackSnapshotGroupParams) SetRequestBody(requestBody *models.SnapshotGroupRollbackParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackSnapshotGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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