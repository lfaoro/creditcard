package creditcard

import (
	"crypto/sha512"
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
