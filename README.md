# Credit Card
Credit Card structure, helpers and validation.

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

// Luhn checksum a number. https://en.wikipedia.org/wiki/Luhn_algorithm
if yes := creditcard.Luhn("1234567891234567"); !yes {
    log.Println("Luhn validation failed.")
}

// Encrypt the CreditCard.Number and get the bytes blob.
b := card.Encrypt("mySuperSecureSalt")

// Common helpers.
fmt.Println(card.First6()) // 123456
fmt.Println(card.Last4()) // 4567
fmt.Println(card.Issuer()) // other, visa, amex, ...
fmt.Println(card.ToJSON()) // {"name":"Leonardo Faoro","number":"1234567891234567","cvv_2":"123","expiry":"06/2019"}
```

## Maxmind's minFraud (https://www.maxmind.com/en/minfraud-services)
minFraud uses a combination of data points across over a billion transactions to provide an accurate risk score for every credit or debit card payment you process. Several variables go into the assessment and scoring, including geolocation, IIN information, reputation checks, IP checks, and device tracking.
### Benefits
- Significantly reduce or eliminate the number of fraudulent transactions your business processes.
- Significantly reduce or eliminate chargebacks from customers (e.g. if they have been the victim of identity theft and their card has been used in a criminal way).
### Quick start
```go
riskScore, err := maxmind.RiskCheck(card.First6(), card.Last4(), "8.8.8.8", "test@test.com")
if err != nil {
    log.Println("riskCheck error: ", err)
}
```



## TODO
- ~~add helpers for first6, last4, issuer and encryption/decryption~~
- ~~add issuer identification~~
- add fraud check patterns using maxMind API
