// Code generated by go-swagger; DO NOT EDIT.

package user_audit_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user audit log API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user audit log API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetUserAuditLogs(params *GetUserAuditLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAuditLogsOK, error)

	GetUserAuditLogsConnection(params *GetUserAuditLogsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAuditLogsConnectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetUserAuditLogs get user audit logs API
*/
func (a *Client) GetUserAuditLogs(params *GetUserAuditLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAuditLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserAuditLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUserAuditLogs",
		Method:             "POST",
		PathPattern:        "/get-user-audit-logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserAuditLogsReader{formats: a.formats},
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
	success, ok := result.(*GetUserAuditLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUserAuditLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetUserAuditLogsConnection get user audit logs connection API
*/
func (a *Client) GetUserAuditLogsConnection(params *GetUserAuditLogsConnectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetUserAuditLogsConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserAuditLogsConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetUserAuditLogsConnection",
		Method:             "POST",
		PathPattern:        "/get-user-audit-logs-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserAuditLogsConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetUserAuditLogsConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetUserAuditLogsConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}