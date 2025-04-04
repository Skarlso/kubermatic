/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

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

package v1

import (
	"encoding/json"

	semverlib "github.com/Masterminds/semver/v3"

	kubermaticv1 "k8c.io/kubermatic/sdk/v2/apis/kubermatic/v1"
)

// ObjectMeta defines the set of fields that objects returned from the API have
// swagger:model ObjectMeta
type ObjectMeta struct {
	// ID unique value that identifies the resource generated by the server. Read-Only.
	ID string `json:"id,omitempty"`

	// Name represents human readable name for the resource
	Name string `json:"name"`

	// Namespace represents the namespace for the resource.
	Namespace string `json:"namespace,omitempty"`

	// Annotations that can be added to the resource
	Annotations map[string]string `json:"annotations,omitempty"`

	// DeletionTimestamp is a timestamp representing the server time when this object was deleted.
	// swagger:strfmt date-time
	DeletionTimestamp *Time `json:"deletionTimestamp,omitempty"`

	// CreationTimestamp is a timestamp representing the server time when this object was created.
	// swagger:strfmt date-time
	CreationTimestamp Time `json:"creationTimestamp,omitempty"`
}

// GCPMachineSizeList represents an array of GCP machine sizes.
// swagger:model GCPMachineSizeList
type GCPMachineSizeList []GCPMachineSize

// GCPMachineSize represents a object of GCP machine size.
// swagger:model GCPMachineSize
type GCPMachineSize struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Memory      int64  `json:"memory"`
	VCPUs       int64  `json:"vcpus"`
}

// GCPNetwork represents a object of GCP networks.
// swagger:model GCPNetwork
type GCPNetwork struct {
	ID                    uint64   `json:"id"`
	Name                  string   `json:"name"`
	AutoCreateSubnetworks bool     `json:"autoCreateSubnetworks"`
	Subnetworks           []string `json:"subnetworks"`
	Kind                  string   `json:"kind"`
	Path                  string   `json:"path"`
}

// GCPSubnetworkList represents an array of GCP subnetworks.
// swagger:model GCPSubnetworkList
type GCPSubnetworkList []GCPSubnetwork

// GCPSubnetwork represents a object of GCP subnetworks.
// swagger:model GCPSubnetwork
type GCPSubnetwork struct {
	ID                    uint64                `json:"id"`
	Name                  string                `json:"name"`
	Network               string                `json:"network"`
	IPCidrRange           string                `json:"ipCidrRange"`
	GatewayAddress        string                `json:"gatewayAddress"`
	Region                string                `json:"region"`
	SelfLink              string                `json:"selfLink"`
	PrivateIPGoogleAccess bool                  `json:"privateIpGoogleAccess"` //nolint:tagliatelle
	Kind                  string                `json:"kind"`
	Path                  string                `json:"path"`
	IPFamily              kubermaticv1.IPFamily `json:"ipFamily"`
}

// VMwareCloudDirectorCatalog represents a VMware Cloud Director catalog.
// swagger:model VMwareCloudDirectorCatalog
type VMwareCloudDirectorCatalog struct {
	Name string `json:"name"`
}

// VMwareCloudDirectorCatalogList represents an array of VMware Cloud Director catalogs.
// swagger:model VMwareCloudDirectorCatalogList
type VMwareCloudDirectorCatalogList []VMwareCloudDirectorCatalog

// VMwareCloudDirectorTemplate represents a VMware Cloud Director template.
// swagger:model VMwareCloudDirectorTemplate
type VMwareCloudDirectorTemplate struct {
	Name string `json:"name"`
}

// VMwareCloudDirectorTemplateList represents an array of VMware Cloud Director templates.
// swagger:model VMwareCloudDirectorTemplateList
type VMwareCloudDirectorTemplateList []VMwareCloudDirectorTemplate

// VMwareCloudDirectorNetwork represents a VMware Cloud Director network.
// swagger:model VMwareCloudDirectorNetwork
type VMwareCloudDirectorNetwork struct {
	Name string `json:"name"`
}

// VMwareCloudDirectorNetworkList represents an array of VMware Cloud Director networks.
// swagger:model VMwareCloudDirectorNetworkList
type VMwareCloudDirectorNetworkList []VMwareCloudDirectorNetwork

// VMwareCloudDirectorStorageProfile represents a VMware Cloud Director storage profile.
// swagger:model VMwareCloudDirectorStorageProfile
type VMwareCloudDirectorStorageProfile struct {
	Name string `json:"name"`
}

// VMwareCloudDirectorStorageProfileList represents an array of VMware Cloud Director storage profiles.
// swagger:model VMwareCloudDirectorStorageProfileList
type VMwareCloudDirectorStorageProfileList []VMwareCloudDirectorStorageProfile

// MasterVersion describes a version of the master components
// swagger:model MasterVersion
type MasterVersion struct {
	Version *semverlib.Version `json:"version"`
	Default bool               `json:"default,omitempty"`

	// If true, then given version control plane version is not compatible
	// with one of the kubelets inside cluster and shouldn't be used.
	RestrictedByKubeletVersion bool `json:"restrictedByKubeletVersion,omitempty"`
}

// Application represents a set of applications that are to be installed for the cluster
// swagger:model Application
type Application struct {
	ObjectMeta `json:",inline"`

	Spec ApplicationSpec `json:"spec"`
}

// ApplicationSpec represents the specification for an application
// swagger:model ApplicationSpec
type ApplicationSpec struct {
	// Namespace describe the desired state of the namespace where application will be created.
	Namespace NamespaceSpec `json:"namespace"`

	// ApplicationRef is a reference to identify which Application should be deployed
	ApplicationRef ApplicationRef `json:"applicationRef"`

	// Values specify values overrides that are passed to helm templating. Comments are not preserved.
	// Deprecated: Use ValuesBlock instead
	Values json.RawMessage `json:"values,omitempty"`

	// ValuesBlock specifies values overrides that are passed to helm templating. Comments are preserved.
	ValuesBlock string `json:"valuesBlock,omitempty"`
}

type NamespaceSpec struct {
	// Name is the namespace to deploy the Application into
	Name string `json:"name" required:"true"`

	// Create defines whether the namespace should be created if it does not exist.
	Create bool `json:"create" required:"true"`

	// Labels of the namespace
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations of the namespace
	Annotations map[string]string `json:"annotations,omitempty"`
}

type ApplicationRef struct {
	// Name of the Application
	Name string `json:"name" required:"true"`

	// Version of the Application. Must be a valid SemVer version
	// NOTE: We are not using Masterminds/semver here, as it keeps data in unexported fields witch causes issues for
	// DeepEqual used in our reconciliation packages. At the same time, we are not using pkg/semver because
	// of the reasons stated in https://github.com/kubermatic/kubermatic/pull/10891.
	Version string `json:"version" required:"true"`
}
