// Copyright 2018 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package creditcard

// CreditCard stores credit card information.
type CreditCard struct {
	Name   string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
	CVV2   string `json:"cvv_2,omitempty"`
	Expiry string `json:"expiry,omitempty"`
}

// New creates a new CreditCard instance.
func New(name, number, cvv2, expiry string) *CreditCard {
	card := &CreditCard{
		Name:   name,
		Number: number,
		CVV2:   cvv2,
		Expiry: expiry,
	}
	return card
}

// NewValidate creates a new CreditCard instance and validates it
// using all the checks implemented so far.
func NewValidate(name, number, cvv2, expiry string) (*CreditCard, error) {
	card := &CreditCard{
		Name:   name,
		Number: number,
		CVV2:   cvv2,
		Expiry: expiry,
	}
	if err := card.Validate(); err != nil {
		return nil, err
	}
	return card, nil
}
