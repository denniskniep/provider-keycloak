package authentication

import "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "authenticationflow"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_authentication_flow", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
	})
	p.AddResourceConfigurator("keycloak_authentication_subflow", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		r.References["parent_flow_alias"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
		}
	})
	p.AddResourceConfigurator("keycloak_authentication_execution", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		r.References["parent_flow_alias"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
		}
	})
	p.AddResourceConfigurator("keycloak_authentication_execution_config", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		r.References["parent_flow_alias"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
		}
	})
	p.AddResourceConfigurator("keycloak_authentication_bindings", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
	})
}
