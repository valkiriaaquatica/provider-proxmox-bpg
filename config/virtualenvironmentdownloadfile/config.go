package virtualenvironmentdownloadfile

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_download_file", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentDownloadFile"
	})
}
