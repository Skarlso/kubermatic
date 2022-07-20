// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// FormFieldType +kubebuilder:validation:Enum=number;boolean;text;text-area
//
// swagger:model FormFieldType
type FormFieldType string

// Validate validates this form field type
func (m FormFieldType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this form field type based on context it is used
func (m FormFieldType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}