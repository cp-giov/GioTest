/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package recipe

const (
	AllowAllKey         = "allow_all"
	AllowAllToPodsKey   = "allow_all_to_pods"
	AllowAllEgressKey   = "allow_all_egress"
	DenyAllKey          = "deny_all"
	DenyAllToPodsKey    = "deny_all_to_pods"
	DenyAllEgressKey    = "deny_all_egress"
	FromOwnNamespaceKey = "from_own_namespace"
	toPodLabelsKey      = "to_pod_labels"
	labelKey            = "key"
	labelValueKey       = "value"
	rulesKey            = "rules"
	portsKey            = "ports"
	portKey             = "port"
	protocolKey         = "protocol"
	ruleSpecKey         = "rule_spec"
	unionKey            = "union"
)
