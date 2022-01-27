// Code generated by go-swagger; DO NOT EDIT.

package deploy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new deploy API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deploy API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetDeploys(params *GetDeploysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploysOK, error)

	GetDeploysConnection(params *GetDeploysConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploysConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetDeploys get deploys API
*/
func (a *Client) GetDeploys(params *GetDeploysParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeploys",
		Method:             "POST",
		PathPattern:        "/get-deploys",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploysReader{formats: a.formats},
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
	success, ok := result.(*GetDeploysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeploys: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetDeploysConnection get deploys connection API
*/
func (a *Client) GetDeploysConnection(params *GetDeploysConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeploysConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeploysConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetDeploysConnection",
		Method:             "POST",
		PathPattern:        "/get-deploys-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetDeploysConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetDeploysConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetDeploysConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}