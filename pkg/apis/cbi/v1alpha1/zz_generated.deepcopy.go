// +build !ignore_autogenerated

/*
Copyright The CBI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildJob) DeepCopyInto(out *BuildJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildJob.
func (in *BuildJob) DeepCopy() *BuildJob {
	if in == nil {
		return nil
	}
	out := new(BuildJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildJobList) DeepCopyInto(out *BuildJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BuildJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildJobList.
func (in *BuildJobList) DeepCopy() *BuildJobList {
	if in == nil {
		return nil
	}
	out := new(BuildJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildJobSpec) DeepCopyInto(out *BuildJobSpec) {
	*out = *in
	out.Registry = in.Registry
	out.Language = in.Language
	out.Context = in.Context
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildJobSpec.
func (in *BuildJobSpec) DeepCopy() *BuildJobSpec {
	if in == nil {
		return nil
	}
	out := new(BuildJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildJobStatus) DeepCopyInto(out *BuildJobStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildJobStatus.
func (in *BuildJobStatus) DeepCopy() *BuildJobStatus {
	if in == nil {
		return nil
	}
	out := new(BuildJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cloudbuild) DeepCopyInto(out *Cloudbuild) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cloudbuild.
func (in *Cloudbuild) DeepCopy() *Cloudbuild {
	if in == nil {
		return nil
	}
	out := new(Cloudbuild)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	out.Git = in.Git
	out.ConfigMapRef = in.ConfigMapRef
	out.HTTP = in.HTTP
	out.Rclone = in.Rclone
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dockerfile) DeepCopyInto(out *Dockerfile) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dockerfile.
func (in *Dockerfile) DeepCopy() *Dockerfile {
	if in == nil {
		return nil
	}
	out := new(Dockerfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Git) DeepCopyInto(out *Git) {
	*out = *in
	out.SSHSecretRef = in.SSHSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Git.
func (in *Git) DeepCopy() *Git {
	if in == nil {
		return nil
	}
	out := new(Git)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTP) DeepCopyInto(out *HTTP) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTP.
func (in *HTTP) DeepCopy() *HTTP {
	if in == nil {
		return nil
	}
	out := new(HTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Language) DeepCopyInto(out *Language) {
	*out = *in
	out.Dockerfile = in.Dockerfile
	out.S2I = in.S2I
	out.Cloudbuild = in.Cloudbuild
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Language.
func (in *Language) DeepCopy() *Language {
	if in == nil {
		return nil
	}
	out := new(Language)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rclone) DeepCopyInto(out *Rclone) {
	*out = *in
	out.SecretRef = in.SecretRef
	out.SSHSecretRef = in.SSHSecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rclone.
func (in *Rclone) DeepCopy() *Rclone {
	if in == nil {
		return nil
	}
	out := new(Rclone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Registry) DeepCopyInto(out *Registry) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Registry.
func (in *Registry) DeepCopy() *Registry {
	if in == nil {
		return nil
	}
	out := new(Registry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S2I) DeepCopyInto(out *S2I) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S2I.
func (in *S2I) DeepCopy() *S2I {
	if in == nil {
		return nil
	}
	out := new(S2I)
	in.DeepCopyInto(out)
	return out
}
