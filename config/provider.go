/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvioronmentacmednsplugin"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentacl"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentacmeplugins"
	"github.com/valkiriaaquatica/provider-proxmox-bpg/config/virtualenvironmentaptrepository"
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
		virtualenvioronmentacmednsplugin.Configure,
		virtualenvironmentacmeplugins.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
