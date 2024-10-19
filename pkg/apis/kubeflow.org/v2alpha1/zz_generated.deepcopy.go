//go:build !ignore_autogenerated

// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v2alpha1

import (
	"k8s.io/api/autoscaling/v2"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/jobset/api/jobset/v1alpha2"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTrainingRuntime) DeepCopyInto(out *ClusterTrainingRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTrainingRuntime.
func (in *ClusterTrainingRuntime) DeepCopy() *ClusterTrainingRuntime {
	if in == nil {
		return nil
	}
	out := new(ClusterTrainingRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTrainingRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTrainingRuntimeList) DeepCopyInto(out *ClusterTrainingRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterTrainingRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTrainingRuntimeList.
func (in *ClusterTrainingRuntimeList) DeepCopy() *ClusterTrainingRuntimeList {
	if in == nil {
		return nil
	}
	out := new(ClusterTrainingRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterTrainingRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerOverride) DeepCopyInto(out *ContainerOverride) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerOverride.
func (in *ContainerOverride) DeepCopy() *ContainerOverride {
	if in == nil {
		return nil
	}
	out := new(ContainerOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoschedulingPodGroupPolicySource) DeepCopyInto(out *CoschedulingPodGroupPolicySource) {
	*out = *in
	if in.ScheduleTimeoutSeconds != nil {
		in, out := &in.ScheduleTimeoutSeconds, &out.ScheduleTimeoutSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoschedulingPodGroupPolicySource.
func (in *CoschedulingPodGroupPolicySource) DeepCopy() *CoschedulingPodGroupPolicySource {
	if in == nil {
		return nil
	}
	out := new(CoschedulingPodGroupPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatasetConfig) DeepCopyInto(out *DatasetConfig) {
	*out = *in
	if in.StorageUri != nil {
		in, out := &in.StorageUri, &out.StorageUri
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatasetConfig.
func (in *DatasetConfig) DeepCopy() *DatasetConfig {
	if in == nil {
		return nil
	}
	out := new(DatasetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InputModel) DeepCopyInto(out *InputModel) {
	*out = *in
	if in.StorageUri != nil {
		in, out := &in.StorageUri, &out.StorageUri
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InputModel.
func (in *InputModel) DeepCopy() *InputModel {
	if in == nil {
		return nil
	}
	out := new(InputModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSetTemplateSpec) DeepCopyInto(out *JobSetTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSetTemplateSpec.
func (in *JobSetTemplateSpec) DeepCopy() *JobSetTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(JobSetTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MLPolicy) DeepCopyInto(out *MLPolicy) {
	*out = *in
	if in.NumNodes != nil {
		in, out := &in.NumNodes, &out.NumNodes
		*out = new(int32)
		**out = **in
	}
	in.MLPolicySource.DeepCopyInto(&out.MLPolicySource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MLPolicy.
func (in *MLPolicy) DeepCopy() *MLPolicy {
	if in == nil {
		return nil
	}
	out := new(MLPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MLPolicySource) DeepCopyInto(out *MLPolicySource) {
	*out = *in
	if in.Torch != nil {
		in, out := &in.Torch, &out.Torch
		*out = new(TorchMLPolicySource)
		(*in).DeepCopyInto(*out)
	}
	if in.MPI != nil {
		in, out := &in.MPI, &out.MPI
		*out = new(MPIMLPolicySource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MLPolicySource.
func (in *MLPolicySource) DeepCopy() *MLPolicySource {
	if in == nil {
		return nil
	}
	out := new(MLPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIMLPolicySource) DeepCopyInto(out *MPIMLPolicySource) {
	*out = *in
	if in.NumProcPerNode != nil {
		in, out := &in.NumProcPerNode, &out.NumProcPerNode
		*out = new(int32)
		**out = **in
	}
	if in.MPIImplementation != nil {
		in, out := &in.MPIImplementation, &out.MPIImplementation
		*out = new(MPIImplementation)
		**out = **in
	}
	if in.SSHAuthMountPath != nil {
		in, out := &in.SSHAuthMountPath, &out.SSHAuthMountPath
		*out = new(string)
		**out = **in
	}
	if in.RunLauncherAsNode != nil {
		in, out := &in.RunLauncherAsNode, &out.RunLauncherAsNode
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIMLPolicySource.
func (in *MPIMLPolicySource) DeepCopy() *MPIMLPolicySource {
	if in == nil {
		return nil
	}
	out := new(MPIMLPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ModelConfig) DeepCopyInto(out *ModelConfig) {
	*out = *in
	if in.Input != nil {
		in, out := &in.Input, &out.Input
		*out = new(InputModel)
		(*in).DeepCopyInto(*out)
	}
	if in.Output != nil {
		in, out := &in.Output, &out.Output
		*out = new(OutputModel)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ModelConfig.
func (in *ModelConfig) DeepCopy() *ModelConfig {
	if in == nil {
		return nil
	}
	out := new(ModelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutputModel) DeepCopyInto(out *OutputModel) {
	*out = *in
	if in.StorageUri != nil {
		in, out := &in.StorageUri, &out.StorageUri
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(v1.SecretReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutputModel.
func (in *OutputModel) DeepCopy() *OutputModel {
	if in == nil {
		return nil
	}
	out := new(OutputModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupPolicy) DeepCopyInto(out *PodGroupPolicy) {
	*out = *in
	in.PodGroupPolicySource.DeepCopyInto(&out.PodGroupPolicySource)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupPolicy.
func (in *PodGroupPolicy) DeepCopy() *PodGroupPolicy {
	if in == nil {
		return nil
	}
	out := new(PodGroupPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodGroupPolicySource) DeepCopyInto(out *PodGroupPolicySource) {
	*out = *in
	if in.Coscheduling != nil {
		in, out := &in.Coscheduling, &out.Coscheduling
		*out = new(CoschedulingPodGroupPolicySource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodGroupPolicySource.
func (in *PodGroupPolicySource) DeepCopy() *PodGroupPolicySource {
	if in == nil {
		return nil
	}
	out := new(PodGroupPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpecOverride) DeepCopyInto(out *PodSpecOverride) {
	*out = *in
	if in.TargetReplicatedJobs != nil {
		in, out := &in.TargetReplicatedJobs, &out.TargetReplicatedJobs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]ContainerOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpecOverride.
func (in *PodSpecOverride) DeepCopy() *PodSpecOverride {
	if in == nil {
		return nil
	}
	out := new(PodSpecOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeRef) DeepCopyInto(out *RuntimeRef) {
	*out = *in
	if in.APIGroup != nil {
		in, out := &in.APIGroup, &out.APIGroup
		*out = new(string)
		**out = **in
	}
	if in.Kind != nil {
		in, out := &in.Kind, &out.Kind
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeRef.
func (in *RuntimeRef) DeepCopy() *RuntimeRef {
	if in == nil {
		return nil
	}
	out := new(RuntimeRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TorchElasticPolicy) DeepCopyInto(out *TorchElasticPolicy) {
	*out = *in
	if in.MaxRestarts != nil {
		in, out := &in.MaxRestarts, &out.MaxRestarts
		*out = new(int32)
		**out = **in
	}
	if in.MinNodes != nil {
		in, out := &in.MinNodes, &out.MinNodes
		*out = new(int32)
		**out = **in
	}
	if in.MaxNodes != nil {
		in, out := &in.MaxNodes, &out.MaxNodes
		*out = new(int32)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v2.MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TorchElasticPolicy.
func (in *TorchElasticPolicy) DeepCopy() *TorchElasticPolicy {
	if in == nil {
		return nil
	}
	out := new(TorchElasticPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TorchMLPolicySource) DeepCopyInto(out *TorchMLPolicySource) {
	*out = *in
	if in.NumProcPerNode != nil {
		in, out := &in.NumProcPerNode, &out.NumProcPerNode
		*out = new(string)
		**out = **in
	}
	if in.ElasticPolicy != nil {
		in, out := &in.ElasticPolicy, &out.ElasticPolicy
		*out = new(TorchElasticPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TorchMLPolicySource.
func (in *TorchMLPolicySource) DeepCopy() *TorchMLPolicySource {
	if in == nil {
		return nil
	}
	out := new(TorchMLPolicySource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainJob) DeepCopyInto(out *TrainJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainJob.
func (in *TrainJob) DeepCopy() *TrainJob {
	if in == nil {
		return nil
	}
	out := new(TrainJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainJobList) DeepCopyInto(out *TrainJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrainJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainJobList.
func (in *TrainJobList) DeepCopy() *TrainJobList {
	if in == nil {
		return nil
	}
	out := new(TrainJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainJobSpec) DeepCopyInto(out *TrainJobSpec) {
	*out = *in
	in.RuntimeRef.DeepCopyInto(&out.RuntimeRef)
	if in.Trainer != nil {
		in, out := &in.Trainer, &out.Trainer
		*out = new(Trainer)
		(*in).DeepCopyInto(*out)
	}
	if in.DatasetConfig != nil {
		in, out := &in.DatasetConfig, &out.DatasetConfig
		*out = new(DatasetConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ModelConfig != nil {
		in, out := &in.ModelConfig, &out.ModelConfig
		*out = new(ModelConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodSpecOverrides != nil {
		in, out := &in.PodSpecOverrides, &out.PodSpecOverrides
		*out = make([]PodSpecOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	if in.ManagedBy != nil {
		in, out := &in.ManagedBy, &out.ManagedBy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainJobSpec.
func (in *TrainJobSpec) DeepCopy() *TrainJobSpec {
	if in == nil {
		return nil
	}
	out := new(TrainJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainJobStatus) DeepCopyInto(out *TrainJobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReplicatedJobsStatus != nil {
		in, out := &in.ReplicatedJobsStatus, &out.ReplicatedJobsStatus
		*out = make([]v1alpha2.ReplicatedJobStatus, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainJobStatus.
func (in *TrainJobStatus) DeepCopy() *TrainJobStatus {
	if in == nil {
		return nil
	}
	out := new(TrainJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trainer) DeepCopyInto(out *Trainer) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NumNodes != nil {
		in, out := &in.NumNodes, &out.NumNodes
		*out = new(int32)
		**out = **in
	}
	if in.ResourcesPerNode != nil {
		in, out := &in.ResourcesPerNode, &out.ResourcesPerNode
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.NumProcPerNode != nil {
		in, out := &in.NumProcPerNode, &out.NumProcPerNode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trainer.
func (in *Trainer) DeepCopy() *Trainer {
	if in == nil {
		return nil
	}
	out := new(Trainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingRuntime) DeepCopyInto(out *TrainingRuntime) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingRuntime.
func (in *TrainingRuntime) DeepCopy() *TrainingRuntime {
	if in == nil {
		return nil
	}
	out := new(TrainingRuntime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainingRuntime) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingRuntimeList) DeepCopyInto(out *TrainingRuntimeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TrainingRuntime, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingRuntimeList.
func (in *TrainingRuntimeList) DeepCopy() *TrainingRuntimeList {
	if in == nil {
		return nil
	}
	out := new(TrainingRuntimeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TrainingRuntimeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrainingRuntimeSpec) DeepCopyInto(out *TrainingRuntimeSpec) {
	*out = *in
	if in.MLPolicy != nil {
		in, out := &in.MLPolicy, &out.MLPolicy
		*out = new(MLPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.PodGroupPolicy != nil {
		in, out := &in.PodGroupPolicy, &out.PodGroupPolicy
		*out = new(PodGroupPolicy)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrainingRuntimeSpec.
func (in *TrainingRuntimeSpec) DeepCopy() *TrainingRuntimeSpec {
	if in == nil {
		return nil
	}
	out := new(TrainingRuntimeSpec)
	in.DeepCopyInto(out)
	return out
}