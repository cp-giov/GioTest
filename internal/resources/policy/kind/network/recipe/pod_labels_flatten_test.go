/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

func TestFlattenPodsLabels(t *testing.T) {
	t.Parallel()

	cases := []struct {
		description string
		input       *policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels
		expected    interface{}
	}{
		{
			description: "check for nil pod labels ",
			input:       nil,
			expected:    nil,
		},
		{
			description: "normal scenario with all values of pod labels data",
			input: &policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels{
				Key:   helper.StringPointer("foo"),
				Value: helper.StringPointer("bar"),
			},
			expected: map[string]interface{}{
				labelKey:      "foo",
				labelValueKey: "bar",
			},
		},
	}

	for _, each := range cases {
		test := each
		t.Run(test.description, func(t *testing.T) {
			actual := flattenToPodLabels(test.input)
			require.Equal(t, test.expected, actual)
		})
	}
}
