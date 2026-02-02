// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/providerconfig"
	environmentclusterfirewall "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtual/environmentclusterfirewall"
	environmentcertificate "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironment/environmentcertificate"
	environmentacl "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentacl/environmentacl"
	environmentacmeaccount "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentacmeaccount/environmentacmeaccount"
	environmentacmednsplugin "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentacmednsplugin/environmentacmednsplugin"
	environmentaptrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentaptrepository/environmentaptrepository"
	environmentaptstandardrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentaptstandardrepository/environmentaptstandardrepository"
	environmentclusterfirewallsecuritygroup "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentclusterfirewallsecuritygroup/environmentclusterfirewallsecuritygroup"
	environmentcontainer "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentcontainer/environmentcontainer"
	environmentdns "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentdns/environmentdns"
	environmentdownloadfile "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentdownloadfile/environmentdownloadfile"
	environmentfile "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentfile/environmentfile"
	environmentfirewallalias "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentfirewallalias/environmentfirewallalias"
	environmentfirewallipset "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentfirewallipset/environmentfirewallipset"
	environmentfirewalloptions "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentfirewalloptions/environmentfirewalloptions"
	environmentfirewallrules "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentfirewallrules/environmentfirewallrules"
	environmentgroup "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentgroup/environmentgroup"
	environmenthagroup "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmenthagroup/environmenthagroup"
	environmentharesource "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentharesource/environmentharesource"
	environmenthosts "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmenthosts/environmenthosts"
	environmentmetricsserver "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentmetricsserver/environmentmetricsserver"
	environmentnetworklinuxbridge "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentnetworklinuxbridge/environmentnetworklinuxbridge"
	environmentnetworklinuxvlan "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentnetworklinuxvlan/environmentnetworklinuxvlan"
	environmentnodefirewall "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentnodefirewall/environmentnodefirewall"
	environmentpool "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentpool/environmentpool"
	environmentrole "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentrole/environmentrole"
	environmenttime "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmenttime/environmenttime"
	environmentuser "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentuser/environmentuser"
	environmentvm "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/cluster/virtualenvironmentvm/environmentvm"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		environmentclusterfirewall.Setup,
		environmentcertificate.Setup,
		environmentacl.Setup,
		environmentacmeaccount.Setup,
		environmentacmednsplugin.Setup,
		environmentaptrepository.Setup,
		environmentaptstandardrepository.Setup,
		environmentclusterfirewallsecuritygroup.Setup,
		environmentcontainer.Setup,
		environmentdns.Setup,
		environmentdownloadfile.Setup,
		environmentfile.Setup,
		environmentfirewallalias.Setup,
		environmentfirewallipset.Setup,
		environmentfirewalloptions.Setup,
		environmentfirewallrules.Setup,
		environmentgroup.Setup,
		environmenthagroup.Setup,
		environmentharesource.Setup,
		environmenthosts.Setup,
		environmentmetricsserver.Setup,
		environmentnetworklinuxbridge.Setup,
		environmentnetworklinuxvlan.Setup,
		environmentnodefirewall.Setup,
		environmentpool.Setup,
		environmentrole.Setup,
		environmenttime.Setup,
		environmentuser.Setup,
		environmentvm.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.SetupGated,
		environmentclusterfirewall.SetupGated,
		environmentcertificate.SetupGated,
		environmentacl.SetupGated,
		environmentacmeaccount.SetupGated,
		environmentacmednsplugin.SetupGated,
		environmentaptrepository.SetupGated,
		environmentaptstandardrepository.SetupGated,
		environmentclusterfirewallsecuritygroup.SetupGated,
		environmentcontainer.SetupGated,
		environmentdns.SetupGated,
		environmentdownloadfile.SetupGated,
		environmentfile.SetupGated,
		environmentfirewallalias.SetupGated,
		environmentfirewallipset.SetupGated,
		environmentfirewalloptions.SetupGated,
		environmentfirewallrules.SetupGated,
		environmentgroup.SetupGated,
		environmenthagroup.SetupGated,
		environmentharesource.SetupGated,
		environmenthosts.SetupGated,
		environmentmetricsserver.SetupGated,
		environmentnetworklinuxbridge.SetupGated,
		environmentnetworklinuxvlan.SetupGated,
		environmentnodefirewall.SetupGated,
		environmentpool.SetupGated,
		environmentrole.SetupGated,
		environmenttime.SetupGated,
		environmentuser.SetupGated,
		environmentvm.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
