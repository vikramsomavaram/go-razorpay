package razorpay

import (
	"bytes"
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
	OrderID        string      `json:"order_id"`
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

//PaymentLink payment link request json data structure
type PaymentLink struct {
	Customer 	Customer 	  `json:"customer"`
	Type        string        `json:"type"`
	ViewLess    int           `json:"view_less"`
	Amount      int           `json:"amount"`
	Currency    string        `json:"currency"`
	Description string        `json:"description"`
	ExpireBy    int           `json:"expire_by,omitempty"`
	CustomerID 	string        `json:"customer_id,omitempty"`
	SmsNotify	string        `json:"sms_notify,omitempty"`
	EmailNotify	string        `json:"email_notify,omitempty"`
	Date		int           `json:"date,omitempty"`
	Terms		string        `json:"terms,omitempty"`
	Notes		[]interface{} `json:"notes,omitempty"`
}

type Customer struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Contact string `json:"contact"`
}

//PaymentLinkResponse link response json data structure
type PaymentLinkResponse struct {
	ID              string      `json:"id"`
	Entity          string      `json:"entity"`
	Receipt         interface{} `json:"receipt"`
	CustomerID      string      `json:"customer_id"`
	CustomerDetails struct {
		Name           string      `json:"name"`
		Email          string      `json:"email"`
		Contact        string      `json:"contact"`
		BillingAddress interface{} `json:"billing_address"`
	} `json:"customer_details"`
	OrderID             string        `json:"order_id"`
	LineItems           []interface{} `json:"line_items"`
	PaymentID           interface{}   `json:"payment_id"`
	Status              string        `json:"status"`
	ExpireBy            interface{}   `json:"expire_by"`
	IssuedAt            int           `json:"issued_at"`
	PaidAt              interface{}   `json:"paid_at"`
	CancelledAt         interface{}   `json:"cancelled_at"`
	ExpiredAt           interface{}   `json:"expired_at"`
	SmsStatus           string        `json:"sms_status"`
	EmailStatus         string        `json:"email_status"`
	Date                int           `json:"date"`
	Terms               interface{}   `json:"terms"`
	PartialPayment      bool          `json:"partial_payment"`
	GrossAmount         int           `json:"gross_amount"`
	TaxAmount           int           `json:"tax_amount"`
	Amount              int           `json:"amount"`
	AmountPaid          int           `json:"amount_paid"`
	AmountDue           int           `json:"amount_due"`
	Currency            string        `json:"currency"`
	Description         string        `json:"description"`
	Notes               []interface{} `json:"notes"`
	Comment             interface{}   `json:"comment"`
	ShortURL            string        `json:"short_url"`
	ViewLess            bool          `json:"view_less"`
	BillingStart        interface{}   `json:"billing_start"`
	BillingEnd          interface{}   `json:"billing_end"`
	Type                string        `json:"type"`
	GroupTaxesDiscounts bool          `json:"group_taxes_discounts"`
	UserID              interface{}   `json:"user_id"`
	CreatedAt           int           `json:"created_at"`
}

type Response struct{
	Success bool `json:"success"`
}

type NewOrder struct {
	Amount         int               `json:"amount"`
	Currency       string            `json:"currency"`
	Receipt        string            `json:"receipt"`
	PaymentCapture bool              `json:"payment_capture"`
	Notes          map[string]string `json:"notes"`
}

type Data struct {
	Amount	int	`json:"amount"`
}



func (r *RazorPay) call(operation string, reqbody []byte, pathparams string, queryparams map[string]string) ([]byte, error) {
	var rurl string
	var rmethod string
	switch operation {
	case "CreateOrder":
		rmethod = "POST"
		rurl = APIURL + "orders"
	case "GetOrders":
		rmethod = "GET"
		rurl = APIURL + "orders"
	case "CreatePaymentLink":
		rmethod = "POST"
		rurl = APIURL + "invoices"
	case "GetPaymentLink":
		rmethod = "GET"
		rurl = APIURL + "invoices/" + pathparams
	case "SendPaymentLink":
		rmethod = "POST"
		rurl = APIURL + "invoices/" + pathparams
	case "CancelPaymentLink":
		rmethod = "POST"
		rurl = APIURL + "invoices/" + pathparams
	case "CapturePayment":
		rmethod = "POST"
		rurl = APIURL + "payments/" + pathparams
	case "GetPaymentByID":
		rmethod = "GET"
		rurl = APIURL + "payment/" + pathparams
	case "GetPayments":
		rmethod = "GET"
		rurl = APIURL + "payments"
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
