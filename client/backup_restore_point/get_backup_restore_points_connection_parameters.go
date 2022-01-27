// Code generated by go-swagger; DO NOT EDIT.

package backup_restore_point

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

// NewGetBackupRestorePointsConnectionParams creates a new GetBackupRestorePointsConnectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetBackupRestorePointsConnectionParams() *GetBackupRestorePointsConnectionParams {
	return &GetBackupRestorePointsConnectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetBackupRestorePointsConnectionParamsWithTimeout creates a new GetBackupRestorePointsConnectionParams object
// with the ability to set a timeout on a request.
func NewGetBackupRestorePointsConnectionParamsWithTimeout(timeout time.Duration) *GetBackupRestorePointsConnectionParams {
	return &GetBackupRestorePointsConnectionParams{
		timeout: timeout,
	}
}

// NewGetBackupRestorePointsConnectionParamsWithContext creates a new GetBackupRestorePointsConnectionParams object
// with the ability to set a context for a request.
func NewGetBackupRestorePointsConnectionParamsWithContext(ctx context.Context) *GetBackupRestorePointsConnectionParams {
	return &GetBackupRestorePointsConnectionParams{
		Context: ctx,
	}
}

// NewGetBackupRestorePointsConnectionParamsWithHTTPClient creates a new GetBackupRestorePointsConnectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetBackupRestorePointsConnectionParamsWithHTTPClient(client *http.Client) *GetBackupRestorePointsConnectionParams {
	return &GetBackupRestorePointsConnectionParams{
		HTTPClient: client,
	}
}

/* GetBackupRestorePointsConnectionParams contains all the parameters to send to the API endpoint
   for the get backup restore points connection operation.

   Typically these are written to a http.Request.
*/
type GetBackupRestorePointsConnectionParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetBackupRestorePointsConnectionRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get backup restore points connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupRestorePointsConnectionParams) WithDefaults() *GetBackupRestorePointsConnectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get backup restore points connection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetBackupRestorePointsConnectionParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetBackupRestorePointsConnectionParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) WithTimeout(timeout time.Duration) *GetBackupRestorePointsConnectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) WithContext(ctx context.Context) *GetBackupRestorePointsConnectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) WithHTTPClient(client *http.Client) *GetBackupRestorePointsConnectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) WithContentLanguage(contentLanguage *string) *GetBackupRestorePointsConnectionParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) WithRequestBody(requestBody *models.GetBackupRestorePointsConnectionRequestBody) *GetBackupRestorePointsConnectionParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get backup restore points connection params
func (o *GetBackupRestorePointsConnectionParams) SetRequestBody(requestBody *models.GetBackupRestorePointsConnectionRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetBackupRestorePointsConnectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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