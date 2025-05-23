// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v1alpha1 "github.com/gardener/gardener/pkg/apis/seedmanagement/v1alpha1"
	seedmanagementv1alpha1 "github.com/gardener/gardener/pkg/client/seedmanagement/applyconfiguration/seedmanagement/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) any {
	switch kind {
	// Group=seedmanagement.gardener.cloud, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithKind("GardenletConfig"):
		return &seedmanagementv1alpha1.GardenletApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("GardenletDeployment"):
		return &seedmanagementv1alpha1.GardenletDeploymentApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Image"):
		return &seedmanagementv1alpha1.ImageApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedSeed"):
		return &seedmanagementv1alpha1.ManagedSeedApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedSeedSet"):
		return &seedmanagementv1alpha1.ManagedSeedSetApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedSeedSetSpec"):
		return &seedmanagementv1alpha1.ManagedSeedSetSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedSeedSetStatus"):
		return &seedmanagementv1alpha1.ManagedSeedSetStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedSeedSpec"):
		return &seedmanagementv1alpha1.ManagedSeedSpecApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedSeedStatus"):
		return &seedmanagementv1alpha1.ManagedSeedStatusApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("ManagedSeedTemplate"):
		return &seedmanagementv1alpha1.ManagedSeedTemplateApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("PendingReplica"):
		return &seedmanagementv1alpha1.PendingReplicaApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("RollingUpdateStrategy"):
		return &seedmanagementv1alpha1.RollingUpdateStrategyApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("Shoot"):
		return &seedmanagementv1alpha1.ShootApplyConfiguration{}
	case v1alpha1.SchemeGroupVersion.WithKind("UpdateStrategy"):
		return &seedmanagementv1alpha1.UpdateStrategyApplyConfiguration{}

	}
	return nil
}
