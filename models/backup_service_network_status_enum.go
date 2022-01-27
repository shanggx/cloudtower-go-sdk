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

// BackupServiceNetworkStatusEnum backup service network status enum
//
// swagger:model BackupServiceNetworkStatusEnum
type BackupServiceNetworkStatusEnum string

func NewBackupServiceNetworkStatusEnum(value BackupServiceNetworkStatusEnum) *BackupServiceNetworkStatusEnum {
	return &value
}

// Pointer returns a pointer to a freshly-allocated BackupServiceNetworkStatusEnum.
func (m BackupServiceNetworkStatusEnum) Pointer() *BackupServiceNetworkStatusEnum {
	return &m
}

const (

	// BackupServiceNetworkStatusEnumCONNECTED captures enum value "CONNECTED"
	BackupServiceNetworkStatusEnumCONNECTED BackupServiceNetworkStatusEnum = "CONNECTED"

	// BackupServiceNetworkStatusEnumDISCONNECTED captures enum value "DISCONNECTED"
	BackupServiceNetworkStatusEnumDISCONNECTED BackupServiceNetworkStatusEnum = "DISCONNECTED"
)

// for schema
var backupServiceNetworkStatusEnumEnum []interface{}

func init() {
	var res []BackupServiceNetworkStatusEnum
	if err := json.Unmarshal([]byte(`["CONNECTED","DISCONNECTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		backupServiceNetworkStatusEnumEnum = append(backupServiceNetworkStatusEnumEnum, v)
	}
}

func (m BackupServiceNetworkStatusEnum) validateBackupServiceNetworkStatusEnumEnum(path, location string, value BackupServiceNetworkStatusEnum) error {
	if err := validate.EnumCase(path, location, value, backupServiceNetworkStatusEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this backup service network status enum
func (m BackupServiceNetworkStatusEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBackupServiceNetworkStatusEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this backup service network status enum based on context it is used
func (m BackupServiceNetworkStatusEnum) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}