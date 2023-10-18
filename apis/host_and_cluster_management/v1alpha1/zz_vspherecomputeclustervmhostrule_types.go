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

type VSphereComputeClusterVmHostRuleObservation struct {

	// When this field is used, virtual machines defined in vm_group_name will be run on the hosts defined in this host group.
	AffinityHostGroupName *string `json:"affinityHostGroupName,omitempty" tf:"affinity_host_group_name,omitempty"`

	// When this field is used, virtual machines defined in vm_group_name will not be run on the hosts defined in this host group.
	AntiAffinityHostGroupName *string `json:"antiAffinityHostGroupName,omitempty" tf:"anti_affinity_host_group_name,omitempty"`

	// The managed object ID of the cluster.
	ComputeClusterID *string `json:"computeClusterId,omitempty" tf:"compute_cluster_id,omitempty"`

	// Enable this rule in the cluster.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When true, prevents any virtual machine operations that may violate this rule.
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// The unique name of the virtual machine group in the cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the virtual machine group to use with this rule.
	VMGroupName *string `json:"vmGroupName,omitempty" tf:"vm_group_name,omitempty"`
}

type VSphereComputeClusterVmHostRuleParameters struct {

	// When this field is used, virtual machines defined in vm_group_name will be run on the hosts defined in this host group.
	// +kubebuilder:validation:Optional
	AffinityHostGroupName *string `json:"affinityHostGroupName,omitempty" tf:"affinity_host_group_name,omitempty"`

	// When this field is used, virtual machines defined in vm_group_name will not be run on the hosts defined in this host group.
	// +kubebuilder:validation:Optional
	AntiAffinityHostGroupName *string `json:"antiAffinityHostGroupName,omitempty" tf:"anti_affinity_host_group_name,omitempty"`

	// The managed object ID of the cluster.
	// +kubebuilder:validation:Optional
	ComputeClusterID *string `json:"computeClusterId,omitempty" tf:"compute_cluster_id,omitempty"`

	// Enable this rule in the cluster.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// When true, prevents any virtual machine operations that may violate this rule.
	// +kubebuilder:validation:Optional
	Mandatory *bool `json:"mandatory,omitempty" tf:"mandatory,omitempty"`

	// The unique name of the virtual machine group in the cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the virtual machine group to use with this rule.
	// +kubebuilder:validation:Optional
	VMGroupName *string `json:"vmGroupName,omitempty" tf:"vm_group_name,omitempty"`
}

// VSphereComputeClusterVmHostRuleSpec defines the desired state of VSphereComputeClusterVmHostRule
type VSphereComputeClusterVmHostRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VSphereComputeClusterVmHostRuleParameters `json:"forProvider"`
}

// VSphereComputeClusterVmHostRuleStatus defines the observed state of VSphereComputeClusterVmHostRule.
type VSphereComputeClusterVmHostRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VSphereComputeClusterVmHostRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereComputeClusterVmHostRule is the Schema for the VSphereComputeClusterVmHostRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vsphere}
type VSphereComputeClusterVmHostRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.computeClusterId)",message="computeClusterId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.vmGroupName)",message="vmGroupName is a required parameter"
	Spec   VSphereComputeClusterVmHostRuleSpec   `json:"spec"`
	Status VSphereComputeClusterVmHostRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VSphereComputeClusterVmHostRuleList contains a list of VSphereComputeClusterVmHostRules
type VSphereComputeClusterVmHostRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VSphereComputeClusterVmHostRule `json:"items"`
}

// Repository type metadata.
var (
	VSphereComputeClusterVmHostRule_Kind             = "VSphereComputeClusterVmHostRule"
	VSphereComputeClusterVmHostRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VSphereComputeClusterVmHostRule_Kind}.String()
	VSphereComputeClusterVmHostRule_KindAPIVersion   = VSphereComputeClusterVmHostRule_Kind + "." + CRDGroupVersion.String()
	VSphereComputeClusterVmHostRule_GroupVersionKind = CRDGroupVersion.WithKind(VSphereComputeClusterVmHostRule_Kind)
)

func init() {
	SchemeBuilder.Register(&VSphereComputeClusterVmHostRule{}, &VSphereComputeClusterVmHostRuleList{})
}
