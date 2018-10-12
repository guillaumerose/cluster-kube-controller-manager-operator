// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerConfig) DeepCopyInto(out *KubeControllerManagerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerConfig.
func (in *KubeControllerManagerConfig) DeepCopy() *KubeControllerManagerConfig {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllerManagerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerOperatorConfig) DeepCopyInto(out *KubeControllerManagerOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerOperatorConfig.
func (in *KubeControllerManagerOperatorConfig) DeepCopy() *KubeControllerManagerOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllerManagerOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerOperatorConfigList) DeepCopyInto(out *KubeControllerManagerOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeControllerManagerOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerOperatorConfigList.
func (in *KubeControllerManagerOperatorConfigList) DeepCopy() *KubeControllerManagerOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllerManagerOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerOperatorConfigSpec) DeepCopyInto(out *KubeControllerManagerOperatorConfigSpec) {
	*out = *in
	out.OperatorSpec = in.OperatorSpec
	in.KubeControllerManagerConfig.DeepCopyInto(&out.KubeControllerManagerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerOperatorConfigSpec.
func (in *KubeControllerManagerOperatorConfigSpec) DeepCopy() *KubeControllerManagerOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllerManagerOperatorConfigStatus) DeepCopyInto(out *KubeControllerManagerOperatorConfigStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllerManagerOperatorConfigStatus.
func (in *KubeControllerManagerOperatorConfigStatus) DeepCopy() *KubeControllerManagerOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KubeControllerManagerOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}