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

// VMCreationParams Vm creation params
//
// swagger:model VmCreationParams
type VMCreationParams struct {

	// cluster id
	// Required: true
	ClusterID *string `json:"cluster_id"`

	// cpu cores
	// Required: true
	CPUCores *int32 `json:"cpu_cores"`

	// cpu sockets
	// Required: true
	CPUSockets *int32 `json:"cpu_sockets"`

	// description
	Description *string `json:"description,omitempty"`

	// firmware
	// Required: true
	Firmware *VMFirmware `json:"firmware"`

	// folder id
	FolderID *string `json:"folder_id,omitempty"`

	// guest os type
	GuestOsType *VMGuestsOperationSystem `json:"guest_os_type,omitempty"`

	// ha
	// Required: true
	Ha *bool `json:"ha"`

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
	// Required: true
	Memory *float64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`

	// status
	// Required: true
	Status *VMStatus `json:"status"`

	// vcpu
	// Required: true
	Vcpu *int32 `json:"vcpu"`

	// vm disks
	// Required: true
	VMDisks *VMDiskParams `json:"vm_disks"`

	// vm nics
	// Required: true
	VMNics []*VMNicParams `json:"vm_nics"`
}

// Validate validates this Vm creation params
func (m *VMCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUCores(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCPUSockets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirmware(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGuestOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHa(formats); err != nil {
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

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVcpu(formats); err != nil {
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

func (m *VMCreationParams) validateClusterID(formats strfmt.Registry) error {

	if err := validate.Required("cluster_id", "body", m.ClusterID); err != nil {
		return err
	}

	return nil
}

func (m *VMCreationParams) validateCPUCores(formats strfmt.Registry) error {

	if err := validate.Required("cpu_cores", "body", m.CPUCores); err != nil {
		return err
	}

	return nil
}

func (m *VMCreationParams) validateCPUSockets(formats strfmt.Registry) error {

	if err := validate.Required("cpu_sockets", "body", m.CPUSockets); err != nil {
		return err
	}

	return nil
}

func (m *VMCreationParams) validateFirmware(formats strfmt.Registry) error {

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
	}

	if err := validate.Required("firmware", "body", m.Firmware); err != nil {
		return err
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

func (m *VMCreationParams) validateGuestOsType(formats strfmt.Registry) error {
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

func (m *VMCreationParams) validateHa(formats strfmt.Registry) error {

	if err := validate.Required("ha", "body", m.Ha); err != nil {
		return err
	}

	return nil
}

func (m *VMCreationParams) validateIoPolicy(formats strfmt.Registry) error {
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

func (m *VMCreationParams) validateMaxBandwidthPolicy(formats strfmt.Registry) error {
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

func (m *VMCreationParams) validateMaxIopsPolicy(formats strfmt.Registry) error {
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

func (m *VMCreationParams) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *VMCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VMCreationParams) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
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

func (m *VMCreationParams) validateVcpu(formats strfmt.Registry) error {

	if err := validate.Required("vcpu", "body", m.Vcpu); err != nil {
		return err
	}

	return nil
}

func (m *VMCreationParams) validateVMDisks(formats strfmt.Registry) error {

	if err := validate.Required("vm_disks", "body", m.VMDisks); err != nil {
		return err
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

func (m *VMCreationParams) validateVMNics(formats strfmt.Registry) error {

	if err := validate.Required("vm_nics", "body", m.VMNics); err != nil {
		return err
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

// ContextValidate validate this Vm creation params based on the context it is used
func (m *VMCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *VMCreationParams) contextValidateFirmware(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreationParams) contextValidateGuestOsType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreationParams) contextValidateIoPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreationParams) contextValidateMaxBandwidthPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreationParams) contextValidateMaxIopsPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreationParams) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreationParams) contextValidateVMDisks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VMCreationParams) contextValidateVMNics(ctx context.Context, formats strfmt.Registry) error {

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
func (m *VMCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMCreationParams) UnmarshalBinary(b []byte) error {
	var res VMCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
