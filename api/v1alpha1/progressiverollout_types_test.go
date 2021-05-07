package v1alpha1

import (
	"testing"
	"time"

	"github.com/Skyscanner/argocd-progressive-rollout/internal/utils"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestOwns(t *testing.T) {
	testCases := []struct {
		name            string
		ownerReferences []metav1.OwnerReference
		expected        bool
	}{{
		name: "owns",
		ownerReferences: []metav1.OwnerReference{{
			APIVersion: "fakeAPIVersion",
			Kind:       "fakeKind",
			Name:       "fakeName",
		}, {
			APIVersion: utils.AppSetAPIGroup,
			Kind:       utils.AppSetKind,
			Name:       "owner-app-set",
		}},
		expected: true,
	}, {
		name: "does not own",
		ownerReferences: []metav1.OwnerReference{{
			APIVersion: "fakeAPIVersion",
			Kind:       "fakeKind",
			Name:       "fakeName",
		}},
		expected: false,
	}}

	ref := utils.AppSetAPIGroup
	pr := ProgressiveRollout{
		ObjectMeta: metav1.ObjectMeta{Name: "pr", Namespace: "namespace"},
		Spec: ProgressiveRolloutSpec{
			SourceRef: corev1.TypedLocalObjectReference{
				APIGroup: &ref,
				Kind:     utils.AppSetKind,
				Name:     "owner-app-set",
			},
			Stages: nil,
		}}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := pr.Owns(testCase.ownerReferences)
			g := NewWithT(t)
			g.Expect(got).To(Equal(testCase.expected))
		})
	}
}

func TestSetStageStatus(t *testing.T) {
	pastTime := metav1.NewTime(time.Now().Add(-1 * time.Hour))
	mockTime := metav1.NewTime(time.Now())

	testCases := []struct {
		name        string
		stageStatus []StageStatus
		newStatus   StageStatus
		expected    []StageStatus
	}{
		{
			name:        "add stage to empty list",
			stageStatus: []StageStatus{},
			newStatus: StageStatus{
				Name:  "stage 1",
				Phase: PhaseProgressing,
			},
			expected: []StageStatus{{
				Name:       "stage 1",
				Phase:      PhaseProgressing,
				Message:    "",
				StartedAt:  &mockTime,
				FinishedAt: nil,
			}},
		}, {
			name: "add stage to list",
			stageStatus: []StageStatus{{
				Name:      "stage 1",
				Phase:     PhaseProgressing,
				StartedAt: &pastTime,
			}},
			newStatus: StageStatus{
				Name:  "stage 2",
				Phase: PhaseProgressing,
			},
			expected: []StageStatus{{
				Name:       "stage 1",
				Phase:      PhaseProgressing,
				Message:    "",
				StartedAt:  &pastTime,
				FinishedAt: nil,
			}, {
				Name:       "stage 2",
				Phase:      PhaseProgressing,
				Message:    "",
				StartedAt:  &mockTime,
				FinishedAt: nil,
			}},
		}, {
			name: "update stage in list",
			stageStatus: []StageStatus{{
				Name:      "stage 1",
				Phase:     PhaseProgressing,
				StartedAt: &pastTime,
			}, {
				Name:      "stage 2",
				Phase:     PhaseProgressing,
				Message:   "old message",
				StartedAt: &pastTime,
			}},
			newStatus: StageStatus{
				Name:      "stage 2",
				Phase:     PhaseSucceeded,
				Message:   "new message",
				StartedAt: &pastTime,
			},
			expected: []StageStatus{{
				Name:       "stage 1",
				Phase:      PhaseProgressing,
				Message:    "",
				StartedAt:  &pastTime,
				FinishedAt: nil,
			}, {
				Name:       "stage 2",
				Phase:      PhaseSucceeded,
				Message:    "new message",
				StartedAt:  &pastTime,
				FinishedAt: &mockTime,
			}},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			ref := utils.AppSetAPIGroup
			pr := ProgressiveRollout{
				ObjectMeta: metav1.ObjectMeta{Name: "pr", Namespace: "namespace"},
				Spec: ProgressiveRolloutSpec{
					SourceRef: corev1.TypedLocalObjectReference{
						APIGroup: &ref,
						Kind:     utils.AppSetKind,
						Name:     "owner-app-set",
					},
					Stages: nil,
				},
				Status: ProgressiveRolloutStatus{
					Stages: testCase.stageStatus,
				},
			}

			pr.SetStageStatus(testCase.newStatus, mockTime)
			g := NewWithT(t)
			g.Expect(pr.Status.Stages).To(Equal(testCase.expected))
		})
	}
}
