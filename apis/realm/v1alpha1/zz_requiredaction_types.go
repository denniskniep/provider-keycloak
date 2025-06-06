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

type RequiredActionInitParameters struct {

	// The alias of the action to attach as a required action.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// The configuration. Keys are specific to each configurable required action and not checked when applying.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// When true, the required action is set as the default action for new users. Defaults to false.
	DefaultAction *bool `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// When false, the required action is not enabled for new users. Defaults to false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the required action.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the required action.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The realm the required action exists in.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

type RequiredActionObservation struct {

	// The alias of the action to attach as a required action.
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// The configuration. Keys are specific to each configurable required action and not checked when applying.
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// When true, the required action is set as the default action for new users. Defaults to false.
	DefaultAction *bool `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// When false, the required action is not enabled for new users. Defaults to false.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the required action.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the required action.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The realm the required action exists in.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`
}

type RequiredActionParameters struct {

	// The alias of the action to attach as a required action.
	// +kubebuilder:validation:Optional
	Alias *string `json:"alias,omitempty" tf:"alias,omitempty"`

	// The configuration. Keys are specific to each configurable required action and not checked when applying.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// When true, the required action is set as the default action for new users. Defaults to false.
	// +kubebuilder:validation:Optional
	DefaultAction *bool `json:"defaultAction,omitempty" tf:"default_action,omitempty"`

	// When false, the required action is not enabled for new users. Defaults to false.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// The name of the required action.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The priority of the required action.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// The realm the required action exists in.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// RequiredActionSpec defines the desired state of RequiredAction
type RequiredActionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RequiredActionParameters `json:"forProvider"`
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
	InitProvider RequiredActionInitParameters `json:"initProvider,omitempty"`
}

// RequiredActionStatus defines the observed state of RequiredAction.
type RequiredActionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RequiredActionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RequiredAction is the Schema for the RequiredActions API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type RequiredAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.alias) || (has(self.initProvider) && has(self.initProvider.alias))",message="spec.forProvider.alias is a required parameter"
	Spec   RequiredActionSpec   `json:"spec"`
	Status RequiredActionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RequiredActionList contains a list of RequiredActions
type RequiredActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RequiredAction `json:"items"`
}

// Repository type metadata.
var (
	RequiredAction_Kind             = "RequiredAction"
	RequiredAction_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RequiredAction_Kind}.String()
	RequiredAction_KindAPIVersion   = RequiredAction_Kind + "." + CRDGroupVersion.String()
	RequiredAction_GroupVersionKind = CRDGroupVersion.WithKind(RequiredAction_Kind)
)

func init() {
	SchemeBuilder.Register(&RequiredAction{}, &RequiredActionList{})
}
