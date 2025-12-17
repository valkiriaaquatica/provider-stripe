package stripeentitlementsfeature

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_entitlements_feature", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeEntitlementsFeature"
		if cvc := ujconfig.GetSchema(r.TerraformResource, "cvc"); cvc != nil {
			cvc.Type = schema.TypeString
		}
	})
}
