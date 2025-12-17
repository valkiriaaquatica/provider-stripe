package stripepromotioncode

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_promotion_code resource for namespaced scope.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_promotion_code", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripePromotionCode"
		r.Kind = "PromotionCode"
	})
}
