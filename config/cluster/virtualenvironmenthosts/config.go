package virtualenvironmenthosts

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures resources for the virtual environment hosts
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_hosts", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentHosts"
	})
}
