// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EnvironmentAptRepositoryInitParameters struct {

	// (Boolean) Indicates the activation status.
	// Indicates the activation status.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The absolute path of the source list file that contains this repository.
	// The absolute path of the source list file that contains this repository.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// (Number) The index within the defining source list file.
	// The index within the defining source list file.
	Index *float64 `json:"index,omitempty" tf:"index,omitempty"`

	// (String) The name of the target Proxmox VE node.
	// The name of the target Proxmox VE node.
	Node *string `json:"node,omitempty" tf:"node,omitempty"`
}

type EnvironmentAptRepositoryObservation struct {

	// (String) The associated comment.
	// The associated comment.
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// (List of String) The list of components.
	// The list of components.
	Components []*string `json:"components,omitempty" tf:"components,omitempty"`

	// (Boolean) Indicates the activation status.
	// Indicates the activation status.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The absolute path of the source list file that contains this repository.
	// The absolute path of the source list file that contains this repository.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// (String) The format of the defining source list file.
	// The format of the defining source list file.
	FileType *string `json:"fileType,omitempty" tf:"file_type,omitempty"`

	// (String) The unique identifier of this APT repository resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) The index within the defining source list file.
	// The index within the defining source list file.
	Index *float64 `json:"index,omitempty" tf:"index,omitempty"`

	// (String) The name of the target Proxmox VE node.
	// The name of the target Proxmox VE node.
	Node *string `json:"node,omitempty" tf:"node,omitempty"`

	// (List of String) The list of package types.
	// The list of package types.
	PackageTypes []*string `json:"packageTypes,omitempty" tf:"package_types,omitempty"`

	// (List of String) The list of package distributions.
	// The list of package distributions.
	Suites []*string `json:"suites,omitempty" tf:"suites,omitempty"`

	// (List of String) The list of repository URIs.
	// The list of repository URIs.
	Uris []*string `json:"uris,omitempty" tf:"uris,omitempty"`
}

type EnvironmentAptRepositoryParameters struct {

	// (Boolean) Indicates the activation status.
	// Indicates the activation status.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) The absolute path of the source list file that contains this repository.
	// The absolute path of the source list file that contains this repository.
	// +kubebuilder:validation:Optional
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// (Number) The index within the defining source list file.
	// The index within the defining source list file.
	// +kubebuilder:validation:Optional
	Index *float64 `json:"index,omitempty" tf:"index,omitempty"`

	// (String) The name of the target Proxmox VE node.
	// The name of the target Proxmox VE node.
	// +kubebuilder:validation:Optional
	Node *string `json:"node,omitempty" tf:"node,omitempty"`
}

// EnvironmentAptRepositorySpec defines the desired state of EnvironmentAptRepository
type EnvironmentAptRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentAptRepositoryParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider EnvironmentAptRepositoryInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentAptRepositoryStatus defines the observed state of EnvironmentAptRepository.
type EnvironmentAptRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentAptRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EnvironmentAptRepository is the Schema for the EnvironmentAptRepositorys API. Manages an APT repository of a Proxmox VE node.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmoxbpg}
type EnvironmentAptRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.filePath) || (has(self.initProvider) && has(self.initProvider.filePath))",message="spec.forProvider.filePath is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.index) || (has(self.initProvider) && has(self.initProvider.index))",message="spec.forProvider.index is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.node) || (has(self.initProvider) && has(self.initProvider.node))",message="spec.forProvider.node is a required parameter"
	Spec   EnvironmentAptRepositorySpec   `json:"spec"`
	Status EnvironmentAptRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentAptRepositoryList contains a list of EnvironmentAptRepositorys
type EnvironmentAptRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentAptRepository `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentAptRepository_Kind             = "EnvironmentAptRepository"
	EnvironmentAptRepository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentAptRepository_Kind}.String()
	EnvironmentAptRepository_KindAPIVersion   = EnvironmentAptRepository_Kind + "." + CRDGroupVersion.String()
	EnvironmentAptRepository_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentAptRepository_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentAptRepository{}, &EnvironmentAptRepositoryList{})
}
