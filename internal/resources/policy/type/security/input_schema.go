/*
Copyright © 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package security

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	policyrecipesecuritymodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/policy/recipe/security"
	reciperesource "github.com/vmware/terraform-provider-tanzu-mission-control/internal/resources/policy/type/security/recipe"
)

var (
	inputSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Input for the security policy, having one of the valid recipes: baseline, custom or strict.",
		Required:    true,
		MaxItems:    1,
		MinItems:    1,
		ForceNew:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				reciperesource.BaselineKey: reciperesource.Baseline,
				reciperesource.CustomKey:   reciperesource.Custom,
				reciperesource.StrictKey:   reciperesource.Strict,
			},
		},
	}
	recipesAllowed = [...]string{reciperesource.BaselineKey, reciperesource.CustomKey, reciperesource.StrictKey}
)

type (
	recipe string
	// InputRecipe is a struct for all types of security policy inputs.
	inputRecipe struct {
		recipe        recipe
		inputBaseline *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Baseline
		inputCustom   *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Custom
		inputStrict   *policyrecipesecuritymodel.VmwareTanzuManageV1alpha1CommonPolicySpecSecurityV1Strict
	}
)

func constructInput(data []interface{}) (inputRecipeData *inputRecipe) {
	if len(data) == 0 || data[0] == nil {
		return inputRecipeData
	}

	inputData, _ := data[0].(map[string]interface{})

	if v, ok := inputData[reciperesource.BaselineKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:        baselineRecipe,
				inputBaseline: reciperesource.ConstructBaseline(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.CustomKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:      customRecipe,
				inputCustom: reciperesource.ConstructCustom(v1),
			}
		}
	}

	if v, ok := inputData[reciperesource.StrictKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			inputRecipeData = &inputRecipe{
				recipe:      strictRecipe,
				inputStrict: reciperesource.ConstructStrict(v1),
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
	case baselineRecipe:
		flattenInputData[reciperesource.BaselineKey] = reciperesource.FlattenBaseline(inputRecipeData.inputBaseline)
	case customRecipe:
		flattenInputData[reciperesource.CustomKey] = reciperesource.FlattenCustom(inputRecipeData.inputCustom)
	case strictRecipe:
		flattenInputData[reciperesource.StrictKey] = reciperesource.FlattenStrict(inputRecipeData.inputStrict)
	case unknownRecipe:
		log.Fatalf("[ERROR]: No valid input recipe block found: minimum one valid input recipe block is required among: %v. Please check the schema.", strings.Join(recipesAllowed[:], `, `))
	}

	return []interface{}{flattenInputData}
}

func validateInput(ctx context.Context, diff *schema.ResourceDiff, i interface{}) error {
	value, ok := diff.GetOk(specKey)
	if !ok {
		return fmt.Errorf("spec: %v is not valid: minimum one valid spec block is required", value)
	}

	data, _ := value.([]interface{})

	if len(data) == 0 || data[0] == nil {
		return fmt.Errorf("spec data: %v is not valid: minimum one valid spec block is required", data)
	}

	specData := data[0].(map[string]interface{})

	v, ok := specData[inputKey]
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

	if v, ok := inputData[reciperesource.BaselineKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.BaselineKey)
		}
	}

	if v, ok := inputData[reciperesource.CustomKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.CustomKey)
		}
	}

	if v, ok := inputData[reciperesource.StrictKey]; ok {
		if v1, ok := v.([]interface{}); ok && len(v1) != 0 {
			recipesFound = append(recipesFound, reciperesource.StrictKey)
		}
	}

	if len(recipesFound) == 0 {
		return fmt.Errorf("no valid input recipe block found: minimum one valid input recipe block is required among: %v", strings.Join(recipesAllowed[:], `, `))
	} else if len(recipesFound) > 1 {
		return fmt.Errorf("found input recipes: %v are not valid: maximum one valid input recipe block is allowed", strings.Join(recipesFound, `, `))
	}

	return nil
}