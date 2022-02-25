// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1EphemeralContainer An EphemeralContainer is a container that may be added temporarily to an existing pod for
// user-initiated activities such as debugging. Ephemeral containers have no resource or
// scheduling guarantees, and they will not be restarted when they exit or when a pod is
// removed or restarted. If an ephemeral container causes a pod to exceed its resource
// allocation, the pod may be evicted.
// Ephemeral containers may not be added by directly updating the pod spec. They must be added
// via the pod's ephemeralcontainers subresource, and they will appear in the pod spec
// once added.
// This is an alpha feature enabled by the EphemeralContainers feature flag.
//
// swagger:model v1EphemeralContainer
type V1EphemeralContainer struct {

	// Ephemeral containers have all of the fields of Container, plus additional fields
	// specific to ephemeral containers. Fields in common with Container are in the
	// following inlined struct so than an EphemeralContainer may easily be converted
	// to a Container.
	EphemeralContainerCommon *V1EphemeralContainerCommon `json:"ephemeralContainerCommon,omitempty"`

	// If set, the name of the container from PodSpec that this ephemeral container targets.
	// The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container.
	// If not set then the ephemeral container is run in whatever namespaces are shared
	// for the pod. Note that the container runtime must support this feature.
	// +optional
	TargetContainerName string `json:"targetContainerName,omitempty"`
}

// Validate validates this v1 ephemeral container
func (m *V1EphemeralContainer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEphemeralContainerCommon(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EphemeralContainer) validateEphemeralContainerCommon(formats strfmt.Registry) error {
	if swag.IsZero(m.EphemeralContainerCommon) { // not required
		return nil
	}

	if m.EphemeralContainerCommon != nil {
		if err := m.EphemeralContainerCommon.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeralContainerCommon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ephemeralContainerCommon")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this v1 ephemeral container based on the context it is used
func (m *V1EphemeralContainer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEphemeralContainerCommon(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1EphemeralContainer) contextValidateEphemeralContainerCommon(ctx context.Context, formats strfmt.Registry) error {

	if m.EphemeralContainerCommon != nil {
		if err := m.EphemeralContainerCommon.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ephemeralContainerCommon")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ephemeralContainerCommon")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1EphemeralContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1EphemeralContainer) UnmarshalBinary(b []byte) error {
	var res V1EphemeralContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}