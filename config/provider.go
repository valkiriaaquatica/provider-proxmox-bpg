/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	clusteracl "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentacl"
	clusteracmeaccount "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentacmeaccount"
	clusteracmednsplugin "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentacmednsplugin"
	clusteraptrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentaptrepository"
	clusteraptstandardrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentaptstandardrepository"
	clustercertificate "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentcertificate"
	clusterclusterfirewall "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentclusterfirewall"
	clusterclusterfirewallsecuritygroup "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentclusterfirewallsecuritygroup"
	clustercontainer "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentcontainer"
	clusterdatastores "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentdatastores"
	clusterdns "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentdns"
	clusterdownloadfile "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentdownloadfile"
	clusterfile "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentfile"
	clusterfirewallalias "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentfirewallalias"
	clusterfirewallipset "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentfirewallipset"
	clusterfirewalloptions "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentfirewalloptions"
	clusterfirewallrules "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentfirewallrules"
	clustergroup "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentgroup"
	clusterhagroup "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmenthagroup"
	clusterharesource "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentharesource"
	clusterhosts "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmenthosts"
	clustermetricsserver "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentmetricsserver"
	clusternetworklinuxbridge "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentnetworklinuxbridge"
	clusternetworklinuxvlan "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentnetworklinuxvlan"
	clusterpool "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentpool"
	clusterrole "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentrole"
	clustertime "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmenttime"
	clusteruser "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentuser"
	clustervm "github.com/valkiriaaquatica/provider-proxmox-bpg/config/cluster/virtualenvironmentvm"

	nsacl "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentacl"
	nsacmeaccount "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentacmeaccount"
	nsacmednsplugin "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentacmednsplugin"
	nsaptrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentaptrepository"
	nsaptstandardrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentaptstandardrepository"
	nscertificate "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentcertificate"
	nsclusterfirewall "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentclusterfirewall"
	nsclusterfirewallsecuritygroup "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentclusterfirewallsecuritygroup"
	nscontainer "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentcontainer"
	nsdatastores "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentdatastores"
	nsdns "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentdns"
	nsdownloadfile "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentdownloadfile"
	nsfile "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentfile"
	nsfirewallalias "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentfirewallalias"
	nsfirewallipset "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentfirewallipset"
	nsfirewalloptions "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentfirewalloptions"
	nsfirewallrules "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentfirewallrules"
	nsgroup "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentgroup"
	nshagroup "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmenthagroup"
	nsharesource "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentharesource"
	nshosts "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmenthosts"
	nsmetricsserver "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentmetricsserver"
	nsnetworklinuxbridge "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentnetworklinuxbridge"
	nsnetworklinuxvlan "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentnetworklinuxvlan"
	nspool "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentpool"
	nsrole "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentrole"
	nstime "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmenttime"
	nsuser "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentuser"
	nsvm "github.com/valkiriaaquatica/provider-proxmox-bpg/config/namespaced/virtualenvironmentvm"
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
		ujconfig.WithRootGroup("proxmoxbpg.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		clusteruser.Configure,
		clusteracl.Configure,
		clusterrole.Configure,
		clusteraptrepository.Configure,
		clusteracmednsplugin.Configure,
		clusteracmeaccount.Configure,
		clusteraptstandardrepository.Configure,
		clustercertificate.Configure,
		clusterclusterfirewall.Configure,
		clusterclusterfirewallsecuritygroup.Configure,
		clustercontainer.Configure,
		clusterdatastores.Configure,
		clusterdns.Configure,
		clusterdownloadfile.Configure,
		clusterfile.Configure,
		clusterfirewallalias.Configure,
		clusterfirewallipset.Configure,
		clusterfirewalloptions.Configure,
		clustergroup.Configure,
		clusterhagroup.Configure,
		clusterfirewallrules.Configure,
		clustermetricsserver.Configure,
		clusternetworklinuxbridge.Configure,
		clusternetworklinuxvlan.Configure,
		clusterpool.Configure,
		clustertime.Configure,
		clustervm.Configure,
		clusterharesource.Configure,
		clusterhosts.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns provider configuration for namespaced resources.
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("proxmoxbpg.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		nsuser.Configure,
		nsacl.Configure,
		nsrole.Configure,
		nsaptrepository.Configure,
		nsacmednsplugin.Configure,
		nsacmeaccount.Configure,
		nsaptstandardrepository.Configure,
		nscertificate.Configure,
		nsclusterfirewall.Configure,
		nsclusterfirewallsecuritygroup.Configure,
		nscontainer.Configure,
		nsdatastores.Configure,
		nsdns.Configure,
		nsdownloadfile.Configure,
		nsfile.Configure,
		nsfirewallalias.Configure,
		nsfirewallipset.Configure,
		nsfirewalloptions.Configure,
		nsgroup.Configure,
		nshagroup.Configure,
		nsfirewallrules.Configure,
		nsmetricsserver.Configure,
		nsnetworklinuxbridge.Configure,
		nsnetworklinuxvlan.Configure,
		nspool.Configure,
		nstime.Configure,
		nsvm.Configure,
		nsharesource.Configure,
		nshosts.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
