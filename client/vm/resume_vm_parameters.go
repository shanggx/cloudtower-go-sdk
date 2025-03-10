// Code generated by go-swagger; DO NOT EDIT.

package vm

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

// NewResumeVMParams creates a new ResumeVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewResumeVMParams() *ResumeVMParams {
	return &ResumeVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewResumeVMParamsWithTimeout creates a new ResumeVMParams object
// with the ability to set a timeout on a request.
func NewResumeVMParamsWithTimeout(timeout time.Duration) *ResumeVMParams {
	return &ResumeVMParams{
		timeout: timeout,
	}
}

// NewResumeVMParamsWithContext creates a new ResumeVMParams object
// with the ability to set a context for a request.
func NewResumeVMParamsWithContext(ctx context.Context) *ResumeVMParams {
	return &ResumeVMParams{
		Context: ctx,
	}
}

// NewResumeVMParamsWithHTTPClient creates a new ResumeVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewResumeVMParamsWithHTTPClient(client *http.Client) *ResumeVMParams {
	return &ResumeVMParams{
		HTTPClient: client,
	}
}

/* ResumeVMParams contains all the parameters to send to the API endpoint
   for the resume Vm operation.

   Typically these are written to a http.Request.
*/
type ResumeVMParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.VMOperateParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the resume Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeVMParams) WithDefaults() *ResumeVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the resume Vm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ResumeVMParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := ResumeVMParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the resume Vm params
func (o *ResumeVMParams) WithTimeout(timeout time.Duration) *ResumeVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the resume Vm params
func (o *ResumeVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the resume Vm params
func (o *ResumeVMParams) WithContext(ctx context.Context) *ResumeVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the resume Vm params
func (o *ResumeVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the resume Vm params
func (o *ResumeVMParams) WithHTTPClient(client *http.Client) *ResumeVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the resume Vm params
func (o *ResumeVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the resume Vm params
func (o *ResumeVMParams) WithContentLanguage(contentLanguage *string) *ResumeVMParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the resume Vm params
func (o *ResumeVMParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the resume Vm params
func (o *ResumeVMParams) WithRequestBody(requestBody *models.VMOperateParams) *ResumeVMParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the resume Vm params
func (o *ResumeVMParams) SetRequestBody(requestBody *models.VMOperateParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *ResumeVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
