package stripeportalconfiguration

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_portal_configuration resource for namespaced scope.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_portal_configuration", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripePortalConfiguration"
	})
}
