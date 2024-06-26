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

type CacheInitParameters struct {

	// Day of the week the entry will become invalid on
	// Day of the week the entry will become invalid on.
	EvictionDay *float64 `json:"evictionDay,omitempty" tf:"eviction_day,omitempty"`

	// Hour of day the entry will become invalid on.
	// Hour of day the entry will become invalid on.
	EvictionHour *float64 `json:"evictionHour,omitempty" tf:"eviction_hour,omitempty"`

	// Minute of day the entry will become invalid on.
	// Minute of day the entry will become invalid on.
	EvictionMinute *float64 `json:"evictionMinute,omitempty" tf:"eviction_minute,omitempty"`

	// Max lifespan of cache entry (duration string).
	// Max lifespan of cache entry (duration string).
	MaxLifespan *string `json:"maxLifespan,omitempty" tf:"max_lifespan,omitempty"`

	// Can be one of DEFAULT, EVICT_DAILY, EVICT_WEEKLY, MAX_LIFESPAN, or NO_CACHE. Defaults to DEFAULT.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type CacheObservation struct {

	// Day of the week the entry will become invalid on
	// Day of the week the entry will become invalid on.
	EvictionDay *float64 `json:"evictionDay,omitempty" tf:"eviction_day,omitempty"`

	// Hour of day the entry will become invalid on.
	// Hour of day the entry will become invalid on.
	EvictionHour *float64 `json:"evictionHour,omitempty" tf:"eviction_hour,omitempty"`

	// Minute of day the entry will become invalid on.
	// Minute of day the entry will become invalid on.
	EvictionMinute *float64 `json:"evictionMinute,omitempty" tf:"eviction_minute,omitempty"`

	// Max lifespan of cache entry (duration string).
	// Max lifespan of cache entry (duration string).
	MaxLifespan *string `json:"maxLifespan,omitempty" tf:"max_lifespan,omitempty"`

	// Can be one of DEFAULT, EVICT_DAILY, EVICT_WEEKLY, MAX_LIFESPAN, or NO_CACHE. Defaults to DEFAULT.
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type CacheParameters struct {

	// Day of the week the entry will become invalid on
	// Day of the week the entry will become invalid on.
	// +kubebuilder:validation:Optional
	EvictionDay *float64 `json:"evictionDay,omitempty" tf:"eviction_day,omitempty"`

	// Hour of day the entry will become invalid on.
	// Hour of day the entry will become invalid on.
	// +kubebuilder:validation:Optional
	EvictionHour *float64 `json:"evictionHour,omitempty" tf:"eviction_hour,omitempty"`

	// Minute of day the entry will become invalid on.
	// Minute of day the entry will become invalid on.
	// +kubebuilder:validation:Optional
	EvictionMinute *float64 `json:"evictionMinute,omitempty" tf:"eviction_minute,omitempty"`

	// Max lifespan of cache entry (duration string).
	// Max lifespan of cache entry (duration string).
	// +kubebuilder:validation:Optional
	MaxLifespan *string `json:"maxLifespan,omitempty" tf:"max_lifespan,omitempty"`

	// Can be one of DEFAULT, EVICT_DAILY, EVICT_WEEKLY, MAX_LIFESPAN, or NO_CACHE. Defaults to DEFAULT.
	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy,omitempty"`
}

type KerberosInitParameters struct {

	// The name of the kerberos realm, e.g. FOO.LOCAL.
	// The name of the kerberos realm, e.g. FOO.LOCAL
	KerberosRealm *string `json:"kerberosRealm,omitempty" tf:"kerberos_realm,omitempty"`

	// Path to the kerberos keytab file on the server with credentials of the service principal.
	// Path to the kerberos keytab file on the server with credentials of the service principal.
	KeyTab *string `json:"keyTab,omitempty" tf:"key_tab,omitempty"`

	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	ServerPrincipal *string `json:"serverPrincipal,omitempty" tf:"server_principal,omitempty"`

	// Use kerberos login module instead of ldap service api. Defaults to false.
	// Use kerberos login module instead of ldap service api. Defaults to `false`.
	UseKerberosForPasswordAuthentication *bool `json:"useKerberosForPasswordAuthentication,omitempty" tf:"use_kerberos_for_password_authentication,omitempty"`
}

type KerberosObservation struct {

	// The name of the kerberos realm, e.g. FOO.LOCAL.
	// The name of the kerberos realm, e.g. FOO.LOCAL
	KerberosRealm *string `json:"kerberosRealm,omitempty" tf:"kerberos_realm,omitempty"`

	// Path to the kerberos keytab file on the server with credentials of the service principal.
	// Path to the kerberos keytab file on the server with credentials of the service principal.
	KeyTab *string `json:"keyTab,omitempty" tf:"key_tab,omitempty"`

	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	ServerPrincipal *string `json:"serverPrincipal,omitempty" tf:"server_principal,omitempty"`

	// Use kerberos login module instead of ldap service api. Defaults to false.
	// Use kerberos login module instead of ldap service api. Defaults to `false`.
	UseKerberosForPasswordAuthentication *bool `json:"useKerberosForPasswordAuthentication,omitempty" tf:"use_kerberos_for_password_authentication,omitempty"`
}

type KerberosParameters struct {

	// The name of the kerberos realm, e.g. FOO.LOCAL.
	// The name of the kerberos realm, e.g. FOO.LOCAL
	// +kubebuilder:validation:Optional
	KerberosRealm *string `json:"kerberosRealm" tf:"kerberos_realm,omitempty"`

	// Path to the kerberos keytab file on the server with credentials of the service principal.
	// Path to the kerberos keytab file on the server with credentials of the service principal.
	// +kubebuilder:validation:Optional
	KeyTab *string `json:"keyTab" tf:"key_tab,omitempty"`

	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	// +kubebuilder:validation:Optional
	ServerPrincipal *string `json:"serverPrincipal" tf:"server_principal,omitempty"`

	// Use kerberos login module instead of ldap service api. Defaults to false.
	// Use kerberos login module instead of ldap service api. Defaults to `false`.
	// +kubebuilder:validation:Optional
	UseKerberosForPasswordAuthentication *bool `json:"useKerberosForPasswordAuthentication,omitempty" tf:"use_kerberos_for_password_authentication,omitempty"`
}

type UserFederationInitParameters struct {

	// The number of users to sync within a single transaction. Defaults to 1000.
	// The number of users to sync within a single transaction.
	BatchSizeForSync *float64 `json:"batchSizeForSync,omitempty" tf:"batch_size_for_sync,omitempty"`

	// Password of LDAP admin. This attribute must be set if bind_dn is set.
	// Password of LDAP admin.
	BindCredentialSecretRef *v1.SecretKeySelector `json:"bindCredentialSecretRef,omitempty" tf:"-"`

	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if bind_credential is set.
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	BindDn *string `json:"bindDn,omitempty" tf:"bind_dn,omitempty"`

	// A block containing the cache settings.
	// Settings regarding cache policy for this realm.
	Cache []CacheInitParameters `json:"cache,omitempty" tf:"cache,omitempty"`

	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod *float64 `json:"changedSyncPeriod,omitempty" tf:"changed_sync_period,omitempty"`

	// LDAP connection timeout in the format of a Go duration string.
	// LDAP connection timeout (duration string)
	ConnectionTimeout *string `json:"connectionTimeout,omitempty" tf:"connection_timeout,omitempty"`

	// Connection URL to the LDAP server.
	// Connection URL to the LDAP server.
	ConnectionURL *string `json:"connectionUrl,omitempty" tf:"connection_url,omitempty"`

	// Additional LDAP filter for filtering searched users. Must begin with ( and end with ).
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	CustomUserSearchFilter *string `json:"customUserSearchFilter,omitempty" tf:"custom_user_search_filter,omitempty"`

	// When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP user federation provider. Defaults to false.
	// When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP user federation provider.
	DeleteDefaultMappers *bool `json:"deleteDefaultMappers,omitempty" tf:"delete_default_mappers,omitempty"`

	// Can be one of READ_ONLY, WRITABLE, or UNSYNCED. UNSYNCED allows user data to be imported but not synced back to LDAP. Defaults to READ_ONLY.
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	EditMode *string `json:"editMode,omitempty" tf:"edit_mode,omitempty"`

	// When false, this provider will not be used when performing queries for users. Defaults to true.
	// When false, this provider will not be used when performing queries for users.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *float64 `json:"fullSyncPeriod,omitempty" tf:"full_sync_period,omitempty"`

	// When true, LDAP users will be imported into the Keycloak database. Defaults to true.
	// When true, LDAP users will be imported into the Keycloak database.
	ImportEnabled *bool `json:"importEnabled,omitempty" tf:"import_enabled,omitempty"`

	// A block containing the kerberos settings.
	// Settings regarding kerberos authentication for this realm.
	Kerberos []KerberosInitParameters `json:"kerberos,omitempty" tf:"kerberos,omitempty"`

	// Display name of the provider when displayed in the console.
	// Display name of the provider when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to true.
	// When true, Keycloak assumes the LDAP server supports pagination.
	Pagination *bool `json:"pagination,omitempty" tf:"pagination,omitempty"`

	// Priority of this provider when looking up users. Lower values are first. Defaults to 0.
	// Priority of this provider when looking up users. Lower values are first.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Name of the LDAP attribute to use as the relative distinguished name.
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute *string `json:"rdnLdapAttribute,omitempty" tf:"rdn_ldap_attribute,omitempty"`

	// LDAP read timeout in the format of a Go duration string.
	// LDAP read timeout (duration string)
	ReadTimeout *string `json:"readTimeout,omitempty" tf:"read_timeout,omitempty"`

	// The realm that this provider will provide user federation for.
	// The realm this provider will provide user federation for.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// Can be one of ONE_LEVEL or SUBTREE:
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	SearchScope *string `json:"searchScope,omitempty" tf:"search_scope,omitempty"`

	// When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
	// When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
	StartTLS *bool `json:"startTls,omitempty" tf:"start_tls,omitempty"`

	// When true, newly created users will be synced back to LDAP. Defaults to false.
	// When true, newly created users will be synced back to LDAP.
	SyncRegistrations *bool `json:"syncRegistrations,omitempty" tf:"sync_registrations,omitempty"`

	// If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
	// If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail *bool `json:"trustEmail,omitempty" tf:"trust_email,omitempty"`

	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UUIDLdapAttribute *string `json:"uuidLdapAttribute,omitempty" tf:"uuid_ldap_attribute,omitempty"`

	// When true, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
	// When `true`, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
	UsePasswordModifyExtendedOp *bool `json:"usePasswordModifyExtendedOp,omitempty" tf:"use_password_modify_extended_op,omitempty"`

	// Can be one of ALWAYS, ONLY_FOR_LDAPS, or NEVER:
	UseTruststoreSpi *string `json:"useTruststoreSpi,omitempty" tf:"use_truststore_spi,omitempty"`

	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	// All values of LDAP objectClass attribute for users in LDAP.
	UserObjectClasses []*string `json:"userObjectClasses,omitempty" tf:"user_object_classes,omitempty"`

	// Name of the LDAP attribute to use as the Keycloak username.
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute *string `json:"usernameLdapAttribute,omitempty" tf:"username_ldap_attribute,omitempty"`

	// Full DN of LDAP tree where your users are.
	// Full DN of LDAP tree where your users are.
	UsersDn *string `json:"usersDn,omitempty" tf:"users_dn,omitempty"`

	// When true, Keycloak will validate passwords using the realm policy before updating it.
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy *bool `json:"validatePasswordPolicy,omitempty" tf:"validate_password_policy,omitempty"`

	// Can be one of OTHER, EDIRECTORY, AD, RHDS, or TIVOLI. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to OTHER.
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type UserFederationObservation struct {

	// The number of users to sync within a single transaction. Defaults to 1000.
	// The number of users to sync within a single transaction.
	BatchSizeForSync *float64 `json:"batchSizeForSync,omitempty" tf:"batch_size_for_sync,omitempty"`

	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if bind_credential is set.
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	BindDn *string `json:"bindDn,omitempty" tf:"bind_dn,omitempty"`

	// A block containing the cache settings.
	// Settings regarding cache policy for this realm.
	Cache []CacheObservation `json:"cache,omitempty" tf:"cache,omitempty"`

	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	ChangedSyncPeriod *float64 `json:"changedSyncPeriod,omitempty" tf:"changed_sync_period,omitempty"`

	// LDAP connection timeout in the format of a Go duration string.
	// LDAP connection timeout (duration string)
	ConnectionTimeout *string `json:"connectionTimeout,omitempty" tf:"connection_timeout,omitempty"`

	// Connection URL to the LDAP server.
	// Connection URL to the LDAP server.
	ConnectionURL *string `json:"connectionUrl,omitempty" tf:"connection_url,omitempty"`

	// Additional LDAP filter for filtering searched users. Must begin with ( and end with ).
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	CustomUserSearchFilter *string `json:"customUserSearchFilter,omitempty" tf:"custom_user_search_filter,omitempty"`

	// When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP user federation provider. Defaults to false.
	// When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP user federation provider.
	DeleteDefaultMappers *bool `json:"deleteDefaultMappers,omitempty" tf:"delete_default_mappers,omitempty"`

	// Can be one of READ_ONLY, WRITABLE, or UNSYNCED. UNSYNCED allows user data to be imported but not synced back to LDAP. Defaults to READ_ONLY.
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	EditMode *string `json:"editMode,omitempty" tf:"edit_mode,omitempty"`

	// When false, this provider will not be used when performing queries for users. Defaults to true.
	// When false, this provider will not be used when performing queries for users.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	FullSyncPeriod *float64 `json:"fullSyncPeriod,omitempty" tf:"full_sync_period,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When true, LDAP users will be imported into the Keycloak database. Defaults to true.
	// When true, LDAP users will be imported into the Keycloak database.
	ImportEnabled *bool `json:"importEnabled,omitempty" tf:"import_enabled,omitempty"`

	// A block containing the kerberos settings.
	// Settings regarding kerberos authentication for this realm.
	Kerberos []KerberosObservation `json:"kerberos,omitempty" tf:"kerberos,omitempty"`

	// Display name of the provider when displayed in the console.
	// Display name of the provider when displayed in the console.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to true.
	// When true, Keycloak assumes the LDAP server supports pagination.
	Pagination *bool `json:"pagination,omitempty" tf:"pagination,omitempty"`

	// Priority of this provider when looking up users. Lower values are first. Defaults to 0.
	// Priority of this provider when looking up users. Lower values are first.
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Name of the LDAP attribute to use as the relative distinguished name.
	// Name of the LDAP attribute to use as the relative distinguished name.
	RdnLdapAttribute *string `json:"rdnLdapAttribute,omitempty" tf:"rdn_ldap_attribute,omitempty"`

	// LDAP read timeout in the format of a Go duration string.
	// LDAP read timeout (duration string)
	ReadTimeout *string `json:"readTimeout,omitempty" tf:"read_timeout,omitempty"`

	// The realm that this provider will provide user federation for.
	// The realm this provider will provide user federation for.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Can be one of ONE_LEVEL or SUBTREE:
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	SearchScope *string `json:"searchScope,omitempty" tf:"search_scope,omitempty"`

	// When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
	// When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
	StartTLS *bool `json:"startTls,omitempty" tf:"start_tls,omitempty"`

	// When true, newly created users will be synced back to LDAP. Defaults to false.
	// When true, newly created users will be synced back to LDAP.
	SyncRegistrations *bool `json:"syncRegistrations,omitempty" tf:"sync_registrations,omitempty"`

	// If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
	// If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
	TrustEmail *bool `json:"trustEmail,omitempty" tf:"trust_email,omitempty"`

	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	UUIDLdapAttribute *string `json:"uuidLdapAttribute,omitempty" tf:"uuid_ldap_attribute,omitempty"`

	// When true, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
	// When `true`, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
	UsePasswordModifyExtendedOp *bool `json:"usePasswordModifyExtendedOp,omitempty" tf:"use_password_modify_extended_op,omitempty"`

	// Can be one of ALWAYS, ONLY_FOR_LDAPS, or NEVER:
	UseTruststoreSpi *string `json:"useTruststoreSpi,omitempty" tf:"use_truststore_spi,omitempty"`

	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	// All values of LDAP objectClass attribute for users in LDAP.
	UserObjectClasses []*string `json:"userObjectClasses,omitempty" tf:"user_object_classes,omitempty"`

	// Name of the LDAP attribute to use as the Keycloak username.
	// Name of the LDAP attribute to use as the Keycloak username.
	UsernameLdapAttribute *string `json:"usernameLdapAttribute,omitempty" tf:"username_ldap_attribute,omitempty"`

	// Full DN of LDAP tree where your users are.
	// Full DN of LDAP tree where your users are.
	UsersDn *string `json:"usersDn,omitempty" tf:"users_dn,omitempty"`

	// When true, Keycloak will validate passwords using the realm policy before updating it.
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	ValidatePasswordPolicy *bool `json:"validatePasswordPolicy,omitempty" tf:"validate_password_policy,omitempty"`

	// Can be one of OTHER, EDIRECTORY, AD, RHDS, or TIVOLI. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to OTHER.
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

type UserFederationParameters struct {

	// The number of users to sync within a single transaction. Defaults to 1000.
	// The number of users to sync within a single transaction.
	// +kubebuilder:validation:Optional
	BatchSizeForSync *float64 `json:"batchSizeForSync,omitempty" tf:"batch_size_for_sync,omitempty"`

	// Password of LDAP admin. This attribute must be set if bind_dn is set.
	// Password of LDAP admin.
	// +kubebuilder:validation:Optional
	BindCredentialSecretRef *v1.SecretKeySelector `json:"bindCredentialSecretRef,omitempty" tf:"-"`

	// DN of LDAP admin, which will be used by Keycloak to access LDAP server. This attribute must be set if bind_credential is set.
	// DN of LDAP admin, which will be used by Keycloak to access LDAP server.
	// +kubebuilder:validation:Optional
	BindDn *string `json:"bindDn,omitempty" tf:"bind_dn,omitempty"`

	// A block containing the cache settings.
	// Settings regarding cache policy for this realm.
	// +kubebuilder:validation:Optional
	Cache []CacheParameters `json:"cache,omitempty" tf:"cache,omitempty"`

	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	// How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users sync.
	// +kubebuilder:validation:Optional
	ChangedSyncPeriod *float64 `json:"changedSyncPeriod,omitempty" tf:"changed_sync_period,omitempty"`

	// LDAP connection timeout in the format of a Go duration string.
	// LDAP connection timeout (duration string)
	// +kubebuilder:validation:Optional
	ConnectionTimeout *string `json:"connectionTimeout,omitempty" tf:"connection_timeout,omitempty"`

	// Connection URL to the LDAP server.
	// Connection URL to the LDAP server.
	// +kubebuilder:validation:Optional
	ConnectionURL *string `json:"connectionUrl,omitempty" tf:"connection_url,omitempty"`

	// Additional LDAP filter for filtering searched users. Must begin with ( and end with ).
	// Additional LDAP filter for filtering searched users. Must begin with '(' and end with ')'.
	// +kubebuilder:validation:Optional
	CustomUserSearchFilter *string `json:"customUserSearchFilter,omitempty" tf:"custom_user_search_filter,omitempty"`

	// When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP user federation provider. Defaults to false.
	// When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP user federation provider.
	// +kubebuilder:validation:Optional
	DeleteDefaultMappers *bool `json:"deleteDefaultMappers,omitempty" tf:"delete_default_mappers,omitempty"`

	// Can be one of READ_ONLY, WRITABLE, or UNSYNCED. UNSYNCED allows user data to be imported but not synced back to LDAP. Defaults to READ_ONLY.
	// READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
	// +kubebuilder:validation:Optional
	EditMode *string `json:"editMode,omitempty" tf:"edit_mode,omitempty"`

	// When false, this provider will not be used when performing queries for users. Defaults to true.
	// When false, this provider will not be used when performing queries for users.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	// How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
	// +kubebuilder:validation:Optional
	FullSyncPeriod *float64 `json:"fullSyncPeriod,omitempty" tf:"full_sync_period,omitempty"`

	// When true, LDAP users will be imported into the Keycloak database. Defaults to true.
	// When true, LDAP users will be imported into the Keycloak database.
	// +kubebuilder:validation:Optional
	ImportEnabled *bool `json:"importEnabled,omitempty" tf:"import_enabled,omitempty"`

	// A block containing the kerberos settings.
	// Settings regarding kerberos authentication for this realm.
	// +kubebuilder:validation:Optional
	Kerberos []KerberosParameters `json:"kerberos,omitempty" tf:"kerberos,omitempty"`

	// Display name of the provider when displayed in the console.
	// Display name of the provider when displayed in the console.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// When true, Keycloak assumes the LDAP server supports pagination. Defaults to true.
	// When true, Keycloak assumes the LDAP server supports pagination.
	// +kubebuilder:validation:Optional
	Pagination *bool `json:"pagination,omitempty" tf:"pagination,omitempty"`

	// Priority of this provider when looking up users. Lower values are first. Defaults to 0.
	// Priority of this provider when looking up users. Lower values are first.
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Name of the LDAP attribute to use as the relative distinguished name.
	// Name of the LDAP attribute to use as the relative distinguished name.
	// +kubebuilder:validation:Optional
	RdnLdapAttribute *string `json:"rdnLdapAttribute,omitempty" tf:"rdn_ldap_attribute,omitempty"`

	// LDAP read timeout in the format of a Go duration string.
	// LDAP read timeout (duration string)
	// +kubebuilder:validation:Optional
	ReadTimeout *string `json:"readTimeout,omitempty" tf:"read_timeout,omitempty"`

	// The realm that this provider will provide user federation for.
	// The realm this provider will provide user federation for.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// Can be one of ONE_LEVEL or SUBTREE:
	// ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
	// +kubebuilder:validation:Optional
	SearchScope *string `json:"searchScope,omitempty" tf:"search_scope,omitempty"`

	// When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
	// When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
	// +kubebuilder:validation:Optional
	StartTLS *bool `json:"startTls,omitempty" tf:"start_tls,omitempty"`

	// When true, newly created users will be synced back to LDAP. Defaults to false.
	// When true, newly created users will be synced back to LDAP.
	// +kubebuilder:validation:Optional
	SyncRegistrations *bool `json:"syncRegistrations,omitempty" tf:"sync_registrations,omitempty"`

	// If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
	// If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
	// +kubebuilder:validation:Optional
	TrustEmail *bool `json:"trustEmail,omitempty" tf:"trust_email,omitempty"`

	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	// Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
	// +kubebuilder:validation:Optional
	UUIDLdapAttribute *string `json:"uuidLdapAttribute,omitempty" tf:"uuid_ldap_attribute,omitempty"`

	// When true, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
	// When `true`, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
	// +kubebuilder:validation:Optional
	UsePasswordModifyExtendedOp *bool `json:"usePasswordModifyExtendedOp,omitempty" tf:"use_password_modify_extended_op,omitempty"`

	// Can be one of ALWAYS, ONLY_FOR_LDAPS, or NEVER:
	// +kubebuilder:validation:Optional
	UseTruststoreSpi *string `json:"useTruststoreSpi,omitempty" tf:"use_truststore_spi,omitempty"`

	// Array of all values of LDAP objectClass attribute for users in LDAP. Must contain at least one.
	// All values of LDAP objectClass attribute for users in LDAP.
	// +kubebuilder:validation:Optional
	UserObjectClasses []*string `json:"userObjectClasses,omitempty" tf:"user_object_classes,omitempty"`

	// Name of the LDAP attribute to use as the Keycloak username.
	// Name of the LDAP attribute to use as the Keycloak username.
	// +kubebuilder:validation:Optional
	UsernameLdapAttribute *string `json:"usernameLdapAttribute,omitempty" tf:"username_ldap_attribute,omitempty"`

	// Full DN of LDAP tree where your users are.
	// Full DN of LDAP tree where your users are.
	// +kubebuilder:validation:Optional
	UsersDn *string `json:"usersDn,omitempty" tf:"users_dn,omitempty"`

	// When true, Keycloak will validate passwords using the realm policy before updating it.
	// When true, Keycloak will validate passwords using the realm policy before updating it.
	// +kubebuilder:validation:Optional
	ValidatePasswordPolicy *bool `json:"validatePasswordPolicy,omitempty" tf:"validate_password_policy,omitempty"`

	// Can be one of OTHER, EDIRECTORY, AD, RHDS, or TIVOLI. When this is selected in the GUI, it provides reasonable defaults for other fields. When used with the Keycloak API, this attribute does nothing, but is still required. Defaults to OTHER.
	// LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
	// +kubebuilder:validation:Optional
	Vendor *string `json:"vendor,omitempty" tf:"vendor,omitempty"`
}

// UserFederationSpec defines the desired state of UserFederation
type UserFederationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserFederationParameters `json:"forProvider"`
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
	InitProvider UserFederationInitParameters `json:"initProvider,omitempty"`
}

// UserFederationStatus defines the observed state of UserFederation.
type UserFederationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserFederationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// UserFederation is the Schema for the UserFederations API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type UserFederation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.connectionUrl) || (has(self.initProvider) && has(self.initProvider.connectionUrl))",message="spec.forProvider.connectionUrl is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.rdnLdapAttribute) || (has(self.initProvider) && has(self.initProvider.rdnLdapAttribute))",message="spec.forProvider.rdnLdapAttribute is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userObjectClasses) || (has(self.initProvider) && has(self.initProvider.userObjectClasses))",message="spec.forProvider.userObjectClasses is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.usernameLdapAttribute) || (has(self.initProvider) && has(self.initProvider.usernameLdapAttribute))",message="spec.forProvider.usernameLdapAttribute is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.usersDn) || (has(self.initProvider) && has(self.initProvider.usersDn))",message="spec.forProvider.usersDn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.uuidLdapAttribute) || (has(self.initProvider) && has(self.initProvider.uuidLdapAttribute))",message="spec.forProvider.uuidLdapAttribute is a required parameter"
	Spec   UserFederationSpec   `json:"spec"`
	Status UserFederationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserFederationList contains a list of UserFederations
type UserFederationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []UserFederation `json:"items"`
}

// Repository type metadata.
var (
	UserFederation_Kind             = "UserFederation"
	UserFederation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: UserFederation_Kind}.String()
	UserFederation_KindAPIVersion   = UserFederation_Kind + "." + CRDGroupVersion.String()
	UserFederation_GroupVersionKind = CRDGroupVersion.WithKind(UserFederation_Kind)
)

func init() {
	SchemeBuilder.Register(&UserFederation{}, &UserFederationList{})
}
