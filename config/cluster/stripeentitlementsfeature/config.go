package stripeentitlementsfeature

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_entitlements_feature", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeEntitlementsFeature"
		r.Kind = "EntitlementsFeature"
	})
}
