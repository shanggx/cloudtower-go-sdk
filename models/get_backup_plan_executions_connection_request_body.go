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

// GetBackupPlanExecutionsConnectionRequestBody get backup plan executions connection request body
// Example: {"after":"backupPlanExecutionsConnection-id-string","before":"backupPlanExecutionsConnection-id-string","first":0,"last":0,"orderBy":"createdAt_ASC","skip":0,"where":{"AND":"BackupPlanExecutionWhereInput[]","NOT":"BackupPlanExecutionWhereInput[]","OR":"BackupPlanExecutionWhereInput[]","backup_plan":"BackupPlanWhereInput","backup_restore_points_every":"BackupRestorePointWhereInput","backup_restore_points_none":"BackupRestorePointWhereInput","backup_restore_points_some":"BackupRestorePointWhereInput","backup_target_executions_every":"BackupTargetExecutionWhereInput","backup_target_executions_none":"BackupTargetExecutionWhereInput","backup_target_executions_some":"BackupTargetExecutionWhereInput","duration":0,"duration_gt":0,"duration_gte":0,"duration_in":[0],"duration_lt":0,"duration_lte":0,"duration_not":0,"duration_not_in":[0],"entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"executed_at":"string","executed_at_gt":"string","executed_at_gte":"string","executed_at_in":["string"],"executed_at_lt":"string","executed_at_lte":"string","executed_at_not":"string","executed_at_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","method":"AUTO","method_in":["AUTO"],"method_not":"AUTO","method_not_in":["AUTO"],"resource_version_gt":0,"resource_version_gte":0,"resource_version_in":[0],"resource_version_lt":0,"resource_version_lte":0,"resource_version_not":0,"resource_version_not_in":[0],"status":"FAILED","status_in":["FAILED"],"status_not":"FAILED","status_not_in":["FAILED"],"success_job_count":0,"success_job_count_gt":0,"success_job_count_gte":0,"success_job_count_in":[0],"success_job_count_lt":0,"success_job_count_lte":0,"success_job_count_not":0,"success_job_count_not_in":[0],"total_job_count":0,"total_job_count_gt":0,"total_job_count_gte":0,"total_job_count_in":[0],"total_job_count_lt":0,"total_job_count_lte":0,"total_job_count_not":0,"total_job_count_not_in":[0],"type":"FULL","type_in":["FULL"],"type_not":"FULL","type_not_in":["FULL"]}}
//
// swagger:model GetBackupPlanExecutionsConnectionRequestBody
type GetBackupPlanExecutionsConnectionRequestBody struct {

	// after
	After *string `json:"after,omitempty"`

	// before
	Before *string `json:"before,omitempty"`

	// first
	First *int32 `json:"first,omitempty"`

	// last
	Last *int32 `json:"last,omitempty"`

	// order by
	OrderBy *BackupPlanExecutionOrderByInput `json:"orderBy,omitempty"`

	// skip
	Skip *int32 `json:"skip,omitempty"`

	// where
	Where *BackupPlanExecutionWhereInput `json:"where,omitempty"`
}

// Validate validates this get backup plan executions connection request body
func (m *GetBackupPlanExecutionsConnectionRequestBody) Validate(formats strfmt.Registry) error {
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

func (m *GetBackupPlanExecutionsConnectionRequestBody) validateOrderBy(formats strfmt.Registry) error {
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

func (m *GetBackupPlanExecutionsConnectionRequestBody) validateWhere(formats strfmt.Registry) error {
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

// ContextValidate validate this get backup plan executions connection request body based on the context it is used
func (m *GetBackupPlanExecutionsConnectionRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *GetBackupPlanExecutionsConnectionRequestBody) contextValidateOrderBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GetBackupPlanExecutionsConnectionRequestBody) contextValidateWhere(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetBackupPlanExecutionsConnectionRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBackupPlanExecutionsConnectionRequestBody) UnmarshalBinary(b []byte) error {
	var res GetBackupPlanExecutionsConnectionRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
