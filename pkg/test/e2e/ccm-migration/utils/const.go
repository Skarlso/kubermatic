/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

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

package utils

import "time"

const (
	ClusterReadinessCheckPeriod = 10 * time.Second
	ClusterReadinessTimeout     = 10 * time.Minute

	MachineDeploymentName      = "ccm-migration-e2e"
	MachineDeploymentNamespace = "kube-system"

	UserClusterPollInterval = 5 * time.Second
	CustomTestTimeout       = 15 * time.Minute
)