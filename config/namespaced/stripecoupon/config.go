package stripecoupon

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_coupon", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeCoupon"
	})
}
