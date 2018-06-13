# creditcard
Credit Card structure and validation.

## Installation
```bash
go get github.com/lfaoro/creditcard
```

## Quick start
```go
card, err := creditcard.New("Leonardo Faoro", "1234567891234567", "123", "06/2019")
if err != nil {
    log.Fatal(err)
}
```

## TODO
- add issuer identification
- add fraud check patterns
