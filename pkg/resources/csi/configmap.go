/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package csi

import (
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/csi/kubevirt"
	"k8c.io/reconciler/pkg/reconciling"
)

// ConfigMapsReconcilers returns the function to create and update the configmaps needed for CSI.
func ConfigMapsReconcilers(data *resources.TemplateData) []reconciling.NamedConfigMapReconcilerFactory {
	creatorGetters := []reconciling.NamedConfigMapReconcilerFactory{}

	switch {
	case data.Cluster().Spec.Cloud.KubeVirt != nil:
		creatorGetters = kubevirt.ConfigMapsReconcilers(data)
	}

	return creatorGetters
}
