package creditcard

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// Validate implements various checks to ensure a credit card is valid.
func (card *CreditCard) Validate() error {
	const shortDate = "01/2006"
	t, err := time.Parse(shortDate, card.Expiry)
	if err != nil {
		return fmt.Errorf("Credit Card date invalid %v", t)
	}
	// Validate: card expired
	if t.Year() < time.Now().Year() {
		return fmt.Errorf("Credit Card expired: %v", card.Expiry)
	}
	// Validate: parsing to integer for number and cvv2
	if _, err := strconv.Atoi(card.Number); err != nil {
		return fmt.Errorf("Credit Card number invalid: %v", card.Number)
	}
	if _, err := strconv.Atoi(card.CVV2); err != nil {
		return fmt.Errorf("Credit Card CVV2 invalid: %v", card.CVV2)
	}
	// Validate: cvv2
	if len(card.CVV2) < 3 || len(card.CVV2) > 4 {
		return fmt.Errorf("Credit Card CVV2 length invalid: %v", card.CVV2)
	}
	// Validate: card length
	if len(card.Number) < 13 || len(card.Number) > 19 {
		return fmt.Errorf("Credit Card number length invalid: %v", card.Number)
	}
	// Validate: test card
	switch card.Number {
	case "4242424242424242",
		"4012888888881881",
		"4000056655665556",
		"5555555555554444",
		"5200828282828210",
		"5105105105105100",
		"378282246310005",
		"371449635398431",
		"6011111111111117",
		"6011000990139424",
		"30569309025904",
		"38520000023237",
		"3530111333300000",
		"3566002020360505":
		return fmt.Errorf("Credit Card is a test card: %v", card.Number)
	}
	// Validate: Luhn algorithm
	yes := Luhn(card.Number)
	if !yes {
		return fmt.Errorf("Credit Card number failed the Luhn algorithm check: %v", card.Number)
	}
	return nil
}

// Luhn implements the Luhn checksum formula that validates
// identification numbers. It was designed to protect against accidental
// errors, not malicious attacks.
//
// https://en.wikipedia.org/wiki/Luhn_algorithm
//
// TODO:
// - refactor
// - benchmark
//
func Luhn(number string) bool {
	sum := 0
	len := len(number)
	flip := false
	for i := len - 1; i > -1; i-- {
		num, err := strconv.Atoi(string(number[i]))
		if err != nil {
			log.Fatal(err)
		}
		flip = !flip
		if flip {
			num = num * 2
			if num > 9 {
				sum += (num)%10 + 1
				continue
			}
			sum += num
			continue
		}
		sum += num
	}
	n := strconv.Itoa(sum * 9)
	x := string(n[2])
	i, _ := strconv.Atoi(x)
	sum += i
	return sum%10 == 0
}
