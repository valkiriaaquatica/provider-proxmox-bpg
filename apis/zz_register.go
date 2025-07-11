// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/v1alpha1"
	v1beta1 "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/v1beta1"
	v1alpha1virtual "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtual/v1alpha1"
	v1alpha1virtualenvironment "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironment/v1alpha1"
	v1alpha1virtualenvironmentacl "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentacl/v1alpha1"
	v1alpha1virtualenvironmentacmeaccount "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentacmeaccount/v1alpha1"
	v1alpha1virtualenvironmentacmednsplugin "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentacmednsplugin/v1alpha1"
	v1alpha1virtualenvironmentaptrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentaptrepository/v1alpha1"
	v1alpha1virtualenvironmentaptstandardrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentaptstandardrepository/v1alpha1"
	v1alpha1virtualenvironmentclusterfirewallsecuritygroup "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentclusterfirewallsecuritygroup/v1alpha1"
	v1alpha1virtualenvironmentcontainer "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentcontainer/v1alpha1"
	v1alpha1virtualenvironmentdns "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentdns/v1alpha1"
	v1alpha1virtualenvironmentdownloadfile "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentdownloadfile/v1alpha1"
	v1alpha1virtualenvironmentfile "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentfile/v1alpha1"
	v1alpha1virtualenvironmentfirewallalias "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentfirewallalias/v1alpha1"
	v1alpha1virtualenvironmentfirewallipset "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentfirewallipset/v1alpha1"
	v1alpha1virtualenvironmentfirewalloptions "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentfirewalloptions/v1alpha1"
	v1alpha1virtualenvironmentfirewallrules "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentfirewallrules/v1alpha1"
	v1alpha1virtualenvironmentgroup "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentgroup/v1alpha1"
	v1alpha1virtualenvironmenthagroup "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmenthagroup/v1alpha1"
	v1alpha1virtualenvironmentharesource "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentharesource/v1alpha1"
	v1alpha1virtualenvironmenthosts "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmenthosts/v1alpha1"
	v1alpha1virtualenvironmentmetricsserver "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentmetricsserver/v1alpha1"
	v1alpha1virtualenvironmentnetworklinuxbridge "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentnetworklinuxbridge/v1alpha1"
	v1alpha1virtualenvironmentnetworklinuxvlan "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentnetworklinuxvlan/v1alpha1"
	v1alpha1virtualenvironmentpool "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentpool/v1alpha1"
	v1alpha1virtualenvironmentrole "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentrole/v1alpha1"
	v1alpha1virtualenvironmenttime "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmenttime/v1alpha1"
	v1alpha1virtualenvironmentuser "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentuser/v1alpha1"
	v1alpha1virtualenvironmentvm "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/virtualenvironmentvm/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1virtual.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironment.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentacl.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentacmeaccount.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentacmednsplugin.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentaptrepository.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentaptstandardrepository.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentclusterfirewallsecuritygroup.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentcontainer.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentdns.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentdownloadfile.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentfile.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentfirewallalias.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentfirewallipset.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentfirewalloptions.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentfirewallrules.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentgroup.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmenthagroup.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentharesource.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmenthosts.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentmetricsserver.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentnetworklinuxbridge.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentnetworklinuxvlan.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentpool.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentrole.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmenttime.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentuser.SchemeBuilder.AddToScheme,
		v1alpha1virtualenvironmentvm.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
