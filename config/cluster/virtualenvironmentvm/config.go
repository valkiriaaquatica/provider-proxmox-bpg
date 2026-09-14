package virtualenvironmentvm

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_vm", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentVm"
		// Allow node_name to be set only in spec.initProvider (used at creation,
		// ignored afterwards): without this filter late-initialization copies the
		// observed node_name into spec.forProvider, which removes it from the
		// ignore_changes set and makes the provider migrate HA-moved VMs back.
		r.LateInitializer = ujconfig.LateInitializer{
			ConditionalIgnoredFields: []string{"node_name"},
		}
	})
}
