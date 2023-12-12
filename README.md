# Cashfree PG Go SDK
![GitHub](https://img.shields.io/github/license/cashfree/cashfree-pg) ![Discord](https://img.shields.io/discord/931125665669972018?label=discord) ![GitHub last commit (branch)](https://img.shields.io/github/last-commit/cashfree/cashfree-pg/main) ![GitHub release (with filter)](https://img.shields.io/github/v/release/cashfree/cashfree-pg?label=latest) ![GitHub forks](https://img.shields.io/github/forks/cashfree/cashfree-pg) [![Coverage Status](https://coveralls.io/repos/github/cashfree/cashfree-pg/badge.svg?branch=main)](https://coveralls.io/github/cashfree/cashfree-pg?branch=main)

The Cashfree PG Go SDK offers a convenient solution to access [Cashfree PG APIs](https://docs.cashfree.com/reference/pg-new-apis-endpoint) from a server-side Go  applications. 



## Documentation

Cashfree's PG API Documentation - https://docs.cashfree.com/reference/pg-new-apis-endpoint

Learn and understand payment gateway workflows at Cashfree Payments [here](https://docs.cashfree.com/docs/payment-gateway)

## Getting Started

### Installation
```bash
go get github.com/cashfree/cashfree-pg/v3
```
### Configuration

```go 
import (
    cashfree "github.com/cashfree/cashfree-pg/v3"
)

clientId := "<x-client-id>"
clientSecret := "<x-client-secret>"
cashfree.XClientId = &clientId
cashfree.XClientSecret = &clientSecret
cashfree.XEnvironment = cashfree.SANDBOX
```

Generate your API keys (x-client-id , x-client-secret) from [Cashfree Merchant Dashboard](https://merchant.cashfree.com/merchants/login)

### Basic Usage
Create Order
```go
version := "2022-09-01"
response, httpResponse, err := cashfree.PGCreateOrder(&version, &request, nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(response)
}
```

Get Order
```go
version := "2022-09-01"
response, httpResponse, err := cashfree.PGFetchOrder(&version, "<order_id>", nil, nil, nil)
if err != nil {
	fmt.Println(err.Error())
} else {
	fmt.Println(httpResponse.StatusCode)
	fmt.Println(response)
}
```

## Licence

Apache Licensed. See [LICENSE.md](LICENSE.md) for more details
