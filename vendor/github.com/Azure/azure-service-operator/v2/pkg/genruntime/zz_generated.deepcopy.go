//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package genruntime

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArbitraryOwnerReference) DeepCopyInto(out *ArbitraryOwnerReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArbitraryOwnerReference.
func (in *ArbitraryOwnerReference) DeepCopy() *ArbitraryOwnerReference {
	if in == nil {
		return nil
	}
	out := new(ArbitraryOwnerReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapReference) DeepCopyInto(out *ConfigMapReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapReference.
func (in *ConfigMapReference) DeepCopy() *ConfigMapReference {
	if in == nil {
		return nil
	}
	out := new(ConfigMapReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KnownResourceReference) DeepCopyInto(out *KnownResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KnownResourceReference.
func (in *KnownResourceReference) DeepCopy() *KnownResourceReference {
	if in == nil {
		return nil
	}
	out := new(KnownResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesOwnerReference) DeepCopyInto(out *KubernetesOwnerReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesOwnerReference.
func (in *KubernetesOwnerReference) DeepCopy() *KubernetesOwnerReference {
	if in == nil {
		return nil
	}
	out := new(KubernetesOwnerReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}
