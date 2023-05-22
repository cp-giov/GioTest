/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindnetwork

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network/recipe"
)

func TestFlattenInput(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *inputRecipe
		expected    []interface{}
	}{
		{
			description: "check for nil input",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with complete input",
			input: &inputRecipe{
				recipe: AllowAllRecipe,
				inputAllowAll: &policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAll{
					FromOwnNamespace: helper.BoolPointer(true),
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					reciperesource.AllowAllKey: []interface{}{
						map[string]interface{}{
							reciperesource.FromOwnNamespaceKey: true,
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenInput(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
