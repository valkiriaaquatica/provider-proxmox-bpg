package virtualenvironmenthagroup

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment high availability groups
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_hagroup", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentHAGroup"
	})
}
