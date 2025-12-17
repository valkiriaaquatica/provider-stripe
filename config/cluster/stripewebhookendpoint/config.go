package stripewebhookendpoint

import ujconfig "github.com/crossplane/upjet/v2/pkg/config"

// Configure configures the stripe_webhook_endpoint resource.
func Configure(p *ujconfig.Provider) {
	p.AddResourceConfigurator("stripe_webhook_endpoint", func(r *ujconfig.Resource) {
		r.ShortGroup = "StripeWebhookEndpoint"
		r.Kind = "WebhookEndpoint"
	})
}
