// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtradbCluster) DeepCopyInto(out *PerconaXtradbCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtradbCluster.
func (in *PerconaXtradbCluster) DeepCopy() *PerconaXtradbCluster {
	if in == nil {
		return nil
	}
	out := new(PerconaXtradbCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtradbCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtradbClusterList) DeepCopyInto(out *PerconaXtradbClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PerconaXtradbCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtradbClusterList.
func (in *PerconaXtradbClusterList) DeepCopy() *PerconaXtradbClusterList {
	if in == nil {
		return nil
	}
	out := new(PerconaXtradbClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PerconaXtradbClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtradbClusterSpec) DeepCopyInto(out *PerconaXtradbClusterSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtradbClusterSpec.
func (in *PerconaXtradbClusterSpec) DeepCopy() *PerconaXtradbClusterSpec {
	if in == nil {
		return nil
	}
	out := new(PerconaXtradbClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PerconaXtradbClusterStatus) DeepCopyInto(out *PerconaXtradbClusterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PerconaXtradbClusterStatus.
func (in *PerconaXtradbClusterStatus) DeepCopy() *PerconaXtradbClusterStatus {
	if in == nil {
		return nil
	}
	out := new(PerconaXtradbClusterStatus)
	in.DeepCopyInto(out)
	return out
}
