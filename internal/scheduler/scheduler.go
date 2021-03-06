package scheduler

import (
	syncv1alpha1 "github.com/Skyscanner/applicationset-progressive-sync/api/v1alpha1"
	"github.com/Skyscanner/applicationset-progressive-sync/internal/utils"
	argov1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/gitops-engine/pkg/health"
	"github.com/go-logr/logr"
)

// Scheduler returns a list of apps to sync for a given stage
func Scheduler(log logr.Logger, apps []argov1alpha1.Application, stage syncv1alpha1.ProgressiveSyncStage, pss utils.ProgressiveSyncState) []argov1alpha1.Application {

	/*
		The Scheduler takes:
			- a ProgressiveSyncStage object
			- the Applications selected with the clusters selector

		The Scheduler splits the Applications in the following groups:
			- OutOfSync Applications: those are the Applications to update during the stage.
			- syncedInCurrentStage Applications: those are Application that synced during the current stage. Those Applications count against the number of clusters to update.
			- progressingApps: those are Applications that are still in progress updating. Those Applications count against the number of clusters to update in parallel.

		Why does the Scheduler need an annotation?
		Consider the scenario where we have 5 Applications - 4 OutOfSync and 1 Synced - and a stage with maxTargets = 3.
		If we don't keep track on which stage the Application synced, we can't compute how many applications we have to update in the current stage.
		Without the annotation, it would be impossible for the scheduler to know if the Application synced at this stage - and so we have only 2 Applications left to sync.
	*/

	log = log.WithName("scheduler")

	var scheduledApps []argov1alpha1.Application
	outOfSyncApps := utils.GetAppsBySyncStatusCode(apps, argov1alpha1.SyncStatusCodeOutOfSync)
	// If there are no OutOfSync Applications, return
	if len(outOfSyncApps) == 0 {
		return scheduledApps
	}

	log.Info("fetched out-of-sync apps", "apps", utils.GetAppsName(outOfSyncApps))

	healthyApps := utils.GetAppsByHealthStatusCode(apps, health.HealthStatusHealthy)
	syncedInCurrentStage := utils.GetSyncedAppsByStage(healthyApps, stage, pss)
	log.Info("fetched synced-in-current-stage apps", "apps", utils.GetAppsName(syncedInCurrentStage))

	progressingApps := utils.GetAppsByHealthStatusCode(apps, health.HealthStatusProgressing)
	log.Info("fetched progressing apps", "apps", utils.GetAppsName(progressingApps))

	maxTargets := pss.GetMaxTargets(stage)
	maxParallel := pss.GetMaxParallel(stage)

	// If we already synced the desired number of Applications, return
	if maxTargets == len(syncedInCurrentStage) {
		return scheduledApps
	}

	// Because of eventual consistency, there might be a time where
	// maxParallel-len(progressingApps) might actually be greater than len(outOfSyncApps)
	// causing the runtime to panic
	p := maxParallel - len(progressingApps)
	if p > len(outOfSyncApps) {
		p = len(outOfSyncApps)
	}

	// Consider the following scenario
	//
	// maxTargets = 3
	// maxParallel = 3
	// outOfSyncApps = 4
	// syncedInCurrentStage = 2
	// progressingApps = 1
	// p = 2
	//
	// Without the following logic we have p=2, so we would end up with a total of 4 applications synced in the stage
	if p+len(syncedInCurrentStage) > maxTargets {
		p = maxTargets - len(syncedInCurrentStage)
	}

	for i := 0; i < p; i++ {
		scheduledApps = append(scheduledApps, outOfSyncApps[i])
	}

	// To recover from a case where something triggers an Application sync, the scheulder also return
	// all the progressing apps but still out of sync, so we can mark them as synced and take back control of the app

	progressingOutOfSyncApps := utils.GetAppsBySyncStatusCode(progressingApps, argov1alpha1.SyncStatusCodeOutOfSync)
	scheduledApps = append(scheduledApps, progressingOutOfSyncApps...)

	return scheduledApps
}

// IsStageFailed returns true if at least one app is failed in the given stage
func IsStageFailed(apps []argov1alpha1.Application, stage syncv1alpha1.ProgressiveSyncStage, pss utils.ProgressiveSyncState) bool {
	// A stage is failed if any of its applications has:
	// - its Health Status Code == Degraded
	// - its Sync Status Code == Synced

	if apps == nil {
		return false
	}

	degradedApps := utils.GetAppsByHealthStatusCode(apps, health.HealthStatusDegraded)
	degradedSyncedApps := utils.GetSyncedAppsByStage(degradedApps, stage, pss)
	return len(degradedSyncedApps) > 0
}

// IsStageInProgress returns true if at least one app is is in progress
func IsStageInProgress(apps []argov1alpha1.Application, stage syncv1alpha1.ProgressiveSyncStage, pss utils.ProgressiveSyncState) bool {
	// An stage is in progress if:
	// - there is at least one app with Health Status Code == Progressing
	// - the number of apps synced so far is less than the apps to sync

	if apps == nil {
		return false
	}

	progressingApps := utils.GetAppsByHealthStatusCode(apps, health.HealthStatusProgressing)
	progressingSyncedApps := utils.GetSyncedAppsByStage(progressingApps, stage, pss)

	appsSyncedSoFar := utils.GetSyncedAppsByStage(apps, stage, pss)

	return len(progressingSyncedApps) > 0 || len(appsSyncedSoFar) < pss.GetMaxTargets(stage)
}

// IsStageComplete returns true if all applications are Synced and Healthy
func IsStageComplete(apps []argov1alpha1.Application, stage syncv1alpha1.ProgressiveSyncStage, pss utils.ProgressiveSyncState) bool {
	// An app is complete if:
	// - its Health Status Code is Healthy
	// - its Sync Status Code is Synced

	if apps == nil {
		return true
	}

	healthyApps := utils.GetAppsByHealthStatusCode(apps, health.HealthStatusHealthy)
	healthySyncedApps := utils.GetSyncedAppsByStage(healthyApps, stage, pss)

	appsToCompleteStage := pss.GetMaxTargets(stage)

	return len(healthySyncedApps) == appsToCompleteStage || appsToCompleteStage == 0
}
