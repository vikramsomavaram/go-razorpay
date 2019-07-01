package razorpay

import (
	"encoding/json"
	"strconv"
	"time"
)

//CapturePayment Capture a payment
func (r *RazorPay) CapturePayment(id string, data Data) (*Payment, error) {
	paymentresp := new(Payment)
	createreqjson, err := json.Marshal(data)
	if err != nil{
		return nil, err
	}
	resp, err := r.call("CapturePayment", createreqjson, id+"/capture", nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, paymentresp)
	if err != nil{
		return nil, err
	}
	return paymentresp, err
}

//GetPaymentByID Fetch a Payment by ID
func (r *RazorPay) GetPaymentByID(id string) (*Payment, error) {
	paymentresp := new(Payment)
	resp, err := r.call("GetPaymentByID", nil, id, nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, paymentresp)
	if err != nil{
		return nil, err
	}
	return paymentresp, err
}

//GetPayments Fetch multiple payments
func (r *RazorPay) GetPayments(from time.Time, to time.Time, count int, skip int) (*[]Payments, error) {
	paymentresp := new([]Payments)
	queryparams := make(map[string]string)
	queryparams["from"]= strconv.Itoa(int(from.Unix()))
	queryparams["to"]= strconv.Itoa(int(to.Unix()))
	queryparams["count"]= strconv.Itoa(count)
	queryparams["skip"]= strconv.Itoa(skip)
	resp, err := r.call("GetPayments", nil, "", queryparams)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, paymentresp)
	if err != nil{
		return nil, err
	}
	return paymentresp, err
}
