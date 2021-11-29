// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// DockerfileScanSeverity The permissible dockerfile scan severity in the pipeline
// swagger:model DockerfileScanSeverity
type DockerfileScanSeverity string

const (

	// DockerfileScanSeverityINFO captures enum value "INFO"
	DockerfileScanSeverityINFO DockerfileScanSeverity = "INFO"

	// DockerfileScanSeverityWARN captures enum value "WARN"
	DockerfileScanSeverityWARN DockerfileScanSeverity = "WARN"

	// DockerfileScanSeverityFATAL captures enum value "FATAL"
	DockerfileScanSeverityFATAL DockerfileScanSeverity = "FATAL"
)

// for schema
var dockerfileScanSeverityEnum []interface{}

func init() {
	var res []DockerfileScanSeverity
	if err := json.Unmarshal([]byte(`["INFO","WARN","FATAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dockerfileScanSeverityEnum = append(dockerfileScanSeverityEnum, v)
	}
}

func (m DockerfileScanSeverity) validateDockerfileScanSeverityEnum(path, location string, value DockerfileScanSeverity) error {
	if err := validate.Enum(path, location, value, dockerfileScanSeverityEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this dockerfile scan severity
func (m DockerfileScanSeverity) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDockerfileScanSeverityEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
