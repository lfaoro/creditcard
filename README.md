# creditcard
Credit Card structure and validation.

## Installation
```bash
go get github.com/lfaoro/creditcard
```

## Quick start
```go
// Create a new CreditCard instance.
card := creditcard.New("Leonardo Faoro", "1234567891234567", "123", "06/2019")

// Create a new CreditCard instance and validate it
// using all the checks implemented so far.
card, err := creditcard.NewValidate("Leonardo Faoro", "1234567891234567", "123", "06/2019")
if err != nil {
    log.Fatal(err)
}

// Validate a CreditCard using all the checks implemented so far.
if err := card.Validate(); err != nil {
    log.Fatal(err)
}

// Luhn checksum a number.
if yes := creditcard.Luhn("1234567891234567"); !yes {
    log.Println("Luhn validation failed.")
}

b := card.Encrypt("mySuperSecureSalt")

// Common helpers.
fmt.Println(card.First6()) // 123456
fmt.Println(card.Last4()) // 4567
fmt.Println(card.Issuer()) // other, visa, amex, ...
```

## TODO
- ~~add helpers for first6, last4, issuer and encryption/decryption~~
- ~~add issuer identification~~
- add fraud check patterns using maxMind API
