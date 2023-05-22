/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/helper"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	policyrecipenetworkcommonmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network/common"
)

var AllowAllToPods = &schema.Schema{
	Type:        schema.TypeList,
	Description: "The input schema for network policy allow-all-to-pods recipe version v1",
	Optional:    true,
	ForceNew:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			FromOwnNamespaceKey: {
				Type:        schema.TypeBool,
				Description: "Allow traffic only from own namespace. Allow traffic only from pods in the same namespace as the destination pod.",
				Optional:    true,
				Default:     false,
			},
			toPodLabelsKey: podLabels,
		},
	},
}

func ConstructAllowAllToPods(data []interface{}) (allowAllToPods *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAllToPods) {
	if len(data) == 0 || data[0] == nil {
		return allowAllToPods
	}

	allowAllToPodsData, _ := data[0].(map[string]interface{})

	allowAllToPods = &policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAllToPods{}

	if v, ok := allowAllToPodsData[FromOwnNamespaceKey]; ok {
		allowAllToPods.FromOwnNamespace = helper.BoolPointer(v.(bool))
	}

	if v, ok := allowAllToPodsData[toPodLabelsKey]; ok {
		if vs, ok := v.([]interface{}); ok {
			if len(vs) != 0 && vs[0] != nil {
				allowAllToPods.ToPodLabels = make([]*policyrecipenetworkcommonmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1ToPodLabels, 0)

				for _, raw := range vs {
					allowAllToPods.ToPodLabels = append(allowAllToPods.ToPodLabels, expandToPodLabels(raw))
				}
			}
		}
	}

	return allowAllToPods
}

func FlattenAllowAllToPods(allowAllToPods *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAllToPods) (data []interface{}) {
	if allowAllToPods == nil {
		return data
	}

	flattenAllToPods := make(map[string]interface{})

	flattenAllToPods[FromOwnNamespaceKey] = *allowAllToPods.FromOwnNamespace

	if allowAllToPods.ToPodLabels != nil {
		var podLabels []interface{}

		for _, podLabel := range allowAllToPods.ToPodLabels {
			podLabels = append(podLabels, flattenToPodLabels(podLabel))
		}

		flattenAllToPods[toPodLabelsKey] = podLabels
	}

	return []interface{}{flattenAllToPods}
}
