/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

func TestFlattenAllowAllToPods(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAllToPods
		expected    []interface{}
	}{
		{
			description: "check for nil allow-all-to-pods recipe network policy ",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with valid allow-all-to-pods recipe network policy ",
			input: &policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAllToPods{
				FromOwnNamespace: helper.BoolPointer(true),
				ToPodLabels: []*policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels{
					{
						Key:   helper.StringPointer("foo"),
						Value: helper.StringPointer("bar"),
					},
				},
			},
			expected: []interface{}{
				map[string]interface{}{
					FromOwnNamespaceKey: true,
					toPodLabelsKey: []interface{}{
						map[string]interface{}{
							labelKey:      "foo",
							labelValueKey: "bar",
						},
					},
				},
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := FlattenAllowAllToPods(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
