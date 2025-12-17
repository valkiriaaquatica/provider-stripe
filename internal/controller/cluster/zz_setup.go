// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/providerconfig"
	card "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripecard/card"
	coupon "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripecoupon/coupon"
	customer "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripecustomer/customer"
	feature "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripeentitlementsfeature/feature"
	file "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripefile/file"
	meter "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripemeter/meter"
	configuration "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripeportalconfiguration/configuration"
	price "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripeprice/price"
	product "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripeproduct/product"
	featurestripeproductfeature "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripeproductfeature/feature"
	code "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripepromotioncode/code"
	rate "github.com/valkiriaaquatica/provider-stripe/internal/controller/cluster/stripeshippingrate/rate"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		card.Setup,
		coupon.Setup,
		customer.Setup,
		feature.Setup,
		file.Setup,
		meter.Setup,
		configuration.Setup,
		price.Setup,
		product.Setup,
		featurestripeproductfeature.Setup,
		code.Setup,
		rate.Setup,
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
		feature.SetupGated,
		file.SetupGated,
		meter.SetupGated,
		configuration.SetupGated,
		price.SetupGated,
		product.SetupGated,
		featurestripeproductfeature.SetupGated,
		code.SetupGated,
		rate.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
