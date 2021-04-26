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

// CloudConnectionUpdate cloud connection update
// swagger:model CloudConnectionUpdate
type CloudConnectionUpdate struct {

	// classic
	Classic *CloudConnectionEndpointClassicUpdate `json:"classic,omitempty"`

	// enable global routing for this cloud connection (default=false)
	GlobalRouting *bool `json:"globalRouting,omitempty"`

	// enable metered for this cloud connection (default=false)
	Metered *bool `json:"metered,omitempty"`

	// name of the cloud connection
	Name *string `json:"name,omitempty"`

	// speed of the cloud connection (speed in megabits per second)
	// Enum: [50 100 200 500 1000 2000 5000 10000]
	Speed *int64 `json:"speed,omitempty"`

	// vpc
	Vpc *CloudConnectionEndpointVPC `json:"vpc,omitempty"`
}

// Validate validates this cloud connection update
func (m *CloudConnectionUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClassic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVpc(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudConnectionUpdate) validateClassic(formats strfmt.Registry) error {

	if swag.IsZero(m.Classic) { // not required
		return nil
	}

	if m.Classic != nil {
		if err := m.Classic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("classic")
			}
			return err
		}
	}

	return nil
}

var cloudConnectionUpdateTypeSpeedPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[50,100,200,500,1000,2000,5000,10000]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloudConnectionUpdateTypeSpeedPropEnum = append(cloudConnectionUpdateTypeSpeedPropEnum, v)
	}
}

// prop value enum
func (m *CloudConnectionUpdate) validateSpeedEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, cloudConnectionUpdateTypeSpeedPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CloudConnectionUpdate) validateSpeed(formats strfmt.Registry) error {

	if swag.IsZero(m.Speed) { // not required
		return nil
	}

	// value enum
	if err := m.validateSpeedEnum("speed", "body", *m.Speed); err != nil {
		return err
	}

	return nil
}

func (m *CloudConnectionUpdate) validateVpc(formats strfmt.Registry) error {

	if swag.IsZero(m.Vpc) { // not required
		return nil
	}

	if m.Vpc != nil {
		if err := m.Vpc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vpc")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudConnectionUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudConnectionUpdate) UnmarshalBinary(b []byte) error {
	var res CloudConnectionUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
