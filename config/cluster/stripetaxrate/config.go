package stripetaxrate

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_tax_rate resource.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_tax_rate", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeTaxRate"
	})
}
