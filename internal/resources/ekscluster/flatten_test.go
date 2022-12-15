/*
Copyright 2022 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package ekscluster

import (
	"testing"

	"github.com/stretchr/testify/require"

	eksmodel "github.com/vmware/terraform-provider-tanzu-mission-control/internal/models/ekscluster"
)

func TestFlattenCluterSpec(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec
		expected    []interface{}
	}{
		{
			description: "nil spec",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description: "full cluster spec",
			getInput:    getClusterSpec,
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "test-cg",
					"proxy":         "test-prooxy",
					"config": []interface{}{
						map[string]interface{}{
							"kubernetes_network_config": []interface{}{
								map[string]interface{}{
									"service_cidr": "10.0.0.0/10",
								},
							},

							"kubernetes_version": "1.12",
							"logging": []interface{}{
								map[string]interface{}{
									"api_server":         false,
									"audit":              false,
									"authenticator":      false,
									"controller_manager": true,
									"scheduler":          true,
								},
							},
							"role_arn": "role-arn",
							"tags": map[string]string{
								"tag1": "tag2",
							},
							"vpc": []interface{}{
								map[string]interface{}{
									"enable_private_access": false,
									"enable_public_access":  false,
									"public_access_cidrs": []string{
										"0.0.0.0/1",
										"1.0.0.0/1",
									},
									"security_groups": []string{
										"sg-1",
										"sg-2",
									},
									"subnet_ids": []string{
										"subnet-1",
										"subnet-2",
									},
								},
							},
						},
					},
					"nodepool": []interface{}{
						map[string]interface{}{
							"info": []interface{}{
								map[string]interface{}{
									"description": "test np",
									"name":        "test-np",
								},
							},
							"spec": []interface{}{
								map[string]interface{}{
									"ami_type":      "AL2_x86_64",
									"capacity_type": "ON_DEMAND",
									"instance_types": []string{
										"t3.medium",
										"m3.large",
									},
									"launch_template": []interface{}{
										map[string]interface{}{
											"id":      "",
											"name":    "templ",
											"version": "7",
										},
									},
									"node_labels": map[string]string{
										"key1": "val1",
									},
									"remote_access": []interface{}{
										map[string]interface{}{
											"security_groups": []string{
												"sg-0a6768722e9716768",
											},
											"ssh_key": "test-key",
										},
									},
									"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
									"root_disk_size": int32(20),
									"scaling_config": []interface{}{
										map[string]interface{}{
											"desired_size": int32(8),
											"max_size":     int32(16),
											"min_size":     int32(3),
										},
									},
									"subnet_ids": []string{
										"subnet-0a184f9301ae39a86",
										"subnet-0b495d7c212fc92a1",
										"subnet-0c86ec9ecde7b9bf7",
										"subnet-06497e6063c209f4d",
									},
									"tags": map[string]string{
										"tg1": "tv1",
									},
									"taints": []interface{}{
										map[string]interface{}{
											"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
											"key":    "tkey",
											"value":  "tvalue",
										},
									},
									"update_config": []interface{}{
										map[string]interface{}{
											"max_unavailable_nodes":      "10",
											"max_unavailable_percentage": "12",
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "empty nodepools",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec {
				spec := getClusterSpec()
				spec.NodePools = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "test-cg",
					"config": []interface{}{
						map[string]interface{}{
							"kubernetes_network_config": []interface{}{
								map[string]interface{}{
									"service_cidr": "10.0.0.0/10",
								},
							},

							"kubernetes_version": "1.12",
							"logging": []interface{}{
								map[string]interface{}{
									"api_server":         false,
									"audit":              false,
									"authenticator":      false,
									"controller_manager": true,
									"scheduler":          true,
								},
							},
							"role_arn": "role-arn",
							"tags": map[string]string{
								"tag1": "tag2",
							},
							"vpc": []interface{}{
								map[string]interface{}{
									"enable_private_access": false,
									"enable_public_access":  false,
									"public_access_cidrs": []string{
										"0.0.0.0/1",
										"1.0.0.0/1",
									},
									"security_groups": []string{
										"sg-1",
										"sg-2",
									},
									"subnet_ids": []string{
										"subnet-1",
										"subnet-2",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "empty proxy",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec {
				spec := getClusterSpec()
				spec.ProxyName = ""
				spec.NodePools = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "test-cg",
					"proxy":         "test-prooxy",
					"config": []interface{}{
						map[string]interface{}{
							"kubernetes_network_config": []interface{}{
								map[string]interface{}{
									"service_cidr": "10.0.0.0/10",
								},
							},

							"kubernetes_version": "1.12",
							"logging": []interface{}{
								map[string]interface{}{
									"api_server":         false,
									"audit":              false,
									"authenticator":      false,
									"controller_manager": true,
									"scheduler":          true,
								},
							},
							"role_arn": "role-arn",
							"tags": map[string]string{
								"tag1": "tag2",
							},
							"vpc": []interface{}{
								map[string]interface{}{
									"enable_private_access": false,
									"enable_public_access":  false,
									"public_access_cidrs": []string{
										"0.0.0.0/1",
										"1.0.0.0/1",
									},
									"security_groups": []string{
										"sg-1",
										"sg-2",
									},
									"subnet_ids": []string{
										"subnet-1",
										"subnet-2",
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "empty config",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec {
				spec := getClusterSpec()
				spec.ProxyName = ""
				spec.Config = nil
				spec.NodePools = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"cluster_group": "test-cg",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenClusterSpec(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func TestFlattenConfig(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig
		expected    []interface{}
	}{
		{
			description: "nil config",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description: "full config",
			getInput:    getConfig,
			expected: []interface{}{
				map[string]interface{}{
					"kubernetes_network_config": []interface{}{
						map[string]interface{}{
							"service_cidr": "10.0.0.0/10",
						},
					},
					"kubernetes_version": "1.12",
					"logging": []interface{}{
						map[string]interface{}{
							"api_server":         false,
							"audit":              false,
							"authenticator":      false,
							"controller_manager": true,
							"scheduler":          true,
						},
					},
					"role_arn": "role-arn",
					"tags": map[string]string{
						"tag1": "tag2",
					},
					"vpc": []interface{}{
						map[string]interface{}{
							"enable_private_access": false,
							"enable_public_access":  false,
							"public_access_cidrs": []string{
								"0.0.0.0/1",
								"1.0.0.0/1",
							},
							"security_groups": []string{
								"sg-1",
								"sg-2",
							},
							"subnet_ids": []string{
								"subnet-1",
								"subnet-2",
							},
						},
					},
				},
			},
		},
		{
			description: "k8s network config is nil config",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				config := getConfig()
				config.KubernetesNetworkConfig = nil
				return config
			},
			expected: []interface{}{
				map[string]interface{}{
					"kubernetes_version": "1.12",
					"logging": []interface{}{
						map[string]interface{}{
							"api_server":         false,
							"audit":              false,
							"authenticator":      false,
							"controller_manager": true,
							"scheduler":          true,
						},
					},
					"role_arn": "role-arn",
					"tags": map[string]string{
						"tag1": "tag2",
					},
					"vpc": []interface{}{
						map[string]interface{}{
							"enable_private_access": false,
							"enable_public_access":  false,
							"public_access_cidrs": []string{
								"0.0.0.0/1",
								"1.0.0.0/1",
							},
							"security_groups": []string{
								"sg-1",
								"sg-2",
							},
							"subnet_ids": []string{
								"subnet-1",
								"subnet-2",
							},
						},
					},
				},
			},
		},
		{
			description: "logging is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				config := getConfig()
				config.Logging = nil
				return config
			},
			expected: []interface{}{
				map[string]interface{}{
					"kubernetes_network_config": []interface{}{
						map[string]interface{}{
							"service_cidr": "10.0.0.0/10",
						},
					},
					"role_arn":           "role-arn",
					"kubernetes_version": "1.12",
					"tags": map[string]string{
						"tag1": "tag2",
					},
					"vpc": []interface{}{
						map[string]interface{}{
							"enable_private_access": false,
							"enable_public_access":  false,
							"public_access_cidrs": []string{
								"0.0.0.0/1",
								"1.0.0.0/1",
							},
							"security_groups": []string{
								"sg-1",
								"sg-2",
							},
							"subnet_ids": []string{
								"subnet-1",
								"subnet-2",
							},
						},
					},
				},
			},
		},
		{
			description: "vpc is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
				config := getConfig()
				config.Vpc = nil
				return config
			},
			expected: []interface{}{
				map[string]interface{}{
					"kubernetes_network_config": []interface{}{
						map[string]interface{}{
							"service_cidr": "10.0.0.0/10",
						},
					},
					"kubernetes_version": "1.12",
					"logging": []interface{}{
						map[string]interface{}{
							"api_server":         false,
							"audit":              false,
							"authenticator":      false,
							"controller_manager": true,
							"scheduler":          true,
						},
					},
					"role_arn": "role-arn",
					"tags": map[string]string{
						"tag1": "tag2",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenConfig(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func TestFlattenNodepools(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition
		expected    []interface{}
	}{
		{
			description: "nil list",
			getInput: func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description: "single nodepool",
			getInput: func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
					getNodepoolDef("test-np"),
				}
			},
			expected: []interface{}{
				map[string]interface{}{
					"info": []interface{}{
						map[string]interface{}{
							"description": "test np",
							"name":        "test-np",
						},
					},
					"spec": []interface{}{
						map[string]interface{}{
							"ami_type":      "AL2_x86_64",
							"capacity_type": "ON_DEMAND",
							"instance_types": []string{
								"t3.medium",
								"m3.large",
							},
							"launch_template": []interface{}{
								map[string]interface{}{
									"id":      "",
									"name":    "templ",
									"version": "7",
								},
							},
							"node_labels": map[string]string{
								"key1": "val1",
							},
							"remote_access": []interface{}{
								map[string]interface{}{
									"security_groups": []string{
										"sg-0a6768722e9716768",
									},
									"ssh_key": "test-key",
								},
							},
							"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
							"root_disk_size": int32(20),
							"scaling_config": []interface{}{
								map[string]interface{}{
									"desired_size": int32(8),
									"max_size":     int32(16),
									"min_size":     int32(3),
								},
							},
							"subnet_ids": []string{
								"subnet-0a184f9301ae39a86",
								"subnet-0b495d7c212fc92a1",
								"subnet-0c86ec9ecde7b9bf7",
								"subnet-06497e6063c209f4d",
							},
							"tags": map[string]string{
								"tg1": "tv1",
							},
							"taints": []interface{}{
								map[string]interface{}{
									"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
									"key":    "tkey",
									"value":  "tvalue",
								},
							},
							"update_config": []interface{}{
								map[string]interface{}{
									"max_unavailable_nodes":      "10",
									"max_unavailable_percentage": "12",
								},
							},
						},
					},
				},
			},
		},
		{
			description: "multiple nodepool",
			getInput: func() []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
					getNodepoolDef("test-np"),
					getNodepoolDef("test-np-2"),
				}
			},
			expected: []interface{}{
				map[string]interface{}{
					"info": []interface{}{
						map[string]interface{}{
							"description": "test np",
							"name":        "test-np",
						},
					},
					"spec": []interface{}{
						map[string]interface{}{
							"ami_type":      "AL2_x86_64",
							"capacity_type": "ON_DEMAND",
							"instance_types": []string{
								"t3.medium",
								"m3.large",
							},
							"launch_template": []interface{}{
								map[string]interface{}{
									"id":      "",
									"name":    "templ",
									"version": "7",
								},
							},
							"node_labels": map[string]string{
								"key1": "val1",
							},
							"remote_access": []interface{}{
								map[string]interface{}{
									"security_groups": []string{
										"sg-0a6768722e9716768",
									},
									"ssh_key": "test-key",
								},
							},
							"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
							"root_disk_size": int32(20),
							"scaling_config": []interface{}{
								map[string]interface{}{
									"desired_size": int32(8),
									"max_size":     int32(16),
									"min_size":     int32(3),
								},
							},
							"subnet_ids": []string{
								"subnet-0a184f9301ae39a86",
								"subnet-0b495d7c212fc92a1",
								"subnet-0c86ec9ecde7b9bf7",
								"subnet-06497e6063c209f4d",
							},
							"tags": map[string]string{
								"tg1": "tv1",
							},
							"taints": []interface{}{
								map[string]interface{}{
									"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
									"key":    "tkey",
									"value":  "tvalue",
								},
							},
							"update_config": []interface{}{
								map[string]interface{}{
									"max_unavailable_nodes":      "10",
									"max_unavailable_percentage": "12",
								},
							},
						},
					},
				},
				map[string]interface{}{
					"info": []interface{}{
						map[string]interface{}{
							"description": "test np",
							"name":        "test-np-2",
						},
					},
					"spec": []interface{}{
						map[string]interface{}{
							"ami_type":      "AL2_x86_64",
							"capacity_type": "ON_DEMAND",
							"instance_types": []string{
								"t3.medium",
								"m3.large",
							},
							"launch_template": []interface{}{
								map[string]interface{}{
									"id":      "",
									"name":    "templ",
									"version": "7",
								},
							},
							"node_labels": map[string]string{
								"key1": "val1",
							},
							"remote_access": []interface{}{
								map[string]interface{}{
									"security_groups": []string{
										"sg-0a6768722e9716768",
									},
									"ssh_key": "test-key",
								},
							},
							"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
							"root_disk_size": int32(20),
							"scaling_config": []interface{}{
								map[string]interface{}{
									"desired_size": int32(8),
									"max_size":     int32(16),
									"min_size":     int32(3),
								},
							},
							"subnet_ids": []string{
								"subnet-0a184f9301ae39a86",
								"subnet-0b495d7c212fc92a1",
								"subnet-0c86ec9ecde7b9bf7",
								"subnet-06497e6063c209f4d",
							},
							"tags": map[string]string{
								"tg1": "tv1",
							},
							"taints": []interface{}{
								map[string]interface{}{
									"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
									"key":    "tkey",
									"value":  "tvalue",
								},
							},
							"update_config": []interface{}{
								map[string]interface{}{
									"max_unavailable_nodes":      "10",
									"max_unavailable_percentage": "12",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenNodePools(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func TestFlattenNodepool(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition
		expected    map[string]interface{}
	}{
		{
			description: "nil np",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return nil
			},
			expected: map[string]interface{}{},
		},
		{
			description: "full nodepool",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
				return getNodepoolDef("test-np")
			},
			expected: map[string]interface{}{
				"info": []interface{}{
					map[string]interface{}{
						"description": "test np",
						"name":        "test-np",
					},
				},
				"spec": []interface{}{
					map[string]interface{}{
						"ami_type":      "AL2_x86_64",
						"capacity_type": "ON_DEMAND",
						"instance_types": []string{
							"t3.medium",
							"m3.large",
						},
						"launch_template": []interface{}{
							map[string]interface{}{
								"id":      "",
								"name":    "templ",
								"version": "7",
							},
						},
						"node_labels": map[string]string{
							"key1": "val1",
						},
						"remote_access": []interface{}{
							map[string]interface{}{
								"security_groups": []string{
									"sg-0a6768722e9716768",
								},
								"ssh_key": "test-key",
							},
						},
						"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
						"root_disk_size": int32(20),
						"scaling_config": []interface{}{
							map[string]interface{}{
								"desired_size": int32(8),
								"max_size":     int32(16),
								"min_size":     int32(3),
							},
						},
						"subnet_ids": []string{
							"subnet-0a184f9301ae39a86",
							"subnet-0b495d7c212fc92a1",
							"subnet-0c86ec9ecde7b9bf7",
							"subnet-06497e6063c209f4d",
						},
						"tags": map[string]string{
							"tg1": "tv1",
						},
						"taints": []interface{}{
							map[string]interface{}{
								"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
								"key":    "tkey",
								"value":  "tvalue",
							},
						},
						"update_config": []interface{}{
							map[string]interface{}{
								"max_unavailable_nodes":      "10",
								"max_unavailable_percentage": "12",
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenNodePool(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func TestFlattenNodepoolSpec(t *testing.T) {
	tests := []struct {
		description string
		getInput    func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec
		expected    []interface{}
	}{
		{
			description: "nil spec",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				return nil
			},
			expected: []interface{}{},
		},
		{
			description: "full spec",
			getInput:    getNodepoolSpec,
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
		{
			description: "launch template with id",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.LaunchTemplate = &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolLaunchTemplate{
					ID:      "lt-id",
					Name:    "",
					Version: "7",
				}
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "lt-id",
							"name":    "",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
		{
			description: "remote access is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RemoteAccess = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
		{
			description: "root disk size is 0",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RootDiskSize = 0
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn": "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
		{
			description: "scaling config is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.ScalingConfig = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":       "AL2_x86_64",
					"capacity_type":  "ON_DEMAND",
					"root_disk_size": int32(20),
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn": "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
		{
			description: "subnet ids are nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.SubnetIds = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
		{
			description: "taints are nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.Taints = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
		{
			description: "update config is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.UpdateConfig = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"security_groups": []string{
								"sg-0a6768722e9716768",
							},
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
				},
			},
		},
		{
			description: "sg in remote access is nil",
			getInput: func() *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolSpec {
				spec := getNodepoolSpec()
				spec.RemoteAccess.SecurityGroups = nil
				return spec
			},
			expected: []interface{}{
				map[string]interface{}{
					"ami_type":      "AL2_x86_64",
					"capacity_type": "ON_DEMAND",
					"instance_types": []string{
						"t3.medium",
						"m3.large",
					},
					"launch_template": []interface{}{
						map[string]interface{}{
							"id":      "",
							"name":    "templ",
							"version": "7",
						},
					},
					"node_labels": map[string]string{
						"key1": "val1",
					},
					"remote_access": []interface{}{
						map[string]interface{}{
							"ssh_key": "test-key",
						},
					},
					"role_arn":       "arn:aws:iam::000000000000:role/control-plane.1234567890123467890.eks.tmc.cloud.vmware.com",
					"root_disk_size": int32(20),
					"scaling_config": []interface{}{
						map[string]interface{}{
							"desired_size": int32(8),
							"max_size":     int32(16),
							"min_size":     int32(3),
						},
					},
					"subnet_ids": []string{
						"subnet-0a184f9301ae39a86",
						"subnet-0b495d7c212fc92a1",
						"subnet-0c86ec9ecde7b9bf7",
						"subnet-06497e6063c209f4d",
					},
					"tags": map[string]string{
						"tg1": "tv1",
					},
					"taints": []interface{}{
						map[string]interface{}{
							"effect": eksmodel.NewVmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffect(eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolTaintEffectPREFERNOSCHEDULE),
							"key":    "tkey",
							"value":  "tvalue",
						},
					},
					"update_config": []interface{}{
						map[string]interface{}{
							"max_unavailable_nodes":      "10",
							"max_unavailable_percentage": "12",
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			output := flattenSpec(test.getInput())
			require.Equal(t, test.expected, output)
		})
	}
}

func getClusterSpec() *eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec {
	return &eksmodel.VmwareTanzuManageV1alpha1EksclusterSpec{
		ClusterGroupName: "test-cg",
		ProxyName:        "test-prooxy",
		Config:           getConfig(),
		NodePools: []*eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
			{
				Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
					Description: "test np",
					Name:        "test-np",
				},
				Spec: getNodepoolSpec(),
			},
		},
	}
}

func getConfig() *eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig {
	return &eksmodel.VmwareTanzuManageV1alpha1EksclusterControlPlaneConfig{
		KubernetesNetworkConfig: &eksmodel.VmwareTanzuManageV1alpha1EksclusterKubernetesNetworkConfig{
			ServiceCidr: "10.0.0.0/10",
		},
		Logging: &eksmodel.VmwareTanzuManageV1alpha1EksclusterLogging{
			APIServer:         false,
			Audit:             false,
			Authenticator:     false,
			ControllerManager: true,
			Scheduler:         true,
		},
		RoleArn: "role-arn",
		Tags: map[string]string{
			"tag1": "tag2",
		},
		Version: "1.12",
		Vpc: &eksmodel.VmwareTanzuManageV1alpha1EksclusterVPCConfig{
			EnablePrivateAccess: false,
			EnablePublicAccess:  false,
			PublicAccessCidrs: []string{
				"0.0.0.0/1",
				"1.0.0.0/1",
			},
			SecurityGroups: []string{
				"sg-1",
				"sg-2",
			},
			SubnetIds: []string{
				"subnet-1",
				"subnet-2",
			},
		},
	}
}

func getNodepoolDef(name string) *eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition {
	return &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolDefinition{
		Info: &eksmodel.VmwareTanzuManageV1alpha1EksclusterNodepoolInfo{
			Description: "test np",
			Name:        name,
		},
		Spec: getNodepoolSpec(),
	}
}
