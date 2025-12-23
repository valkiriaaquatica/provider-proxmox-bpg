// SPDX-FileCopyrightText: 2025 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	clusterv1beta1 "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/valkiriaaquatica/provider-proxmox-bpg/apis/namespaced/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal proxmox credentials as JSON"
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
// returns Terraform provider setup configuration.
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, c client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		credsSpec, err := resolveProviderConfig(ctx, c, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		creds, err := loadCredentials(ctx, c, credsSpec)
		if err != nil {
			return terraform.Setup{}, err
		}

		ps.Configuration = buildConfiguration(creds)
		return ps, nil
	}
}

func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*clusterv1beta1.ProviderCredentials, error) {
	modern, ok := mg.(resource.ModernManaged)
	if !ok {
		return nil, errors.New("resource is not a managed resource")
	}
	return resolveModern(ctx, crClient, modern)
}

func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*clusterv1beta1.ProviderCredentials, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, "unknown GVK for ProviderConfig")
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		return nil, errors.New("provider config type is not a client.Object")
	}

	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		t := resource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return nil, errors.Wrap(err, errTrackUsage)
		}
		if pc.Spec.Credentials.SecretRef != nil {
			pc.Spec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
		return &clusterv1beta1.ProviderCredentials{
			Source:                    pc.Spec.Credentials.Source,
			CommonCredentialSelectors: pc.Spec.Credentials.CommonCredentialSelectors,
		}, nil
	case *namespacedv1beta1.ClusterProviderConfig:
		t := resource.NewProviderConfigUsageTracker(crClient, &namespacedv1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return nil, errors.Wrap(err, errTrackUsage)
		}
		return &clusterv1beta1.ProviderCredentials{
			Source:                    pc.Spec.Credentials.Source,
			CommonCredentialSelectors: pc.Spec.Credentials.CommonCredentialSelectors,
		}, nil
	case *clusterv1beta1.ProviderConfig:
		t := resource.NewProviderConfigUsageTracker(crClient, &clusterv1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return nil, errors.Wrap(err, errTrackUsage)
		}
		return &pc.Spec.Credentials, nil
	default:
		return nil, errors.New("unknown provider config type")
	}
}

func loadCredentials(ctx context.Context, c client.Client, credsSpec *clusterv1beta1.ProviderCredentials) (map[string]string, error) {
	data, err := resource.CommonCredentialExtractor(ctx, credsSpec.Source, c, credsSpec.CommonCredentialSelectors)
	if err != nil {
		return nil, errors.Wrap(err, errExtractCredentials)
	}

	creds := map[string]string{}
	if err := json.Unmarshal(data, &creds); err != nil {
		return nil, errors.Wrap(err, errUnmarshalCredentials)
	}

	// If the Secret does not contain API token, but contains auth ticket and
	// CSRF token, then set the API token to empty string so that the provider
	// will use the auth ticket and CSRF token.
	if creds[keyAPIToken] == "" && creds[keyAuthTicket] != "" && creds[keyCSRFPreventionToken] != "" {
		creds[keyAPIToken] = ""
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
