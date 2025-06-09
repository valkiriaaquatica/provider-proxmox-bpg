// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	providerconfig "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/providerconfig"
	environmentacl "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentacl/environmentacl"
	environmentacmeaccount "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentacmeaccount/environmentacmeaccount"
	environmentacmednsplugin "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentacmednsplugin/environmentacmednsplugin"
	environmentaptrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentaptrepository/environmentaptrepository"
	environmentaptstandardrepository "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentaptstandardrepository/environmentaptstandardrepository"
	environmentrole "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentrole/environmentrole"
	environmentuser "github.com/valkiriaaquatica/provider-proxmox-bpg/internal/controller/virtualenvironmentuser/environmentuser"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		environmentacl.Setup,
		environmentacmeaccount.Setup,
		environmentacmednsplugin.Setup,
		environmentaptrepository.Setup,
		environmentaptstandardrepository.Setup,
		environmentrole.Setup,
		environmentuser.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
