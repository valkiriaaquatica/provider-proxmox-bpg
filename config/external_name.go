/*
Copyright 2022 Upbound Inc.
*/

package config

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Import requires using a randomly generated ID from provider: nl-2e21sda
	"proxmox_virtual_environment_user":                            config.IdentifierFromProvider,
	"proxmox_virtual_environment_acl":                             config.IdentifierFromProvider,
	"proxmox_virtual_environment_role":                            config.IdentifierFromProvider,
	"proxmox_virtual_environment_apt_repository":                  config.IdentifierFromProvider,
	"proxmox_virtual_environment_acme_dns_plugin":                 config.IdentifierFromProvider,
	"proxmox_virtual_environment_acme_account":                    config.IdentifierFromProvider,
	"proxmox_virtual_environment_apt_standard_repository":         config.IdentifierFromProvider,
	"proxmox_virtual_environment_certificate":                     config.IdentifierFromProvider,
	"proxmox_virtual_environment_cluster_firewall":                config.IdentifierFromProvider,
	"proxmox_virtual_environment_cluster_firewall_security_group": config.IdentifierFromProvider,
	"proxmox_virtual_environment_container":                       config.IdentifierFromProvider,
	"proxmox_virtual_environment_datastores":                      config.IdentifierFromProvider,
	"proxmox_virtual_environment_dns":                             config.IdentifierFromProvider,
	"proxmox_virtual_environment_download_file":                   config.IdentifierFromProvider,
	"proxmox_virtual_environment_file":                            config.IdentifierFromProvider,
	"proxmox_virtual_environment_firewall_alias":                  config.IdentifierFromProvider,
	"proxmox_virtual_environment_firewall_ipset":                  config.IdentifierFromProvider,
	"proxmox_virtual_environment_firewall_options":                config.IdentifierFromProvider,
	"proxmox_virtual_environment_group":                           config.IdentifierFromProvider,
	"proxmox_virtual_environment_hagroup":                         config.IdentifierFromProvider,
	"proxmox_virtual_environment_firewall_rules":                  config.IdentifierFromProvider,
	"proxmox_virtual_environment_metrics_server":                  config.IdentifierFromProvider,
	"proxmox_virtual_environment_network_linux_bridge":            config.IdentifierFromProvider,
	"proxmox_virtual_environment_network_linux_vlan":              config.IdentifierFromProvider,
	"proxmox_virtual_environment_pool":                            config.IdentifierFromProvider,
	"proxmox_virtual_environment_time":                            config.IdentifierFromProvider,
	"proxmox_virtual_environment_vm":                              config.IdentifierFromProvider,
	"proxmox_virtual_environment_haresource":                      haResourceExternalName(),
	"proxmox_virtual_environment_hosts":                           config.IdentifierFromProvider,
	"proxmox_virtual_environment_node_firewall":                   config.IdentifierFromProvider,
}

// haResourceExternalName derives the Terraform ID from spec.forProvider.resourceId
// before creation, because the plugin-framework haresource errors on a blank ID
// during upjet's pre-create refresh (SDKv2 treated blank IDs as "not found").
// The Terraform ID of this resource is identical to its resource_id.
func haResourceExternalName() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetIDFn = func(_ context.Context, externalName string, parameters map[string]any, _ map[string]any) (string, error) {
		if externalName != "" {
			return externalName, nil
		}
		if rid, ok := parameters["resource_id"].(string); ok {
			return rid, nil
		}
		return "", nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
