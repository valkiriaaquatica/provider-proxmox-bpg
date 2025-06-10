/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentacl"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentacmeaccount"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentacmednsplugin"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentaptrepository"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentaptstandardrepository"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentcertificate"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentclusterfirewall"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentclusterfirewallsecuritygroup"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentcontainer"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentdatastores"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentdns"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentdownloadfile"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentfile"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentrole"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentuser"
)

const (
	resourcePrefix = "proxmoxbpg"
	modulePath     = "github.com/valkiriaaquatica/provider-proxmox-bpg"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		virtualenvironmentuser.Configure,
		virtualenvironmentacl.Configure,
		virtualenvironmentrole.Configure,
		virtualenvironmentaptrepository.Configure,
		virtualenvironmentacmednsplugin.Configure,
		virtualenvironmentacmeaccount.Configure,
		virtualenvironmentaptstandardrepository.Configure,
		virtualenvironmentcertificate.Configure,
		virtualenvironmentclusterfirewall.Configure,
		virtualenvironmentclusterfirewallsecuritygroup.Configure,
		virtualenvironmentcontainer.Configure,
		virtualenvironmentdatastores.Configure,
		virtualenvironmentdns.Configure,
		virtualenvironmentdownloadfile.Configure,
		virtualenvironmentfile.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
