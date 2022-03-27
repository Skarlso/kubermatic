// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VGPUDisplayOptions v g p u display options
//
// swagger:model VGPUDisplayOptions
type VGPUDisplayOptions struct {

	// Enabled determines if a display addapter backed by a vGPU should be enabled or disabled on the guest.
	// Defaults to true.
	// +optional
	Enabled bool `json:"enabled,omitempty"`

	// ram f b
	RAMFB *FeatureState `json:"ramFB,omitempty"`
}

// Validate validates this v g p u display options
func (m *VGPUDisplayOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRAMFB(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VGPUDisplayOptions) validateRAMFB(formats strfmt.Registry) error {
	if swag.IsZero(m.RAMFB) { // not required
		return nil
	}

	if m.RAMFB != nil {
		if err := m.RAMFB.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ramFB")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v g p u display options based on the context it is used
func (m *VGPUDisplayOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRAMFB(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VGPUDisplayOptions) contextValidateRAMFB(ctx context.Context, formats strfmt.Registry) error {

	if m.RAMFB != nil {
		if err := m.RAMFB.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ramFB")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VGPUDisplayOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VGPUDisplayOptions) UnmarshalBinary(b []byte) error {
	var res VGPUDisplayOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}