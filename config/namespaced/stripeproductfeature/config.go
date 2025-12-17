package stripeproductfeature

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_product_feature resource for namespaced scope.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_product_feature", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeProductFeature"
		r.Kind = "ProductFeature"
	})
}
