/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindnetwork

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy"
	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
)

var SpecSchema = &schema.Schema{
	Type:        schema.TypeList,
	Description: "Spec for the network policy",
	Required:    true,
	MaxItems:    1,
	Elem: &schema.Resource{
		Schema: map[string]*schema.Schema{
			policy.InputKey:             inputSchema,
			policy.NamespaceSelectorKey: policy.NamespaceSelector,
		},
	},
}

func ConstructSpec(d *schema.ResourceData) (spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) {
	value, ok := d.GetOk(policy.SpecKey)
	if !ok {
		return spec
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return spec
	}

	specData := data[0].(map[string]interface{})

	spec = &policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec{
		Type:          typePolicy,
		RecipeVersion: policy.RecipeVersionDefaultValue,
	}

	input, ok := specData[policy.InputKey]
	if !ok {
		return spec
	}

	inputData, ok := input.([]interface{})
	if !ok {
		return spec
	}

	inputRecipeData := constructInput(inputData)

	if inputRecipeData == nil || inputRecipeData.recipe == "" {
		return spec
	}

	spec.Recipe = strings.ReplaceAll(string(inputRecipeData.recipe), "_", "-")

	switch inputRecipeData.recipe {
	case AllowAllRecipe:
		if inputRecipeData.inputAllowAll != nil {
			spec.Input = *inputRecipeData.inputAllowAll
		}
	case UnknownRecipe:
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	if v, ok := specData[policy.NamespaceSelectorKey]; ok {
		if v1, ok := v.([]interface{}); ok {
			spec.NamespaceSelector = policy.ConstructNamespaceSelector(v1)
		}
	}

	return spec
}

func FlattenSpec(spec *policymodel.VmwareTanzuManageV1alpha1CommonPolicySpec) (data []interface{}) {
	if spec == nil {
		return data
	}

	flattenSpecData := make(map[string]interface{})

	if spec.Input == nil {
		return data
	}

	input, ok := spec.Input.(map[string]interface{})
	if !ok {
		return data
	}

	var inputRecipeData *inputRecipe

	byteSlice, err := json.Marshal(input)
	if err != nil {
		return data
	}

	switch strings.ReplaceAll(spec.Recipe, "-", "_") {
	case string(AllowAllRecipe):
		var allowAllRecipeInput policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAll

		err = allowAllRecipeInput.UnmarshalBinary(byteSlice)
		if err != nil {
			return data
		}

		inputRecipeData = &inputRecipe{
			recipe:        AllowAllRecipe,
			inputAllowAll: &allowAllRecipeInput,
		}
	case string(UnknownRecipe):
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	flattenSpecData[policy.InputKey] = flattenInput(inputRecipeData)

	if spec.NamespaceSelector != nil {
		flattenSpecData[policy.NamespaceSelectorKey] = policy.FlattenNamespaceSelector(spec.NamespaceSelector)
	}

	return []interface{}{flattenSpecData}
}
