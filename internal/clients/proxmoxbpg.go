/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/valkiriaaquatica/provider-proxmox-bpg/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal proxmox  credentials as JSON"
)

const (
	keyEndpoint            = "endpoint"
	keyUsername            = "username"
	keyPassword            = "password"
	keyAPIToken            = "api_token"
	keyAuthTicket          = "auth_ticket"
	keyCSRFPreventionToken = "csrf_prevention_token"
	keyInsecure            = "insecure"
	keySSHUsername         = "ssh_username"
	keySSHPassword         = "ssh_password"
	keySSHPrivateKey       = "ssh_private_key"
	keyTmpDir              = "tmp_dir"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration. Only include
		// keys that have a non-empty value to avoid validation errors from
		// the Terraform provider.
		cfg := map[string]any{}

		if v := creds[keyEndpoint]; v != "" {
			cfg[keyEndpoint] = v
		}
		if v := creds[keyUsername]; v != "" {
			cfg[keyUsername] = v
		}
		if v := creds[keyPassword]; v != "" {
			cfg[keyPassword] = v
		}
		if v := creds[keyAPIToken]; v != "" {
			cfg[keyAPIToken] = v
		}
		if v := creds[keyAuthTicket]; v != "" {
			cfg[keyAuthTicket] = v
		}
		if v := creds[keyCSRFPreventionToken]; v != "" {
			cfg[keyCSRFPreventionToken] = v
		}
		if v := creds[keyInsecure]; v != "" {
			cfg[keyInsecure] = v
		}
		if v := creds[keyTmpDir]; v != "" {
			cfg[keyTmpDir] = v
		}

		sshCfg := map[string]any{}
		if v := creds[keySSHUsername]; v != "" {
			sshCfg["username"] = v
		}
		if v := creds[keySSHPassword]; v != "" {
			sshCfg["password"] = v
		}
		if v := creds[keySSHPrivateKey]; v != "" {
			sshCfg["private_key"] = v
		}
		if len(sshCfg) > 0 {
			cfg["ssh"] = sshCfg
		}

		ps.Configuration = cfg
		return ps, nil
	}
}
