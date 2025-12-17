package stripemeter

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_meter resource.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_meter", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeMeter"
	})
}
