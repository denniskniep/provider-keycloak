package authentication

import (
	"github.com/crossplane-contrib/provider-keycloak/config/common"
	"github.com/crossplane/upjet/pkg/config"
)

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
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
	})
	p.AddResourceConfigurator("keycloak_authentication_execution", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		r.References["parent_flow_alias"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
	})
	p.AddResourceConfigurator("keycloak_authentication_execution_config", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		r.References["parent_flow_alias"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
		r.References["execution_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Execution",
		}
	})
	p.AddResourceConfigurator("keycloak_authentication_bindings", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		r.References["browser_flow"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
		r.References["registration_flow"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
		r.References["direct_grant_flow"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
		r.References["reset_credentials_flow"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
		r.References["client_authentication_flow"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
		r.References["docker_authentication_flow"] = config.Reference{
			Type:      "github.com/crossplane-contrib/provider-keycloak/apis/authenticationflow/v1alpha1.Flow",
			Extractor: common.PathAuthenticationFlowAliasExtractor,
		}
	})
}
