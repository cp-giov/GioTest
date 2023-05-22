/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var DenyAllToPods = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy deny-all-to-pods recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			toPodLabelsKey: podLabels,
		},
	},
}

func ConstructDenyAllToPods(data []interface{}) (denyAllToPods *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1DenyAllToPods) {
	if len(data) == 0 || data[0] == nil {
		return denyAllToPods
	}

	denyAllToPodsData, _ := data[0].(map[string]interface{})

	denyAllToPods = &policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1DenyAllToPods{}

	if v, ok := denyAllToPodsData[toPodLabelsKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				denyAllToPods.ToPodLabels = make([]*policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels, 0)

				for _, raw := range vs {
					denyAllToPods.ToPodLabels = append(denyAllToPods.ToPodLabels, expandToPodLabels(raw))
				}
			}
		}
	}

	return denyAllToPods
}

func FlattenDenyAllToPods(denyAllToPods *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1DenyAllToPods) (data []interface{}) {
	if denyAllToPods == nil {
		return data
	}

	flattenDenyAllToPods := make(map[string]interface{})

	if denyAllToPods.ToPodLabels != nil {
		var podLabels []interface{}

		for _, podLabel := range denyAllToPods.ToPodLabels {
			podLabels = append(podLabels, flattenToPodLabels(podLabel))
		}

		flattenDenyAllToPods[toPodLabelsKey] = podLabels
	}

	return []interface{}{flattenDenyAllToPods}
}
