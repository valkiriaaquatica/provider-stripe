package stripecard

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Configure configures resources for the virtual environment group
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_customer", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeCustomer"

		// The provider marks CVC as sensitive but defines it as a number,
		// which upjet cannot map to a Secret. Treat it as a string to
		// generate the CRD successfully.
		if cvc := ujconfig.GetSchema(r.TerraformResource, "cvc"); cvc != nil {
			cvc.Type = schema.TypeString
		}
	})
}
