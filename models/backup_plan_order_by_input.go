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

// BackupPlanOrderByInput backup plan order by input
//
// swagger:model BackupPlanOrderByInput
type BackupPlanOrderByInput string

func NewBackupPlanOrderByInput(value BackupPlanOrderByInput) *BackupPlanOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated BackupPlanOrderByInput.
func (m BackupPlanOrderByInput) Pointer() *BackupPlanOrderByInput {
	return &m
}

const (

	// BackupPlanOrderByInputBackupDelayOptionASC captures enum value "backup_delay_option_ASC"
	BackupPlanOrderByInputBackupDelayOptionASC BackupPlanOrderByInput = "backup_delay_option_ASC"

	// BackupPlanOrderByInputBackupDelayOptionDESC captures enum value "backup_delay_option_DESC"
	BackupPlanOrderByInputBackupDelayOptionDESC BackupPlanOrderByInput = "backup_delay_option_DESC"

	// BackupPlanOrderByInputBackupRestorePointCountASC captures enum value "backup_restore_point_count_ASC"
	BackupPlanOrderByInputBackupRestorePointCountASC BackupPlanOrderByInput = "backup_restore_point_count_ASC"

	// BackupPlanOrderByInputBackupRestorePointCountDESC captures enum value "backup_restore_point_count_DESC"
	BackupPlanOrderByInputBackupRestorePointCountDESC BackupPlanOrderByInput = "backup_restore_point_count_DESC"

	// BackupPlanOrderByInputBackupTotalSizeASC captures enum value "backup_total_size_ASC"
	BackupPlanOrderByInputBackupTotalSizeASC BackupPlanOrderByInput = "backup_total_size_ASC"

	// BackupPlanOrderByInputBackupTotalSizeDESC captures enum value "backup_total_size_DESC"
	BackupPlanOrderByInputBackupTotalSizeDESC BackupPlanOrderByInput = "backup_total_size_DESC"

	// BackupPlanOrderByInputCompressionASC captures enum value "compression_ASC"
	BackupPlanOrderByInputCompressionASC BackupPlanOrderByInput = "compression_ASC"

	// BackupPlanOrderByInputCompressionDESC captures enum value "compression_DESC"
	BackupPlanOrderByInputCompressionDESC BackupPlanOrderByInput = "compression_DESC"

	// BackupPlanOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	BackupPlanOrderByInputCreatedAtASC BackupPlanOrderByInput = "createdAt_ASC"

	// BackupPlanOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	BackupPlanOrderByInputCreatedAtDESC BackupPlanOrderByInput = "createdAt_DESC"

	// BackupPlanOrderByInputDeleteStrategyASC captures enum value "delete_strategy_ASC"
	BackupPlanOrderByInputDeleteStrategyASC BackupPlanOrderByInput = "delete_strategy_ASC"

	// BackupPlanOrderByInputDeleteStrategyDESC captures enum value "delete_strategy_DESC"
	BackupPlanOrderByInputDeleteStrategyDESC BackupPlanOrderByInput = "delete_strategy_DESC"

	// BackupPlanOrderByInputDescriptionASC captures enum value "description_ASC"
	BackupPlanOrderByInputDescriptionASC BackupPlanOrderByInput = "description_ASC"

	// BackupPlanOrderByInputDescriptionDESC captures enum value "description_DESC"
	BackupPlanOrderByInputDescriptionDESC BackupPlanOrderByInput = "description_DESC"

	// BackupPlanOrderByInputEnableWindowASC captures enum value "enable_window_ASC"
	BackupPlanOrderByInputEnableWindowASC BackupPlanOrderByInput = "enable_window_ASC"

	// BackupPlanOrderByInputEnableWindowDESC captures enum value "enable_window_DESC"
	BackupPlanOrderByInputEnableWindowDESC BackupPlanOrderByInput = "enable_window_DESC"

	// BackupPlanOrderByInputEntityAsyncStatusASC captures enum value "entityAsyncStatus_ASC"
	BackupPlanOrderByInputEntityAsyncStatusASC BackupPlanOrderByInput = "entityAsyncStatus_ASC"

	// BackupPlanOrderByInputEntityAsyncStatusDESC captures enum value "entityAsyncStatus_DESC"
	BackupPlanOrderByInputEntityAsyncStatusDESC BackupPlanOrderByInput = "entityAsyncStatus_DESC"

	// BackupPlanOrderByInputFullIntervalASC captures enum value "full_interval_ASC"
	BackupPlanOrderByInputFullIntervalASC BackupPlanOrderByInput = "full_interval_ASC"

	// BackupPlanOrderByInputFullIntervalDESC captures enum value "full_interval_DESC"
	BackupPlanOrderByInputFullIntervalDESC BackupPlanOrderByInput = "full_interval_DESC"

	// BackupPlanOrderByInputFullPeriodASC captures enum value "full_period_ASC"
	BackupPlanOrderByInputFullPeriodASC BackupPlanOrderByInput = "full_period_ASC"

	// BackupPlanOrderByInputFullPeriodDESC captures enum value "full_period_DESC"
	BackupPlanOrderByInputFullPeriodDESC BackupPlanOrderByInput = "full_period_DESC"

	// BackupPlanOrderByInputFullTimePointASC captures enum value "full_time_point_ASC"
	BackupPlanOrderByInputFullTimePointASC BackupPlanOrderByInput = "full_time_point_ASC"

	// BackupPlanOrderByInputFullTimePointDESC captures enum value "full_time_point_DESC"
	BackupPlanOrderByInputFullTimePointDESC BackupPlanOrderByInput = "full_time_point_DESC"

	// BackupPlanOrderByInputIDASC captures enum value "id_ASC"
	BackupPlanOrderByInputIDASC BackupPlanOrderByInput = "id_ASC"

	// BackupPlanOrderByInputIDDESC captures enum value "id_DESC"
	BackupPlanOrderByInputIDDESC BackupPlanOrderByInput = "id_DESC"

	// BackupPlanOrderByInputIncrementalPeriodASC captures enum value "incremental_period_ASC"
	BackupPlanOrderByInputIncrementalPeriodASC BackupPlanOrderByInput = "incremental_period_ASC"

	// BackupPlanOrderByInputIncrementalPeriodDESC captures enum value "incremental_period_DESC"
	BackupPlanOrderByInputIncrementalPeriodDESC BackupPlanOrderByInput = "incremental_period_DESC"

	// BackupPlanOrderByInputIncrementalTimePointsASC captures enum value "incremental_time_points_ASC"
	BackupPlanOrderByInputIncrementalTimePointsASC BackupPlanOrderByInput = "incremental_time_points_ASC"

	// BackupPlanOrderByInputIncrementalTimePointsDESC captures enum value "incremental_time_points_DESC"
	BackupPlanOrderByInputIncrementalTimePointsDESC BackupPlanOrderByInput = "incremental_time_points_DESC"

	// BackupPlanOrderByInputKeepPolicyASC captures enum value "keep_policy_ASC"
	BackupPlanOrderByInputKeepPolicyASC BackupPlanOrderByInput = "keep_policy_ASC"

	// BackupPlanOrderByInputKeepPolicyDESC captures enum value "keep_policy_DESC"
	BackupPlanOrderByInputKeepPolicyDESC BackupPlanOrderByInput = "keep_policy_DESC"

	// BackupPlanOrderByInputKeepPolicyValueASC captures enum value "keep_policy_value_ASC"
	BackupPlanOrderByInputKeepPolicyValueASC BackupPlanOrderByInput = "keep_policy_value_ASC"

	// BackupPlanOrderByInputKeepPolicyValueDESC captures enum value "keep_policy_value_DESC"
	BackupPlanOrderByInputKeepPolicyValueDESC BackupPlanOrderByInput = "keep_policy_value_DESC"

	// BackupPlanOrderByInputLastExecuteStatusASC captures enum value "last_execute_status_ASC"
	BackupPlanOrderByInputLastExecuteStatusASC BackupPlanOrderByInput = "last_execute_status_ASC"

	// BackupPlanOrderByInputLastExecuteStatusDESC captures enum value "last_execute_status_DESC"
	BackupPlanOrderByInputLastExecuteStatusDESC BackupPlanOrderByInput = "last_execute_status_DESC"

	// BackupPlanOrderByInputLastExecuteStatusMessageASC captures enum value "last_execute_status_message_ASC"
	BackupPlanOrderByInputLastExecuteStatusMessageASC BackupPlanOrderByInput = "last_execute_status_message_ASC"

	// BackupPlanOrderByInputLastExecuteStatusMessageDESC captures enum value "last_execute_status_message_DESC"
	BackupPlanOrderByInputLastExecuteStatusMessageDESC BackupPlanOrderByInput = "last_execute_status_message_DESC"

	// BackupPlanOrderByInputLastExecuteSuccessJobCountASC captures enum value "last_execute_success_job_count_ASC"
	BackupPlanOrderByInputLastExecuteSuccessJobCountASC BackupPlanOrderByInput = "last_execute_success_job_count_ASC"

	// BackupPlanOrderByInputLastExecuteSuccessJobCountDESC captures enum value "last_execute_success_job_count_DESC"
	BackupPlanOrderByInputLastExecuteSuccessJobCountDESC BackupPlanOrderByInput = "last_execute_success_job_count_DESC"

	// BackupPlanOrderByInputLastExecuteTotalJobCountASC captures enum value "last_execute_total_job_count_ASC"
	BackupPlanOrderByInputLastExecuteTotalJobCountASC BackupPlanOrderByInput = "last_execute_total_job_count_ASC"

	// BackupPlanOrderByInputLastExecuteTotalJobCountDESC captures enum value "last_execute_total_job_count_DESC"
	BackupPlanOrderByInputLastExecuteTotalJobCountDESC BackupPlanOrderByInput = "last_execute_total_job_count_DESC"

	// BackupPlanOrderByInputLastExecutedAtASC captures enum value "last_executed_at_ASC"
	BackupPlanOrderByInputLastExecutedAtASC BackupPlanOrderByInput = "last_executed_at_ASC"

	// BackupPlanOrderByInputLastExecutedAtDESC captures enum value "last_executed_at_DESC"
	BackupPlanOrderByInputLastExecutedAtDESC BackupPlanOrderByInput = "last_executed_at_DESC"

	// BackupPlanOrderByInputNameASC captures enum value "name_ASC"
	BackupPlanOrderByInputNameASC BackupPlanOrderByInput = "name_ASC"

	// BackupPlanOrderByInputNameDESC captures enum value "name_DESC"
	BackupPlanOrderByInputNameDESC BackupPlanOrderByInput = "name_DESC"

	// BackupPlanOrderByInputNamespaceASC captures enum value "namespace_ASC"
	BackupPlanOrderByInputNamespaceASC BackupPlanOrderByInput = "namespace_ASC"

	// BackupPlanOrderByInputNamespaceDESC captures enum value "namespace_DESC"
	BackupPlanOrderByInputNamespaceDESC BackupPlanOrderByInput = "namespace_DESC"

	// BackupPlanOrderByInputNextExecuteTimeASC captures enum value "next_execute_time_ASC"
	BackupPlanOrderByInputNextExecuteTimeASC BackupPlanOrderByInput = "next_execute_time_ASC"

	// BackupPlanOrderByInputNextExecuteTimeDESC captures enum value "next_execute_time_DESC"
	BackupPlanOrderByInputNextExecuteTimeDESC BackupPlanOrderByInput = "next_execute_time_DESC"

	// BackupPlanOrderByInputResourceVersionASC captures enum value "resource_version_ASC"
	BackupPlanOrderByInputResourceVersionASC BackupPlanOrderByInput = "resource_version_ASC"

	// BackupPlanOrderByInputResourceVersionDESC captures enum value "resource_version_DESC"
	BackupPlanOrderByInputResourceVersionDESC BackupPlanOrderByInput = "resource_version_DESC"

	// BackupPlanOrderByInputStatusASC captures enum value "status_ASC"
	BackupPlanOrderByInputStatusASC BackupPlanOrderByInput = "status_ASC"

	// BackupPlanOrderByInputStatusDESC captures enum value "status_DESC"
	BackupPlanOrderByInputStatusDESC BackupPlanOrderByInput = "status_DESC"

	// BackupPlanOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	BackupPlanOrderByInputUpdatedAtASC BackupPlanOrderByInput = "updatedAt_ASC"

	// BackupPlanOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	BackupPlanOrderByInputUpdatedAtDESC BackupPlanOrderByInput = "updatedAt_DESC"

	// BackupPlanOrderByInputWindowEndASC captures enum value "window_end_ASC"
	BackupPlanOrderByInputWindowEndASC BackupPlanOrderByInput = "window_end_ASC"

	// BackupPlanOrderByInputWindowEndDESC captures enum value "window_end_DESC"
	BackupPlanOrderByInputWindowEndDESC BackupPlanOrderByInput = "window_end_DESC"

	// BackupPlanOrderByInputWindowStartASC captures enum value "window_start_ASC"
	BackupPlanOrderByInputWindowStartASC BackupPlanOrderByInput = "window_start_ASC"

	// BackupPlanOrderByInputWindowStartDESC captures enum value "window_start_DESC"
	BackupPlanOrderByInputWindowStartDESC BackupPlanOrderByInput = "window_start_DESC"
)

// for schema
var backupPlanOrderByInputEnum []interface{}

func init() {
	var res []BackupPlanOrderByInput
	if err := json.Unmarshal([]byte(`["backup_delay_option_ASC","backup_delay_option_DESC","backup_restore_point_count_ASC","backup_restore_point_count_DESC","backup_total_size_ASC","backup_total_size_DESC","compression_ASC","compression_DESC","createdAt_ASC","createdAt_DESC","delete_strategy_ASC","delete_strategy_DESC","description_ASC","description_DESC","enable_window_ASC","enable_window_DESC","entityAsyncStatus_ASC","entityAsyncStatus_DESC","full_interval_ASC","full_interval_DESC","full_period_ASC","full_period_DESC","full_time_point_ASC","full_time_point_DESC","id_ASC","id_DESC","incremental_period_ASC","incremental_period_DESC","incremental_time_points_ASC","incremental_time_points_DESC","keep_policy_ASC","keep_policy_DESC","keep_policy_value_ASC","keep_policy_value_DESC","last_execute_status_ASC","last_execute_status_DESC","last_execute_status_message_ASC","last_execute_status_message_DESC","last_execute_success_job_count_ASC","last_execute_success_job_count_DESC","last_execute_total_job_count_ASC","last_execute_total_job_count_DESC","last_executed_at_ASC","last_executed_at_DESC","name_ASC","name_DESC","namespace_ASC","namespace_DESC","next_execute_time_ASC","next_execute_time_DESC","resource_version_ASC","resource_version_DESC","status_ASC","status_DESC","updatedAt_ASC","updatedAt_DESC","window_end_ASC","window_end_DESC","window_start_ASC","window_start_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupPlanOrderByInputEnum = append(backupPlanOrderByInputEnum, v)
	}
}

func (m BackupPlanOrderByInput) validateBackupPlanOrderByInputEnum(path, location string, value BackupPlanOrderByInput) error {
	if err := validate.EnumCase(path, location, value, backupPlanOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup plan order by input
func (m BackupPlanOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupPlanOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup plan order by input based on context it is used
func (m BackupPlanOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
