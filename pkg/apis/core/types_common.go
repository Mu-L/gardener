// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package core

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ErrorCode is a string alias.
type ErrorCode string

const (
	// ErrorInfraUnauthenticated indicates that the last error occurred due to the client request not being completed because it lacks valid authentication credentials for the requested resource.
	// It is classified as a non-retryable error code.
	ErrorInfraUnauthenticated ErrorCode = "ERR_INFRA_UNAUTHENTICATED"
	// ErrorInfraUnauthorized indicates that the last error occurred due to the server understanding the request but refusing to authorize it.
	// It is classified as a non-retryable error code.
	ErrorInfraUnauthorized ErrorCode = "ERR_INFRA_UNAUTHORIZED"
	// ErrorInfraQuotaExceeded indicates that the last error occurred due to infrastructure quota limits.
	// It is classified as a non-retryable error code.
	ErrorInfraQuotaExceeded ErrorCode = "ERR_INFRA_QUOTA_EXCEEDED"
	// ErrorInfraRateLimitsExceeded indicates that the last error occurred due to exceeded infrastructure request rate limits.
	ErrorInfraRateLimitsExceeded ErrorCode = "ERR_INFRA_RATE_LIMITS_EXCEEDED"
	// ErrorInfraDependencies indicates that the last error occurred due to dependent objects on the infrastructure level.
	// It is classified as a non-retryable error code.
	ErrorInfraDependencies ErrorCode = "ERR_INFRA_DEPENDENCIES"
	// ErrorRetryableInfraDependencies indicates that the last error occurred due to dependent objects on the infrastructure level, but operation should be retried.
	ErrorRetryableInfraDependencies ErrorCode = "ERR_RETRYABLE_INFRA_DEPENDENCIES"
	// ErrorInfraResourcesDepleted indicates that the last error occurred due to depleted resource in the infrastructure.
	ErrorInfraResourcesDepleted ErrorCode = "ERR_INFRA_RESOURCES_DEPLETED"
	// ErrorCleanupClusterResources indicates that the last error occurred due to resources in the cluster that are stuck in deletion.
	ErrorCleanupClusterResources ErrorCode = "ERR_CLEANUP_CLUSTER_RESOURCES"
	// ErrorConfigurationProblem indicates that the last error occurred due to a configuration problem.
	// It is classified as a non-retryable error code.
	ErrorConfigurationProblem ErrorCode = "ERR_CONFIGURATION_PROBLEM"
	// ErrorRetryableConfigurationProblem indicates that the last error occurred due to a retryable configuration problem.
	ErrorRetryableConfigurationProblem ErrorCode = "ERR_RETRYABLE_CONFIGURATION_PROBLEM"
	// ErrorProblematicWebhook indicates that the last error occurred due to a webhook not following the Kubernetes
	// best practices (https://kubernetes.io/docs/reference/access-authn-authz/extensible-admission-controllers/#best-practices-and-warnings).
	// It is classified as a non-retryable error code.
	ErrorProblematicWebhook ErrorCode = "ERR_PROBLEMATIC_WEBHOOK"
)

// LastError indicates the last occurred error for an operation on a resource.
type LastError struct {
	// A human readable message indicating details about the last error.
	Description string
	// ID of the task which caused this last error
	TaskID *string
	// Well-defined error codes of the last error(s).
	// +optional
	Codes []ErrorCode
	// Last time the error was reported
	LastUpdateTime *metav1.Time
}

// LastOperationType is a string alias.
type LastOperationType string

const (
	// LastOperationTypeCreate indicates a 'create' operation.
	LastOperationTypeCreate LastOperationType = "Create"
	// LastOperationTypeReconcile indicates a 'reconcile' operation.
	LastOperationTypeReconcile LastOperationType = "Reconcile"
	// LastOperationTypeDelete indicates a 'delete' operation.
	LastOperationTypeDelete LastOperationType = "Delete"
	// LastOperationTypeRestore indicates a 'restore' operation.
	LastOperationTypeRestore LastOperationType = "Restore"
	// LastOperationTypeMigrate indicates a 'migrate' operation.
	LastOperationTypeMigrate LastOperationType = "Migrate"
)

// LastOperationState is a string alias.
type LastOperationState string

const (
	// LastOperationStateProcessing indicates that an operation is ongoing.
	LastOperationStateProcessing LastOperationState = "Processing"
	// LastOperationStateSucceeded indicates that an operation has completed successfully.
	LastOperationStateSucceeded LastOperationState = "Succeeded"
	// LastOperationStateError indicates that an operation is completed with errors and will be retried.
	LastOperationStateError LastOperationState = "Error"
	// LastOperationStateFailed indicates that an operation is completed with errors and won't be retried.
	LastOperationStateFailed LastOperationState = "Failed"
	// LastOperationStatePending indicates that an operation cannot be done now, but will be tried in future.
	LastOperationStatePending LastOperationState = "Pending"
	// LastOperationStateAborted indicates that an operation has been aborted.
	LastOperationStateAborted LastOperationState = "Aborted"
)

// LastOperation indicates the type and the state of the last operation, along with a description
// message and a progress indicator.
type LastOperation struct {
	// A human readable message indicating details about the last operation.
	Description string
	// Last time the operation state transitioned from one to another.
	LastUpdateTime metav1.Time
	// The progress in percentage (0-100) of the last operation.
	Progress int32
	// Status of the last operation, one of Aborted, Processing, Succeeded, Error, Failed.
	State LastOperationState
	// Type of the last operation, one of Create, Reconcile, Delete, Migrate, Restore.
	Type LastOperationType
}

// Gardener holds the information about the Gardener.
type Gardener struct {
	// ID is the container id of the Gardener which last acted on a Shoot cluster.
	ID string
	// Name is the hostname (pod name) of the Gardener which last acted on a Shoot cluster.
	Name string
	// Version is the version of the Gardener which last acted on a Shoot cluster.
	Version string
}

// HighAvailability specifies the configuration settings for high availability for a resource. Typical
// usages could be to configure HA for shoot control plane or for seed system components.
type HighAvailability struct {
	// FailureTolerance specifies the failure tolerance level in a high availability setup.
	FailureTolerance FailureTolerance
}

// FailureTolerance holds information about failure tolerance configuration.
type FailureTolerance struct {
	// Type specifies the type of failure that the highly available resource can tolerate.
	Type FailureToleranceType
}

// FailureToleranceType is the type of failure that a highly available setup can tolerate.
type FailureToleranceType string

const (
	// FailureToleranceTypeNode specifies that a highly available resource can tolerate the
	// failure of one or more nodes within a single-zone setup and still be available.
	FailureToleranceTypeNode FailureToleranceType = "node"
	// FailureToleranceTypeZone specifies that a highly available resource can tolerate the
	// failure of one or more zones within a multi-zone setup and still be available.
	FailureToleranceTypeZone FailureToleranceType = "zone"
)

const (
	// GardenerName is the value in a Garden resource's `.metadata.finalizers[]` array on which the Gardener will react
	// when performing a delete request on a resource.
	GardenerName = "gardener"
)
