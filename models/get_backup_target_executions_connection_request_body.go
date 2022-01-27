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

// GetBackupTargetExecutionsConnectionRequestBody get backup target executions connection request body
// Example: {"after":"backupTargetExecutionsConnection-id-string","before":"backupTargetExecutionsConnection-id-string","first":0,"last":0,"orderBy":"backup_group_ASC","skip":0,"where":{"AND":"BackupTargetExecutionWhereInput[]","NOT":"BackupTargetExecutionWhereInput[]","OR":"BackupTargetExecutionWhereInput[]","backup_group":"string","backup_group_contains":"string","backup_group_ends_with":"string","backup_group_gt":"string","backup_group_gte":"string","backup_group_in":["string"],"backup_group_lt":"string","backup_group_lte":"string","backup_group_not":"string","backup_group_not_contains":"string","backup_group_not_ends_with":"string","backup_group_not_in":["string"],"backup_group_not_starts_with":"string","backup_group_starts_with":"string","backup_plan_execution":"BackupPlanExecutionWhereInput","backup_restore_point":"BackupRestorePointWhereInput","cluster_local_id":"string","cluster_local_id_contains":"string","cluster_local_id_ends_with":"string","cluster_local_id_gt":"string","cluster_local_id_gte":"string","cluster_local_id_in":["string"],"cluster_local_id_lt":"string","cluster_local_id_lte":"string","cluster_local_id_not":"string","cluster_local_id_not_contains":"string","cluster_local_id_not_ends_with":"string","cluster_local_id_not_in":["string"],"cluster_local_id_not_starts_with":"string","cluster_local_id_starts_with":"string","duration":0,"duration_gt":0,"duration_gte":0,"duration_in":[0],"duration_lt":0,"duration_lte":0,"duration_not":0,"duration_not_in":[0],"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"executed_at":"string","executed_at_gt":"string","executed_at_gte":"string","executed_at_in":["string"],"executed_at_lt":"string","executed_at_lte":"string","executed_at_not":"string","executed_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","parent_backup":"string","parent_backup_contains":"string","parent_backup_ends_with":"string","parent_backup_gt":"string","parent_backup_gte":"string","parent_backup_in":["string"],"parent_backup_lt":"string","parent_backup_lte":"string","parent_backup_not":"string","parent_backup_not_contains":"string","parent_backup_not_ends_with":"string","parent_backup_not_in":["string"],"parent_backup_not_starts_with":"string","parent_backup_starts_with":"string","read_bytes":0,"read_bytes_gt":0,"read_bytes_gte":0,"read_bytes_in":[0],"read_bytes_lt":0,"read_bytes_lte":0,"read_bytes_not":0,"read_bytes_not_in":[0],"resource_version_gt":0,"resource_version_gte":0,"resource_version_in":[0],"resource_version_lt":0,"resource_version_lte":0,"resource_version_not":0,"resource_version_not_in":[0],"status":"ABORTED","status_in":["ABORTED"],"status_not":"ABORTED","status_not_in":["ABORTED"],"total_bytes":0,"total_bytes_gt":0,"total_bytes_gte":0,"total_bytes_in":[0],"total_bytes_lt":0,"total_bytes_lte":0,"total_bytes_not":0,"total_bytes_not_in":[0],"type":"FULL","type_in":["FULL"],"type_not":"FULL","type_not_in":["FULL"],"vm":"VmWhereInput","vm_local_id":"string","vm_local_id_contains":"string","vm_local_id_ends_with":"string","vm_local_id_gt":"string","vm_local_id_gte":"string","vm_local_id_in":["string"],"vm_local_id_lt":"string","vm_local_id_lte":"string","vm_local_id_not":"string","vm_local_id_not_contains":"string","vm_local_id_not_ends_with":"string","vm_local_id_not_in":["string"],"vm_local_id_not_starts_with":"string","vm_local_id_starts_with":"string","vm_name":"string","vm_name_contains":"string","vm_name_ends_with":"string","vm_name_gt":"string","vm_name_gte":"string","vm_name_in":["string"],"vm_name_lt":"string","vm_name_lte":"string","vm_name_not":"string","vm_name_not_contains":"string","vm_name_not_ends_with":"string","vm_name_not_in":["string"],"vm_name_not_starts_with":"string","vm_name_starts_with":"string"}}
//
// swagger:model GetBackupTargetExecutionsConnectionRequestBody
type GetBackupTargetExecutionsConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *BackupTargetExecutionOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *BackupTargetExecutionWhereInput `json:"where,omitempty"`
}

// Validate validates this get backup target executions connection request body
func (m *GetBackupTargetExecutionsConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetBackupTargetExecutionsConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetBackupTargetExecutionsConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get backup target executions connection request body based on the context it is used
func (m *GetBackupTargetExecutionsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetBackupTargetExecutionsConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetBackupTargetExecutionsConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetBackupTargetExecutionsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupTargetExecutionsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetBackupTargetExecutionsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}