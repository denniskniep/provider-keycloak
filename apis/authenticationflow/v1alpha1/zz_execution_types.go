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

type ExecutionInitParameters struct {

	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator *string `json:"authenticator,omitempty" tf:"authenticator,omitempty"`

	// When true, the authentication execution with the specified authenticator inside the authentication flow with the specified alias parent_flow_alias is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with authentication executions that Keycloak creates automatically during realm creation, such as browser/identity-provider-redirector and registration/registration-user-creation. Note, that the execution will not be removed during destruction if import is true.
	Import *bool `json:"import,omitempty" tf:"import,omitempty"`

	// The alias of the flow this execution is attached to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.AuthenticationFlowAliasExtractor()
	// +crossplane:generate:reference:refFieldName=ParentFlowAliasRef
	// +crossplane:generate:reference:selectorFieldName=ParentFlowAliasSelector
	ParentFlowAlias *string `json:"parentFlowAlias,omitempty" tf:"parent_flow_alias,omitempty"`

	// Reference to a Flow in authenticationflow to populate parentFlowAlias.
	// +kubebuilder:validation:Optional
	ParentFlowAliasRef *v1.Reference `json:"parentFlowAliasRef,omitempty" tf:"-"`

	// Selector for a Flow in authenticationflow to populate parentFlowAlias.
	// +kubebuilder:validation:Optional
	ParentFlowAliasSelector *v1.Selector `json:"parentFlowAliasSelector,omitempty" tf:"-"`

	// The realm the authentication execution exists in.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The requirement setting, which can be one of REQUIRED, ALTERNATIVE, OPTIONAL, CONDITIONAL, or DISABLED. Defaults to DISABLED.
	Requirement *string `json:"requirement,omitempty" tf:"requirement,omitempty"`
}

type ExecutionObservation struct {

	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	Authenticator *string `json:"authenticator,omitempty" tf:"authenticator,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When true, the authentication execution with the specified authenticator inside the authentication flow with the specified alias parent_flow_alias is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with authentication executions that Keycloak creates automatically during realm creation, such as browser/identity-provider-redirector and registration/registration-user-creation. Note, that the execution will not be removed during destruction if import is true.
	Import *bool `json:"import,omitempty" tf:"import,omitempty"`

	// The alias of the flow this execution is attached to.
	ParentFlowAlias *string `json:"parentFlowAlias,omitempty" tf:"parent_flow_alias,omitempty"`

	// The realm the authentication execution exists in.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// The requirement setting, which can be one of REQUIRED, ALTERNATIVE, OPTIONAL, CONDITIONAL, or DISABLED. Defaults to DISABLED.
	Requirement *string `json:"requirement,omitempty" tf:"requirement,omitempty"`
}

type ExecutionParameters struct {

	// The name of the authenticator. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools.
	// +kubebuilder:validation:Optional
	Authenticator *string `json:"authenticator,omitempty" tf:"authenticator,omitempty"`

	// When true, the authentication execution with the specified authenticator inside the authentication flow with the specified alias parent_flow_alias is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with authentication executions that Keycloak creates automatically during realm creation, such as browser/identity-provider-redirector and registration/registration-user-creation. Note, that the execution will not be removed during destruction if import is true.
	// +kubebuilder:validation:Optional
	Import *bool `json:"import,omitempty" tf:"import,omitempty"`

	// The alias of the flow this execution is attached to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.AuthenticationFlowAliasExtractor()
	// +crossplane:generate:reference:refFieldName=ParentFlowAliasRef
	// +crossplane:generate:reference:selectorFieldName=ParentFlowAliasSelector
	// +kubebuilder:validation:Optional
	ParentFlowAlias *string `json:"parentFlowAlias,omitempty" tf:"parent_flow_alias,omitempty"`

	// Reference to a Flow in authenticationflow to populate parentFlowAlias.
	// +kubebuilder:validation:Optional
	ParentFlowAliasRef *v1.Reference `json:"parentFlowAliasRef,omitempty" tf:"-"`

	// Selector for a Flow in authenticationflow to populate parentFlowAlias.
	// +kubebuilder:validation:Optional
	ParentFlowAliasSelector *v1.Selector `json:"parentFlowAliasSelector,omitempty" tf:"-"`

	// The realm the authentication execution exists in.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The requirement setting, which can be one of REQUIRED, ALTERNATIVE, OPTIONAL, CONDITIONAL, or DISABLED. Defaults to DISABLED.
	// +kubebuilder:validation:Optional
	Requirement *string `json:"requirement,omitempty" tf:"requirement,omitempty"`
}

// ExecutionSpec defines the desired state of Execution
type ExecutionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ExecutionParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ExecutionInitParameters `json:"initProvider,omitempty"`
}

// ExecutionStatus defines the observed state of Execution.
type ExecutionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ExecutionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Execution is the Schema for the Executions API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type Execution struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.authenticator) || (has(self.initProvider) && has(self.initProvider.authenticator))",message="spec.forProvider.authenticator is a required parameter"
	Spec   ExecutionSpec   `json:"spec"`
	Status ExecutionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExecutionList contains a list of Executions
type ExecutionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Execution `json:"items"`
}

// Repository type metadata.
var (
	Execution_Kind             = "Execution"
	Execution_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Execution_Kind}.String()
	Execution_KindAPIVersion   = Execution_Kind + "." + CRDGroupVersion.String()
	Execution_GroupVersionKind = CRDGroupVersion.WithKind(Execution_Kind)
)

func init() {
	SchemeBuilder.Register(&Execution{}, &ExecutionList{})
}
