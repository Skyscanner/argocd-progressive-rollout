// +build !ignore_autogenerated

/**
 * Copyright 2021 Skyscanner Limited.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clusters) DeepCopyInto(out *Clusters) {
	*out = *in
	in.Selector.DeepCopyInto(&out.Selector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clusters.
func (in *Clusters) DeepCopy() *Clusters {
	if in == nil {
		return nil
	}
	out := new(Clusters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRollout) DeepCopyInto(out *ProgressiveRollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRollout.
func (in *ProgressiveRollout) DeepCopy() *ProgressiveRollout {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProgressiveRollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutList) DeepCopyInto(out *ProgressiveRolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProgressiveRollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutList.
func (in *ProgressiveRolloutList) DeepCopy() *ProgressiveRolloutList {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProgressiveRolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutSpec) DeepCopyInto(out *ProgressiveRolloutSpec) {
	*out = *in
	in.SourceRef.DeepCopyInto(&out.SourceRef)
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]ProgressiveRolloutStage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutSpec.
func (in *ProgressiveRolloutSpec) DeepCopy() *ProgressiveRolloutSpec {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutStage) DeepCopyInto(out *ProgressiveRolloutStage) {
	*out = *in
	out.MaxParallel = in.MaxParallel
	out.MaxTargets = in.MaxTargets
	in.Targets.DeepCopyInto(&out.Targets)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutStage.
func (in *ProgressiveRolloutStage) DeepCopy() *ProgressiveRolloutStage {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutStatus) DeepCopyInto(out *ProgressiveRolloutStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]StageStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutStatus.
func (in *ProgressiveRolloutStatus) DeepCopy() *ProgressiveRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProgressiveRolloutTargets) DeepCopyInto(out *ProgressiveRolloutTargets) {
	*out = *in
	in.Clusters.DeepCopyInto(&out.Clusters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProgressiveRolloutTargets.
func (in *ProgressiveRolloutTargets) DeepCopy() *ProgressiveRolloutTargets {
	if in == nil {
		return nil
	}
	out := new(ProgressiveRolloutTargets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageStatus) DeepCopyInto(out *StageStatus) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Syncing != nil {
		in, out := &in.Syncing, &out.Syncing
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Requeued != nil {
		in, out := &in.Requeued, &out.Requeued
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Failed != nil {
		in, out := &in.Failed, &out.Failed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.StartedAt.DeepCopyInto(&out.StartedAt)
	in.FinishedAt.DeepCopyInto(&out.FinishedAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageStatus.
func (in *StageStatus) DeepCopy() *StageStatus {
	if in == nil {
		return nil
	}
	out := new(StageStatus)
	in.DeepCopyInto(out)
	return out
}
