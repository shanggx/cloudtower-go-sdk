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

// MountDisksParams mount disks params
//
// swagger:model MountDisksParams
type MountDisksParams struct {

	// boot
	// Required: true
	Boot *int32 `json:"boot"`

	// bus
	// Required: true
	Bus *Bus `json:"bus"`

	// index
	// Required: true
	Index *int32 `json:"index"`

	// key
	Key *int32 `json:"key,omitempty"`

	// max bandwidth
	MaxBandwidth *float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *int32 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// vm volume id
	// Required: true
	VMVolumeID *string `json:"vm_volume_id"`
}

// Validate validates this mount disks params
func (m *MountDisksParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBoot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMVolumeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountDisksParams) validateBoot(formats strfmt.Registry) error {

	if err := validate.Required("boot", "body", m.Boot); err != nil {
		return err
	}

	return nil
}

func (m *MountDisksParams) validateBus(formats strfmt.Registry) error {

	if err := validate.Required("bus", "body", m.Bus); err != nil {
		return err
	}

	if err := validate.Required("bus", "body", m.Bus); err != nil {
		return err
	}

	if m.Bus != nil {
		if err := m.Bus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bus")
			}
			return err
		}
	}

	return nil
}

func (m *MountDisksParams) validateIndex(formats strfmt.Registry) error {

	if err := validate.Required("index", "body", m.Index); err != nil {
		return err
	}

	return nil
}

func (m *MountDisksParams) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicy) { // not required
		return nil
	}

	if err := m.MaxBandwidthPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_bandwidth_policy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("max_bandwidth_policy")
		}
		return err
	}

	return nil
}

func (m *MountDisksParams) validateMaxIopsPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicy) { // not required
		return nil
	}

	if err := m.MaxIopsPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_iops_policy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("max_iops_policy")
		}
		return err
	}

	return nil
}

func (m *MountDisksParams) validateVMVolumeID(formats strfmt.Registry) error {

	if err := validate.Required("vm_volume_id", "body", m.VMVolumeID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mount disks params based on the context it is used
func (m *MountDisksParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MountDisksParams) contextValidateBus(ctx context.Context, formats strfmt.Registry) error {

	if m.Bus != nil {
		if err := m.Bus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bus")
			}
			return err
		}
	}

	return nil
}

func (m *MountDisksParams) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaxBandwidthPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_bandwidth_policy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("max_bandwidth_policy")
		}
		return err
	}

	return nil
}

func (m *MountDisksParams) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

	if err := m.MaxIopsPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("max_iops_policy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("max_iops_policy")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MountDisksParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MountDisksParams) UnmarshalBinary(b []byte) error {
	var res MountDisksParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}