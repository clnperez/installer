// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Network network
// swagger:model Network
type Network struct {

	// Network in CIDR notation (192.168.0.0/24)
	// Required: true
	Cidr *string `json:"cidr"`

	// (currently not available) cloud connections this network is attached
	CloudConnections []*NetworkCloudConnectionsItems0 `json:"cloudConnections,omitempty"`

	// DNS Servers
	// Required: true
	DNSServers []string `json:"dnsServers"`

	// Gateway IP Address
	Gateway string `json:"gateway,omitempty"`

	// ip address metrics
	// Required: true
	IPAddressMetrics *NetworkIPAddressMetrics `json:"ipAddressMetrics"`

	// IP Address Ranges
	// Required: true
	IPAddressRanges []*IPAddressRange `json:"ipAddressRanges"`

	// MTU Jumbo Network enabled
	// Required: true
	Jumbo *bool `json:"jumbo"`

	// Network Name
	// Required: true
	Name *string `json:"name"`

	// Unique Network ID
	// Required: true
	NetworkID *string `json:"networkID"`

	// Public IP Address Ranges (for pub-vlan networks)
	PublicIPAddressRanges []*IPAddressRange `json:"publicIPAddressRanges,omitempty"`

	// Type of Network {vlan, pub-vlan}
	// Required: true
	// Enum: [vlan pub-vlan]
	Type *string `json:"type"`

	// VLAN ID
	// Required: true
	VlanID *float64 `json:"vlanID"`
}

// Validate validates this network
func (m *Network) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSServers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddressMetrics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddressRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJumbo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicIPAddressRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Network) validateCidr(formats strfmt.Registry) error {

	if err := validate.Required("cidr", "body", m.Cidr); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateCloudConnections(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudConnections); i++ {
		if swag.IsZero(m.CloudConnections[i]) { // not required
			continue
		}

		if m.CloudConnections[i] != nil {
			if err := m.CloudConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Network) validateDNSServers(formats strfmt.Registry) error {

	if err := validate.Required("dnsServers", "body", m.DNSServers); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateIPAddressMetrics(formats strfmt.Registry) error {

	if err := validate.Required("ipAddressMetrics", "body", m.IPAddressMetrics); err != nil {
		return err
	}

	if m.IPAddressMetrics != nil {
		if err := m.IPAddressMetrics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipAddressMetrics")
			}
			return err
		}
	}

	return nil
}

func (m *Network) validateIPAddressRanges(formats strfmt.Registry) error {

	if err := validate.Required("ipAddressRanges", "body", m.IPAddressRanges); err != nil {
		return err
	}

	for i := 0; i < len(m.IPAddressRanges); i++ {
		if swag.IsZero(m.IPAddressRanges[i]) { // not required
			continue
		}

		if m.IPAddressRanges[i] != nil {
			if err := m.IPAddressRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Network) validateJumbo(formats strfmt.Registry) error {

	if err := validate.Required("jumbo", "body", m.Jumbo); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("networkID", "body", m.NetworkID); err != nil {
		return err
	}

	return nil
}

func (m *Network) validatePublicIPAddressRanges(formats strfmt.Registry) error {

	if swag.IsZero(m.PublicIPAddressRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.PublicIPAddressRanges); i++ {
		if swag.IsZero(m.PublicIPAddressRanges[i]) { // not required
			continue
		}

		if m.PublicIPAddressRanges[i] != nil {
			if err := m.PublicIPAddressRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("publicIPAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var networkTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["vlan","pub-vlan"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkTypeTypePropEnum = append(networkTypeTypePropEnum, v)
	}
}

const (

	// NetworkTypeVlan captures enum value "vlan"
	NetworkTypeVlan string = "vlan"

	// NetworkTypePubVlan captures enum value "pub-vlan"
	NetworkTypePubVlan string = "pub-vlan"
)

// prop value enum
func (m *Network) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, networkTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Network) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Network) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("vlanID", "body", m.VlanID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Network) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Network) UnmarshalBinary(b []byte) error {
	var res Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkCloudConnectionsItems0 network cloud connections items0
// swagger:model NetworkCloudConnectionsItems0
type NetworkCloudConnectionsItems0 struct {

	// the cloud connection id
	CloudConnectionID string `json:"cloudConnectionID,omitempty"`

	// link to the cloud connection resource
	Href string `json:"href,omitempty"`
}

// Validate validates this network cloud connections items0
func (m *NetworkCloudConnectionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCloudConnectionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCloudConnectionsItems0) UnmarshalBinary(b []byte) error {
	var res NetworkCloudConnectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NetworkIPAddressMetrics IP Address Metrics
// swagger:model NetworkIPAddressMetrics
type NetworkIPAddressMetrics struct {

	// Number of available IP addresses
	// Required: true
	Available *float64 `json:"available"`

	// Total number of all IP addresses in all ipAddressRanges
	// Required: true
	Total *float64 `json:"total"`

	// Number of IP addresses currently in use
	// Required: true
	Used *float64 `json:"used"`

	// Utilization of IP addresses in percent form (used / total) [0 - 100]
	// Required: true
	Utilization *float64 `json:"utilization"`
}

// Validate validates this network IP address metrics
func (m *NetworkIPAddressMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtilization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkIPAddressMetrics) validateAvailable(formats strfmt.Registry) error {

	if err := validate.Required("ipAddressMetrics"+"."+"available", "body", m.Available); err != nil {
		return err
	}

	return nil
}

func (m *NetworkIPAddressMetrics) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("ipAddressMetrics"+"."+"total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

func (m *NetworkIPAddressMetrics) validateUsed(formats strfmt.Registry) error {

	if err := validate.Required("ipAddressMetrics"+"."+"used", "body", m.Used); err != nil {
		return err
	}

	return nil
}

func (m *NetworkIPAddressMetrics) validateUtilization(formats strfmt.Registry) error {

	if err := validate.Required("ipAddressMetrics"+"."+"utilization", "body", m.Utilization); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkIPAddressMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkIPAddressMetrics) UnmarshalBinary(b []byte) error {
	var res NetworkIPAddressMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
