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

// VMEntityFilterResult Vm entity filter result
//
// swagger:model VmEntityFilterResult
type VMEntityFilterResult struct {

	// entity filter
	// Required: true
	EntityFilter *NestedEntityFilter `json:"entityFilter"`

	// id
	// Required: true
	ID *string `json:"id"`

	// result
	// Required: true
	Result []float64 `json:"result"`

	// vm
	// Required: true
	VM *NestedVM `json:"vm"`
}

// Validate validates this Vm entity filter result
func (m *VMEntityFilterResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVM(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMEntityFilterResult) validateEntityFilter(formats strfmt.Registry) error {

	if err := validate.Required("entityFilter", "body", m.EntityFilter); err != nil {
		return err
	}

	if m.EntityFilter != nil {
		if err := m.EntityFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityFilter")
			}
			return err
		}
	}

	return nil
}

func (m *VMEntityFilterResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VMEntityFilterResult) validateResult(formats strfmt.Registry) error {

	if err := validate.Required("result", "body", m.Result); err != nil {
		return err
	}

	return nil
}

func (m *VMEntityFilterResult) validateVM(formats strfmt.Registry) error {

	if err := validate.Required("vm", "body", m.VM); err != nil {
		return err
	}

	if m.VM != nil {
		if err := m.VM.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Vm entity filter result based on the context it is used
func (m *VMEntityFilterResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntityFilter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVM(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMEntityFilterResult) contextValidateEntityFilter(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityFilter != nil {
		if err := m.EntityFilter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityFilter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityFilter")
			}
			return err
		}
	}

	return nil
}

func (m *VMEntityFilterResult) contextValidateVM(ctx context.Context, formats strfmt.Registry) error {

	if m.VM != nil {
		if err := m.VM.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMEntityFilterResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMEntityFilterResult) UnmarshalBinary(b []byte) error {
	var res VMEntityFilterResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}