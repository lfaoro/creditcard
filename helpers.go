// Copyright 2018 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package creditcard

import (
	"crypto/sha512"
	"encoding/json"
	"io"
	"log"
)

// First6 returns a string with the first 6 digits
// of the CreditCard.Number.
func (card *CreditCard) First6() string {
	return card.Number[0:6]
}

// Last4 returns a string with the last 4 digits
// of the CreditCard.Number.
func (card *CreditCard) Last4() string {
	return card.Number[len(card.Number)-4 : len(card.Number)]
}

// Encrypt the Credit Card number using the industry standard
// and returns a salted SHA512 hash.
func (card *CreditCard) Encrypt(salt string) []byte {
	h := sha512.New()
	// It is very, very unlikely to have such error, decided to omit
	// adding complexity to method.
	_, err := io.WriteString(h, card.Number+salt)
	if err != nil {
		log.Println("Unable to write encrypted bytes: ", err)
	}
	return h.Sum(nil)
}

// Decrypt the provided data and match it against the creditcard.Number.
// NOT IMPLEMENTED
func (card *CreditCard) Decrypt(data []byte, salt string) bool {
	return true
}

// ToJSON returns a string with the CreditCard json marshalled data.
func (card *CreditCard) ToJSON() string {
	b, err := json.Marshal(card)
	// Data is validated, very unlikely that this error may happen
	// for simplicity, I've decided to omit returning it.
	if err != nil {
		log.Println("Unable to Marshall to JSON this CreditCard", err)
	}
	return string(b)
}

// FromJSON returns a CreditCard from a JSON data bytes.
//
// JSON data example:
// {"name":"Leonardo Faoro","number":"1234567891234567","cvv_2":"123","expiry":"06/2019"}
//
func FromJSON(data []byte) (card *CreditCard, err error) {
	if err := json.Unmarshal(data, card); err != nil {
		return nil, err
	}
	return card, nil
}
