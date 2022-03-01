// Code generated by go-swagger; DO NOT EDIT.

package cluster_upgrade_history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cluster upgrade history API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cluster upgrade history API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetClusterUpgradeHistories(params *GetClusterUpgradeHistoriesParams, opts ...ClientOption) (*GetClusterUpgradeHistoriesOK, error)

	GetClusterUpgradeHistoriesConnection(params *GetClusterUpgradeHistoriesConnectionParams, opts ...ClientOption) (*GetClusterUpgradeHistoriesConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetClusterUpgradeHistories get cluster upgrade histories API
*/
func (a *Client) GetClusterUpgradeHistories(params *GetClusterUpgradeHistoriesParams, opts ...ClientOption) (*GetClusterUpgradeHistoriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterUpgradeHistoriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClusterUpgradeHistories",
		Method:             "POST",
		PathPattern:        "/get-cluster-upgrade-histories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterUpgradeHistoriesReader{formats: a.formats},
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
	success, ok := result.(*GetClusterUpgradeHistoriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetClusterUpgradeHistories: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetClusterUpgradeHistoriesConnection get cluster upgrade histories connection API
*/
func (a *Client) GetClusterUpgradeHistoriesConnection(params *GetClusterUpgradeHistoriesConnectionParams, opts ...ClientOption) (*GetClusterUpgradeHistoriesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterUpgradeHistoriesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetClusterUpgradeHistoriesConnection",
		Method:             "POST",
		PathPattern:        "/get-cluster-upgrade-histories-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetClusterUpgradeHistoriesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetClusterUpgradeHistoriesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetClusterUpgradeHistoriesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
