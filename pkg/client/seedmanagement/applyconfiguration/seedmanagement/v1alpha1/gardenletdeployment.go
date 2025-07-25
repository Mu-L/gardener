// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	"maps"

	v1 "k8s.io/api/core/v1"
)

// GardenletDeploymentApplyConfiguration represents an declarative configuration of the GardenletDeployment type for use
// with apply.
type GardenletDeploymentApplyConfiguration struct {
	ReplicaCount           *int32                   `json:"replicaCount,omitempty"`
	RevisionHistoryLimit   *int32                   `json:"revisionHistoryLimit,omitempty"`
	ServiceAccountName     *string                  `json:"serviceAccountName,omitempty"`
	Image                  *ImageApplyConfiguration `json:"image,omitempty"`
	Resources              *v1.ResourceRequirements `json:"resources,omitempty"`
	PodLabels              map[string]string        `json:"podLabels,omitempty"`
	PodAnnotations         map[string]string        `json:"podAnnotations,omitempty"`
	AdditionalVolumes      []v1.Volume              `json:"additionalVolumes,omitempty"`
	AdditionalVolumeMounts []v1.VolumeMount         `json:"additionalVolumeMounts,omitempty"`
	Env                    []v1.EnvVar              `json:"env,omitempty"`
	VPA                    *bool                    `json:"vpa,omitempty"`
}

// GardenletDeploymentApplyConfiguration constructs an declarative configuration of the GardenletDeployment type for use with
// apply.
func GardenletDeployment() *GardenletDeploymentApplyConfiguration {
	return &GardenletDeploymentApplyConfiguration{}
}

// WithReplicaCount sets the ReplicaCount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReplicaCount field is set to the value of the last call.
func (b *GardenletDeploymentApplyConfiguration) WithReplicaCount(value int32) *GardenletDeploymentApplyConfiguration {
	b.ReplicaCount = &value
	return b
}

// WithRevisionHistoryLimit sets the RevisionHistoryLimit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RevisionHistoryLimit field is set to the value of the last call.
func (b *GardenletDeploymentApplyConfiguration) WithRevisionHistoryLimit(value int32) *GardenletDeploymentApplyConfiguration {
	b.RevisionHistoryLimit = &value
	return b
}

// WithServiceAccountName sets the ServiceAccountName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccountName field is set to the value of the last call.
func (b *GardenletDeploymentApplyConfiguration) WithServiceAccountName(value string) *GardenletDeploymentApplyConfiguration {
	b.ServiceAccountName = &value
	return b
}

// WithImage sets the Image field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Image field is set to the value of the last call.
func (b *GardenletDeploymentApplyConfiguration) WithImage(value *ImageApplyConfiguration) *GardenletDeploymentApplyConfiguration {
	b.Image = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *GardenletDeploymentApplyConfiguration) WithResources(value v1.ResourceRequirements) *GardenletDeploymentApplyConfiguration {
	b.Resources = &value
	return b
}

// WithPodLabels puts the entries into the PodLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the PodLabels field,
// overwriting an existing map entries in PodLabels field with the same key.
func (b *GardenletDeploymentApplyConfiguration) WithPodLabels(entries map[string]string) *GardenletDeploymentApplyConfiguration {
	if b.PodLabels == nil && len(entries) > 0 {
		b.PodLabels = make(map[string]string, len(entries))
	}
	maps.Copy(b.PodLabels, entries)
	return b
}

// WithPodAnnotations puts the entries into the PodAnnotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the PodAnnotations field,
// overwriting an existing map entries in PodAnnotations field with the same key.
func (b *GardenletDeploymentApplyConfiguration) WithPodAnnotations(entries map[string]string) *GardenletDeploymentApplyConfiguration {
	if b.PodAnnotations == nil && len(entries) > 0 {
		b.PodAnnotations = make(map[string]string, len(entries))
	}
	maps.Copy(b.PodAnnotations, entries)
	return b
}

// WithAdditionalVolumes adds the given value to the AdditionalVolumes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalVolumes field.
func (b *GardenletDeploymentApplyConfiguration) WithAdditionalVolumes(values ...v1.Volume) *GardenletDeploymentApplyConfiguration {
	for i := range values {
		b.AdditionalVolumes = append(b.AdditionalVolumes, values[i])
	}
	return b
}

// WithAdditionalVolumeMounts adds the given value to the AdditionalVolumeMounts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the AdditionalVolumeMounts field.
func (b *GardenletDeploymentApplyConfiguration) WithAdditionalVolumeMounts(values ...v1.VolumeMount) *GardenletDeploymentApplyConfiguration {
	for i := range values {
		b.AdditionalVolumeMounts = append(b.AdditionalVolumeMounts, values[i])
	}
	return b
}

// WithEnv adds the given value to the Env field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Env field.
func (b *GardenletDeploymentApplyConfiguration) WithEnv(values ...v1.EnvVar) *GardenletDeploymentApplyConfiguration {
	for i := range values {
		b.Env = append(b.Env, values[i])
	}
	return b
}

// WithVPA sets the VPA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VPA field is set to the value of the last call.
func (b *GardenletDeploymentApplyConfiguration) WithVPA(value bool) *GardenletDeploymentApplyConfiguration {
	b.VPA = &value
	return b
}
