// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Config config
// swagger:model Config
type Config struct {

	// Compatability Level
	// Enum: [BACKWARD BACKWARD_TRANSITIVE FORWARD FORWARD_TRANSITIVE FULL FULL_TRANSITIVE NONE]
	CompatibilityLevel string `json:"compatibilityLevel,omitempty"`
}

// Validate validates this config
func (m *Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompatibilityLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var configTypeCompatibilityLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKWARD","BACKWARD_TRANSITIVE","FORWARD","FORWARD_TRANSITIVE","FULL","FULL_TRANSITIVE","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configTypeCompatibilityLevelPropEnum = append(configTypeCompatibilityLevelPropEnum, v)
	}
}

const (

	// ConfigCompatibilityLevelBACKWARD captures enum value "BACKWARD"
	ConfigCompatibilityLevelBACKWARD string = "BACKWARD"

	// ConfigCompatibilityLevelBACKWARDTRANSITIVE captures enum value "BACKWARD_TRANSITIVE"
	ConfigCompatibilityLevelBACKWARDTRANSITIVE string = "BACKWARD_TRANSITIVE"

	// ConfigCompatibilityLevelFORWARD captures enum value "FORWARD"
	ConfigCompatibilityLevelFORWARD string = "FORWARD"

	// ConfigCompatibilityLevelFORWARDTRANSITIVE captures enum value "FORWARD_TRANSITIVE"
	ConfigCompatibilityLevelFORWARDTRANSITIVE string = "FORWARD_TRANSITIVE"

	// ConfigCompatibilityLevelFULL captures enum value "FULL"
	ConfigCompatibilityLevelFULL string = "FULL"

	// ConfigCompatibilityLevelFULLTRANSITIVE captures enum value "FULL_TRANSITIVE"
	ConfigCompatibilityLevelFULLTRANSITIVE string = "FULL_TRANSITIVE"

	// ConfigCompatibilityLevelNONE captures enum value "NONE"
	ConfigCompatibilityLevelNONE string = "NONE"
)

// prop value enum
func (m *Config) validateCompatibilityLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, configTypeCompatibilityLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Config) validateCompatibilityLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.CompatibilityLevel) { // not required
		return nil
	}

	// value enum
	if err := m.validateCompatibilityLevelEnum("compatibilityLevel", "body", m.CompatibilityLevel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Config) UnmarshalBinary(b []byte) error {
	var res Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
