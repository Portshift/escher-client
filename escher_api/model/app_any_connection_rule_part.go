// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AppAnyConnectionRulePart app any connection rule part
// swagger:model AppAnyConnectionRulePart
type AppAnyConnectionRulePart struct {

	// environments
	Environments []string `json:"environments"`
}

// ConnectionRulePartType gets the connection rule part type of this subtype
func (m *AppAnyConnectionRulePart) ConnectionRulePartType() string {
	return "AppAnyConnectionRulePart"
}

// SetConnectionRulePartType sets the connection rule part type of this subtype
func (m *AppAnyConnectionRulePart) SetConnectionRulePartType(val string) {

}

// Environments gets the environments of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *AppAnyConnectionRulePart) UnmarshalJSON(raw []byte) error {
	var data struct {

		// environments
		Environments []string `json:"environments"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		ConnectionRulePartType string `json:"connectionRulePartType"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result AppAnyConnectionRulePart

	if base.ConnectionRulePartType != result.ConnectionRulePartType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid connectionRulePartType value: %q", base.ConnectionRulePartType)
	}

	result.Environments = data.Environments

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m AppAnyConnectionRulePart) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// environments
		Environments []string `json:"environments"`
	}{

		Environments: m.Environments,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		ConnectionRulePartType string `json:"connectionRulePartType"`
	}{

		ConnectionRulePartType: m.ConnectionRulePartType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this app any connection rule part
func (m *AppAnyConnectionRulePart) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AppAnyConnectionRulePart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppAnyConnectionRulePart) UnmarshalBinary(b []byte) error {
	var res AppAnyConnectionRulePart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}