// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/crossplane-contrib/provider-keycloak/apis/client/v1alpha1"
	v1alpha1defaults "github.com/crossplane-contrib/provider-keycloak/apis/defaults/v1alpha1"
	v1alpha1group "github.com/crossplane-contrib/provider-keycloak/apis/group/v1alpha1"
	v1alpha1oidc "github.com/crossplane-contrib/provider-keycloak/apis/oidc/v1alpha1"
	v1alpha1openidclient "github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1"
	v1alpha1openidgroup "github.com/crossplane-contrib/provider-keycloak/apis/openidgroup/v1alpha1"
	v1alpha1realm "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1"
	v1alpha1role "github.com/crossplane-contrib/provider-keycloak/apis/role/v1alpha1"
	v1alpha1saml "github.com/crossplane-contrib/provider-keycloak/apis/saml/v1alpha1"
	v1alpha1samlclient "github.com/crossplane-contrib/provider-keycloak/apis/samlclient/v1alpha1"
	v1alpha1user "github.com/crossplane-contrib/provider-keycloak/apis/user/v1alpha1"
	v1alpha1apis "github.com/crossplane-contrib/provider-keycloak/apis/v1alpha1"
	v1beta1 "github.com/crossplane-contrib/provider-keycloak/apis/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1defaults.SchemeBuilder.AddToScheme,
		v1alpha1group.SchemeBuilder.AddToScheme,
		v1alpha1oidc.SchemeBuilder.AddToScheme,
		v1alpha1openidclient.SchemeBuilder.AddToScheme,
		v1alpha1openidgroup.SchemeBuilder.AddToScheme,
		v1alpha1realm.SchemeBuilder.AddToScheme,
		v1alpha1role.SchemeBuilder.AddToScheme,
		v1alpha1saml.SchemeBuilder.AddToScheme,
		v1alpha1samlclient.SchemeBuilder.AddToScheme,
		v1alpha1user.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
