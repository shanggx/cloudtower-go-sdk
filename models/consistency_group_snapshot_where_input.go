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

// ConsistencyGroupSnapshotWhereInput consistency group snapshot where input
// Example: {"AND":"ConsistencyGroupSnapshotWhereInput[]","Iscsi_lun_snapshots_every":"IscsiLunSnapshotWhereInput","Iscsi_lun_snapshots_none":"IscsiLunSnapshotWhereInput","Iscsi_lun_snapshots_some":"IscsiLunSnapshotWhereInput","NOT":"ConsistencyGroupSnapshotWhereInput[]","OR":"ConsistencyGroupSnapshotWhereInput[]","consistency_group":"ConsistencyGroupWhereInput","entityAsyncStatus":"CREATING","entityAsyncStatus_in":["CREATING"],"entityAsyncStatus_not":"CREATING","entityAsyncStatus_not_in":["CREATING"],"id":"string","id_contains":"string","id_ends_with":"string","id_gt":"string","id_gte":"string","id_in":["string"],"id_lt":"string","id_lte":"string","id_not":"string","id_not_contains":"string","id_not_ends_with":"string","id_not_in":["string"],"id_not_starts_with":"string","id_starts_with":"string","labels_every":"LabelWhereInput","labels_none":"LabelWhereInput","labels_some":"LabelWhereInput","local_created_at":"string","local_created_at_gt":"string","local_created_at_gte":"string","local_created_at_in":["string"],"local_created_at_lt":"string","local_created_at_lte":"string","local_created_at_not":"string","local_created_at_not_in":["string"],"local_id":"string","local_id_contains":"string","local_id_ends_with":"string","local_id_gt":"string","local_id_gte":"string","local_id_in":["string"],"local_id_lt":"string","local_id_lte":"string","local_id_not":"string","local_id_not_contains":"string","local_id_not_ends_with":"string","local_id_not_in":["string"],"local_id_not_starts_with":"string","local_id_starts_with":"string","name":"string","name_contains":"string","name_ends_with":"string","name_gt":"string","name_gte":"string","name_in":["string"],"name_lt":"string","name_lte":"string","name_not":"string","name_not_contains":"string","name_not_ends_with":"string","name_not_in":["string"],"name_not_starts_with":"string","name_starts_with":"string","nvmf_namespace_snapshots_every":"NvmfNamespaceSnapshotWhereInput","nvmf_namespace_snapshots_none":"NvmfNamespaceSnapshotWhereInput","nvmf_namespace_snapshots_some":"NvmfNamespaceSnapshotWhereInput","unique_size":0,"unique_size_gt":0,"unique_size_gte":0,"unique_size_in":[0],"unique_size_lt":0,"unique_size_lte":0,"unique_size_not":0,"unique_size_not_in":[0]}
//
// swagger:model ConsistencyGroupSnapshotWhereInput
type ConsistencyGroupSnapshotWhereInput struct {

	// a n d
	AND []*ConsistencyGroupSnapshotWhereInput `json:"AND,omitempty"`

	// iscsi lun snapshots every
	IscsiLunSnapshotsEvery *IscsiLunSnapshotWhereInput `json:"Iscsi_lun_snapshots_every,omitempty"`

	// iscsi lun snapshots none
	IscsiLunSnapshotsNone *IscsiLunSnapshotWhereInput `json:"Iscsi_lun_snapshots_none,omitempty"`

	// iscsi lun snapshots some
	IscsiLunSnapshotsSome *IscsiLunSnapshotWhereInput `json:"Iscsi_lun_snapshots_some,omitempty"`

	// n o t
	NOT []*ConsistencyGroupSnapshotWhereInput `json:"NOT,omitempty"`

	// o r
	OR []*ConsistencyGroupSnapshotWhereInput `json:"OR,omitempty"`

	// consistency group
	ConsistencyGroup *ConsistencyGroupWhereInput `json:"consistency_group,omitempty"`

	// entity async status
	EntityAsyncStatus *EntityAsyncStatus `json:"entityAsyncStatus,omitempty"`

	// entity async status in
	EntityAsyncStatusIn []EntityAsyncStatus `json:"entityAsyncStatus_in,omitempty"`

	// entity async status not
	EntityAsyncStatusNot *EntityAsyncStatus `json:"entityAsyncStatus_not,omitempty"`

	// entity async status not in
	EntityAsyncStatusNotIn []EntityAsyncStatus `json:"entityAsyncStatus_not_in,omitempty"`

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

	// labels every
	LabelsEvery *LabelWhereInput `json:"labels_every,omitempty"`

	// labels none
	LabelsNone *LabelWhereInput `json:"labels_none,omitempty"`

	// labels some
	LabelsSome *LabelWhereInput `json:"labels_some,omitempty"`

	// local created at
	LocalCreatedAt *string `json:"local_created_at,omitempty"`

	// local created at gt
	LocalCreatedAtGt *string `json:"local_created_at_gt,omitempty"`

	// local created at gte
	LocalCreatedAtGte *string `json:"local_created_at_gte,omitempty"`

	// local created at in
	LocalCreatedAtIn []string `json:"local_created_at_in,omitempty"`

	// local created at lt
	LocalCreatedAtLt *string `json:"local_created_at_lt,omitempty"`

	// local created at lte
	LocalCreatedAtLte *string `json:"local_created_at_lte,omitempty"`

	// local created at not
	LocalCreatedAtNot *string `json:"local_created_at_not,omitempty"`

	// local created at not in
	LocalCreatedAtNotIn []string `json:"local_created_at_not_in,omitempty"`

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

	// nvmf namespace snapshots every
	NvmfNamespaceSnapshotsEvery *NvmfNamespaceSnapshotWhereInput `json:"nvmf_namespace_snapshots_every,omitempty"`

	// nvmf namespace snapshots none
	NvmfNamespaceSnapshotsNone *NvmfNamespaceSnapshotWhereInput `json:"nvmf_namespace_snapshots_none,omitempty"`

	// nvmf namespace snapshots some
	NvmfNamespaceSnapshotsSome *NvmfNamespaceSnapshotWhereInput `json:"nvmf_namespace_snapshots_some,omitempty"`

	// unique size
	UniqueSize *float64 `json:"unique_size,omitempty"`

	// unique size gt
	UniqueSizeGt *float64 `json:"unique_size_gt,omitempty"`

	// unique size gte
	UniqueSizeGte *float64 `json:"unique_size_gte,omitempty"`

	// unique size in
	UniqueSizeIn []float64 `json:"unique_size_in,omitempty"`

	// unique size lt
	UniqueSizeLt *float64 `json:"unique_size_lt,omitempty"`

	// unique size lte
	UniqueSizeLte *float64 `json:"unique_size_lte,omitempty"`

	// unique size not
	UniqueSizeNot *float64 `json:"unique_size_not,omitempty"`

	// unique size not in
	UniqueSizeNotIn []float64 `json:"unique_size_not_in,omitempty"`
}

// Validate validates this consistency group snapshot where input
func (m *ConsistencyGroupSnapshotWhereInput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAND(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiLunSnapshotsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiLunSnapshotsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIscsiLunSnapshotsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNOT(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOR(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsistencyGroup(formats); err != nil {
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

	if err := m.validateLabelsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabelsSome(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfNamespaceSnapshotsEvery(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfNamespaceSnapshotsNone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNvmfNamespaceSnapshotsSome(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) validateAND(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateIscsiLunSnapshotsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiLunSnapshotsEvery) { // not required
		return nil
	}

	if m.IscsiLunSnapshotsEvery != nil {
		if err := m.IscsiLunSnapshotsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Iscsi_lun_snapshots_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Iscsi_lun_snapshots_every")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) validateIscsiLunSnapshotsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiLunSnapshotsNone) { // not required
		return nil
	}

	if m.IscsiLunSnapshotsNone != nil {
		if err := m.IscsiLunSnapshotsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Iscsi_lun_snapshots_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Iscsi_lun_snapshots_none")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) validateIscsiLunSnapshotsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.IscsiLunSnapshotsSome) { // not required
		return nil
	}

	if m.IscsiLunSnapshotsSome != nil {
		if err := m.IscsiLunSnapshotsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Iscsi_lun_snapshots_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Iscsi_lun_snapshots_some")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) validateNOT(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateOR(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateConsistencyGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.ConsistencyGroup) { // not required
		return nil
	}

	if m.ConsistencyGroup != nil {
		if err := m.ConsistencyGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) validateEntityAsyncStatus(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateEntityAsyncStatusIn(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateEntityAsyncStatusNot(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateEntityAsyncStatusNotIn(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateLabelsEvery(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateLabelsNone(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateLabelsSome(formats strfmt.Registry) error {
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

func (m *ConsistencyGroupSnapshotWhereInput) validateNvmfNamespaceSnapshotsEvery(formats strfmt.Registry) error {
	if swag.IsZero(m.NvmfNamespaceSnapshotsEvery) { // not required
		return nil
	}

	if m.NvmfNamespaceSnapshotsEvery != nil {
		if err := m.NvmfNamespaceSnapshotsEvery.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_namespace_snapshots_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_namespace_snapshots_every")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) validateNvmfNamespaceSnapshotsNone(formats strfmt.Registry) error {
	if swag.IsZero(m.NvmfNamespaceSnapshotsNone) { // not required
		return nil
	}

	if m.NvmfNamespaceSnapshotsNone != nil {
		if err := m.NvmfNamespaceSnapshotsNone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_namespace_snapshots_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_namespace_snapshots_none")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) validateNvmfNamespaceSnapshotsSome(formats strfmt.Registry) error {
	if swag.IsZero(m.NvmfNamespaceSnapshotsSome) { // not required
		return nil
	}

	if m.NvmfNamespaceSnapshotsSome != nil {
		if err := m.NvmfNamespaceSnapshotsSome.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_namespace_snapshots_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_namespace_snapshots_some")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group snapshot where input based on the context it is used
func (m *ConsistencyGroupSnapshotWhereInput) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAND(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiLunSnapshotsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiLunSnapshotsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIscsiLunSnapshotsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNOT(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOR(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConsistencyGroup(ctx, formats); err != nil {
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

	if err := m.contextValidateLabelsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLabelsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfNamespaceSnapshotsEvery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfNamespaceSnapshotsNone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNvmfNamespaceSnapshotsSome(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateAND(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateIscsiLunSnapshotsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.IscsiLunSnapshotsEvery != nil {
		if err := m.IscsiLunSnapshotsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Iscsi_lun_snapshots_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Iscsi_lun_snapshots_every")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateIscsiLunSnapshotsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.IscsiLunSnapshotsNone != nil {
		if err := m.IscsiLunSnapshotsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Iscsi_lun_snapshots_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Iscsi_lun_snapshots_none")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateIscsiLunSnapshotsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.IscsiLunSnapshotsSome != nil {
		if err := m.IscsiLunSnapshotsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Iscsi_lun_snapshots_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Iscsi_lun_snapshots_some")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateNOT(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateOR(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateConsistencyGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.ConsistencyGroup != nil {
		if err := m.ConsistencyGroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("consistency_group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("consistency_group")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateEntityAsyncStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateEntityAsyncStatusIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateEntityAsyncStatusNot(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateEntityAsyncStatusNotIn(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateLabelsEvery(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateLabelsNone(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateLabelsSome(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateNvmfNamespaceSnapshotsEvery(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfNamespaceSnapshotsEvery != nil {
		if err := m.NvmfNamespaceSnapshotsEvery.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_namespace_snapshots_every")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_namespace_snapshots_every")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateNvmfNamespaceSnapshotsNone(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfNamespaceSnapshotsNone != nil {
		if err := m.NvmfNamespaceSnapshotsNone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_namespace_snapshots_none")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_namespace_snapshots_none")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupSnapshotWhereInput) contextValidateNvmfNamespaceSnapshotsSome(ctx context.Context, formats strfmt.Registry) error {

	if m.NvmfNamespaceSnapshotsSome != nil {
		if err := m.NvmfNamespaceSnapshotsSome.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nvmf_namespace_snapshots_some")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nvmf_namespace_snapshots_some")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotWhereInput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupSnapshotWhereInput) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupSnapshotWhereInput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}