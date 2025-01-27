/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SubnetNetworkSecurityGroupAssociationObservation struct {

	// The ID of the Subnet.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the Network Security Group which should be associated with the Subnet. Changing this forces a new resource to be created.
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// The ID of the Subnet. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`
}

type SubnetNetworkSecurityGroupAssociationParameters struct {

	// The ID of the Network Security Group which should be associated with the Subnet. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=SecurityGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupID *string `json:"networkSecurityGroupId,omitempty" tf:"network_security_group_id,omitempty"`

	// Reference to a SecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDRef *v1.Reference `json:"networkSecurityGroupIdRef,omitempty" tf:"-"`

	// Selector for a SecurityGroup to populate networkSecurityGroupId.
	// +kubebuilder:validation:Optional
	NetworkSecurityGroupIDSelector *v1.Selector `json:"networkSecurityGroupIdSelector,omitempty" tf:"-"`

	// The ID of the Subnet. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=Subnet
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-azure/apis/rconfig.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`
}

// SubnetNetworkSecurityGroupAssociationSpec defines the desired state of SubnetNetworkSecurityGroupAssociation
type SubnetNetworkSecurityGroupAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetNetworkSecurityGroupAssociationParameters `json:"forProvider"`
}

// SubnetNetworkSecurityGroupAssociationStatus defines the observed state of SubnetNetworkSecurityGroupAssociation.
type SubnetNetworkSecurityGroupAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetNetworkSecurityGroupAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetNetworkSecurityGroupAssociation is the Schema for the SubnetNetworkSecurityGroupAssociations API. Associates a
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SubnetNetworkSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetNetworkSecurityGroupAssociationSpec   `json:"spec"`
	Status            SubnetNetworkSecurityGroupAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetNetworkSecurityGroupAssociationList contains a list of SubnetNetworkSecurityGroupAssociations
type SubnetNetworkSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubnetNetworkSecurityGroupAssociation `json:"items"`
}

// Repository type metadata.
var (
	SubnetNetworkSecurityGroupAssociation_Kind             = "SubnetNetworkSecurityGroupAssociation"
	SubnetNetworkSecurityGroupAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubnetNetworkSecurityGroupAssociation_Kind}.String()
	SubnetNetworkSecurityGroupAssociation_KindAPIVersion   = SubnetNetworkSecurityGroupAssociation_Kind + "." + CRDGroupVersion.String()
	SubnetNetworkSecurityGroupAssociation_GroupVersionKind = CRDGroupVersion.WithKind(SubnetNetworkSecurityGroupAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&SubnetNetworkSecurityGroupAssociation{}, &SubnetNetworkSecurityGroupAssociationList{})
}
