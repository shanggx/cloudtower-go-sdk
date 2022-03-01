// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertNotifierUpdationParams alert notifier updation params
//
// swagger:model AlertNotifierUpdationParams
type AlertNotifierUpdationParams struct {

	// disabled
	Disabled *bool `json:"disabled,omitempty"`

	// email from
	EmailFrom *string `json:"email_from,omitempty"`

	// email tos
	EmailTos []string `json:"email_tos,omitempty"`

	// language code
	LanguageCode *NotifierLanguageCode `json:"language_code,omitempty"`

	// notice severities
	NoticeSeverities []string `json:"notice_severities,omitempty"`
}

// Validate validates this alert notifier updation params
func (m *AlertNotifierUpdationParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLanguageCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoticeSeverities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertNotifierUpdationParams) validateLanguageCode(formats strfmt.Registry) error {
	if swag.IsZero(m.LanguageCode) { // not required
		return nil
	}

	if m.LanguageCode != nil {
		if err := m.LanguageCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

var alertNotifierUpdationParamsNoticeSeveritiesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CRITICAL","NOTICE","INFO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		alertNotifierUpdationParamsNoticeSeveritiesItemsEnum = append(alertNotifierUpdationParamsNoticeSeveritiesItemsEnum, v)
	}
}

func (m *AlertNotifierUpdationParams) validateNoticeSeveritiesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, alertNotifierUpdationParamsNoticeSeveritiesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AlertNotifierUpdationParams) validateNoticeSeverities(formats strfmt.Registry) error {
	if swag.IsZero(m.NoticeSeverities) { // not required
		return nil
	}

	for i := 0; i < len(m.NoticeSeverities); i++ {

		// value enum
		if err := m.validateNoticeSeveritiesItemsEnum("notice_severities"+"."+strconv.Itoa(i), "body", m.NoticeSeverities[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validate this alert notifier updation params based on the context it is used
func (m *AlertNotifierUpdationParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLanguageCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertNotifierUpdationParams) contextValidateLanguageCode(ctx context.Context, formats strfmt.Registry) error {

	if m.LanguageCode != nil {
		if err := m.LanguageCode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("language_code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("language_code")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertNotifierUpdationParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertNotifierUpdationParams) UnmarshalBinary(b []byte) error {
	var res AlertNotifierUpdationParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
