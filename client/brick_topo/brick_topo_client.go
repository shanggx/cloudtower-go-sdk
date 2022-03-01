// Code generated by go-swagger; DO NOT EDIT.

package brick_topo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new brick topo API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for brick topo API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBrickTopo(params *CreateBrickTopoParams, opts ...ClientOption) (*CreateBrickTopoOK, error)

	DeleteBrickTopo(params *DeleteBrickTopoParams, opts ...ClientOption) (*DeleteBrickTopoOK, error)

	GetBrickTopoes(params *GetBrickTopoesParams, opts ...ClientOption) (*GetBrickTopoesOK, error)

	GetBrickTopoesConnection(params *GetBrickTopoesConnectionParams, opts ...ClientOption) (*GetBrickTopoesConnectionOK, error)

	MoveBrickTopo(params *MoveBrickTopoParams, opts ...ClientOption) (*MoveBrickTopoOK, error)

	UpdateBrickTopo(params *UpdateBrickTopoParams, opts ...ClientOption) (*UpdateBrickTopoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBrickTopo create brick topo API
*/
func (a *Client) CreateBrickTopo(params *CreateBrickTopoParams, opts ...ClientOption) (*CreateBrickTopoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBrickTopoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateBrickTopo",
		Method:             "POST",
		PathPattern:        "/create-brick-topo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBrickTopoReader{formats: a.formats},
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
	success, ok := result.(*CreateBrickTopoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for CreateBrickTopo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBrickTopo delete brick topo API
*/
func (a *Client) DeleteBrickTopo(params *DeleteBrickTopoParams, opts ...ClientOption) (*DeleteBrickTopoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBrickTopoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteBrickTopo",
		Method:             "POST",
		PathPattern:        "/delete-brick-topo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBrickTopoReader{formats: a.formats},
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
	success, ok := result.(*DeleteBrickTopoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteBrickTopo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBrickTopoes get brick topoes API
*/
func (a *Client) GetBrickTopoes(params *GetBrickTopoesParams, opts ...ClientOption) (*GetBrickTopoesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBrickTopoesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBrickTopoes",
		Method:             "POST",
		PathPattern:        "/get-brick-topoes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBrickTopoesReader{formats: a.formats},
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
	success, ok := result.(*GetBrickTopoesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBrickTopoes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBrickTopoesConnection get brick topoes connection API
*/
func (a *Client) GetBrickTopoesConnection(params *GetBrickTopoesConnectionParams, opts ...ClientOption) (*GetBrickTopoesConnectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBrickTopoesConnectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBrickTopoesConnection",
		Method:             "POST",
		PathPattern:        "/get-brick-topoes-connection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBrickTopoesConnectionReader{formats: a.formats},
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
	success, ok := result.(*GetBrickTopoesConnectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBrickTopoesConnection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MoveBrickTopo move brick topo API
*/
func (a *Client) MoveBrickTopo(params *MoveBrickTopoParams, opts ...ClientOption) (*MoveBrickTopoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMoveBrickTopoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "MoveBrickTopo",
		Method:             "POST",
		PathPattern:        "/move-brick-topo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MoveBrickTopoReader{formats: a.formats},
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
	success, ok := result.(*MoveBrickTopoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for MoveBrickTopo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBrickTopo update brick topo API
*/
func (a *Client) UpdateBrickTopo(params *UpdateBrickTopoParams, opts ...ClientOption) (*UpdateBrickTopoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBrickTopoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateBrickTopo",
		Method:             "POST",
		PathPattern:        "/update-brick-topo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBrickTopoReader{formats: a.formats},
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
	success, ok := result.(*UpdateBrickTopoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for UpdateBrickTopo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
