package stripeprice

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_price resource.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_price", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripePrice"
	})
}
