// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubernetesEnvironment kubernetes environment
// swagger:model KubernetesEnvironment
type KubernetesEnvironment struct {

	// id
	// Format: uuid
	ID strfmt.UUID `json:"id,omitempty"`

	// kubernetes cluster
	// Required: true
	// Format: uuid
	KubernetesCluster *strfmt.UUID `json:"kubernetesCluster"`

	// kubernetes cluster name
	KubernetesClusterName string `json:"kubernetesClusterName,omitempty"`

	// namespace labels
	NamespaceLabels []*Label `json:"namespaceLabels"`

	// namespaces
	Namespaces []string `json:"namespaces"`

	// preserve original source Ip
	// Read Only: true
	PreserveOriginalSourceIP *bool `json:"preserveOriginalSourceIp,omitempty"`
}

// Validate validates this kubernetes environment
func (m *KubernetesEnvironment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespaceLabels(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesEnvironment) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesEnvironment) validateKubernetesCluster(formats strfmt.Registry) error {

	if err := validate.Required("kubernetesCluster", "body", m.KubernetesCluster); err != nil {
		return err
	}

	if err := validate.FormatOf("kubernetesCluster", "body", "uuid", m.KubernetesCluster.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *KubernetesEnvironment) validateNamespaceLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.NamespaceLabels) { // not required
		return nil
	}

	for i := 0; i < len(m.NamespaceLabels); i++ {
		if swag.IsZero(m.NamespaceLabels[i]) { // not required
			continue
		}

		if m.NamespaceLabels[i] != nil {
			if err := m.NamespaceLabels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("namespaceLabels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesEnvironment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesEnvironment) UnmarshalBinary(b []byte) error {
	var res KubernetesEnvironment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
