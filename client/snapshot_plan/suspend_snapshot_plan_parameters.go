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

	"github.com/Sczlog/cloudtower-go-sdk/models"
)

// NewSuspendSnapshotPlanParams creates a new SuspendSnapshotPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSuspendSnapshotPlanParams() *SuspendSnapshotPlanParams {
	return &SuspendSnapshotPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSuspendSnapshotPlanParamsWithTimeout creates a new SuspendSnapshotPlanParams object
// with the ability to set a timeout on a request.
func NewSuspendSnapshotPlanParamsWithTimeout(timeout time.Duration) *SuspendSnapshotPlanParams {
	return &SuspendSnapshotPlanParams{
		timeout: timeout,
	}
}

// NewSuspendSnapshotPlanParamsWithContext creates a new SuspendSnapshotPlanParams object
// with the ability to set a context for a request.
func NewSuspendSnapshotPlanParamsWithContext(ctx context.Context) *SuspendSnapshotPlanParams {
	return &SuspendSnapshotPlanParams{
		Context: ctx,
	}
}

// NewSuspendSnapshotPlanParamsWithHTTPClient creates a new SuspendSnapshotPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewSuspendSnapshotPlanParamsWithHTTPClient(client *http.Client) *SuspendSnapshotPlanParams {
	return &SuspendSnapshotPlanParams{
		HTTPClient: client,
	}
}

/* SuspendSnapshotPlanParams contains all the parameters to send to the API endpoint
   for the suspend snapshot plan operation.

   Typically these are written to a http.Request.
*/
type SuspendSnapshotPlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.SnapshotPlanSuspendedParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the suspend snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendSnapshotPlanParams) WithDefaults() *SuspendSnapshotPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the suspend snapshot plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SuspendSnapshotPlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := SuspendSnapshotPlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) WithTimeout(timeout time.Duration) *SuspendSnapshotPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) WithContext(ctx context.Context) *SuspendSnapshotPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) WithHTTPClient(client *http.Client) *SuspendSnapshotPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) WithContentLanguage(contentLanguage *string) *SuspendSnapshotPlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) WithRequestBody(requestBody *models.SnapshotPlanSuspendedParams) *SuspendSnapshotPlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the suspend snapshot plan params
func (o *SuspendSnapshotPlanParams) SetRequestBody(requestBody *models.SnapshotPlanSuspendedParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *SuspendSnapshotPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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