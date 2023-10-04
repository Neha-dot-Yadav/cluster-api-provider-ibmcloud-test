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
	"github.com/go-openapi/validate"
)

// DisasterRecoveryLocations disaster recovery locations
//
// swagger:model DisasterRecoveryLocations
type DisasterRecoveryLocations struct {

	// The list of Disaster Recovery Locations
	// Required: true
	DisasterRecoveryLocations []*DisasterRecoveryLocation `json:"disasterRecoveryLocations"`
}

// Validate validates this disaster recovery locations
func (m *DisasterRecoveryLocations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisasterRecoveryLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisasterRecoveryLocations) validateDisasterRecoveryLocations(formats strfmt.Registry) error {

	if err := validate.Required("disasterRecoveryLocations", "body", m.DisasterRecoveryLocations); err != nil {
		return err
	}

	for i := 0; i < len(m.DisasterRecoveryLocations); i++ {
		if swag.IsZero(m.DisasterRecoveryLocations[i]) { // not required
			continue
		}

		if m.DisasterRecoveryLocations[i] != nil {
			if err := m.DisasterRecoveryLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disasterRecoveryLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disasterRecoveryLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this disaster recovery locations based on the context it is used
func (m *DisasterRecoveryLocations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisasterRecoveryLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DisasterRecoveryLocations) contextValidateDisasterRecoveryLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisasterRecoveryLocations); i++ {

		if m.DisasterRecoveryLocations[i] != nil {

			if swag.IsZero(m.DisasterRecoveryLocations[i]) { // not required
				return nil
			}

			if err := m.DisasterRecoveryLocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disasterRecoveryLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disasterRecoveryLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DisasterRecoveryLocations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DisasterRecoveryLocations) UnmarshalBinary(b []byte) error {
	var res DisasterRecoveryLocations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
