// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	providerconfig "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/providerconfig"
	coupon "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripe/coupon"
	customer "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripe/customer"
	card "github.com/valkiriaaquatica/provider-stripe/internal/controller/namespaced/stripecard/card"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		providerconfig.Setup,
		coupon.Setup,
		customer.Setup,
		card.Setup,
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
		coupon.SetupGated,
		customer.SetupGated,
		card.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
