package scheduler

import (
	deploymentskyscannernetv1alpha1 "github.com/Skyscanner/argocd-progressive-rollout/api/v1alpha1"
	"github.com/Skyscanner/argocd-progressive-rollout/internal/utils"
	argov1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/gitops-engine/pkg/health"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"testing"
)

const (
	SchedulerTestNamespace = "test-scheduler"
	StageName = "stage"
)

func TestScheduler(t *testing.T) {
	testCases := []struct {
		name     string
		apps     []argov1alpha1.Application
		stage    deploymentskyscannernetv1alpha1.ProgressiveRolloutStage
		expected []argov1alpha1.Application
	}{
		{
			name: "Applications: outOfSync 3, syncedInCurrentStage 0, progressing 0, | Stage: maxTargets 2, maxParallel 2 | Expected: scheduled 2",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("2"),
				MaxTargets:  intstr.Parse("3"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},
		{
			name: "Applications: outOfSync 3, syncedInCurrentStage 1, progressing 1, | Stage: maxTargets 5, maxParallel 2 | Expected: scheduled 1",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "app-four",
						Namespace:   SchedulerTestNamespace,
						Annotations: map[string]string{utils.ProgressiveRolloutSyncedAtStageKey: StageName},
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeSynced,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "app-five",
						Namespace:   SchedulerTestNamespace,
						Annotations: map[string]string{utils.ProgressiveRolloutSyncedAtStageKey: StageName},
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeSynced,
						},
						Health: argov1alpha1.HealthStatus{Status: health.HealthStatusProgressing},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("2"),
				MaxTargets:  intstr.Parse("3"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},
		{
			name: "Applications: outOfSync 5, syncedInCurrentStage 0, progressing 0, | Stage: maxTargets 3, maxParallel 2 | Expected: scheduled 2",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("2"),
				MaxTargets:  intstr.Parse("3"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},
		{
			name: "Applications: outOfSync 5, syncedInCurrentStage 0, progressing 0, | Stage: maxTargets 50%, maxParallel 100% | Expected: scheduled 2",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("100%"),
				MaxTargets:  intstr.Parse("50%"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},
		{
			name: "Applications: outOfSync 2, syncedInCurrentStage 3, progressing 0, | Stage: maxTargets 3, maxParallel 1 | Expected: scheduled 0",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
						Annotations: map[string]string{
							utils.ProgressiveRolloutSyncedAtStageKey: StageName,
						},
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeSynced,
						},
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusHealthy},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
						Annotations: map[string]string{
							utils.ProgressiveRolloutSyncedAtStageKey: StageName,
						},
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeSynced,
						},
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusHealthy},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
						Annotations: map[string]string{
							utils.ProgressiveRolloutSyncedAtStageKey: StageName,
						},
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeSynced,
						},
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusHealthy},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("1"),
				MaxTargets:  intstr.Parse("3"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: nil,
		},
		{
			name: "Applications: outOfSync 1, syncedInCurrentStage 0, progressing 0, | Stage: maxTargets 1, maxParallel 1 | Expected: scheduled 1",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("1"),
				MaxTargets:  intstr.Parse("1"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},
		{
			name: "Applications: outOfSync 5, syncedInCurrentStage 0, progressing 0, | Stage: maxTargets 3, maxParallel 3 | Expected: scheduled 3",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("3"),
				MaxTargets:  intstr.Parse("3"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},
		{
			name: "Applications: outOfSync 2, syncedInCurrentStage 0, progressing 0, | Stage: maxTargets 10%, maxParallel 10% | Expected: scheduled 1",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("10%"),
				MaxTargets:  intstr.Parse("10%"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},

		{
			name: "Applications: outOfSync 2, syncedInCurrentStage 0, progressing 0, syncedInPreviousStage 1 | Stage: maxTargets 2, maxParallel 2 | Expected: scheduled 2",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:        "app-three",
						Namespace:   SchedulerTestNamespace,
						Annotations: map[string]string{utils.ProgressiveRolloutSyncedAtStageKey: "previous-stage"},
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeSynced,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("2"),
				MaxTargets:  intstr.Parse("2"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Sync: argov1alpha1.SyncStatus{
							Status: argov1alpha1.SyncStatusCodeOutOfSync,
						},
					},
				},
			},
		},

		{
			name: "Applications: outOfSync 0, syncedInCurrentStage 0, progressing 0, | Stage: maxTargets 3, maxParallel 3 | Expected: scheduled 0",
			apps: nil,
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name:        StageName,
				MaxParallel: intstr.Parse("3"),
				MaxTargets:  intstr.Parse("3"),
				Targets:     deploymentskyscannernetv1alpha1.ProgressiveRolloutTargets{},
			},
			expected: nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			utils.SortAppsByName(&testCase.apps)
			got := Scheduler(testCase.apps, testCase.stage)
			g := NewWithT(t)
			g.Expect(got).To(Equal(testCase.expected))
		})
	}
}

func TestIsStageFailed(t *testing.T) {
	testCases := []struct{
		name string
		apps []argov1alpha1.Application
		stage deploymentskyscannernetv1alpha1.ProgressiveRolloutStage
		expected bool
	}{
		{
			name: "stage failed",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusUnknown,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusHealthy,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusProgressing,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusMissing,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusSuspended,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-six",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusDegraded,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name: StageName,
			},
			expected: true,
		},
		{
			name: "stage not failed",
			apps: []argov1alpha1.Application{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-one",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusUnknown,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-two",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusHealthy,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-three",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusProgressing,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-four",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusMissing,
						},
					},
				},
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "app-five",
						Namespace: SchedulerTestNamespace,
					},
					Status: argov1alpha1.ApplicationStatus{
						Health: argov1alpha1.HealthStatus{
							Status: health.HealthStatusSuspended,
						},
					},
				},
			},
			stage: deploymentskyscannernetv1alpha1.ProgressiveRolloutStage{
				Name: StageName,
			},
			expected: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := IsStageFailed(testCase.apps, testCase.stage)
			g := NewWithT(t)
			g.Expect(got).To(Equal(testCase.expected))
		})
	}

}
