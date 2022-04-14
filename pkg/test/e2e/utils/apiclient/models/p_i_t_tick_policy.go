// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// PITTickPolicy PITTickPolicy determines what happens when QEMU misses a deadline for injecting a tick to the guest.
//
// swagger:model PITTickPolicy
type PITTickPolicy string

// Validate validates this p i t tick policy
func (m PITTickPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this p i t tick policy based on context it is used
func (m PITTickPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}