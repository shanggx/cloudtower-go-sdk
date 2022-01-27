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

// ElfDataStoreOrderByInput elf data store order by input
//
// swagger:model ElfDataStoreOrderByInput
type ElfDataStoreOrderByInput string

func NewElfDataStoreOrderByInput(value ElfDataStoreOrderByInput) *ElfDataStoreOrderByInput {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ElfDataStoreOrderByInput.
func (m ElfDataStoreOrderByInput) Pointer() *ElfDataStoreOrderByInput {
	return &m
}

const (

	// ElfDataStoreOrderByInputCreatedAtASC captures enum value "createdAt_ASC"
	ElfDataStoreOrderByInputCreatedAtASC ElfDataStoreOrderByInput = "createdAt_ASC"

	// ElfDataStoreOrderByInputCreatedAtDESC captures enum value "createdAt_DESC"
	ElfDataStoreOrderByInputCreatedAtDESC ElfDataStoreOrderByInput = "createdAt_DESC"

	// ElfDataStoreOrderByInputDescriptionASC captures enum value "description_ASC"
	ElfDataStoreOrderByInputDescriptionASC ElfDataStoreOrderByInput = "description_ASC"

	// ElfDataStoreOrderByInputDescriptionDESC captures enum value "description_DESC"
	ElfDataStoreOrderByInputDescriptionDESC ElfDataStoreOrderByInput = "description_DESC"

	// ElfDataStoreOrderByInputExternalUseASC captures enum value "external_use_ASC"
	ElfDataStoreOrderByInputExternalUseASC ElfDataStoreOrderByInput = "external_use_ASC"

	// ElfDataStoreOrderByInputExternalUseDESC captures enum value "external_use_DESC"
	ElfDataStoreOrderByInputExternalUseDESC ElfDataStoreOrderByInput = "external_use_DESC"

	// ElfDataStoreOrderByInputIDASC captures enum value "id_ASC"
	ElfDataStoreOrderByInputIDASC ElfDataStoreOrderByInput = "id_ASC"

	// ElfDataStoreOrderByInputIDDESC captures enum value "id_DESC"
	ElfDataStoreOrderByInputIDDESC ElfDataStoreOrderByInput = "id_DESC"

	// ElfDataStoreOrderByInputInternalASC captures enum value "internal_ASC"
	ElfDataStoreOrderByInputInternalASC ElfDataStoreOrderByInput = "internal_ASC"

	// ElfDataStoreOrderByInputInternalDESC captures enum value "internal_DESC"
	ElfDataStoreOrderByInputInternalDESC ElfDataStoreOrderByInput = "internal_DESC"

	// ElfDataStoreOrderByInputIPWhitelistASC captures enum value "ip_whitelist_ASC"
	ElfDataStoreOrderByInputIPWhitelistASC ElfDataStoreOrderByInput = "ip_whitelist_ASC"

	// ElfDataStoreOrderByInputIPWhitelistDESC captures enum value "ip_whitelist_DESC"
	ElfDataStoreOrderByInputIPWhitelistDESC ElfDataStoreOrderByInput = "ip_whitelist_DESC"

	// ElfDataStoreOrderByInputLocalIDASC captures enum value "local_id_ASC"
	ElfDataStoreOrderByInputLocalIDASC ElfDataStoreOrderByInput = "local_id_ASC"

	// ElfDataStoreOrderByInputLocalIDDESC captures enum value "local_id_DESC"
	ElfDataStoreOrderByInputLocalIDDESC ElfDataStoreOrderByInput = "local_id_DESC"

	// ElfDataStoreOrderByInputNameASC captures enum value "name_ASC"
	ElfDataStoreOrderByInputNameASC ElfDataStoreOrderByInput = "name_ASC"

	// ElfDataStoreOrderByInputNameDESC captures enum value "name_DESC"
	ElfDataStoreOrderByInputNameDESC ElfDataStoreOrderByInput = "name_DESC"

	// ElfDataStoreOrderByInputReplicaNumASC captures enum value "replica_num_ASC"
	ElfDataStoreOrderByInputReplicaNumASC ElfDataStoreOrderByInput = "replica_num_ASC"

	// ElfDataStoreOrderByInputReplicaNumDESC captures enum value "replica_num_DESC"
	ElfDataStoreOrderByInputReplicaNumDESC ElfDataStoreOrderByInput = "replica_num_DESC"

	// ElfDataStoreOrderByInputThinProvisionASC captures enum value "thin_provision_ASC"
	ElfDataStoreOrderByInputThinProvisionASC ElfDataStoreOrderByInput = "thin_provision_ASC"

	// ElfDataStoreOrderByInputThinProvisionDESC captures enum value "thin_provision_DESC"
	ElfDataStoreOrderByInputThinProvisionDESC ElfDataStoreOrderByInput = "thin_provision_DESC"

	// ElfDataStoreOrderByInputTypeASC captures enum value "type_ASC"
	ElfDataStoreOrderByInputTypeASC ElfDataStoreOrderByInput = "type_ASC"

	// ElfDataStoreOrderByInputTypeDESC captures enum value "type_DESC"
	ElfDataStoreOrderByInputTypeDESC ElfDataStoreOrderByInput = "type_DESC"

	// ElfDataStoreOrderByInputUpdatedAtASC captures enum value "updatedAt_ASC"
	ElfDataStoreOrderByInputUpdatedAtASC ElfDataStoreOrderByInput = "updatedAt_ASC"

	// ElfDataStoreOrderByInputUpdatedAtDESC captures enum value "updatedAt_DESC"
	ElfDataStoreOrderByInputUpdatedAtDESC ElfDataStoreOrderByInput = "updatedAt_DESC"
)

// for schema
var elfDataStoreOrderByInputEnum []interface{}

func init() {
	var res []ElfDataStoreOrderByInput
	if err := json.Unmarshal([]byte(`["createdAt_ASC","createdAt_DESC","description_ASC","description_DESC","external_use_ASC","external_use_DESC","id_ASC","id_DESC","internal_ASC","internal_DESC","ip_whitelist_ASC","ip_whitelist_DESC","local_id_ASC","local_id_DESC","name_ASC","name_DESC","replica_num_ASC","replica_num_DESC","thin_provision_ASC","thin_provision_DESC","type_ASC","type_DESC","updatedAt_ASC","updatedAt_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		elfDataStoreOrderByInputEnum = append(elfDataStoreOrderByInputEnum, v)
	}
}

func (m ElfDataStoreOrderByInput) validateElfDataStoreOrderByInputEnum(path, location string, value ElfDataStoreOrderByInput) error {
	if err := validate.EnumCase(path, location, value, elfDataStoreOrderByInputEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this elf data store order by input
func (m ElfDataStoreOrderByInput) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateElfDataStoreOrderByInputEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this elf data store order by input based on context it is used
func (m ElfDataStoreOrderByInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}