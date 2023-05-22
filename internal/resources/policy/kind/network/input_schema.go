/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package policykindnetwork

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipenetworkmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/network"
	"github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/kind/network/recipe"
)

var (
	inputSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Input for the network policy, having one of the valid recipes: allowed-all, custom-egress, custom-ingress .",
		Required:    true,
		MaxItems:    1,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				reciperesource.AllowAllKey:       reciperesource.AllowAll,
				reciperesource.AllowAllToPodsKey: reciperesource.AllowAllToPods,
				reciperesource.AllowAllEgressKey: reciperesource.AllowAllEgress,
				reciperesource.DenyAllKey:        reciperesource.DenyAll,
				reciperesource.DenyAllToPodsKey:  reciperesource.DenyAllToPods,
				reciperesource.DenyAllEgressKey:  reciperesource.DenyAllEgress,
			},
		},
	}
	RecipesAllowed = [...]string{reciperesource.AllowAllKey, reciperesource.AllowAllToPodsKey, reciperesource.AllowAllEgressKey, reciperesource.DenyAllKey, reciperesource.DenyAllToPodsKey, reciperesource.DenyAllEgressKey}
)

type (
	Recipe string
	// InputRecipe is a struct for all types of network policy inputs.
	inputRecipe struct {
		recipe              Recipe
		inputAllowAll       *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAll
		inputAllowAllToPods *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1AllowAllToPods
		inputDenyAllToPods  *policyrecipenetworkmodel.VmwareTanzuManageV1alpha1CommonPolicySpecNetworkV1DenyAllToPods
	}
)

func constructInput(data []interface{}) (inputRecipeData *inputRecipe) {
	if len(data) == 0 || data[0] == nil {
		return inputRecipeData
	}

	inputData, _ := data[0].(map[string]interface{})

	if v, ok := inputData[reciperesource.AllowAllKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:        AllowAllRecipe,
				inputAllowAll: reciperesource.ConstructAllowAll(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.AllowAllToPodsKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:              AllowAllToPodsRecipe,
				inputAllowAllToPods: reciperesource.ConstructAllowAllToPods(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.AllowAllEgressKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe: AllowAllEgressRecipe,
			}
		}
	}

	if v, ok := inputData[reciperesource.DenyAllKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe: DenyAllRecipe,
			}
		}
	}

	if v, ok := inputData[reciperesource.DenyAllToPodsKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:             DenyAllToPodsRecipe,
				inputDenyAllToPods: reciperesource.ConstructDenyAllToPods(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.DenyAllEgressKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe: DenyAllEgressRecipe,
			}
		}
	}

	return inputRecipeData
}

func flattenInput(inputRecipeData *inputRecipe) (data []interface{}) {
	if inputRecipeData == nil {
		return data
	}

	flattenInputData := make(map[string]interface{})

	switch inputRecipeData.recipe {
	case AllowAllRecipe:
		flattenInputData[reciperesource.AllowAllKey] = reciperesource.FlattenAllowAll(inputRecipeData.inputAllowAll)
	case AllowAllToPodsRecipe:
		flattenInputData[reciperesource.AllowAllToPodsKey] = reciperesource.FlattenAllowAllToPods(inputRecipeData.inputAllowAllToPods)
	case AllowAllEgressRecipe:
		flattenInputData[reciperesource.AllowAllEgressKey] = []interface{}{make(map[string]interface{})}
	case DenyAllRecipe:
		flattenInputData[reciperesource.DenyAllKey] = []interface{}{make(map[string]interface{})}
	case DenyAllToPodsRecipe:
		flattenInputData[reciperesource.DenyAllToPodsKey] = reciperesource.FlattenDenyAllToPods(inputRecipeData.inputDenyAllToPods)
	case DenyAllEgressRecipe:
		flattenInputData[reciperesource.DenyAllEgressKey] = []interface{}{make(map[string]interface{})}
	case UnknownRecipe:
		fmt.Printf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(RecipesAllowed[:], `, `))
	}

	return []interface{}{flattenInputData}
}

func ValidateInput(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(policy.SpecKey)
	if !ok {
		return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("spec data: %v is not valid: minimum one valid spec block is required among: %v", data, strings.Join(RecipesAllowed[:], `, `))
	}

	specData := data[0].(map[string]interface{})

	v, ok := specData[policy.InputKey]
	if !ok {
		return fmt.Errorf("input: %v is not valid: minimum one valid input block is required", v)
	}

	v1, ok := v.([]interface{})
	if !ok {
		return fmt.Errorf("type of input block data: %v is not valid", v1)
	}

	if len(v1) == 0 || v1[0] == nil {
		return fmt.Errorf("input data: %v is not valid: minimum one valid input block is required", v1)
	}

	inputData, _ := v1[0].(map[string]interface{})
	recipesFound := make([]string, 0)

	if v, ok := inputData[reciperesource.AllowAllKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.AllowAllKey)
		}
	}

	if v, ok := inputData[reciperesource.AllowAllToPodsKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.AllowAllToPodsKey)
		}
	}

	if v, ok := inputData[reciperesource.AllowAllEgressKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.AllowAllEgressKey)
		}
	}

	if v, ok := inputData[reciperesource.DenyAllKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.DenyAllKey)
		}
	}

	if v, ok := inputData[reciperesource.DenyAllToPodsKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.DenyAllToPodsKey)
		}
	}

	if v, ok := inputData[reciperesource.DenyAllEgressKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.DenyAllEgressKey)
		}
	}

	if len(recipesFound) == 0 {
		return fmt.Errorf("no valid input recipe block found: minimum one valid input recipe block is required among: %v", strings.Join(RecipesAllowed[:], `, `))
	} else if len(recipesFound) > 1 {
		return fmt.Errorf("found input recipes: %v are not valid: maximum one valid input recipe block is allowed", strings.Join(recipesFound, `, `))
	}

	return nil
}
