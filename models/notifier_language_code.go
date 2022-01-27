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

// NotifierLanguageCode notifier language code
//
// swagger:model NotifierLanguageCode
type NotifierLanguageCode string

func NewNotifierLanguageCode(value NotifierLanguageCode) *NotifierLanguageCode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated NotifierLanguageCode.
func (m NotifierLanguageCode) Pointer() *NotifierLanguageCode {
	return &m
}

const (

	// NotifierLanguageCodeENUS captures enum value "EN_US"
	NotifierLanguageCodeENUS NotifierLanguageCode = "EN_US"

	// NotifierLanguageCodeZHCN captures enum value "ZH_CN"
	NotifierLanguageCodeZHCN NotifierLanguageCode = "ZH_CN"
)

// for schema
var notifierLanguageCodeEnum []interface{}

func init() {
	var res []NotifierLanguageCode
	if err := json.Unmarshal([]byte(`["EN_US","ZH_CN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notifierLanguageCodeEnum = append(notifierLanguageCodeEnum, v)
	}
}

func (m NotifierLanguageCode) validateNotifierLanguageCodeEnum(path, location string, value NotifierLanguageCode) error {
	if err := validate.EnumCase(path, location, value, notifierLanguageCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this notifier language code
func (m NotifierLanguageCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNotifierLanguageCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this notifier language code based on context it is used
func (m NotifierLanguageCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}