// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// V1AzureFileVolumeSource AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
//
// swagger:model v1AzureFileVolumeSource
type V1AzureFileVolumeSource struct {

	// Defaults to false (read/write). ReadOnly here will force
	// the ReadOnly setting in VolumeMounts.
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`

	// the name of secret that contains Azure Storage Account Name and Key
	SecretName string `json:"secretName,omitempty"`

	// Share Name
	ShareName string `json:"shareName,omitempty"`
}

// Validate validates this v1 azure file volume source
func (m *V1AzureFileVolumeSource) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this v1 azure file volume source based on context it is used
func (m *V1AzureFileVolumeSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V1AzureFileVolumeSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1AzureFileVolumeSource) UnmarshalBinary(b []byte) error {
	var res V1AzureFileVolumeSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}