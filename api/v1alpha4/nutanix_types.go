/*
Copyright 2022 Nutanix

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

package v1alpha4

// NutanixIdentifierType is an enumeration of different resource identifier types.
type NutanixIdentifierType string

const (
	// NutanixIdentifierUUID is a resource identifier identifying the object by UUID.
	NutanixIdentifierUUID NutanixIdentifierType = "uuid"

	// NutanixIdentifierName is a resource identifier identifying the object by Name.
	NutanixIdentifierName NutanixIdentifierType = "name"

	// NutanixIdentifierBootTypeLegacy is a resource identifier identifying the legacy boot type for virtual machines.
	NutanixIdentifierBootTypeLegacy NutanixIdentifierType = "legacy"

	// NutanixIdentifierBootTypeUEFI is a resource identifier identifying the UEFI boot type for virtual machines.
	NutanixIdentifierBootTypeUEFI NutanixIdentifierType = "uefi"

	// DefaultCAPICategoryPrefix is the default category prefix used for CAPI clusters.
	DefaultCAPICategoryPrefix = "kubernetes-io-cluster-"

	// DefaultCAPICategoryDescription is the default category description used for CAPI clusters.
	DefaultCAPICategoryDescription = "Managed by CAPX"

	// DefaultCAPICategoryOwnedValue is the default category value used for CAPI clusters.
	DefaultCAPICategoryOwnedValue = "owned"
)

// NutanixResourceIdentifier holds the identity of a Nutanix PC resource (cluster, image, subnet, etc.)
// +union
type NutanixResourceIdentifier struct {
	// Type is the identifier type to use for this resource.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum:=uuid;name
	Type NutanixIdentifierType `json:"type"`

	// uuid is the UUID of the resource in the PC.
	// +optional
	UUID *string `json:"uuid,omitempty"`

	// name is the resource name in the PC
	// +optional
	Name *string `json:"name,omitempty"`
}

type NutanixCategoryIdentifier struct {
	// key is the Key of category in PC.
	// +optional
	Key string `json:"key,omitempty"`

	// value is the category value linked to the category key in PC
	// +optional
	Value string `json:"value,omitempty"`
}
