package virtualenvironmentharesource

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment high availability resources
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_haresource", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentHAResource"
	})
}
