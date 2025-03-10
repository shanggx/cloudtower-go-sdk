// Code generated by go-swagger; DO NOT EDIT.

package iscsi_lun

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

// NewRollbackIscsiLunFromSnapshotParams creates a new RollbackIscsiLunFromSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRollbackIscsiLunFromSnapshotParams() *RollbackIscsiLunFromSnapshotParams {
	return &RollbackIscsiLunFromSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackIscsiLunFromSnapshotParamsWithTimeout creates a new RollbackIscsiLunFromSnapshotParams object
// with the ability to set a timeout on a request.
func NewRollbackIscsiLunFromSnapshotParamsWithTimeout(timeout time.Duration) *RollbackIscsiLunFromSnapshotParams {
	return &RollbackIscsiLunFromSnapshotParams{
		timeout: timeout,
	}
}

// NewRollbackIscsiLunFromSnapshotParamsWithContext creates a new RollbackIscsiLunFromSnapshotParams object
// with the ability to set a context for a request.
func NewRollbackIscsiLunFromSnapshotParamsWithContext(ctx context.Context) *RollbackIscsiLunFromSnapshotParams {
	return &RollbackIscsiLunFromSnapshotParams{
		Context: ctx,
	}
}

// NewRollbackIscsiLunFromSnapshotParamsWithHTTPClient creates a new RollbackIscsiLunFromSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewRollbackIscsiLunFromSnapshotParamsWithHTTPClient(client *http.Client) *RollbackIscsiLunFromSnapshotParams {
	return &RollbackIscsiLunFromSnapshotParams{
		HTTPClient: client,
	}
}

/* RollbackIscsiLunFromSnapshotParams contains all the parameters to send to the API endpoint
   for the rollback iscsi lun from snapshot operation.

   Typically these are written to a http.Request.
*/
type RollbackIscsiLunFromSnapshotParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.IscsiLunRollbackParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rollback iscsi lun from snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackIscsiLunFromSnapshotParams) WithDefaults() *RollbackIscsiLunFromSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rollback iscsi lun from snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RollbackIscsiLunFromSnapshotParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := RollbackIscsiLunFromSnapshotParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) WithTimeout(timeout time.Duration) *RollbackIscsiLunFromSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) WithContext(ctx context.Context) *RollbackIscsiLunFromSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) WithHTTPClient(client *http.Client) *RollbackIscsiLunFromSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) WithContentLanguage(contentLanguage *string) *RollbackIscsiLunFromSnapshotParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) WithRequestBody(requestBody []*models.IscsiLunRollbackParams) *RollbackIscsiLunFromSnapshotParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the rollback iscsi lun from snapshot params
func (o *RollbackIscsiLunFromSnapshotParams) SetRequestBody(requestBody []*models.IscsiLunRollbackParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackIscsiLunFromSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
