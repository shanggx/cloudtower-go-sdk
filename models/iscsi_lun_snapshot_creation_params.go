// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IscsiLunSnapshotCreationParams iscsi lun snapshot creation params
//
// swagger:model IscsiLunSnapshotCreationParams
type IscsiLunSnapshotCreationParams struct {

	// iscsi lun id
	// Required: true
	IscsiLunID *string `json:"iscsi_lun_id"`

	// iscsi target id
	// Required: true
	IscsiTargetID *string `json:"iscsi_target_id"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this iscsi lun snapshot creation params
func (m *IscsiLunSnapshotCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIscsiLunID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiTargetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IscsiLunSnapshotCreationParams) validateIscsiLunID(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_lun_id", "body", m.IscsiLunID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLunSnapshotCreationParams) validateIscsiTargetID(formats strfmt.Registry) error {

	if err := validate.Required("iscsi_target_id", "body", m.IscsiTargetID); err != nil {
		return err
	}

	return nil
}

func (m *IscsiLunSnapshotCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this iscsi lun snapshot creation params based on context it is used
func (m *IscsiLunSnapshotCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IscsiLunSnapshotCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IscsiLunSnapshotCreationParams) UnmarshalBinary(b []byte) error {
	var res IscsiLunSnapshotCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}