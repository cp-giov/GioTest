/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
Code generated by go-swagger; DO NOT EDIT.
*/

package policyrecipenetworkcommonmodel

import "github.com/go-openapi/swag"

// VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels Label Key-Value Pair
//
// swagger:model VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels
type VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels struct {

	// Label Key
	// Required: true
	// Max Length: 316
	// Pattern: ^([[:alnum:]][[:alnum:]-]*[[:alnum:]](.[[:alnum:]][[:alnum:]-]*[[:alnum:]])*/)?[[:alnum:]][[:alnum:]._-]{0,61}[[:alnum:]]$
	Key *string `json:"key"`

	// Label Value
	// Required: true
	// Max Length: 63
	// Pattern: ^([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9]$
	Value *string `json:"value"`
}

// MarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}

	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels) UnmarshalBinary(b []byte) error {
	var res VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*m = res

	return nil
}