// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMRebuildParams Vm rebuild params
//
// swagger:model VmRebuildParams
type VMRebuildParams struct {

	// cluster id
	ClusterID *string `json:"cluster_id,omitempty"`

	// cpu cores
	CPUCores *int32 `json:"cpu_cores,omitempty"`

	// cpu sockets
	CPUSockets *int32 `json:"cpu_sockets,omitempty"`

	// description
	Description *string `json:"description,omitempty"`

	// firmware
	Firmware *VMFirmware `json:"firmware,omitempty"`

	// folder id
	FolderID *string `json:"folder_id,omitempty"`

	// guest os type
	GuestOsType *VMGuestsOperationSystem `json:"guest_os_type,omitempty"`

	// ha
	Ha *bool `json:"ha,omitempty"`

	// host id
	HostID *string `json:"host_id,omitempty"`

	// io policy
	IoPolicy *VMDiskIoPolicy `json:"io_policy,omitempty"`

	// max bandwidth
	MaxBandwidth *float64 `json:"max_bandwidth,omitempty"`

	// max bandwidth policy
	MaxBandwidthPolicy *VMDiskIoRestrictType `json:"max_bandwidth_policy,omitempty"`

	// max iops
	MaxIops *int32 `json:"max_iops,omitempty"`

	// max iops policy
	MaxIopsPolicy *VMDiskIoRestrictType `json:"max_iops_policy,omitempty"`

	// memory
	Memory *float64 `json:"memory,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// rebuild from snapshot id
	// Required: true
	RebuildFromSnapshotID *string `json:"rebuild_from_snapshot_id"`

	// status
	Status *VMStatus `json:"status,omitempty"`

	// vcpu
	Vcpu *int32 `json:"vcpu,omitempty"`

	// vm disks
	VMDisks *VMDiskParams `json:"vm_disks,omitempty"`

	// vm nics
	VMNics []*VMNicParams `json:"vm_nics,omitempty"`
}

// Validate validates this Vm rebuild params
func (m *VMRebuildParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxBandwidthPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaxIopsPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRebuildFromSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVMNics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMRebuildParams) validateFirmware(formats strfmt.Registry) error {
	if swag.IsZero(m.Firmware) { // not required
		return nil
	}

	if m.Firmware != nil {
		if err := m.Firmware.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) validateGuestOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.GuestOsType) { // not required
		return nil
	}

	if m.GuestOsType != nil {
		if err := m.GuestOsType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guest_os_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guest_os_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) validateIoPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.IoPolicy) { // not required
		return nil
	}

	if m.IoPolicy != nil {
		if err := m.IoPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxBandwidthPolicy) { // not required
		return nil
	}

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) validateMaxIopsPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MaxIopsPolicy) { // not required
		return nil
	}

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMRebuildParams) validateRebuildFromSnapshotID(formats strfmt.Registry) error {

	if err := validate.Required("rebuild_from_snapshot_id", "body", m.RebuildFromSnapshotID); err != nil {
		return err
	}

	return nil
}

func (m *VMRebuildParams) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) validateVMDisks(formats strfmt.Registry) error {
	if swag.IsZero(m.VMDisks) { // not required
		return nil
	}

	if m.VMDisks != nil {
		if err := m.VMDisks.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_disks")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) validateVMNics(formats strfmt.Registry) error {
	if swag.IsZero(m.VMNics) { // not required
		return nil
	}

	for i := 0; i < len(m.VMNics); i++ {
		if swag.IsZero(m.VMNics[i]) { // not required
			continue
		}

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Vm rebuild params based on the context it is used
func (m *VMRebuildParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFirmware(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGuestOsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxBandwidthPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMaxIopsPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMDisks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVMNics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMRebuildParams) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

	if m.Firmware != nil {
		if err := m.Firmware.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("firmware")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("firmware")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) contextValidateGuestOsType(ctx context.Context, formats strfmt.Registry) error {

	if m.GuestOsType != nil {
		if err := m.GuestOsType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("guest_os_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("guest_os_type")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.IoPolicy != nil {
		if err := m.IoPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("io_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("io_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxBandwidthPolicy != nil {
		if err := m.MaxBandwidthPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_bandwidth_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_bandwidth_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MaxIopsPolicy != nil {
		if err := m.MaxIopsPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("max_iops_policy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("max_iops_policy")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

	if m.VMDisks != nil {
		if err := m.VMDisks.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vm_disks")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vm_disks")
			}
			return err
		}
	}

	return nil
}

func (m *VMRebuildParams) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VMNics); i++ {

		if m.VMNics[i] != nil {
			if err := m.VMNics[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vm_nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMRebuildParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMRebuildParams) UnmarshalBinary(b []byte) error {
	var res VMRebuildParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
