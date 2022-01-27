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

// VdsWhereInput vds where input
// Example: {"AND":"VdsWhereInput[]","NOT":"VdsWhereInput[]","OR":"VdsWhereInput[]","bond_mode":"string","bond_mode_contains":"string","bond_mode_ends_with":"string","bond_mode_gt":"string","bond_mode_gte":"string","bond_mode_in":["string"],"bond_mode_lt":"string","bond_mode_lte":"string","bond_mode_not":"string","bond_mode_not_contains":"string","bond_mode_not_ends_with":"string","bond_mode_not_in":["string"],"bond_mode_not_starts_with":"string","bond_mode_starts_with":"string","cluster":"ClusterWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"everoute_cluster":"EverouteClusterWhereInput","id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","internal":false,"internal_not":false,"labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","nics_every":"NicWhereInput","nics_none":"NicWhereInput","nics_some":"NicWhereInput","ovsbr_name":"string","ovsbr_name_contains":"string","ovsbr_name_ends_with":"string","ovsbr_name_gt":"string","ovsbr_name_gte":"string","ovsbr_name_in":["string"],"ovsbr_name_lt":"string","ovsbr_name_lte":"string","ovsbr_name_not":"string","ovsbr_name_not_contains":"string","ovsbr_name_not_ends_with":"string","ovsbr_name_not_in":["string"],"ovsbr_name_not_starts_with":"string","ovsbr_name_starts_with":"string","type":"ACCESS","type_in":["ACCESS"],"type_not":"ACCESS","type_not_in":["ACCESS"],"vlans_every":"VlanWhereInput","vlans_none":"VlanWhereInput","vlans_num":0,"vlans_num_gt":0,"vlans_num_gte":0,"vlans_num_in":[0],"vlans_num_lt":0,"vlans_num_lte":0,"vlans_num_not":0,"vlans_num_not_in":[0],"vlans_some":"VlanWhereInput","work_mode":"string","work_mode_contains":"string","work_mode_ends_with":"string","work_mode_gt":"string","work_mode_gte":"string","work_mode_in":["string"],"work_mode_lt":"string","work_mode_lte":"string","work_mode_not":"string","work_mode_not_contains":"string","work_mode_not_ends_with":"string","work_mode_not_in":["string"],"work_mode_not_starts_with":"string","work_mode_starts_with":"string"}
//
// swagger:model VdsWhereInput
type VdsWhereInput struct {

	// a n d
	AND []*VdsWhereInput `json:"AND,omitempty"`

	// n o t
	NOT []*VdsWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*VdsWhereInput `json:"OR,omitempty"`

	// bond mode
	BondMode *string `json:"bond_mode,omitempty"`

	// bond mode contains
	BondModeContains *string `json:"bond_mode_contains,omitempty"`

	// bond mode ends with
	BondModeEndsWith *string `json:"bond_mode_ends_with,omitempty"`

	// bond mode gt
	BondModeGt *string `json:"bond_mode_gt,omitempty"`

	// bond mode gte
	BondModeGte *string `json:"bond_mode_gte,omitempty"`

	// bond mode in
	BondModeIn []string `json:"bond_mode_in,omitempty"`

	// bond mode lt
	BondModeLt *string `json:"bond_mode_lt,omitempty"`

	// bond mode lte
	BondModeLte *string `json:"bond_mode_lte,omitempty"`

	// bond mode not
	BondModeNot *string `json:"bond_mode_not,omitempty"`

	// bond mode not contains
	BondModeNotContains *string `json:"bond_mode_not_contains,omitempty"`

	// bond mode not ends with
	BondModeNotEndsWith *string `json:"bond_mode_not_ends_with,omitempty"`

	// bond mode not in
	BondModeNotIn []string `json:"bond_mode_not_in,omitempty"`

	// bond mode not starts with
	BondModeNotStartsWith *string `json:"bond_mode_not_starts_with,omitempty"`

	// bond mode starts with
	BondModeStartsWith *string `json:"bond_mode_starts_with,omitempty"`

	// cluster
	Cluster *ClusterWhereInput `json:"cluster,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

	// everoute cluster
	EverouteCluster *EverouteClusterWhereInput `json:"everoute_cluster,omitempty"`

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

	// internal
	Internal *bool `json:"internal,omitempty"`

	// internal not
	InternalNot *bool `json:"internal_not,omitempty"`

	// labels every
	LabelsEvery *LabelWhereInput `json:"labels_every,omitempty"`

	// labels none
	LabelsNone *LabelWhereInput `json:"labels_none,omitempty"`

	// labels some
	LabelsSome *LabelWhereInput `json:"labels_some,omitempty"`

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

	// name
	Name *string `json:"name,omitempty"`

	// name contains
	NameContains *string `json:"name_contains,omitempty"`

	// name ends with
	NameEndsWith *string `json:"name_ends_with,omitempty"`

	// name gt
	NameGt *string `json:"name_gt,omitempty"`

	// name gte
	NameGte *string `json:"name_gte,omitempty"`

	// name in
	NameIn []string `json:"name_in,omitempty"`

	// name lt
	NameLt *string `json:"name_lt,omitempty"`

	// name lte
	NameLte *string `json:"name_lte,omitempty"`

	// name not
	NameNot *string `json:"name_not,omitempty"`

	// name not contains
	NameNotContains *string `json:"name_not_contains,omitempty"`

	// name not ends with
	NameNotEndsWith *string `json:"name_not_ends_with,omitempty"`

	// name not in
	NameNotIn []string `json:"name_not_in,omitempty"`

	// name not starts with
	NameNotStartsWith *string `json:"name_not_starts_with,omitempty"`

	// name starts with
	NameStartsWith *string `json:"name_starts_with,omitempty"`

	// nics every
	NicsEvery *NicWhereInput `json:"nics_every,omitempty"`

	// nics none
	NicsNone *NicWhereInput `json:"nics_none,omitempty"`

	// nics some
	NicsSome *NicWhereInput `json:"nics_some,omitempty"`

	// ovsbr name
	OvsbrName *string `json:"ovsbr_name,omitempty"`

	// ovsbr name contains
	OvsbrNameContains *string `json:"ovsbr_name_contains,omitempty"`

	// ovsbr name ends with
	OvsbrNameEndsWith *string `json:"ovsbr_name_ends_with,omitempty"`

	// ovsbr name gt
	OvsbrNameGt *string `json:"ovsbr_name_gt,omitempty"`

	// ovsbr name gte
	OvsbrNameGte *string `json:"ovsbr_name_gte,omitempty"`

	// ovsbr name in
	OvsbrNameIn []string `json:"ovsbr_name_in,omitempty"`

	// ovsbr name lt
	OvsbrNameLt *string `json:"ovsbr_name_lt,omitempty"`

	// ovsbr name lte
	OvsbrNameLte *string `json:"ovsbr_name_lte,omitempty"`

	// ovsbr name not
	OvsbrNameNot *string `json:"ovsbr_name_not,omitempty"`

	// ovsbr name not contains
	OvsbrNameNotContains *string `json:"ovsbr_name_not_contains,omitempty"`

	// ovsbr name not ends with
	OvsbrNameNotEndsWith *string `json:"ovsbr_name_not_ends_with,omitempty"`

	// ovsbr name not in
	OvsbrNameNotIn []string `json:"ovsbr_name_not_in,omitempty"`

	// ovsbr name not starts with
	OvsbrNameNotStartsWith *string `json:"ovsbr_name_not_starts_with,omitempty"`

	// ovsbr name starts with
	OvsbrNameStartsWith *string `json:"ovsbr_name_starts_with,omitempty"`

	// type
	Type *NetworkType `json:"type,omitempty"`

	// type in
	TypeIn []NetworkType `json:"type_in,omitempty"`

	// type not
	TypeNot *NetworkType `json:"type_not,omitempty"`

	// type not in
	TypeNotIn []NetworkType `json:"type_not_in,omitempty"`

	// vlans every
	VlansEvery *VlanWhereInput `json:"vlans_every,omitempty"`

	// vlans none
	VlansNone *VlanWhereInput `json:"vlans_none,omitempty"`

	// vlans num
	VlansNum *int32 `json:"vlans_num,omitempty"`

	// vlans num gt
	VlansNumGt *int32 `json:"vlans_num_gt,omitempty"`

	// vlans num gte
	VlansNumGte *int32 `json:"vlans_num_gte,omitempty"`

	// vlans num in
	VlansNumIn []int32 `json:"vlans_num_in,omitempty"`

	// vlans num lt
	VlansNumLt *int32 `json:"vlans_num_lt,omitempty"`

	// vlans num lte
	VlansNumLte *int32 `json:"vlans_num_lte,omitempty"`

	// vlans num not
	VlansNumNot *int32 `json:"vlans_num_not,omitempty"`

	// vlans num not in
	VlansNumNotIn []int32 `json:"vlans_num_not_in,omitempty"`

	// vlans some
	VlansSome *VlanWhereInput `json:"vlans_some,omitempty"`

	// work mode
	WorkMode *string `json:"work_mode,omitempty"`

	// work mode contains
	WorkModeContains *string `json:"work_mode_contains,omitempty"`

	// work mode ends with
	WorkModeEndsWith *string `json:"work_mode_ends_with,omitempty"`

	// work mode gt
	WorkModeGt *string `json:"work_mode_gt,omitempty"`

	// work mode gte
	WorkModeGte *string `json:"work_mode_gte,omitempty"`

	// work mode in
	WorkModeIn []string `json:"work_mode_in,omitempty"`

	// work mode lt
	WorkModeLt *string `json:"work_mode_lt,omitempty"`

	// work mode lte
	WorkModeLte *string `json:"work_mode_lte,omitempty"`

	// work mode not
	WorkModeNot *string `json:"work_mode_not,omitempty"`

	// work mode not contains
	WorkModeNotContains *string `json:"work_mode_not_contains,omitempty"`

	// work mode not ends with
	WorkModeNotEndsWith *string `json:"work_mode_not_ends_with,omitempty"`

	// work mode not in
	WorkModeNotIn []string `json:"work_mode_not_in,omitempty"`

	// work mode not starts with
	WorkModeNotStartsWith *string `json:"work_mode_not_starts_with,omitempty"`

	// work mode starts with
	WorkModeStartsWith *string `json:"work_mode_starts_with,omitempty"`
}

// Validate validates this vds where input
func (m *VdsWhereInput) Validate(formats strfmt.Registry) error {
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

	if err := m.validateEverouteCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNicsSome(formats); err != nil {
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

	if err := m.validateVlansEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlansNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlansSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VdsWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateCluster(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateEverouteCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.EverouteCluster) { // not required
		return nil
	}

	if m.EverouteCluster != nil {
		if err := m.EverouteCluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("everoute_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("everoute_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateLabelsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsEvery) { // not required
		return nil
	}

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateLabelsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsNone) { // not required
		return nil
	}

	if m.LabelsNone != nil {
		if err := m.LabelsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateLabelsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.LabelsSome) { // not required
		return nil
	}

	if m.LabelsSome != nil {
		if err := m.LabelsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateNicsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.NicsEvery) { // not required
		return nil
	}

	if m.NicsEvery != nil {
		if err := m.NicsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics_every")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateNicsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.NicsNone) { // not required
		return nil
	}

	if m.NicsNone != nil {
		if err := m.NicsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics_none")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateNicsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.NicsSome) { // not required
		return nil
	}

	if m.NicsSome != nil {
		if err := m.NicsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics_some")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateType(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateTypeIn(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateTypeNot(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateTypeNotIn(formats strfmt.Registry) error {
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

func (m *VdsWhereInput) validateVlansEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.VlansEvery) { // not required
		return nil
	}

	if m.VlansEvery != nil {
		if err := m.VlansEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans_every")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateVlansNone(formats strfmt.Registry) error {
	if swag.IsZero(m.VlansNone) { // not required
		return nil
	}

	if m.VlansNone != nil {
		if err := m.VlansNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans_none")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) validateVlansSome(formats strfmt.Registry) error {
	if swag.IsZero(m.VlansSome) { // not required
		return nil
	}

	if m.VlansSome != nil {
		if err := m.VlansSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vds where input based on the context it is used
func (m *VdsWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateEverouteCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNicsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNicsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNicsSome(ctx, formats); err != nil {
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

	if err := m.contextValidateVlansEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlansNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVlansSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VdsWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateEverouteCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.EverouteCluster != nil {
		if err := m.EverouteCluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("everoute_cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("everoute_cluster")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateLabelsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsEvery != nil {
		if err := m.LabelsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_every")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateLabelsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsNone != nil {
		if err := m.LabelsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_none")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateLabelsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.LabelsSome != nil {
		if err := m.LabelsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("labels_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("labels_some")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateNicsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.NicsEvery != nil {
		if err := m.NicsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics_every")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateNicsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.NicsNone != nil {
		if err := m.NicsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics_none")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateNicsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.NicsSome != nil {
		if err := m.NicsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nics_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nics_some")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateTypeIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateTypeNot(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateTypeNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *VdsWhereInput) contextValidateVlansEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.VlansEvery != nil {
		if err := m.VlansEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans_every")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateVlansNone(ctx context.Context, formats strfmt.Registry) error {

	if m.VlansNone != nil {
		if err := m.VlansNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans_none")
			}
			return err
		}
	}

	return nil
}

func (m *VdsWhereInput) contextValidateVlansSome(ctx context.Context, formats strfmt.Registry) error {

	if m.VlansSome != nil {
		if err := m.VlansSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vlans_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vlans_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VdsWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VdsWhereInput) UnmarshalBinary(b []byte) error {
	var res VdsWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}