// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetDisksRequestBody get disks request body
// Example: {"after":"disks-id-string","before":"disks-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"DiskWhereInput[]","NOT":"DiskWhereInput[]","OR":"DiskWhereInput[]","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"firmware":"string","firmware_contains":"string","firmware_ends_with":"string","firmware_gt":"string","firmware_gte":"string","firmware_in":["string"],"firmware_lt":"string","firmware_lte":"string","firmware_not":"string","firmware_not_contains":"string","firmware_not_ends_with":"string","firmware_not_in":["string"],"firmware_not_starts_with":"string","firmware_starts_with":"string","function":"CACHE","function_in":["CACHE"],"function_not":"CACHE","function_not_in":["CACHE"],"health_status":"HEALTHY","health_status_in":["HEALTHY"],"health_status_not":"HEALTHY","health_status_not_in":["HEALTHY"],"healthy":false,"healthy_not":false,"host":"HostWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","model":"string","model_contains":"string","model_ends_with":"string","model_gt":"string","model_gte":"string","model_in":["string"],"model_lt":"string","model_lte":"string","model_not":"string","model_not_contains":"string","model_not_ends_with":"string","model_not_in":["string"],"model_not_starts_with":"string","model_starts_with":"string","mounted":false,"mounted_not":false,"name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","numa_node":0,"numa_node_gt":0,"numa_node_gte":0,"numa_node_in":[0],"numa_node_lt":0,"numa_node_lte":0,"numa_node_not":0,"numa_node_not_in":[0],"offline":false,"offline_not":false,"path":"string","path_contains":"string","path_ends_with":"string","path_gt":"string","path_gte":"string","path_in":["string"],"path_lt":"string","path_lte":"string","path_not":"string","path_not_contains":"string","path_not_ends_with":"string","path_not_in":["string"],"path_not_starts_with":"string","path_starts_with":"string","persistent_memory_type":"string","persistent_memory_type_contains":"string","persistent_memory_type_ends_with":"string","persistent_memory_type_gt":"string","persistent_memory_type_gte":"string","persistent_memory_type_in":["string"],"persistent_memory_type_lt":"string","persistent_memory_type_lte":"string","persistent_memory_type_not":"string","persistent_memory_type_not_contains":"string","persistent_memory_type_not_ends_with":"string","persistent_memory_type_not_in":["string"],"persistent_memory_type_not_starts_with":"string","persistent_memory_type_starts_with":"string","physical_slot_on_brick":0,"physical_slot_on_brick_gt":0,"physical_slot_on_brick_gte":0,"physical_slot_on_brick_in":[0],"physical_slot_on_brick_lt":0,"physical_slot_on_brick_lte":0,"physical_slot_on_brick_not":0,"physical_slot_on_brick_not_in":[0],"pmem_dimms_every":"PmemDimmWhereInput","pmem_dimms_none":"PmemDimmWhereInput","pmem_dimms_some":"PmemDimmWhereInput","recommended_usage":"BOOT","recommended_usage_in":["BOOT"],"recommended_usage_not":"BOOT","recommended_usage_not_in":["BOOT"],"remaining_life_percent":0,"remaining_life_percent_gt":0,"remaining_life_percent_gte":0,"remaining_life_percent_in":[0],"remaining_life_percent_lt":0,"remaining_life_percent_lte":0,"remaining_life_percent_not":0,"remaining_life_percent_not_in":[0],"serial":"string","serial_contains":"string","serial_ends_with":"string","serial_gt":"string","serial_gte":"string","serial_in":["string"],"serial_lt":"string","serial_lte":"string","serial_not":"string","serial_not_contains":"string","serial_not_ends_with":"string","serial_not_in":["string"],"serial_not_starts_with":"string","serial_starts_with":"string","size":0,"size_gt":0,"size_gte":0,"size_in":[0],"size_lt":0,"size_lte":0,"size_not":0,"size_not_in":[0],"type":"HDD","type_in":["HDD"],"type_not":"HDD","type_not_in":["HDD"],"usage":"BOOT","usage_in":["BOOT"],"usage_not":"BOOT","usage_not_in":["BOOT"],"usage_status":"ISOLATED","usage_status_in":["ISOLATED"],"usage_status_not":"ISOLATED","usage_status_not_in":["ISOLATED"]}}
//
// swagger:model GetDisksRequestBody
type GetDisksRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *DiskOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *DiskWhereInput `json:"where,omitempty"`
}

// Validate validates this get disks request body
func (m *GetDisksRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDisksRequestBody) validateOrderBy(formats strfmt.Registry) error {
	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	if m.OrderBy != nil {
		if err := m.OrderBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetDisksRequestBody) validateWhere(formats strfmt.Registry) error {
	if swag.IsZero(m.Where) { // not required
		return nil
	}

	if m.Where != nil {
		if err := m.Where.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get disks request body based on the context it is used
func (m *GetDisksRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOrderBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhere(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDisksRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

	if m.OrderBy != nil {
		if err := m.OrderBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orderBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("orderBy")
			}
			return err
		}
	}

	return nil
}

func (m *GetDisksRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

	if m.Where != nil {
		if err := m.Where.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("where")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("where")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDisksRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDisksRequestBody) UnmarshalBinary(b []byte) error {
	var res GetDisksRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}