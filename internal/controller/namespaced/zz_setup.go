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
	feature "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripeentitlementsfeature/feature"
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
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
