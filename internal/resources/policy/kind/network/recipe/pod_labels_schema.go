/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var podLabels = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Pod Labels on which traffic should be allowed/denied. Use a label selector to identify the pods to which the policy applies.",
	Optional:    true,
	MinItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			labelKey: {
				Type:         schema.TypeString,
				Description:  "Label key",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 316),
			},
			labelValueKey: {
				Type:         schema.TypeString,
				Description:  "Label Value",
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 63),
			},
		},
	},
}

func expandToPodLabels(data interface{}) (podLabels *policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels) {
	if data == nil {
		return podLabels
	}

	podLabelsData, ok := data.(map[string]interface{})
	if !ok {
		return podLabels
	}

	podLabels = &policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels{}

	if v, ok := podLabelsData[labelKey]; ok {
		podLabels.Key = helper.StringPointer(v.(string))
	}

	if v, ok := podLabelsData[labelValueKey]; ok {
		podLabels.Value = helper.StringPointer(v.(string))
	}

	return podLabels
}

func flattenToPodLabels(podLabels *policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels) (data interface{}) {
	if podLabels == nil {
		return data
	}

	flattenPodLabels := make(map[string]interface{})

	flattenPodLabels[labelKey] = *podLabels.Key
	flattenPodLabels[labelValueKey] = *podLabels.Value

	return flattenPodLabels
}
