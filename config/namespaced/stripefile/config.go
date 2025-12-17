package stripefile

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_file resource for namespaced scope.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_file", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeFile"
	})
}
