// SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/gardener/gardener/plugin/pkg/shoot/dnsrewriting/apis/shootdnsrewriting"
	"github.com/gardener/gardener/plugin/pkg/shoot/dnsrewriting/apis/shootdnsrewriting/v1alpha1"
)

// Install registers the API group and adds types to a scheme.
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(shootdnsrewriting.AddToScheme(scheme))
	utilruntime.Must(v1alpha1.AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(v1alpha1.SchemeGroupVersion))
}
