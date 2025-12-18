// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/providerconfig"
	card "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripecard/card"
	coupon "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripecoupon/coupon"
	customer "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripecustomer/customer"
	entitlementsfeature "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripeentitlementsfeature/entitlementsfeature"
	file "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripefile/file"
	meter "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripemeter/meter"
	configuration "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripeportalconfiguration/configuration"
	price "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripeprice/price"
	product "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripeproduct/product"
	productfeature "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripeproductfeature/productfeature"
	promotioncode "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripepromotioncode/promotioncode"
	shippingrate "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripeshippingrate/shippingrate"
	rate "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripetaxrate/rate"
	webhookendpoint "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripewebhookendpoint/webhookendpoint"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		card.Setup,
		coupon.Setup,
		customer.Setup,
		entitlementsfeature.Setup,
		file.Setup,
		meter.Setup,
		configuration.Setup,
		price.Setup,
		product.Setup,
		productfeature.Setup,
		promotioncode.Setup,
		shippingrate.Setup,
		rate.Setup,
		webhookendpoint.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.SetupGated,
		card.SetupGated,
		coupon.SetupGated,
		customer.SetupGated,
		entitlementsfeature.SetupGated,
		file.SetupGated,
		meter.SetupGated,
		configuration.SetupGated,
		price.SetupGated,
		product.SetupGated,
		productfeature.SetupGated,
		promotioncode.SetupGated,
		shippingrate.SetupGated,
		rate.SetupGated,
		webhookendpoint.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
