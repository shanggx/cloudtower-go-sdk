// Code generated by go-swagger; DO NOT EDIT.

package vm_placement_group

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

// NewGetVMPlacementGroupsParams creates a new GetVMPlacementGroupsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetVMPlacementGroupsParams() *GetVMPlacementGroupsParams {
	return &GetVMPlacementGroupsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetVMPlacementGroupsParamsWithTimeout creates a new GetVMPlacementGroupsParams object
// with the ability to set a timeout on a request.
func NewGetVMPlacementGroupsParamsWithTimeout(timeout time.Duration) *GetVMPlacementGroupsParams {
	return &GetVMPlacementGroupsParams{
		timeout: timeout,
	}
}

// NewGetVMPlacementGroupsParamsWithContext creates a new GetVMPlacementGroupsParams object
// with the ability to set a context for a request.
func NewGetVMPlacementGroupsParamsWithContext(ctx context.Context) *GetVMPlacementGroupsParams {
	return &GetVMPlacementGroupsParams{
		Context: ctx,
	}
}

// NewGetVMPlacementGroupsParamsWithHTTPClient creates a new GetVMPlacementGroupsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetVMPlacementGroupsParamsWithHTTPClient(client *http.Client) *GetVMPlacementGroupsParams {
	return &GetVMPlacementGroupsParams{
		HTTPClient: client,
	}
}

/* GetVMPlacementGroupsParams contains all the parameters to send to the API endpoint
   for the get Vm placement groups operation.

   Typically these are written to a http.Request.
*/
type GetVMPlacementGroupsParams struct {

	// ContentLanguage.
	//
	// Default: "en-US"
	ContentLanguage *string

	// RequestBody.
	RequestBody *models.GetVMPlacementGroupsRequestBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Vm placement groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMPlacementGroupsParams) WithDefaults() *GetVMPlacementGroupsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Vm placement groups params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetVMPlacementGroupsParams) SetDefaults() {
	var (
		contentLanguageDefault = string("en-US")
	)

	val := GetVMPlacementGroupsParams{
		ContentLanguage: &contentLanguageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) WithTimeout(timeout time.Duration) *GetVMPlacementGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) WithContext(ctx context.Context) *GetVMPlacementGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) WithHTTPClient(client *http.Client) *GetVMPlacementGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentLanguage adds the contentLanguage to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) WithContentLanguage(contentLanguage *string) *GetVMPlacementGroupsParams {
	o.SetContentLanguage(contentLanguage)
	return o
}

// SetContentLanguage adds the contentLanguage to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) SetContentLanguage(contentLanguage *string) {
	o.ContentLanguage = contentLanguage
}

// WithRequestBody adds the requestBody to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) WithRequestBody(requestBody *models.GetVMPlacementGroupsRequestBody) *GetVMPlacementGroupsParams {
	o.SetRequestBody(requestBody)
	return o
}

// SetRequestBody adds the requestBody to the get Vm placement groups params
func (o *GetVMPlacementGroupsParams) SetRequestBody(requestBody *models.GetVMPlacementGroupsRequestBody) {
	o.RequestBody = requestBody
}

// WriteToRequest writes these params to a swagger request
func (o *GetVMPlacementGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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