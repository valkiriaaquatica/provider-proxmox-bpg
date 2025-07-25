//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentNetworkLinuxBridge) DeepCopyInto(out *EnvironmentNetworkLinuxBridge) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentNetworkLinuxBridge.
func (in *EnvironmentNetworkLinuxBridge) DeepCopy() *EnvironmentNetworkLinuxBridge {
	if in == nil {
		return nil
	}
	out := new(EnvironmentNetworkLinuxBridge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentNetworkLinuxBridge) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentNetworkLinuxBridgeInitParameters) DeepCopyInto(out *EnvironmentNetworkLinuxBridgeInitParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Address6 != nil {
		in, out := &in.Address6, &out.Address6
		*out = new(string)
		**out = **in
	}
	if in.Autostart != nil {
		in, out := &in.Autostart, &out.Autostart
		*out = new(bool)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Gateway6 != nil {
		in, out := &in.Gateway6, &out.Gateway6
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VlanAware != nil {
		in, out := &in.VlanAware, &out.VlanAware
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentNetworkLinuxBridgeInitParameters.
func (in *EnvironmentNetworkLinuxBridgeInitParameters) DeepCopy() *EnvironmentNetworkLinuxBridgeInitParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentNetworkLinuxBridgeInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentNetworkLinuxBridgeList) DeepCopyInto(out *EnvironmentNetworkLinuxBridgeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EnvironmentNetworkLinuxBridge, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentNetworkLinuxBridgeList.
func (in *EnvironmentNetworkLinuxBridgeList) DeepCopy() *EnvironmentNetworkLinuxBridgeList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentNetworkLinuxBridgeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentNetworkLinuxBridgeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentNetworkLinuxBridgeObservation) DeepCopyInto(out *EnvironmentNetworkLinuxBridgeObservation) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Address6 != nil {
		in, out := &in.Address6, &out.Address6
		*out = new(string)
		**out = **in
	}
	if in.Autostart != nil {
		in, out := &in.Autostart, &out.Autostart
		*out = new(bool)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Gateway6 != nil {
		in, out := &in.Gateway6, &out.Gateway6
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VlanAware != nil {
		in, out := &in.VlanAware, &out.VlanAware
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentNetworkLinuxBridgeObservation.
func (in *EnvironmentNetworkLinuxBridgeObservation) DeepCopy() *EnvironmentNetworkLinuxBridgeObservation {
	if in == nil {
		return nil
	}
	out := new(EnvironmentNetworkLinuxBridgeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentNetworkLinuxBridgeParameters) DeepCopyInto(out *EnvironmentNetworkLinuxBridgeParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.Address6 != nil {
		in, out := &in.Address6, &out.Address6
		*out = new(string)
		**out = **in
	}
	if in.Autostart != nil {
		in, out := &in.Autostart, &out.Autostart
		*out = new(bool)
		**out = **in
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(string)
		**out = **in
	}
	if in.Gateway6 != nil {
		in, out := &in.Gateway6, &out.Gateway6
		*out = new(string)
		**out = **in
	}
	if in.Mtu != nil {
		in, out := &in.Mtu, &out.Mtu
		*out = new(float64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NodeName != nil {
		in, out := &in.NodeName, &out.NodeName
		*out = new(string)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.VlanAware != nil {
		in, out := &in.VlanAware, &out.VlanAware
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentNetworkLinuxBridgeParameters.
func (in *EnvironmentNetworkLinuxBridgeParameters) DeepCopy() *EnvironmentNetworkLinuxBridgeParameters {
	if in == nil {
		return nil
	}
	out := new(EnvironmentNetworkLinuxBridgeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentNetworkLinuxBridgeSpec) DeepCopyInto(out *EnvironmentNetworkLinuxBridgeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentNetworkLinuxBridgeSpec.
func (in *EnvironmentNetworkLinuxBridgeSpec) DeepCopy() *EnvironmentNetworkLinuxBridgeSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentNetworkLinuxBridgeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentNetworkLinuxBridgeStatus) DeepCopyInto(out *EnvironmentNetworkLinuxBridgeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentNetworkLinuxBridgeStatus.
func (in *EnvironmentNetworkLinuxBridgeStatus) DeepCopy() *EnvironmentNetworkLinuxBridgeStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentNetworkLinuxBridgeStatus)
	in.DeepCopyInto(out)
	return out
}
