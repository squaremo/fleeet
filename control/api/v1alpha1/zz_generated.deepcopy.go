// +build !ignore_autogenerated

/*
Copyright 2021 Michael Bridgen <mikeb@squaremobius.net>.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/squaremo/fleeet/pkg/api"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapModule) DeepCopyInto(out *BootstrapModule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapModule.
func (in *BootstrapModule) DeepCopy() *BootstrapModule {
	if in == nil {
		return nil
	}
	out := new(BootstrapModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BootstrapModule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapModuleList) DeepCopyInto(out *BootstrapModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BootstrapModule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapModuleList.
func (in *BootstrapModuleList) DeepCopy() *BootstrapModuleList {
	if in == nil {
		return nil
	}
	out := new(BootstrapModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BootstrapModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapModuleSpec) DeepCopyInto(out *BootstrapModuleSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Sync.DeepCopyInto(&out.Sync)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapModuleSpec.
func (in *BootstrapModuleSpec) DeepCopy() *BootstrapModuleSpec {
	if in == nil {
		return nil
	}
	out := new(BootstrapModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapModuleStatus) DeepCopyInto(out *BootstrapModuleStatus) {
	*out = *in
	if in.ObservedSync != nil {
		in, out := &in.ObservedSync, &out.ObservedSync
		*out = new(api.Sync)
		(*in).DeepCopyInto(*out)
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(SyncSummary)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapModuleStatus.
func (in *BootstrapModuleStatus) DeepCopy() *BootstrapModuleStatus {
	if in == nil {
		return nil
	}
	out := new(BootstrapModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalKubeconfigReference) DeepCopyInto(out *LocalKubeconfigReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalKubeconfigReference.
func (in *LocalKubeconfigReference) DeepCopy() *LocalKubeconfigReference {
	if in == nil {
		return nil
	}
	out := new(LocalKubeconfigReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Module) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleList) DeepCopyInto(out *ModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleList.
func (in *ModuleList) DeepCopy() *ModuleList {
	if in == nil {
		return nil
	}
	out := new(ModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleSpec) DeepCopyInto(out *ModuleSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Sync.DeepCopyInto(&out.Sync)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleSpec.
func (in *ModuleSpec) DeepCopy() *ModuleSpec {
	if in == nil {
		return nil
	}
	out := new(ModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModuleStatus) DeepCopyInto(out *ModuleStatus) {
	*out = *in
	if in.ObservedSync != nil {
		in, out := &in.ObservedSync, &out.ObservedSync
		*out = new(api.Sync)
		(*in).DeepCopyInto(*out)
	}
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = new(SyncSummary)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModuleStatus.
func (in *ModuleStatus) DeepCopy() *ModuleStatus {
	if in == nil {
		return nil
	}
	out := new(ModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteAssemblage) DeepCopyInto(out *RemoteAssemblage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteAssemblage.
func (in *RemoteAssemblage) DeepCopy() *RemoteAssemblage {
	if in == nil {
		return nil
	}
	out := new(RemoteAssemblage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteAssemblage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteAssemblageList) DeepCopyInto(out *RemoteAssemblageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RemoteAssemblage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteAssemblageList.
func (in *RemoteAssemblageList) DeepCopy() *RemoteAssemblageList {
	if in == nil {
		return nil
	}
	out := new(RemoteAssemblageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteAssemblageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteAssemblageSpec) DeepCopyInto(out *RemoteAssemblageSpec) {
	*out = *in
	out.KubeconfigRef = in.KubeconfigRef
	in.Assemblage.DeepCopyInto(&out.Assemblage)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteAssemblageSpec.
func (in *RemoteAssemblageSpec) DeepCopy() *RemoteAssemblageSpec {
	if in == nil {
		return nil
	}
	out := new(RemoteAssemblageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteAssemblageStatus) DeepCopyInto(out *RemoteAssemblageStatus) {
	*out = *in
	if in.Syncs != nil {
		in, out := &in.Syncs, &out.Syncs
		*out = make([]api.SyncStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteAssemblageStatus.
func (in *RemoteAssemblageStatus) DeepCopy() *RemoteAssemblageStatus {
	if in == nil {
		return nil
	}
	out := new(RemoteAssemblageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SyncSummary) DeepCopyInto(out *SyncSummary) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SyncSummary.
func (in *SyncSummary) DeepCopy() *SyncSummary {
	if in == nil {
		return nil
	}
	out := new(SyncSummary)
	in.DeepCopyInto(out)
	return out
}
