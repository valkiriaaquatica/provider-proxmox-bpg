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
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}
		if err := populateProviderConfig(ctx, c, mg, &ps); err != nil {
			return ps, err
		}
		return ps, nil
	}
}

func populateProviderConfig(ctx context.Context, c client.Client, mg resource.Managed, ps *terraform.Setup) error {
	ref := mg.GetProviderConfigReference()
	if ref == nil {
		return errors.New(errNoProviderConfig)
	}

	pc := &v1beta1.ProviderConfig{}
	if err := c.Get(ctx, types.NamespacedName{Name: ref.Name}, pc); err != nil {
		return errors.Wrap(err, errGetProviderConfig)
	}

	if err := trackProviderUsage(ctx, c, mg); err != nil {
		return err
	}

	creds, err := loadCredentials(ctx, c, pc)
	if err != nil {
		return err
	}

	ps.Configuration = buildConfiguration(creds)
	return nil
}

func trackProviderUsage(ctx context.Context, c client.Client, mg resource.Managed) error {
	tracker := resource.NewProviderConfigUsageTracker(c, &v1beta1.ProviderConfigUsage{})
	return errors.Wrap(tracker.Track(ctx, mg), errTrackUsage)
}

func loadCredentials(ctx context.Context, c client.Client, pc *v1beta1.ProviderConfig) (map[string]string, error) {
	data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, c, pc.Spec.Credentials.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errExtractCredentials)
	}
	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}
	return creds, nil
}

func buildConfiguration(creds map[string]string) map[string]any {
	cfg := map[string]any{}

	add := func(k string) {
		if v := creds[k]; v != "" {
			cfg[k] = v
		}
	}

	for _, k := range []string{
		keyEndpoint, keyUsername, keyPassword, keyAPIToken,
		keyAuthTicket, keyCSRFPreventionToken, keyInsecure, keyTmpDir,
	} {
		add(k)
	}

	ssh := map[string]any{}
	if v := creds[keySSHUsername]; v != "" {
		ssh["username"] = v
	}
	if v := creds[keySSHPassword]; v != "" {
		ssh["password"] = v
	}
	if v := creds[keySSHPrivateKey]; v != "" {
		ssh["private_key"] = v
	}
	if len(ssh) > 0 {
		cfg["ssh"] = ssh
	}

	return cfg
}
