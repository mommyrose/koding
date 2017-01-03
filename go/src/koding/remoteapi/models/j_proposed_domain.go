package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// JProposedDomain j proposed domain
// swagger:model JProposedDomain
type JProposedDomain struct {

	// domain
	Domain string `json:"domain,omitempty"`

	// group
	Group string `json:"group,omitempty"`

	// machines
	Machines []interface{} `json:"machines"`

	// meta
	Meta interface{} `json:"meta,omitempty"`

	// proxy
	Proxy *JProposedDomainProxy `json:"proxy,omitempty"`
}

// Validate validates this j proposed domain
func (m *JProposedDomain) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachines(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JProposedDomain) validateMachines(formats strfmt.Registry) error {

	if swag.IsZero(m.Machines) { // not required
		return nil
	}

	for i := 0; i < len(m.Machines); i++ {

	}

	return nil
}

func (m *JProposedDomain) validateProxy(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {

		if err := m.Proxy.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

// JProposedDomainProxy j proposed domain proxy
// swagger:model JProposedDomainProxy
type JProposedDomainProxy struct {

	// full Url
	FullURL string `json:"fullUrl,omitempty"`

	// key
	Key string `json:"key,omitempty"`

	// mode
	Mode string `json:"mode,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this j proposed domain proxy
func (m *JProposedDomainProxy) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
