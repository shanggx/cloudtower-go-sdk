// Code generated by go-swagger; DO NOT EDIT.

package elf_image

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

// NewUpdateElfImageParams creates a new UpdateElfImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateElfImageParams() *UpdateElfImageParams {
	return &UpdateElfImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateElfImageParamsWithTimeout creates a new UpdateElfImageParams object
// with the ability to set a timeout on a request.
func NewUpdateElfImageParamsWithTimeout(timeout time.Duration) *UpdateElfImageParams {
	return &UpdateElfImageParams{
		timeout: timeout,
	}
}

// NewUpdateElfImageParamsWithContext creates a new UpdateElfImageParams object
// with the ability to set a context for a request.
func NewUpdateElfImageParamsWithContext(ctx context.Context) *UpdateElfImageParams {
	return &UpdateElfImageParams{
		Context: ctx,
	}
}

// NewUpdateElfImageParamsWithHTTPClient creates a new UpdateElfImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateElfImageParamsWithHTTPClient(client *http.Client) *UpdateElfImageParams {
	return &UpdateElfImageParams{
		HTTPClient: client,
	}
}

/* UpdateElfImageParams contains all the parameters to send to the API endpoint
   for the update elf image operation.

   Typically these are written to a http.Request.
*/
type UpdateElfImageParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.ElfImageUpdationParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update elf image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateElfImageParams) WithDefaults() *UpdateElfImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update elf image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateElfImageParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := UpdateElfImageParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the update elf image params
func (o *UpdateElfImageParams) WithTimeout(timeout time.Duration) *UpdateElfImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update elf image params
func (o *UpdateElfImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update elf image params
func (o *UpdateElfImageParams) WithContext(ctx context.Context) *UpdateElfImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update elf image params
func (o *UpdateElfImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update elf image params
func (o *UpdateElfImageParams) WithHTTPClient(client *http.Client) *UpdateElfImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update elf image params
func (o *UpdateElfImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the update elf image params
func (o *UpdateElfImageParams) WithContentLanguage(contentLanguage *string) *UpdateElfImageParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the update elf image params
func (o *UpdateElfImageParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the update elf image params
func (o *UpdateElfImageParams) WithRequestBody(requestBody *models.ElfImageUpdationParams) *UpdateElfImageParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the update elf image params
func (o *UpdateElfImageParams) SetRequestBody(requestBody *models.ElfImageUpdationParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateElfImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
