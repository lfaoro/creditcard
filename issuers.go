// Copyright 2018 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package creditcard

import "regexp"

// Issuers list.
const (
	Visa       = "visa"
	MasterCard = "mastercard"
	Amex       = "american express"
	Diners     = "diners"
	Discover   = "discover"
	JCB        = "jcb"
	Other      = "other"
)

// Issuer attempts to identify the Credit Card issuer by recognizing
// their numeric patterns.
func (card *CreditCard) Issuer() string {
	regVisa, _ := regexp.Compile(`^4[0-9]{12}(?:[0-9]{3})?$`)
	regMaster, _ := regexp.Compile(`^5[1-5][0-9]{14}$`)
	regAmex, _ := regexp.Compile(`^3[47][0-9]{13}$`)
	regDiners, _ := regexp.Compile(`^3(?:0[0-5]|[68][0-9])[0-9]{11}$`)
	regDiscover, _ := regexp.Compile(`^6(?:011|5[0-9]{2})[0-9]{12}$`)
	regJCB, _ := regexp.Compile(`^(?:2131|1800|35\d{3})\d{11}$`)
	reg := map[string]interface{}{
		Visa:       regVisa,
		MasterCard: regMaster,
		Amex:       regAmex,
		Diners:     regDiners,
		Discover:   regDiscover,
		JCB:        regJCB,
	}
	for t, r := range reg {
		if r.(*regexp.Regexp).MatchString(card.Number) {
			return t
		}
	}
	return Other
}
