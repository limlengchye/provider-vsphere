/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type VSphereComputeClusterVmDependencyRuleObservation struct {

	// The managed object ID of the cluster.
	ComputeClusterID *string `json:"computeClusterId,omitempty" tf:"compute_cluster_id,omitempty"`

	// The name of the VM group that this rule depends on. The VMs defined in the group specified by vm_group_name will not be started until the VMs in this group are started.
	DependencyVMGroupName *string `json:"dependencyVmGroupName,omitempty" tf:"dependency_vm_group_name,omitempty"`

	// Enable this rule in the cluster.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When true, prevents any virtual machine operations that may violate this rule.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// The unique name of the virtual machine group in the cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the VM group that is the subject of this rule. The VMs defined in this group will not be started until the VMs in the group specified by dependency_vm_group_name are started.
	VMGroupName *string `json:"vmGroupName,omitempty" tf:"vm_group_name,omitempty"`
}

type VSphereComputeClusterVmDependencyRuleParameters struct {

	// The managed object ID of the cluster.
	// +kubebuilder:validation:Optional
	ComputeClusterID *string `json:"computeClusterId,omitempty" tf:"compute_cluster_id,omitempty"`

	// The name of the VM group that this rule depends on. The VMs defined in the group specified by vm_group_name will not be started until the VMs in this group are started.
	// +kubebuilder:validation:Optional
	DependencyVMGroupName *string `json:"dependencyVmGroupName,omitempty" tf:"dependency_vm_group_name,omitempty"`

	// Enable this rule in the cluster.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// When true, prevents any virtual machine operations that may violate this rule.
	// +kubebuilder:validation:Optional
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// The unique name of the virtual machine group in the cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the VM group that is the subject of this rule. The VMs defined in this group will not be started until the VMs in the group specified by dependency_vm_group_name are started.
	// +kubebuilder:validation:Optional
	VMGroupName *string `json:"vmGroupName,omitempty" tf:"vm_group_name,omitempty"`
}

// VSphereComputeClusterVmDependencyRuleSpec defines the desired state of VSphereComputeClusterVmDependencyRule
type VSphereComputeClusterVmDependencyRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereComputeClusterVmDependencyRuleParameters `json:"forProvider"`
}

// VSphereComputeClusterVmDependencyRuleStatus defines the observed state of VSphereComputeClusterVmDependencyRule.
type VSphereComputeClusterVmDependencyRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereComputeClusterVmDependencyRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereComputeClusterVmDependencyRule is the Schema for the VSphereComputeClusterVmDependencyRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereComputeClusterVmDependencyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.computeClusterId)",message="computeClusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.dependencyVmGroupName)",message="dependencyVmGroupName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vmGroupName)",message="vmGroupName is a required parameter"
	Spec   VSphereComputeClusterVmDependencyRuleSpec   `json:"spec"`
	Status VSphereComputeClusterVmDependencyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereComputeClusterVmDependencyRuleList contains a list of VSphereComputeClusterVmDependencyRules
type VSphereComputeClusterVmDependencyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereComputeClusterVmDependencyRule `json:"items"`
}

// Repository type metadata.
var (
	VSphereComputeClusterVmDependencyRule_Kind             = "VSphereComputeClusterVmDependencyRule"
	VSphereComputeClusterVmDependencyRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereComputeClusterVmDependencyRule_Kind}.String()
	VSphereComputeClusterVmDependencyRule_KindAPIVersion   = VSphereComputeClusterVmDependencyRule_Kind + "." + CRDGroupVersion.String()
	VSphereComputeClusterVmDependencyRule_GroupVersionKind = CRDGroupVersion.WithKind(VSphereComputeClusterVmDependencyRule_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereComputeClusterVmDependencyRule{}, &VSphereComputeClusterVmDependencyRuleList{})
}
