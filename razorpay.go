package razorpay

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	//APIURL RazorPay Endpoint
	APIURL = "https://api.razorpay.com/v1/"
)

//RazorPay ...
type RazorPay struct {
	apikey    string
	apisecret string
}

//Payments ...
type Payments struct {
	Count  int    `json:"count"`
	Entity string `json:"entity"`
	Items  []Payment
}

//Payment ...
type Payment struct {
	ID             string      `json:"id"`
	Entity         string      `json:"entity"`
	Amount         int         `json:"amount"`
	Currency       string      `json:"currency"`
	Status         string      `json:"status"`
	OrderID        interface{} `json:"order_id"`
	InvoiceID      interface{} `json:"invoice_id"`
	International  bool        `json:"international"`
	Method         string      `json:"method"`
	AmountRefunded int         `json:"amount_refunded"`
	RefundStatus   interface{} `json:"refund_status"`
	Captured       bool        `json:"captured"`
	Description    string      `json:"description"`
	CardID         interface{} `json:"card_id"`
	Bank           interface{} `json:"bank"`
	Wallet         string      `json:"wallet"`
	Vpa            interface{} `json:"vpa"`
	Email          string      `json:"email"`
	Contact        string      `json:"contact"`
	Notes          struct {
		MerchantOrderID string `json:"merchant_order_id"`
	} `json:"notes"`
	Fee              int         `json:"fee"`
	Tax              int         `json:"tax"`
	ErrorCode        interface{} `json:"error_code"`
	ErrorDescription interface{} `json:"error_description"`
	CreatedAt        int         `json:"created_at"`
}

//Refund ...
type Refund struct {
	ID        string `json:"id"`
	Entity    string `json:"entity"`
	Amount    int    `json:"amount"`
	Currency  string `json:"currency"`
	PaymentID string `json:"payment_id"`
	Notes     struct {
	} `json:"notes"`
	CreatedAt int `json:"created_at"`
}

//Refunds ...
type Refunds struct {
	Count  int    `json:"count"`
	Entity string `json:"entity"`
	Items  []struct {
		ID        string `json:"id"`
		Entity    string `json:"entity"`
		Amount    int    `json:"amount"`
		Currency  string `json:"currency"`
		PaymentID string `json:"payment_id"`
		Notes     struct {
		} `json:"notes"`
		CreatedAt int `json:"created_at"`
	} `json:"items"`
}

type Order struct {
	ID        string        `json:"id"`
	Entity    string        `json:"entity"`
	Amount    int           `json:"amount"`
	Currency  string        `json:"currency"`
	Receipt   string        `json:"receipt"`
	Status    string        `json:"status"`
	Attempts  int           `json:"attempts"`
	Notes     []interface{} `json:"notes"`
	CreatedAt int           `json:"created_at"`
}

type Orders struct {
	Entity string  `json:"entity"`
	Count  int     `json:"count"`
	Items  []Order `json:"items"`
}

//NewRazorPay creates new instance of client
func NewRazorPay(APIKey string, apisecret string) *RazorPay {
	return &RazorPay{apikey: APIKey}
}

//GetPayments Retrieve multiple payments
func (r *RazorPay) GetPayments(from time.Time, to time.Time, count int, skip int) []Payment {
	return payments
}

//GetPaymentByID Retrieve a Payment by ID
func (r *RazorPay) GetPaymentByID(ID string) {

}

//CapturePayment Capture a payment
func (r *RazorPay) CapturePayment() {

}

//Refund issue refund
func (r *RazorPay) Refund() {

}

//GetRefundsByPaymentID Retrieve multiple refunds of a payment
func (r *RazorPay) GetRefundsByPaymentID() {

}

func (r *RazorPay) GetRefundByPaymentID() {

}

func (r *RazorPay) GetRefunds() {

}

func (r *RazorPay) GetRefundByID() {

}

func (r *RazorPay) CreateOrder() {

}

func (r *RazorPay) GetOrders() {

}

func (r *RazorPay) GetOrderByID() {

}

func (r *RazorPay) call(operation string, reqbody []byte) ([]byte, error) {
	var rurl string
	switch operation {
	case "CreateOrder":
		rurl = APIURL + "orders"
	default:
		err := errors.New("Invalid Operation, Must one of OTP,AUTH,KYC or BFD")
		return nil, err
	}
	log.Debugln("[Razopay] Request URL:", rurl)
	log.Debug("[Razopay] Request Body:", string(reqbody))
	req, err := http.NewRequest("POST", rurl, bytes.NewBuffer(reqbody))
	req.Header.Add("Content-Type", "text/xml")
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	defer resp.Body.Close()
	log.Debugln("[Razopay] Response Status:", resp.Status)
	log.Debugln("[Razopay] Response Headers:", resp.Header)
	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	log.Debug("[Razopay] Response Body:", string(respbody))
	return respbody, nil
}
