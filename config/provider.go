package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	stripecard "github.com/valkiriaaquatica/provider-stripe/config/cluster/stripecard"
	stripecoupon "github.com/valkiriaaquatica/provider-stripe/config/cluster/stripecoupon"
	stripecustomer "github.com/valkiriaaquatica/provider-stripe/config/cluster/stripecustomer"
	stripeentitlementsfeature "github.com/valkiriaaquatica/provider-stripe/config/cluster/stripeentitlementsfeature"

	namespacedstripecard "github.com/valkiriaaquatica/provider-stripe/config/namespaced/stripecard"
	namespacedstripecoupon "github.com/valkiriaaquatica/provider-stripe/config/namespaced/stripecoupon"
	namespacedstripecustomer "github.com/valkiriaaquatica/provider-stripe/config/namespaced/stripecustomer"
	namespacedstripeentitlementsfeature "github.com/valkiriaaquatica/provider-stripe/config/namespaced/stripeentitlementsfeature"
)

const (
	resourcePrefix = "stripe"
	modulePath     = "github.com/valkiriaaquatica/provider-stripe"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stripe.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		stripecard.Configure,
		stripecustomer.Configure,
		stripecoupon.Configure,
		stripeentitlementsfeature.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("stripe.m.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		namespacedstripecard.Configure,
		namespacedstripecustomer.Configure,
		namespacedstripecoupon.Configure,
		namespacedstripeentitlementsfeature.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
