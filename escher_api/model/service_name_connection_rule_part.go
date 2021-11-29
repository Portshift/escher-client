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

// ServiceNameConnectionRulePart service name connection rule part
// swagger:model ServiceNameConnectionRulePart
type ServiceNameConnectionRulePart struct {

	// environments
	Environments []string `json:"environments"`

	// services
	Services []string `json:"services"`
}

// ConnectionRulePartType gets the connection rule part type of this subtype
func (m *ServiceNameConnectionRulePart) ConnectionRulePartType() string {
	return "ServiceNameConnectionRulePart"
}

// SetConnectionRulePartType sets the connection rule part type of this subtype
func (m *ServiceNameConnectionRulePart) SetConnectionRulePartType(val string) {

}

// Environments gets the environments of this subtype

// Services gets the services of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *ServiceNameConnectionRulePart) UnmarshalJSON(raw []byte) error {
	var data struct {

		// environments
		Environments []string `json:"environments"`

		// services
		Services []string `json:"services"`
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

	var result ServiceNameConnectionRulePart

	if base.ConnectionRulePartType != result.ConnectionRulePartType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid connectionRulePartType value: %q", base.ConnectionRulePartType)
	}

	result.Environments = data.Environments

	result.Services = data.Services

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m ServiceNameConnectionRulePart) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// environments
		Environments []string `json:"environments"`

		// services
		Services []string `json:"services"`
	}{

		Environments: m.Environments,

		Services: m.Services,
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

// Validate validates this service name connection rule part
func (m *ServiceNameConnectionRulePart) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceNameConnectionRulePart) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceNameConnectionRulePart) UnmarshalBinary(b []byte) error {
	var res ServiceNameConnectionRulePart
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}