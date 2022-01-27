// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedResourceMeta nested resource meta
//
// swagger:model NestedResourceMeta
type NestedResourceMeta struct {

	// fields
	// Required: true
	Fields []string `json:"fields"`

	// filter
	// Required: true
	Filter interface{} `json:"filter"`

	// name
	// Required: true
	Name *string `json:"name"`

	// type
	// Required: true
	Type *ReportResourceInputEnum `json:"type"`
}

// Validate validates this nested resource meta
func (m *NestedResourceMeta) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedResourceMeta) validateFields(formats strfmt.Registry) error {

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	return nil
}

func (m *NestedResourceMeta) validateFilter(formats strfmt.Registry) error {

	if m.Filter == nil {
		return errors.Required("filter", "body", nil)
	}

	return nil
}

func (m *NestedResourceMeta) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NestedResourceMeta) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this nested resource meta based on the context it is used
func (m *NestedResourceMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedResourceMeta) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedResourceMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedResourceMeta) UnmarshalBinary(b []byte) error {
	var res NestedResourceMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}