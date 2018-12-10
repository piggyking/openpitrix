// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixReplaceAttachmentResponse openpitrix replace attachment response
// swagger:model openpitrixReplaceAttachmentResponse
type OpenpitrixReplaceAttachmentResponse struct {

	// attachment id
	AttachmentID string `json:"attachment_id,omitempty"`

	// filename
	Filename []string `json:"filename"`
}

// Validate validates this openpitrix replace attachment response
func (m *OpenpitrixReplaceAttachmentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilename(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenpitrixReplaceAttachmentResponse) validateFilename(formats strfmt.Registry) error {

	if swag.IsZero(m.Filename) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixReplaceAttachmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixReplaceAttachmentResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixReplaceAttachmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}