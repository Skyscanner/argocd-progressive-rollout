package utils

import (
	argov1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/gitops-engine/pkg/health"
	corev1 "k8s.io/api/core/v1"
	"sort"
)

// IsArgoCDCluster returns true if one of the labels is the ArgoCD secret label with the secret type cluster as value
func IsArgoCDCluster(labels map[string]string) bool {
	val, ok := labels[ArgoCDSecretTypeLabel]
	return val == ArgoCDSecretTypeCluster && ok
}

// SortSecretsByName sort the SecretList in place by the secrets name
func SortSecretsByName(secrets *corev1.SecretList) {
	sort.SliceStable(secrets.Items, func(i, j int) bool { return secrets.Items[i].Name < secrets.Items[j].Name })
}

// SortSecretsByName sort the Application slice in place by the app name
func SortAppsByName(apps *[]argov1alpha1.Application) {
	sort.SliceStable(*apps, func(i, j int) bool { return (*apps)[i].Name < (*apps)[j].Name })
}

// GetAppsBySyncStatusCode returns the Applications matching the specified sync status code
func GetAppsBySyncStatusCode(apps []argov1alpha1.Application, code argov1alpha1.SyncStatusCode) []argov1alpha1.Application {
	var result []argov1alpha1.Application

	for _, app := range apps {
		if app.Status.Sync.Status == code {
			result = append(result, app)
		}
	}

	return result
}

// GetAppsByHealthStatusCode returns the Applications matching the specified sync status code
func GetAppsByHealthStatusCode(apps []argov1alpha1.Application, code health.HealthStatusCode) []argov1alpha1.Application {
	var result []argov1alpha1.Application

	for _, app := range apps {
		if app.Status.Health.Status == code {
			result = append(result, app)
		}
	}

	return result
}