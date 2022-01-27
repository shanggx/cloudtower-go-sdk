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
)

// BackupLicenseWhereInput backup license where input
// Example: {"AND":"BackupLicenseWhereInput[]","NOT":"BackupLicenseWhereInput[]","OR":"BackupLicenseWhereInput[]","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"expire_date":"string","expire_date_gt":"string","expire_date_gte":"string","expire_date_in":["string"],"expire_date_lt":"string","expire_date_lte":"string","expire_date_not":"string","expire_date_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","license_serial":"string","license_serial_contains":"string","license_serial_ends_with":"string","license_serial_gt":"string","license_serial_gte":"string","license_serial_in":["string"],"license_serial_lt":"string","license_serial_lte":"string","license_serial_not":"string","license_serial_not_contains":"string","license_serial_not_ends_with":"string","license_serial_not_in":["string"],"license_serial_not_starts_with":"string","license_serial_starts_with":"string","max_capacity":0,"max_capacity_gt":0,"max_capacity_gte":0,"max_capacity_in":[0],"max_capacity_lt":0,"max_capacity_lte":0,"max_capacity_not":0,"max_capacity_not_in":[0],"sign_date":"string","sign_date_gt":"string","sign_date_gte":"string","sign_date_in":["string"],"sign_date_lt":"string","sign_date_lte":"string","sign_date_not":"string","sign_date_not_in":["string"],"software_edition":"COMMUNITY","software_edition_in":["COMMUNITY"],"software_edition_not":"COMMUNITY","software_edition_not_in":["COMMUNITY"],"type":"PERPETUAL","type_in":["PERPETUAL"],"type_not":"PERPETUAL","type_not_in":["PERPETUAL"]}
//
// swagger:model BackupLicenseWhereInput
type BackupLicenseWhereInput struct {

	// a n d
	AND []*BackupLicenseWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*BackupLicenseWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*BackupLicenseWhereInput `json:"OR,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// expire date
	ExpireDate *string `json:"expire_date,omitempty"`

	// expire date gt
	ExpireDateGt *string `json:"expire_date_gt,omitempty"`

	// expire date gte
	ExpireDateGte *string `json:"expire_date_gte,omitempty"`

	// expire date in
	ExpireDateIn []string `json:"expire_date_in,omitempty"`

	// expire date lt
	ExpireDateLt *string `json:"expire_date_lt,omitempty"`

	// expire date lte
	ExpireDateLte *string `json:"expire_date_lte,omitempty"`

	// expire date not
	ExpireDateNot *string `json:"expire_date_not,omitempty"`

	// expire date not in
	ExpireDateNotIn []string `json:"expire_date_not_in,omitempty"`

	// id
	ID *string `json:"id,omitempty"`

	// id contains
	IDContains *string `json:"id_contains,omitempty"`

	// id ends with
	IDEndsWith *string `json:"id_ends_with,omitempty"`

	// id gt
	IDGt *string `json:"id_gt,omitempty"`

	// id gte
	IDGte *string `json:"id_gte,omitempty"`

	// id in
	IDIn []string `json:"id_in,omitempty"`

	// id lt
	IDLt *string `json:"id_lt,omitempty"`

	// id lte
	IDLte *string `json:"id_lte,omitempty"`

	// id not
	IDNot *string `json:"id_not,omitempty"`

	// id not contains
	IDNotContains *string `json:"id_not_contains,omitempty"`

	// id not ends with
	IDNotEndsWith *string `json:"id_not_ends_with,omitempty"`

	// id not in
	IDNotIn []string `json:"id_not_in,omitempty"`

	// id not starts with
	IDNotStartsWith *string `json:"id_not_starts_with,omitempty"`

	// id starts with
	IDStartsWith *string `json:"id_starts_with,omitempty"`

	// license serial
	LicenseSerial *string `json:"license_serial,omitempty"`

	// license serial contains
	LicenseSerialContains *string `json:"license_serial_contains,omitempty"`

	// license serial ends with
	LicenseSerialEndsWith *string `json:"license_serial_ends_with,omitempty"`

	// license serial gt
	LicenseSerialGt *string `json:"license_serial_gt,omitempty"`

	// license serial gte
	LicenseSerialGte *string `json:"license_serial_gte,omitempty"`

	// license serial in
	LicenseSerialIn []string `json:"license_serial_in,omitempty"`

	// license serial lt
	LicenseSerialLt *string `json:"license_serial_lt,omitempty"`

	// license serial lte
	LicenseSerialLte *string `json:"license_serial_lte,omitempty"`

	// license serial not
	LicenseSerialNot *string `json:"license_serial_not,omitempty"`

	// license serial not contains
	LicenseSerialNotContains *string `json:"license_serial_not_contains,omitempty"`

	// license serial not ends with
	LicenseSerialNotEndsWith *string `json:"license_serial_not_ends_with,omitempty"`

	// license serial not in
	LicenseSerialNotIn []string `json:"license_serial_not_in,omitempty"`

	// license serial not starts with
	LicenseSerialNotStartsWith *string `json:"license_serial_not_starts_with,omitempty"`

	// license serial starts with
	LicenseSerialStartsWith *string `json:"license_serial_starts_with,omitempty"`

	// max capacity
	MaxCapacity *float64 `json:"max_capacity,omitempty"`

	// max capacity gt
	MaxCapacityGt *float64 `json:"max_capacity_gt,omitempty"`

	// max capacity gte
	MaxCapacityGte *float64 `json:"max_capacity_gte,omitempty"`

	// max capacity in
	MaxCapacityIn []float64 `json:"max_capacity_in,omitempty"`

	// max capacity lt
	MaxCapacityLt *float64 `json:"max_capacity_lt,omitempty"`

	// max capacity lte
	MaxCapacityLte *float64 `json:"max_capacity_lte,omitempty"`

	// max capacity not
	MaxCapacityNot *float64 `json:"max_capacity_not,omitempty"`

	// max capacity not in
	MaxCapacityNotIn []float64 `json:"max_capacity_not_in,omitempty"`

	// sign date
	SignDate *string `json:"sign_date,omitempty"`

	// sign date gt
	SignDateGt *string `json:"sign_date_gt,omitempty"`

	// sign date gte
	SignDateGte *string `json:"sign_date_gte,omitempty"`

	// sign date in
	SignDateIn []string `json:"sign_date_in,omitempty"`

	// sign date lt
	SignDateLt *string `json:"sign_date_lt,omitempty"`

	// sign date lte
	SignDateLte *string `json:"sign_date_lte,omitempty"`

	// sign date not
	SignDateNot *string `json:"sign_date_not,omitempty"`

	// sign date not in
	SignDateNotIn []string `json:"sign_date_not_in,omitempty"`

	// software edition
	SoftwareEdition *SoftwareEdition `json:"software_edition,omitempty"`

	// software edition in
	SoftwareEditionIn []SoftwareEdition `json:"software_edition_in,omitempty"`

	// software edition not
	SoftwareEditionNot *SoftwareEdition `json:"software_edition_not,omitempty"`

	// software edition not in
	SoftwareEditionNotIn []SoftwareEdition `json:"software_edition_not_in,omitempty"`

	// type
	Type *LicenseType `json:"type,omitempty"`

	// type in
	TypeIn []LicenseType `json:"type_in,omitempty"`

	// type not
	TypeNot *LicenseType `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []LicenseType `json:"type_not_in,omitempty"`
}

// Validate validates this backup license where input
func (m *BackupLicenseWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityAsyncStatusNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEdition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEditionIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEditionNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareEditionNotIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeIn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeNotIn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupLicenseWhereInput) validateAND(formats strfmt.Registry) error {
	if swag.IsZero(m.AND) { // not required
		return nil
	}

	for i := 0; i < len(m.AND); i++ {
		if swag.IsZero(m.AND[i]) { // not required
			continue
		}

		if m.AND[i] != nil {
			if err := m.AND[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateNOT(formats strfmt.Registry) error {
	if swag.IsZero(m.NOT) { // not required
		return nil
	}

	for i := 0; i < len(m.NOT); i++ {
		if swag.IsZero(m.NOT[i]) { // not required
			continue
		}

		if m.NOT[i] != nil {
			if err := m.NOT[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateOR(formats strfmt.Registry) error {
	if swag.IsZero(m.OR) { // not required
		return nil
	}

	for i := 0; i < len(m.OR); i++ {
		if swag.IsZero(m.OR[i]) { // not required
			continue
		}

		if m.OR[i] != nil {
			if err := m.OR[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatus) { // not required
		return nil
	}

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNot) { // not required
		return nil
	}

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityAsyncStatusNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateSoftwareEdition(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEdition) { // not required
		return nil
	}

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) validateSoftwareEditionIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareEditionIn); i++ {

		if err := m.SoftwareEditionIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateSoftwareEditionNot(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionNot) { // not required
		return nil
	}

	if m.SoftwareEditionNot != nil {
		if err := m.SoftwareEditionNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) validateSoftwareEditionNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareEditionNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.SoftwareEditionNotIn); i++ {

		if err := m.SoftwareEditionNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) validateTypeIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) validateTypeNot(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNot) { // not required
		return nil
	}

	if m.TypeNot != nil {
		if err := m.TypeNot.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
	if swag.IsZero(m.TypeNotIn) { // not required
		return nil
	}

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this backup license where input based on the context it is used
func (m *BackupLicenseWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntityAsyncStatusNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEdition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEditionIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEditionNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareEditionNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNot(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTypeNotIn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupLicenseWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AND); i++ {

		if m.AND[i] != nil {
			if err := m.AND[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AND" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AND" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NOT); i++ {

		if m.NOT[i] != nil {
			if err := m.NOT[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("NOT" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("NOT" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OR); i++ {

		if m.OR[i] != nil {
			if err := m.OR[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("OR" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("OR" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatus != nil {
		if err := m.EntityAsyncStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusIn); i++ {

		if err := m.EntityAsyncStatusIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

	if m.EntityAsyncStatusNot != nil {
		if err := m.EntityAsyncStatusNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EntityAsyncStatusNotIn); i++ {

		if err := m.EntityAsyncStatusNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entityAsyncStatus_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateSoftwareEdition(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareEdition != nil {
		if err := m.SoftwareEdition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateSoftwareEditionIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SoftwareEditionIn); i++ {

		if err := m.SoftwareEditionIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateSoftwareEditionNot(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareEditionNot != nil {
		if err := m.SoftwareEditionNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateSoftwareEditionNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SoftwareEditionNotIn); i++ {

		if err := m.SoftwareEditionNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("software_edition_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeIn); i++ {

		if err := m.TypeIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateTypeNot(ctx context.Context, formats strfmt.Registry) error {

	if m.TypeNot != nil {
		if err := m.TypeNot.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not")
			}
			return err
		}
	}

	return nil
}

func (m *BackupLicenseWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TypeNotIn); i++ {

		if err := m.TypeNotIn[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type_not_in" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupLicenseWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupLicenseWhereInput) UnmarshalBinary(b []byte) error {
	var res BackupLicenseWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}