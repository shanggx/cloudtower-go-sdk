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

// ZoneOrderByInput zone order by input
//
// swagger:model ZoneOrderByInput
type ZoneOrderByInput string

func NewZoneOrderByInput(value ZoneOrderByInput) *ZoneOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ZoneOrderByInput.
func (m ZoneOrderByInput) Pointer() *ZoneOrderByInput {
	return &m
}

const (

	// ZoneOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ZoneOrderByInputCreatedAtASC ZoneOrderByInput = "createdAt_ASC"

	// ZoneOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ZoneOrderByInputCreatedAtDESC ZoneOrderByInput = "createdAt_DESC"

	// ZoneOrderByInputFailureDataSpaceASC captures enum value "failure_data_space_ASC"
	ZoneOrderByInputFailureDataSpaceASC ZoneOrderByInput = "failure_data_space_ASC"

	// ZoneOrderByInputFailureDataSpaceDESC captures enum value "failure_data_space_DESC"
	ZoneOrderByInputFailureDataSpaceDESC ZoneOrderByInput = "failure_data_space_DESC"

	// ZoneOrderByInputHostNumASC captures enum value "host_num_ASC"
	ZoneOrderByInputHostNumASC ZoneOrderByInput = "host_num_ASC"

	// ZoneOrderByInputHostNumDESC captures enum value "host_num_DESC"
	ZoneOrderByInputHostNumDESC ZoneOrderByInput = "host_num_DESC"

	// ZoneOrderByInputIDASC captures enum value "id_ASC"
	ZoneOrderByInputIDASC ZoneOrderByInput = "id_ASC"

	// ZoneOrderByInputIDDESC captures enum value "id_DESC"
	ZoneOrderByInputIDDESC ZoneOrderByInput = "id_DESC"

	// ZoneOrderByInputIsPreferredASC captures enum value "is_preferred_ASC"
	ZoneOrderByInputIsPreferredASC ZoneOrderByInput = "is_preferred_ASC"

	// ZoneOrderByInputIsPreferredDESC captures enum value "is_preferred_DESC"
	ZoneOrderByInputIsPreferredDESC ZoneOrderByInput = "is_preferred_DESC"

	// ZoneOrderByInputLocalIDASC captures enum value "local_id_ASC"
	ZoneOrderByInputLocalIDASC ZoneOrderByInput = "local_id_ASC"

	// ZoneOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	ZoneOrderByInputLocalIDDESC ZoneOrderByInput = "local_id_DESC"

	// ZoneOrderByInputProvisionedCPUCoresASC captures enum value "provisioned_cpu_cores_ASC"
	ZoneOrderByInputProvisionedCPUCoresASC ZoneOrderByInput = "provisioned_cpu_cores_ASC"

	// ZoneOrderByInputProvisionedCPUCoresDESC captures enum value "provisioned_cpu_cores_DESC"
	ZoneOrderByInputProvisionedCPUCoresDESC ZoneOrderByInput = "provisioned_cpu_cores_DESC"

	// ZoneOrderByInputProvisionedCPUCoresForActiveVMASC captures enum value "provisioned_cpu_cores_for_active_vm_ASC"
	ZoneOrderByInputProvisionedCPUCoresForActiveVMASC ZoneOrderByInput = "provisioned_cpu_cores_for_active_vm_ASC"

	// ZoneOrderByInputProvisionedCPUCoresForActiveVMDESC captures enum value "provisioned_cpu_cores_for_active_vm_DESC"
	ZoneOrderByInputProvisionedCPUCoresForActiveVMDESC ZoneOrderByInput = "provisioned_cpu_cores_for_active_vm_DESC"

	// ZoneOrderByInputProvisionedDataSpaceASC captures enum value "provisioned_data_space_ASC"
	ZoneOrderByInputProvisionedDataSpaceASC ZoneOrderByInput = "provisioned_data_space_ASC"

	// ZoneOrderByInputProvisionedDataSpaceDESC captures enum value "provisioned_data_space_DESC"
	ZoneOrderByInputProvisionedDataSpaceDESC ZoneOrderByInput = "provisioned_data_space_DESC"

	// ZoneOrderByInputProvisionedMemoryBytesASC captures enum value "provisioned_memory_bytes_ASC"
	ZoneOrderByInputProvisionedMemoryBytesASC ZoneOrderByInput = "provisioned_memory_bytes_ASC"

	// ZoneOrderByInputProvisionedMemoryBytesDESC captures enum value "provisioned_memory_bytes_DESC"
	ZoneOrderByInputProvisionedMemoryBytesDESC ZoneOrderByInput = "provisioned_memory_bytes_DESC"

	// ZoneOrderByInputRunningVMNumASC captures enum value "running_vm_num_ASC"
	ZoneOrderByInputRunningVMNumASC ZoneOrderByInput = "running_vm_num_ASC"

	// ZoneOrderByInputRunningVMNumDESC captures enum value "running_vm_num_DESC"
	ZoneOrderByInputRunningVMNumDESC ZoneOrderByInput = "running_vm_num_DESC"

	// ZoneOrderByInputStoppedVMNumASC captures enum value "stopped_vm_num_ASC"
	ZoneOrderByInputStoppedVMNumASC ZoneOrderByInput = "stopped_vm_num_ASC"

	// ZoneOrderByInputStoppedVMNumDESC captures enum value "stopped_vm_num_DESC"
	ZoneOrderByInputStoppedVMNumDESC ZoneOrderByInput = "stopped_vm_num_DESC"

	// ZoneOrderByInputSuspendedVMNumASC captures enum value "suspended_vm_num_ASC"
	ZoneOrderByInputSuspendedVMNumASC ZoneOrderByInput = "suspended_vm_num_ASC"

	// ZoneOrderByInputSuspendedVMNumDESC captures enum value "suspended_vm_num_DESC"
	ZoneOrderByInputSuspendedVMNumDESC ZoneOrderByInput = "suspended_vm_num_DESC"

	// ZoneOrderByInputTotalCacheCapacityASC captures enum value "total_cache_capacity_ASC"
	ZoneOrderByInputTotalCacheCapacityASC ZoneOrderByInput = "total_cache_capacity_ASC"

	// ZoneOrderByInputTotalCacheCapacityDESC captures enum value "total_cache_capacity_DESC"
	ZoneOrderByInputTotalCacheCapacityDESC ZoneOrderByInput = "total_cache_capacity_DESC"

	// ZoneOrderByInputTotalCPUCoresASC captures enum value "total_cpu_cores_ASC"
	ZoneOrderByInputTotalCPUCoresASC ZoneOrderByInput = "total_cpu_cores_ASC"

	// ZoneOrderByInputTotalCPUCoresDESC captures enum value "total_cpu_cores_DESC"
	ZoneOrderByInputTotalCPUCoresDESC ZoneOrderByInput = "total_cpu_cores_DESC"

	// ZoneOrderByInputTotalCPUHzASC captures enum value "total_cpu_hz_ASC"
	ZoneOrderByInputTotalCPUHzASC ZoneOrderByInput = "total_cpu_hz_ASC"

	// ZoneOrderByInputTotalCPUHzDESC captures enum value "total_cpu_hz_DESC"
	ZoneOrderByInputTotalCPUHzDESC ZoneOrderByInput = "total_cpu_hz_DESC"

	// ZoneOrderByInputTotalDataCapacityASC captures enum value "total_data_capacity_ASC"
	ZoneOrderByInputTotalDataCapacityASC ZoneOrderByInput = "total_data_capacity_ASC"

	// ZoneOrderByInputTotalDataCapacityDESC captures enum value "total_data_capacity_DESC"
	ZoneOrderByInputTotalDataCapacityDESC ZoneOrderByInput = "total_data_capacity_DESC"

	// ZoneOrderByInputTotalMemoryBytesASC captures enum value "total_memory_bytes_ASC"
	ZoneOrderByInputTotalMemoryBytesASC ZoneOrderByInput = "total_memory_bytes_ASC"

	// ZoneOrderByInputTotalMemoryBytesDESC captures enum value "total_memory_bytes_DESC"
	ZoneOrderByInputTotalMemoryBytesDESC ZoneOrderByInput = "total_memory_bytes_DESC"

	// ZoneOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	ZoneOrderByInputUpdatedAtASC ZoneOrderByInput = "updatedAt_ASC"

	// ZoneOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	ZoneOrderByInputUpdatedAtDESC ZoneOrderByInput = "updatedAt_DESC"

	// ZoneOrderByInputUsedDataSpaceASC captures enum value "used_data_space_ASC"
	ZoneOrderByInputUsedDataSpaceASC ZoneOrderByInput = "used_data_space_ASC"

	// ZoneOrderByInputUsedDataSpaceDESC captures enum value "used_data_space_DESC"
	ZoneOrderByInputUsedDataSpaceDESC ZoneOrderByInput = "used_data_space_DESC"

	// ZoneOrderByInputValidDataSpaceASC captures enum value "valid_data_space_ASC"
	ZoneOrderByInputValidDataSpaceASC ZoneOrderByInput = "valid_data_space_ASC"

	// ZoneOrderByInputValidDataSpaceDESC captures enum value "valid_data_space_DESC"
	ZoneOrderByInputValidDataSpaceDESC ZoneOrderByInput = "valid_data_space_DESC"

	// ZoneOrderByInputVMNumASC captures enum value "vm_num_ASC"
	ZoneOrderByInputVMNumASC ZoneOrderByInput = "vm_num_ASC"

	// ZoneOrderByInputVMNumDESC captures enum value "vm_num_DESC"
	ZoneOrderByInputVMNumDESC ZoneOrderByInput = "vm_num_DESC"
)

// for schema
var zoneOrderByInputEnum []interface{}

func init() {
	var res []ZoneOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","failure_data_space_ASC","failure_data_space_DESC","host_num_ASC","host_num_DESC","id_ASC","id_DESC","is_preferred_ASC","is_preferred_DESC","local_id_ASC","local_id_DESC","provisioned_cpu_cores_ASC","provisioned_cpu_cores_DESC","provisioned_cpu_cores_for_active_vm_ASC","provisioned_cpu_cores_for_active_vm_DESC","provisioned_data_space_ASC","provisioned_data_space_DESC","provisioned_memory_bytes_ASC","provisioned_memory_bytes_DESC","running_vm_num_ASC","running_vm_num_DESC","stopped_vm_num_ASC","stopped_vm_num_DESC","suspended_vm_num_ASC","suspended_vm_num_DESC","total_cache_capacity_ASC","total_cache_capacity_DESC","total_cpu_cores_ASC","total_cpu_cores_DESC","total_cpu_hz_ASC","total_cpu_hz_DESC","total_data_capacity_ASC","total_data_capacity_DESC","total_memory_bytes_ASC","total_memory_bytes_DESC","updatedAt_ASC","updatedAt_DESC","used_data_space_ASC","used_data_space_DESC","valid_data_space_ASC","valid_data_space_DESC","vm_num_ASC","vm_num_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zoneOrderByInputEnum = append(zoneOrderByInputEnum, v)
	}
}

func (m ZoneOrderByInput) validateZoneOrderByInputEnum(path, location string, value ZoneOrderByInput) error {
	if err := validate.EnumCase(path, location, value, zoneOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this zone order by input
func (m ZoneOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateZoneOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this zone order by input based on context it is used
func (m ZoneOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}