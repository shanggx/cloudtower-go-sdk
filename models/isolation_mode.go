// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IsolationMode isolation mode
//
// swagger:model IsolationMode
type IsolationMode string

func NewIsolationMode(value IsolationMode) *IsolationMode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IsolationMode.
func (m IsolationMode) Pointer() *IsolationMode {
	return &m
}

const (

	// IsolationModeALL captures enum value "ALL"
	IsolationModeALL IsolationMode = "ALL"

	// IsolationModePARTIAL captures enum value "PARTIAL"
	IsolationModePARTIAL IsolationMode = "PARTIAL"
)

// for schema
var isolationModeEnum []interface{}

func init() {
	var res []IsolationMode
	if err := json.Unmarshal([]byte(`["ALL","PARTIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		isolationModeEnum = append(isolationModeEnum, v)
	}
}

func (m IsolationMode) validateIsolationModeEnum(path, location string, value IsolationMode) error {
	if err := validate.EnumCase(path, location, value, isolationModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this isolation mode
func (m IsolationMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIsolationModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this isolation mode based on context it is used
func (m IsolationMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}