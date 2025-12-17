package stripeproduct

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_product resource for namespaced scope.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_product", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeProduct"
	})
}
