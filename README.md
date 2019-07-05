# RazorPay Golang SDK

The Razorpay Payment Gateway enables you to accept payments via debit card, credit card, netbanking (supports 3D Secure), UPI or through any of our supported wallets.

Currently only supports few operations like payments, orders and refunds and rest are easy to implement (PR's are welcome).

I know that this library could be structed / designed in a better way but the current design serves the purpose of our internal projects.

## Installation

`go get -u github.com/itsbalamurali/go-razorpay`

With go modules add `github.com/itsbalamurali/go-razorpay v0.1.1-alpha` to your project `go.mod` file

## Usage

```go
rzp := razorpay.New("APIKey", "apiSecret")
paymentDetails := rzp.GetPaymentByID("pay_khadf87343hjkheqw")
```

## Development

Pull requests are welcome!

## License

The Razorpay Golang SDK is released under the MIT License.
