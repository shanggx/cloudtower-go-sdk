// Code generated by go-swagger; DO NOT EDIT.

package graph

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new graph API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for graph API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateGraph(params *CreateGraphParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGraphOK, error)

	DeleteGraph(params *DeleteGraphParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGraphOK, error)

	GetGraphs(params *GetGraphsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGraphsOK, error)

	GetGraphsConnection(params *GetGraphsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGraphsConnectionOK, error)

	UpdateGraph(params *UpdateGraphParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGraphOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateGraph create graph API
*/
func (a *Client) CreateGraph(params *CreateGraphParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateGraphOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateGraphParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateGraph",
		Method:             "POST",
		PathPattern:        "/create-graph",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateGraphReader{formats: a.formats},
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
	success, ok := result.(*CreateGraphOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateGraph: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteGraph delete graph API
*/
func (a *Client) DeleteGraph(params *DeleteGraphParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteGraphOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGraphParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteGraph",
		Method:             "POST",
		PathPattern:        "/delete-graph",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteGraphReader{formats: a.formats},
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
	success, ok := result.(*DeleteGraphOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteGraph: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGraphs get graphs API
*/
func (a *Client) GetGraphs(params *GetGraphsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGraphsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGraphsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGraphs",
		Method:             "POST",
		PathPattern:        "/get-graphs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGraphsReader{formats: a.formats},
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
	success, ok := result.(*GetGraphsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGraphs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetGraphsConnection get graphs connection API
*/
func (a *Client) GetGraphsConnection(params *GetGraphsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetGraphsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetGraphsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetGraphsConnection",
		Method:             "POST",
		PathPattern:        "/get-graphs-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetGraphsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetGraphsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetGraphsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateGraph update graph API
*/
func (a *Client) UpdateGraph(params *UpdateGraphParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateGraphOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateGraphParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateGraph",
		Method:             "POST",
		PathPattern:        "/update-graph",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateGraphReader{formats: a.formats},
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
	success, ok := result.(*UpdateGraphOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateGraph: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}