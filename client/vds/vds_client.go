// Code generated by go-swagger; DO NOT EDIT.

package vds

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new vds API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vds API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVds(params *CreateVdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVdsOK, error)

	CreateVdsWithAccessVlan(params *CreateVdsWithAccessVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVdsWithAccessVlanOK, error)

	CreateVdsWithMigrateVlan(params *CreateVdsWithMigrateVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVdsWithMigrateVlanOK, error)

	DeleteVds(params *DeleteVdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVdsOK, error)

	GetVdses(params *GetVdsesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVdsesOK, error)

	GetVdsesConnection(params *GetVdsesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVdsesConnectionOK, error)

	UpdateVds(params *UpdateVdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVdsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVds create vds API
*/
func (a *Client) CreateVds(params *CreateVdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVds",
		Method:             "POST",
		PathPattern:        "/create-vds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVdsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateVds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateVdsWithAccessVlan create vds with access vlan API
*/
func (a *Client) CreateVdsWithAccessVlan(params *CreateVdsWithAccessVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVdsWithAccessVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVdsWithAccessVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVdsWithAccessVlan",
		Method:             "POST",
		PathPattern:        "/create-vds-with-access-vlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVdsWithAccessVlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVdsWithAccessVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateVdsWithAccessVlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CreateVdsWithMigrateVlan create vds with migrate vlan API
*/
func (a *Client) CreateVdsWithMigrateVlan(params *CreateVdsWithMigrateVlanParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVdsWithMigrateVlanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVdsWithMigrateVlanParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVdsWithMigrateVlan",
		Method:             "POST",
		PathPattern:        "/create-vds-with-migrate-vlan",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVdsWithMigrateVlanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVdsWithMigrateVlanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateVdsWithMigrateVlan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVds delete vds API
*/
func (a *Client) DeleteVds(params *DeleteVdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteVds",
		Method:             "POST",
		PathPattern:        "/delete-vds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVdsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteVds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVdses get vdses API
*/
func (a *Client) GetVdses(params *GetVdsesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVdsesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVdsesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVdses",
		Method:             "POST",
		PathPattern:        "/get-vdses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVdsesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVdsesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVdses: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetVdsesConnection get vdses connection API
*/
func (a *Client) GetVdsesConnection(params *GetVdsesConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetVdsesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVdsesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetVdsesConnection",
		Method:             "POST",
		PathPattern:        "/get-vdses-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVdsesConnectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVdsesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetVdsesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVds update vds API
*/
func (a *Client) UpdateVds(params *UpdateVdsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVdsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVdsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateVds",
		Method:             "POST",
		PathPattern:        "/update-vds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVdsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateVdsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateVds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}