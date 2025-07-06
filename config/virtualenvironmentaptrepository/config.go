package virtualenvironmentaptrepository

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_apt_repository", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentAptRepository"
	})
}
