package virtualenvironmentgroup

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_group", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentGroup"
	})
}
