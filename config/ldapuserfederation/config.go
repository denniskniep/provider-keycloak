package ldapuserfederation

import "github.com/crossplane/upjet/pkg/config"

const (
	// Group is the short group for this provider.
	Group = "ldapuserfederation"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("keycloak_ldap_user_federation", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
	})
	p.AddResourceConfigurator("keycloak_ldap_user_attribute_mapper", func(r *config.Resource) {
		r.ShortGroup = Group
		r.References["realm_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm",
		}
		r.References["ldap_user_federation_id"] = config.Reference{
			Type: "github.com/crossplane-contrib/provider-keycloak/apis/ldapuserfederation/v1alpha1.UserFederation",
		}
	})
}
