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

type EnvironmentACLInitParameters struct {

	// (String) The group the ACL should apply to (mutually exclusive with token_id and user_id)
	// The group the ACL should apply to (mutually exclusive with `token_id` and `user_id`)
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (String) Access control path
	// Access control path
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Boolean) Allow to propagate (inherit) permissions.
	// Allow to propagate (inherit) permissions.
	Propagate *bool `json:"propagate,omitempty" tf:"propagate,omitempty"`

	// (String) The role to apply
	// The role to apply
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (String) The token the ACL should apply to (mutually exclusive with group_id and user_id)
	// The token the ACL should apply to (mutually exclusive with `group_id` and `user_id`)
	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`

	// (String) The user the ACL should apply to (mutually exclusive with group_id and token_id)
	// The user the ACL should apply to (mutually exclusive with `group_id` and `token_id`)
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type EnvironmentACLObservation struct {

	// (String) The group the ACL should apply to (mutually exclusive with token_id and user_id)
	// The group the ACL should apply to (mutually exclusive with `token_id` and `user_id`)
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (String) The unique identifier of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Access control path
	// Access control path
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Boolean) Allow to propagate (inherit) permissions.
	// Allow to propagate (inherit) permissions.
	Propagate *bool `json:"propagate,omitempty" tf:"propagate,omitempty"`

	// (String) The role to apply
	// The role to apply
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (String) The token the ACL should apply to (mutually exclusive with group_id and user_id)
	// The token the ACL should apply to (mutually exclusive with `group_id` and `user_id`)
	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`

	// (String) The user the ACL should apply to (mutually exclusive with group_id and token_id)
	// The user the ACL should apply to (mutually exclusive with `group_id` and `token_id`)
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

type EnvironmentACLParameters struct {

	// (String) The group the ACL should apply to (mutually exclusive with token_id and user_id)
	// The group the ACL should apply to (mutually exclusive with `token_id` and `user_id`)
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// (String) Access control path
	// Access control path
	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// (Boolean) Allow to propagate (inherit) permissions.
	// Allow to propagate (inherit) permissions.
	// +kubebuilder:validation:Optional
	Propagate *bool `json:"propagate,omitempty" tf:"propagate,omitempty"`

	// (String) The role to apply
	// The role to apply
	// +kubebuilder:validation:Optional
	RoleID *string `json:"roleId,omitempty" tf:"role_id,omitempty"`

	// (String) The token the ACL should apply to (mutually exclusive with group_id and user_id)
	// The token the ACL should apply to (mutually exclusive with `group_id` and `user_id`)
	// +kubebuilder:validation:Optional
	TokenID *string `json:"tokenId,omitempty" tf:"token_id,omitempty"`

	// (String) The user the ACL should apply to (mutually exclusive with group_id and token_id)
	// The user the ACL should apply to (mutually exclusive with `group_id` and `token_id`)
	// +kubebuilder:validation:Optional
	UserID *string `json:"userId,omitempty" tf:"user_id,omitempty"`
}

// EnvironmentACLSpec defines the desired state of EnvironmentACL
type EnvironmentACLSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentACLParameters `json:"forProvider"`
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
	InitProvider EnvironmentACLInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentACLStatus defines the observed state of EnvironmentACL.
type EnvironmentACLStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentACLObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EnvironmentACL is the Schema for the EnvironmentACLs API. Manages ACLs on the Proxmox cluster. ACLs are used to control access to resources in the Proxmox cluster. Each ACL consists of a path, a user, group or token, a role, and a flag to allow propagation of permissions.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,proxmoxbpg}
type EnvironmentACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.path) || (has(self.initProvider) && has(self.initProvider.path))",message="spec.forProvider.path is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleId) || (has(self.initProvider) && has(self.initProvider.roleId))",message="spec.forProvider.roleId is a required parameter"
	Spec   EnvironmentACLSpec   `json:"spec"`
	Status EnvironmentACLStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentACLList contains a list of EnvironmentACLs
type EnvironmentACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentACL `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentACL_Kind             = "EnvironmentACL"
	EnvironmentACL_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentACL_Kind}.String()
	EnvironmentACL_KindAPIVersion   = EnvironmentACL_Kind + "." + CRDGroupVersion.String()
	EnvironmentACL_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentACL_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentACL{}, &EnvironmentACLList{})
}
