// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PodSecurityPolicyValidation When the rule action is ALLOW, onViolationAction is mandatory
// swagger:model PodSecurityPolicyValidation
type PodSecurityPolicyValidation struct {

	// on violation action
	OnViolationAction OnViolationAction `json:"onViolationAction,omitempty"`

	// pod security policy Id
	// Required: true
	// Format: uuid
	PodSecurityPolicyID *strfmt.UUID `json:"podSecurityPolicyId"`

	// pod security policy name
	PodSecurityPolicyName string `json:"podSecurityPolicyName"`

	// should mutate
	ShouldMutate *bool `json:"shouldMutate,omitempty"`
}

// Validate validates this pod security policy validation
func (m *PodSecurityPolicyValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOnViolationAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodSecurityPolicyID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodSecurityPolicyValidation) validateOnViolationAction(formats strfmt.Registry) error {

	if swag.IsZero(m.OnViolationAction) { // not required
		return nil
	}

	if err := m.OnViolationAction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("onViolationAction")
		}
		return err
	}

	return nil
}

func (m *PodSecurityPolicyValidation) validatePodSecurityPolicyID(formats strfmt.Registry) error {

	if err := validate.Required("podSecurityPolicyId", "body", m.PodSecurityPolicyID); err != nil {
		return err
	}

	if err := validate.FormatOf("podSecurityPolicyId", "body", "uuid", m.PodSecurityPolicyID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodSecurityPolicyValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodSecurityPolicyValidation) UnmarshalBinary(b []byte) error {
	var res PodSecurityPolicyValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
