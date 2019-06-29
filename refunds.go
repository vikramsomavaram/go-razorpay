package razorpay

import "encoding/json"

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