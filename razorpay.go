package razorpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	//APIURL RazorPay Endpoint
	APIURL = "https://api.razorpay.com/v1/"
)

//RazorPay ...
type RazorPay struct {
	apikey     string //rzp_test_eq0OycgzlfwDIs
	apisecret  string //vt9YFHVw8EgMb2phcKz9SZjd
	merchantID string
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

//New creates new instance of client
func New(APIKey string, apiSecret string) *RazorPay {
	return &RazorPay{apikey: APIKey, apisecret: apiSecret}
}

/*
//GetPayments Retrieve multiple payments
func (r *RazorPay) GetPayments(from time.Time, to time.Time, count int, skip int) []Payment {
	paymentresp := new(Payment)
	resp, err := r.call("GetPayments", nil, urlparams)
	json.Unmarshal(resp, paymentresp)
	return paymentresp, nil
}



//CapturePayment Capture a payment
func (r *RazorPay) CapturePayment() (*Payment, error) {
	orderresp := new(Order)
	orderreq := order
	createreqjson, err := json.Marshal(&orderreq)
	resp, err := r.call("CreateOrder", createreqjson, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}

//Refund issue refund
func (r *RazorPay) Refund() (*Refund, error) {
	orderresp := new(Order)
	orderreq := order
	createreqjson, err := json.Marshal(&orderreq)
	resp, err := r.call("CreateOrder", createreqjson, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}

//GetRefundsByPaymentID Retrieve multiple refunds of a payment
func (r *RazorPay) GetRefundsByPaymentID(id string) (*Refunds, error) {
	orderresp := new(Order)
	orderreq := order
	createreqjson, err := json.Marshal(&orderreq)
	resp, err := r.call("CreateOrder", createreqjson, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}

func (r *RazorPay) GetRefundByPaymentID(id string) (*Refund, error) {
	orderresp := new(Order)
	orderreq := order
	createreqjson, err := json.Marshal(&orderreq)
	resp, err := r.call("CreateOrder", createreqjson, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}

func (r *RazorPay) GetRefunds() (*Refunds, error) {
	orderresp := new(Order)
	orderreq := order
	createreqjson, err := json.Marshal(&orderreq)
	resp, err := r.call("CreateOrder", createreqjson, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}

func (r *RazorPay) GetRefundByID(id string) (*Refund, error) {
	orderresp := new(Order)
	orderreq := order
	createreqjson, err := json.Marshal(&orderreq)
	resp, err := r.call("CreateOrder", createreqjson, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}


func (r *RazorPay) GetOrders() (*Order, error) {
	orderresp := new(Order)
	resp, err := r.call("GetOrders", nil, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}

func (r *RazorPay) GetOrderByID() (*Order, error) {
	orderresp := new(Order)
	resp, err := r.call("GetOrderByID", nil, nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, nil
}
*/

type NewOrder struct {
	Amount         int    `json:"amount"`
	Currency       string `json:"currency"`
	Receipt        string `json:"receipt"`
	PaymentCapture bool   `json:"payment_capture"`
	Notes          string `json:"notes"`
}

//CreateOrder Create an order
func (r *RazorPay) CreateOrder(order NewOrder) (*Order, error) {
	orderresp := new(Order)
	orderreq := order
	createreqjson, err := json.Marshal(&orderreq)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := r.call("CreateOrder", createreqjson, "", nil)
	json.Unmarshal(resp, orderresp)
	return orderresp, err
}

//GetPaymentByID Retrieve a Payment by ID
func (r *RazorPay) GetPaymentByID(id string) (*Payment, error) {
	paymentresp := new(Payment)
	resp, err := r.call("GetPaymentByID", nil, id, nil)
	json.Unmarshal(resp, paymentresp)
	return paymentresp, err
}

func (r *RazorPay) call(operation string, reqbody []byte, pathparams string, queryparams map[string]string) ([]byte, error) {
	var rurl string
	var rmethod string
	switch operation {
	case "CreateOrder":
		rmethod = "POST"
		rurl = APIURL + "orders"
	case "GetPaymentByID":
		rmethod = "GET"
		rurl = APIURL + "payments/" + pathparams
	default:
		err := errors.New("Invalid Method/Operation")
		return nil, err
	}
	log.Debugln("[Razopay] Request URL:", rurl)
	log.Debugln("[Razopay] Path Params:", pathparams)
	log.Debugln("[Razopay] Query Params:", queryparams)
	log.Debugln("[Razopay] Request Body:", string(reqbody))
	req, err := http.NewRequest(rmethod, rurl, bytes.NewBuffer(reqbody))
	req.SetBasicAuth(r.apikey, r.apisecret)
	req.Header.Add("Content-Type", "application/json")
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
