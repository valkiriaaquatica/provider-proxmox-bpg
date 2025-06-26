// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/providerconfig"
	environmentclusterfirewall "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtual/environmentclusterfirewall"
	environmentcertificate "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironment/environmentcertificate"
	environmentacl "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentacl/environmentacl"
	environmentacmeaccount "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentacmeaccount/environmentacmeaccount"
	environmentacmednsplugin "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentacmednsplugin/environmentacmednsplugin"
	environmentaptrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentaptrepository/environmentaptrepository"
	environmentaptstandardrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentaptstandardrepository/environmentaptstandardrepository"
	environmentclusterfirewallsecuritygroup "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentclusterfirewallsecuritygroup/environmentclusterfirewallsecuritygroup"
	environmentcontainer "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentcontainer/environmentcontainer"
	environmentdns "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentdns/environmentdns"
	environmentdownloadfile "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentdownloadfile/environmentdownloadfile"
	environmentfile "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentfile/environmentfile"
	environmentfirewallalias "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentfirewallalias/environmentfirewallalias"
	environmentfirewallipset "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentfirewallipset/environmentfirewallipset"
	environmentfirewalloptions "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentfirewalloptions/environmentfirewalloptions"
	environmentfirewallrules "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentfirewallrules/environmentfirewallrules"
	environmentgroup "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentgroup/environmentgroup"
	environmenthagroup "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmenthagroup/environmenthagroup"
	environmentharesource "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentharesource/environmentharesource"
	environmenthosts "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmenthosts/environmenthosts"
	environmentmetricsserver "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentmetricsserver/environmentmetricsserver"
	environmentnetworklinuxbridge "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentnetworklinuxbridge/environmentnetworklinuxbridge"
	environmentnetworklinuxvlan "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentnetworklinuxvlan/environmentnetworklinuxvlan"
	environmentpool "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentpool/environmentpool"
	environmentrole "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentrole/environmentrole"
	environmenttime "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmenttime/environmenttime"
	environmentuser "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentuser/environmentuser"
	environmentvm "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentvm/environmentvm"
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
