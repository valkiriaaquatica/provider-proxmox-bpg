package virtualenvironmentnetworklinuxvlan

import (
	ujconfig "github.com/crossplane/upjet/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("proxmox_virtual_environment_network_linux_vlan", func(r *ujconfig.Resource) {
		r.ShortGroup = "VirtualEnvironmentNetworkLinuxVlan"
		// 	r.ExternalName = ujconfig.TemplatedStringAsIdentifier(
		// 		"",
		// 		"{{ .parameters.node_name }}:{{ .parameters.name }}",
		// 	)
	})
}
