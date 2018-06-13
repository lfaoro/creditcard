package creditcard

// Last4 returns a string with the last 4 digits
// of the CreditCard.Number.
func (card *CreditCard) Last4() string {
	return card.Number[len(card.Number)-4 : len(card.Number)]
}

// First6 returns a string with the first 6 digits
// of the CreditCard.Number.
func (card *CreditCard) First6() string {
	return card.Number[0:6]
}
