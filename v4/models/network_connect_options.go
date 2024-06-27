// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkConnectOptions NetworkConnectOptions describes options for connecting
// a container to a network
//
// swagger:model NetworkConnectOptions
type NetworkConnectOptions struct {

	// Aliases contains a list of names which the dns server should resolve
	// to this container. Should only be set when DNSEnabled is true on the Network.
	// If aliases are set but there is no dns support for this network the
	// network interface implementation should ignore this and NOT error.
	// Optional.
	Aliases []string `json:"aliases"`

	// container
	Container string `json:"container,omitempty"`

	// InterfaceName for this container. Required in the backend.
	// Optional in the frontend. Will be filled with ethX (where X is a integer) when empty.
	InterfaceName string `json:"interface_name,omitempty"`

	// StaticIPs for this container. Optional.
	StaticIPs []IP `json:"static_ips"`

	// static mac
	StaticMac MacAddress `json:"static_mac,omitempty"`
}

// Validate validates this network connect options
func (m *NetworkConnectOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStaticIPs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaticMac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkConnectOptions) validateStaticIPs(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticIPs) { // not required
		return nil
	}

	for i := 0; i < len(m.StaticIPs); i++ {

		if err := m.StaticIPs[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static_ips" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static_ips" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *NetworkConnectOptions) validateStaticMac(formats strfmt.Registry) error {
	if swag.IsZero(m.StaticMac) { // not required
		return nil
	}

	if err := m.StaticMac.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_mac")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_mac")
		}
		return err
	}

	return nil
}

// ContextValidate validate this network connect options based on the context it is used
func (m *NetworkConnectOptions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStaticIPs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStaticMac(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkConnectOptions) contextValidateStaticIPs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.StaticIPs); i++ {

		if err := m.StaticIPs[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("static_ips" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("static_ips" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *NetworkConnectOptions) contextValidateStaticMac(ctx context.Context, formats strfmt.Registry) error {

	if err := m.StaticMac.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("static_mac")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("static_mac")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkConnectOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkConnectOptions) UnmarshalBinary(b []byte) error {
	var res NetworkConnectOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
