package stripeshippingrate

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_shipping_rate resource for namespaced scope.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_shipping_rate", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeShippingRate"
		r.Kind = "ShippingRate"
	})
}
