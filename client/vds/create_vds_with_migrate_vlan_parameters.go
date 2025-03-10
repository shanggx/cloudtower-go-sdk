// Code generated by go-swagger; DO NOT EDIT.

package vds

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

// NewCreateVdsWithMigrateVlanParams creates a new CreateVdsWithMigrateVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateVdsWithMigrateVlanParams() *CreateVdsWithMigrateVlanParams {
	return &CreateVdsWithMigrateVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateVdsWithMigrateVlanParamsWithTimeout creates a new CreateVdsWithMigrateVlanParams object
// with the ability to set a timeout on a request.
func NewCreateVdsWithMigrateVlanParamsWithTimeout(timeout time.Duration) *CreateVdsWithMigrateVlanParams {
	return &CreateVdsWithMigrateVlanParams{
		timeout: timeout,
	}
}

// NewCreateVdsWithMigrateVlanParamsWithContext creates a new CreateVdsWithMigrateVlanParams object
// with the ability to set a context for a request.
func NewCreateVdsWithMigrateVlanParamsWithContext(ctx context.Context) *CreateVdsWithMigrateVlanParams {
	return &CreateVdsWithMigrateVlanParams{
		Context: ctx,
	}
}

// NewCreateVdsWithMigrateVlanParamsWithHTTPClient creates a new CreateVdsWithMigrateVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateVdsWithMigrateVlanParamsWithHTTPClient(client *http.Client) *CreateVdsWithMigrateVlanParams {
	return &CreateVdsWithMigrateVlanParams{
		HTTPClient: client,
	}
}

/* CreateVdsWithMigrateVlanParams contains all the parameters to send to the API endpoint
   for the create vds with migrate vlan operation.

   Typically these are written to a http.Request.
*/
type CreateVdsWithMigrateVlanParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody []*models.VdsCreationWithMigrateVlanParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create vds with migrate vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVdsWithMigrateVlanParams) WithDefaults() *CreateVdsWithMigrateVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create vds with migrate vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateVdsWithMigrateVlanParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := CreateVdsWithMigrateVlanParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) WithTimeout(timeout time.Duration) *CreateVdsWithMigrateVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) WithContext(ctx context.Context) *CreateVdsWithMigrateVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) WithHTTPClient(client *http.Client) *CreateVdsWithMigrateVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) WithContentLanguage(contentLanguage *string) *CreateVdsWithMigrateVlanParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) WithRequestBody(requestBody []*models.VdsCreationWithMigrateVlanParams) *CreateVdsWithMigrateVlanParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the create vds with migrate vlan params
func (o *CreateVdsWithMigrateVlanParams) SetRequestBody(requestBody []*models.VdsCreationWithMigrateVlanParams) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *CreateVdsWithMigrateVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
