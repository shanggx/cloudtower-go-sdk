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

// ClusterUpgradeHistoryWhereInput cluster upgrade history where input
// Example: {"AND":"ClusterUpgradeHistoryWhereInput[]","NOT":"ClusterUpgradeHistoryWhereInput[]","OR":"ClusterUpgradeHistoryWhereInput[]","cluster":"ClusterWhereInput","date":"string","date_gt":"string","date_gte":"string","date_in":["string"],"date_lt":"string","date_lte":"string","date_not":"string","date_not_in":["string"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","result":"string","result_contains":"string","result_ends_with":"string","result_gt":"string","result_gte":"string","result_in":["string"],"result_lt":"string","result_lte":"string","result_not":"string","result_not_contains":"string","result_not_ends_with":"string","result_not_in":["string"],"result_not_starts_with":"string","result_starts_with":"string","task_uuid":"string","task_uuid_contains":"string","task_uuid_ends_with":"string","task_uuid_gt":"string","task_uuid_gte":"string","task_uuid_in":["string"],"task_uuid_lt":"string","task_uuid_lte":"string","task_uuid_not":"string","task_uuid_not_contains":"string","task_uuid_not_ends_with":"string","task_uuid_not_in":["string"],"task_uuid_not_starts_with":"string","task_uuid_starts_with":"string","version":"string","version_contains":"string","version_ends_with":"string","version_gt":"string","version_gte":"string","version_in":["string"],"version_lt":"string","version_lte":"string","version_not":"string","version_not_contains":"string","version_not_ends_with":"string","version_not_in":["string"],"version_not_starts_with":"string","version_starts_with":"string"}
//
// swagger:model ClusterUpgradeHistoryWhereInput
type ClusterUpgradeHistoryWhereInput struct {

	// a n d
	AND []*ClusterUpgradeHistoryWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*ClusterUpgradeHistoryWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*ClusterUpgradeHistoryWhereInput `json:"OR,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// date
	Date *string `json:"date,omitempty"`

	// date gt
	DateGt *string `json:"date_gt,omitempty"`

	// date gte
	DateGte *string `json:"date_gte,omitempty"`

	// date in
	DateIn []string `json:"date_in,omitempty"`

	// date lt
	DateLt *string `json:"date_lt,omitempty"`

	// date lte
	DateLte *string `json:"date_lte,omitempty"`

	// date not
	DateNot *string `json:"date_not,omitempty"`

	// date not in
	DateNotIn []string `json:"date_not_in,omitempty"`

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

	// local id
	LocalID *string `json:"local_id,omitempty"`

	// local id contains
	LocalIDContains *string `json:"local_id_contains,omitempty"`

	// local id ends with
	LocalIDEndsWith *string `json:"local_id_ends_with,omitempty"`

	// local id gt
	LocalIDGt *string `json:"local_id_gt,omitempty"`

	// local id gte
	LocalIDGte *string `json:"local_id_gte,omitempty"`

	// local id in
	LocalIDIn []string `json:"local_id_in,omitempty"`

	// local id lt
	LocalIDLt *string `json:"local_id_lt,omitempty"`

	// local id lte
	LocalIDLte *string `json:"local_id_lte,omitempty"`

	// local id not
	LocalIDNot *string `json:"local_id_not,omitempty"`

	// local id not contains
	LocalIDNotContains *string `json:"local_id_not_contains,omitempty"`

	// local id not ends with
	LocalIDNotEndsWith *string `json:"local_id_not_ends_with,omitempty"`

	// local id not in
	LocalIDNotIn []string `json:"local_id_not_in,omitempty"`

	// local id not starts with
	LocalIDNotStartsWith *string `json:"local_id_not_starts_with,omitempty"`

	// local id starts with
	LocalIDStartsWith *string `json:"local_id_starts_with,omitempty"`

	// result
	Result *string `json:"result,omitempty"`

	// result contains
	ResultContains *string `json:"result_contains,omitempty"`

	// result ends with
	ResultEndsWith *string `json:"result_ends_with,omitempty"`

	// result gt
	ResultGt *string `json:"result_gt,omitempty"`

	// result gte
	ResultGte *string `json:"result_gte,omitempty"`

	// result in
	ResultIn []string `json:"result_in,omitempty"`

	// result lt
	ResultLt *string `json:"result_lt,omitempty"`

	// result lte
	ResultLte *string `json:"result_lte,omitempty"`

	// result not
	ResultNot *string `json:"result_not,omitempty"`

	// result not contains
	ResultNotContains *string `json:"result_not_contains,omitempty"`

	// result not ends with
	ResultNotEndsWith *string `json:"result_not_ends_with,omitempty"`

	// result not in
	ResultNotIn []string `json:"result_not_in,omitempty"`

	// result not starts with
	ResultNotStartsWith *string `json:"result_not_starts_with,omitempty"`

	// result starts with
	ResultStartsWith *string `json:"result_starts_with,omitempty"`

	// task uuid
	TaskUUID *string `json:"task_uuid,omitempty"`

	// task uuid contains
	TaskUUIDContains *string `json:"task_uuid_contains,omitempty"`

	// task uuid ends with
	TaskUUIDEndsWith *string `json:"task_uuid_ends_with,omitempty"`

	// task uuid gt
	TaskUUIDGt *string `json:"task_uuid_gt,omitempty"`

	// task uuid gte
	TaskUUIDGte *string `json:"task_uuid_gte,omitempty"`

	// task uuid in
	TaskUUIDIn []string `json:"task_uuid_in,omitempty"`

	// task uuid lt
	TaskUUIDLt *string `json:"task_uuid_lt,omitempty"`

	// task uuid lte
	TaskUUIDLte *string `json:"task_uuid_lte,omitempty"`

	// task uuid not
	TaskUUIDNot *string `json:"task_uuid_not,omitempty"`

	// task uuid not contains
	TaskUUIDNotContains *string `json:"task_uuid_not_contains,omitempty"`

	// task uuid not ends with
	TaskUUIDNotEndsWith *string `json:"task_uuid_not_ends_with,omitempty"`

	// task uuid not in
	TaskUUIDNotIn []string `json:"task_uuid_not_in,omitempty"`

	// task uuid not starts with
	TaskUUIDNotStartsWith *string `json:"task_uuid_not_starts_with,omitempty"`

	// task uuid starts with
	TaskUUIDStartsWith *string `json:"task_uuid_starts_with,omitempty"`

	// version
	Version *string `json:"version,omitempty"`

	// version contains
	VersionContains *string `json:"version_contains,omitempty"`

	// version ends with
	VersionEndsWith *string `json:"version_ends_with,omitempty"`

	// version gt
	VersionGt *string `json:"version_gt,omitempty"`

	// version gte
	VersionGte *string `json:"version_gte,omitempty"`

	// version in
	VersionIn []string `json:"version_in,omitempty"`

	// version lt
	VersionLt *string `json:"version_lt,omitempty"`

	// version lte
	VersionLte *string `json:"version_lte,omitempty"`

	// version not
	VersionNot *string `json:"version_not,omitempty"`

	// version not contains
	VersionNotContains *string `json:"version_not_contains,omitempty"`

	// version not ends with
	VersionNotEndsWith *string `json:"version_not_ends_with,omitempty"`

	// version not in
	VersionNotIn []string `json:"version_not_in,omitempty"`

	// version not starts with
	VersionNotStartsWith *string `json:"version_not_starts_with,omitempty"`

	// version starts with
	VersionStartsWith *string `json:"version_starts_with,omitempty"`
}

// Validate validates this cluster upgrade history where input
func (m *ClusterUpgradeHistoryWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpgradeHistoryWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *ClusterUpgradeHistoryWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *ClusterUpgradeHistoryWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *ClusterUpgradeHistoryWhereInput) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster upgrade history where input based on the context it is used
func (m *ClusterUpgradeHistoryWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUpgradeHistoryWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterUpgradeHistoryWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterUpgradeHistoryWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterUpgradeHistoryWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUpgradeHistoryWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUpgradeHistoryWhereInput) UnmarshalBinary(b []byte) error {
	var res ClusterUpgradeHistoryWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}