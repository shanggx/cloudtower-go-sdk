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

// NvmfNamespaceCreationParams nvmf namespace creation params
//
// swagger:model NvmfNamespaceCreationParams
type NvmfNamespaceCreationParams struct {

	// assigned size
	// Required: true
	AssignedSize *float64 `json:"assigned_size"`

	// group id
	GroupID *string `json:"group_id,omitempty"`

	// is shared
	IsShared *bool `json:"is_shared,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// namespace id
	NamespaceID *int32 `json:"namespace_id,omitempty"`

	// nvmf subsystem id
	// Required: true
	NvmfSubsystemID *string `json:"nvmf_subsystem_id"`

	// replica num
	// Required: true
	ReplicaNum *int32 `json:"replica_num"`

	NvmfNamespaceCommonParams
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NvmfNamespaceCreationParams) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		AssignedSize *float64 `json:"assigned_size"`

		GroupID *string `json:"group_id,omitempty"`

		IsShared *bool `json:"is_shared,omitempty"`

		Name *string `json:"name"`

		NamespaceID *int32 `json:"namespace_id,omitempty"`

		NvmfSubsystemID *string `json:"nvmf_subsystem_id"`

		ReplicaNum *int32 `json:"replica_num"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.AssignedSize = dataAO0.AssignedSize

	m.GroupID = dataAO0.GroupID

	m.IsShared = dataAO0.IsShared

	m.Name = dataAO0.Name

	m.NamespaceID = dataAO0.NamespaceID

	m.NvmfSubsystemID = dataAO0.NvmfSubsystemID

	m.ReplicaNum = dataAO0.ReplicaNum

	// AO1
	var aO1 NvmfNamespaceCommonParams
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NvmfNamespaceCommonParams = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NvmfNamespaceCreationParams) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		AssignedSize *float64 `json:"assigned_size"`

		GroupID *string `json:"group_id,omitempty"`

		IsShared *bool `json:"is_shared,omitempty"`

		Name *string `json:"name"`

		NamespaceID *int32 `json:"namespace_id,omitempty"`

		NvmfSubsystemID *string `json:"nvmf_subsystem_id"`

		ReplicaNum *int32 `json:"replica_num"`
	}

	dataAO0.AssignedSize = m.AssignedSize

	dataAO0.GroupID = m.GroupID

	dataAO0.IsShared = m.IsShared

	dataAO0.Name = m.Name

	dataAO0.NamespaceID = m.NamespaceID

	dataAO0.NvmfSubsystemID = m.NvmfSubsystemID

	dataAO0.ReplicaNum = m.ReplicaNum

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	aO1, err := swag.WriteJSON(m.NvmfNamespaceCommonParams)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this nvmf namespace creation params
func (m *NvmfNamespaceCreationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfSubsystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReplicaNum(formats); err != nil {
		res = append(res, err)
	}

	// validation for a type composition with NvmfNamespaceCommonParams
	if err := m.NvmfNamespaceCommonParams.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NvmfNamespaceCreationParams) validateAssignedSize(formats strfmt.Registry) error {

	if err := validate.Required("assigned_size", "body", m.AssignedSize); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespaceCreationParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespaceCreationParams) validateNvmfSubsystemID(formats strfmt.Registry) error {

	if err := validate.Required("nvmf_subsystem_id", "body", m.NvmfSubsystemID); err != nil {
		return err
	}

	return nil
}

func (m *NvmfNamespaceCreationParams) validateReplicaNum(formats strfmt.Registry) error {

	if err := validate.Required("replica_num", "body", m.ReplicaNum); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nvmf namespace creation params based on the context it is used
func (m *NvmfNamespaceCreationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with NvmfNamespaceCommonParams
	if err := m.NvmfNamespaceCommonParams.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NvmfNamespaceCreationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NvmfNamespaceCreationParams) UnmarshalBinary(b []byte) error {
	var res NvmfNamespaceCreationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
