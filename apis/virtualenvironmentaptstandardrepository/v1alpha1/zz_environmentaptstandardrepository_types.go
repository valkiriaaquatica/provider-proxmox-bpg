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

type EnvironmentAptStandardRepositoryInitParameters struct {

	// quincy-enterprise | ceph-quincy-no-subscription | ceph-quincy-test | ceph-reef-enterprise | ceph-reef-no-subscription | ceph-reef-test | enterprise | no-subscription | test.
	// The handle of the APT standard repository. Must be `ceph-quincy-enterprise` | `ceph-quincy-no-subscription` | `ceph-quincy-test` | `ceph-reef-enterprise` | `ceph-reef-no-subscription` | `ceph-reef-test` | `enterprise` | `no-subscription` | `test`.
	Handle *string `json:"handle,omitempty" tf:"handle,omitempty"`

	// (String) The name of the target Proxmox VE node.
	// The name of the target Proxmox VE node.
	Node *string `json:"node,omitempty" tf:"node,omitempty"`
}

type EnvironmentAptStandardRepositoryObservation struct {

	// (String) The description of the APT standard repository.
	// The description of the APT standard repository.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The absolute path of the source list file that contains this standard repository.
	// The absolute path of the source list file that contains this standard repository.
	FilePath *string `json:"filePath,omitempty" tf:"file_path,omitempty"`

	// quincy-enterprise | ceph-quincy-no-subscription | ceph-quincy-test | ceph-reef-enterprise | ceph-reef-no-subscription | ceph-reef-test | enterprise | no-subscription | test.
	// The handle of the APT standard repository. Must be `ceph-quincy-enterprise` | `ceph-quincy-no-subscription` | `ceph-quincy-test` | `ceph-reef-enterprise` | `ceph-reef-no-subscription` | `ceph-reef-test` | `enterprise` | `no-subscription` | `test`.
	Handle *string `json:"handle,omitempty" tf:"handle,omitempty"`

	// (String) The unique identifier of this APT standard repository resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number) The index within the defining source list file.
	// The index within the defining source list file.
	Index *float64 `json:"index,omitempty" tf:"index,omitempty"`

	// (String) The name of the APT standard repository.
	// The name of the APT standard repository.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The name of the target Proxmox VE node.
	// The name of the target Proxmox VE node.
	Node *string `json:"node,omitempty" tf:"node,omitempty"`

	// (Number) Indicates the activation status.
	// Indicates the activation status.
	Status *float64 `json:"status,omitempty" tf:"status,omitempty"`
}

type EnvironmentAptStandardRepositoryParameters struct {

	// quincy-enterprise | ceph-quincy-no-subscription | ceph-quincy-test | ceph-reef-enterprise | ceph-reef-no-subscription | ceph-reef-test | enterprise | no-subscription | test.
	// The handle of the APT standard repository. Must be `ceph-quincy-enterprise` | `ceph-quincy-no-subscription` | `ceph-quincy-test` | `ceph-reef-enterprise` | `ceph-reef-no-subscription` | `ceph-reef-test` | `enterprise` | `no-subscription` | `test`.
	// +kubebuilder:validation:Optional
	Handle *string `json:"handle,omitempty" tf:"handle,omitempty"`

	// (String) The name of the target Proxmox VE node.
	// The name of the target Proxmox VE node.
	// +kubebuilder:validation:Optional
	Node *string `json:"node,omitempty" tf:"node,omitempty"`
}

// EnvironmentAptStandardRepositorySpec defines the desired state of EnvironmentAptStandardRepository
type EnvironmentAptStandardRepositorySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentAptStandardRepositoryParameters `json:"forProvider"`
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
	InitProvider EnvironmentAptStandardRepositoryInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentAptStandardRepositoryStatus defines the observed state of EnvironmentAptStandardRepository.
type EnvironmentAptStandardRepositoryStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentAptStandardRepositoryObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EnvironmentAptStandardRepository is the Schema for the EnvironmentAptStandardRepositorys API. Manages an APT standard repository of a Proxmox VE node.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmoxbpg}
type EnvironmentAptStandardRepository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.handle) || (has(self.initProvider) && has(self.initProvider.handle))",message="spec.forProvider.handle is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.node) || (has(self.initProvider) && has(self.initProvider.node))",message="spec.forProvider.node is a required parameter"
	Spec   EnvironmentAptStandardRepositorySpec   `json:"spec"`
	Status EnvironmentAptStandardRepositoryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentAptStandardRepositoryList contains a list of EnvironmentAptStandardRepositorys
type EnvironmentAptStandardRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentAptStandardRepository `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentAptStandardRepository_Kind             = "EnvironmentAptStandardRepository"
	EnvironmentAptStandardRepository_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentAptStandardRepository_Kind}.String()
	EnvironmentAptStandardRepository_KindAPIVersion   = EnvironmentAptStandardRepository_Kind + "." + CRDGroupVersion.String()
	EnvironmentAptStandardRepository_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentAptStandardRepository_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentAptStandardRepository{}, &EnvironmentAptStandardRepositoryList{})
}
