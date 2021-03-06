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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const ProgressiveSyncFinalizer = "finalizers.argoproj.skyscanner.net"

// ProgressiveSyncSpec defines the desired state of ProgressiveSync
type ProgressiveSyncSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// SourceRef defines the resource, example an ApplicationSet, which owns ArgoCD Applications
	//+kubebuilder:validation:Required
	SourceRef corev1.TypedLocalObjectReference `json:"sourceRef"`
	// Stages defines a list of Progressive Rollout stages
	//+kubebuilder:validation:Optional
	Stages []ProgressiveSyncStage `json:"stages,omitempty"`
}

// ProgressiveSyncStage defines a rollout stage
type ProgressiveSyncStage struct {
	// Name is a human friendly name for the stage
	//+kubebuilder:validation:Required
	Name string `json:"name"`
	// MaxParallel is how many selected targets to update in parallel
	//+kubebuilder:validation:Minimum:1
	MaxParallel intstr.IntOrString `json:"maxParallel"`
	// MaxTargets is the maximum number of selected targets to update
	//+kubebuilder:validation:Minimum:1
	MaxTargets intstr.IntOrString `json:"maxTargets"`
	// Targets is the targets to update in the stage
	//+kubebuilder:validation:Optional
	Targets ProgressiveSyncTargets `json:"targets,omitempty"`
}

// ProgressiveSyncTargets defines the target of the Progressive Rollout
type ProgressiveSyncTargets struct {
	// Clusters is the a cluster type of targets
	//+kubebuilder:validation:Optional
	Clusters Clusters `json:"clusters"`
}

// Clusters defines a target of type clusters
type Clusters struct {
	// Selector is a label selector to get the clusters for the update
	//+kubebuilder:validation:Required
	Selector metav1.LabelSelector `json:"selector"`
}

// ProgressiveSyncStatus defines the observed state of ProgressiveSync
type ProgressiveSyncStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty"`
	Stages     []StageStatus      `json:"stages,omitempty"`
}

// GetStatusConditions returns a pointer to the Status.Conditions slice
func (in *ProgressiveSync) GetStatusConditions() *[]metav1.Condition {
	return &in.Status.Conditions
}

// NewStatusCondition adds a new Condition
func (in *ProgressiveSync) NewStatusCondition(t string, s metav1.ConditionStatus, r string, m string) metav1.Condition {
	return metav1.Condition{
		Type:               t,
		Status:             s,
		LastTransitionTime: metav1.Now(),
		Reason:             r,
		Message:            m,
	}
}

// Owns returns true if the ProgressiveSync object has a reference to one of the owners
func (in *ProgressiveSync) Owns(owners []metav1.OwnerReference) bool {
	for _, owner := range owners {
		if owner.Kind == in.Spec.SourceRef.Kind && owner.APIVersion == *in.Spec.SourceRef.APIGroup && owner.Name == in.Spec.SourceRef.Name {
			return true
		}
	}
	return false
}

// SetStageStatus sets the corresponding StageStatus in stageStatus to newStatus
// - If a stage doesn't exist, it will be added to StageStatus slice
// - If a stage already exists it will be updated
func (in *ProgressiveSync) SetStageStatus(newStatus StageStatus, updateTime *metav1.Time) {
	// If StartedAt is not set and the stage is in progress, assign StartedAt
	if newStatus.Phase == PhaseProgressing && newStatus.StartedAt.IsZero() {
		newStatus.StartedAt = updateTime
	}
	// If the stage is not progressing it is either completed or failed.
	// If FinishedAt is not set we assign it.
	if (newStatus.Phase == PhaseFailed || newStatus.Phase == PhaseSucceeded) && newStatus.FinishedAt.IsZero() {
		newStatus.FinishedAt = updateTime
	}

	// Get the status if it already exists
	existingStatus := FindStageStatus(in.Status.Stages, newStatus.Name)

	if existingStatus == nil {
		in.Status.Stages = append(in.Status.Stages, newStatus)
		return
	}

	existingStatus.Phase = newStatus.Phase
	existingStatus.Message = newStatus.Message
	existingStatus.StartedAt = newStatus.StartedAt
	existingStatus.FinishedAt = newStatus.FinishedAt
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// ProgressiveSync is the Schema for the progressivesyncs API
type ProgressiveSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProgressiveSyncSpec   `json:"spec,omitempty"`
	Status ProgressiveSyncStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProgressiveSyncList contains a list of ProgressiveSync
type ProgressiveSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProgressiveSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProgressiveSync{}, &ProgressiveSyncList{})
}
