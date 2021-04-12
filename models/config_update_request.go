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

// ConfigUpdateRequest config update request
// swagger:model ConfigUpdateRequest
type ConfigUpdateRequest struct {

	// Compatability Level
	// Enum: [BACKWARD BACKWARD_TRANSITIVE FORWARD FORWARD_TRANSITIVE FULL FULL_TRANSITIVE NONE]
	Compatibility string `json:"compatibility,omitempty"`
}

// Validate validates this config update request
func (m *ConfigUpdateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompatibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var configUpdateRequestTypeCompatibilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BACKWARD","BACKWARD_TRANSITIVE","FORWARD","FORWARD_TRANSITIVE","FULL","FULL_TRANSITIVE","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configUpdateRequestTypeCompatibilityPropEnum = append(configUpdateRequestTypeCompatibilityPropEnum, v)
	}
}

const (

	// ConfigUpdateRequestCompatibilityBACKWARD captures enum value "BACKWARD"
	ConfigUpdateRequestCompatibilityBACKWARD string = "BACKWARD"

	// ConfigUpdateRequestCompatibilityBACKWARDTRANSITIVE captures enum value "BACKWARD_TRANSITIVE"
	ConfigUpdateRequestCompatibilityBACKWARDTRANSITIVE string = "BACKWARD_TRANSITIVE"

	// ConfigUpdateRequestCompatibilityFORWARD captures enum value "FORWARD"
	ConfigUpdateRequestCompatibilityFORWARD string = "FORWARD"

	// ConfigUpdateRequestCompatibilityFORWARDTRANSITIVE captures enum value "FORWARD_TRANSITIVE"
	ConfigUpdateRequestCompatibilityFORWARDTRANSITIVE string = "FORWARD_TRANSITIVE"

	// ConfigUpdateRequestCompatibilityFULL captures enum value "FULL"
	ConfigUpdateRequestCompatibilityFULL string = "FULL"

	// ConfigUpdateRequestCompatibilityFULLTRANSITIVE captures enum value "FULL_TRANSITIVE"
	ConfigUpdateRequestCompatibilityFULLTRANSITIVE string = "FULL_TRANSITIVE"

	// ConfigUpdateRequestCompatibilityNONE captures enum value "NONE"
	ConfigUpdateRequestCompatibilityNONE string = "NONE"
)

// prop value enum
func (m *ConfigUpdateRequest) validateCompatibilityEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, configUpdateRequestTypeCompatibilityPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ConfigUpdateRequest) validateCompatibility(formats strfmt.Registry) error {

	if swag.IsZero(m.Compatibility) { // not required
		return nil
	}

	// value enum
	if err := m.validateCompatibilityEnum("compatibility", "body", m.Compatibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigUpdateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigUpdateRequest) UnmarshalBinary(b []byte) error {
	var res ConfigUpdateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
