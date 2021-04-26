// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VolumesCloneExecute volumes clone execute
// swagger:model VolumesCloneExecute
type VolumesCloneExecute struct {

	// Base name of the new cloned volume(s).
	// Cloned Volume names will be prefixed with 'clone-'
	//     and suffixed with ‘-#####’ (where ##### is a 5 digit random number)
	// If multiple volumes cloned they will be further suffixed with an incremental number starting with 1.
	//   Example volume names using name="volume-abcdef"
	//     single volume clone will be named "clone-volume-abcdef-83081“
	//     multi volume clone will be named "clone-volume-abcdef-73721-1”, "clone-volume-abcdef-73721-2”, ...
	//
	// Required: true
	Name *string `json:"name"`

	// default False, Execute failure rolls back clone activity but leaves prepared snapshot
	// True, Execute failure rolls back clone activity and removes the prepared snapshot
	//
	RollbackPrepare bool `json:"rollbackPrepare,omitempty"`
}

// Validate validates this volumes clone execute
func (m *VolumesCloneExecute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumesCloneExecute) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumesCloneExecute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumesCloneExecute) UnmarshalBinary(b []byte) error {
	var res VolumesCloneExecute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
