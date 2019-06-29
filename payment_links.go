package razorpay

import (
	"encoding/json"
	"log"
)

//CreatePaymentLink Create a new payment link
func (r *RazorPay) CreatePaymentLink(paylink PaymentLink) (*PaymentLinkResponse, error) {
	paylinkresp := new(PaymentLinkResponse) 
	createreqjson, err := json.Marshal(&paylink)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := r.call("CreatePaymentLink", createreqjson, "", nil)
	json.Unmarshal(resp, paylinkresp)
	return paylinkresp, err
}

//GetPaymentLink Get payment link by id
func (r *RazorPay) GetPaymentLink(id string) (*PaymentLinkResponse, error) {
	getidresp := new(PaymentLinkResponse)
	resp, err := r.call("GetPaymentLink", nil, id, nil)
	json.Unmarshal(resp, getidresp)
	return getidresp, err
}

//SendPaymentLink Send/Resend payment link notification to the user 
func (r *RazorPay) SendPaymentLink(id string, medium string) (bool, error) {
	sendresp := new(Response)
	id = id+"/notify_by/"+medium
	resp, err := r.call("SendPaymentLink", nil, id, nil)
	json.Unmarshal(resp, sendresp)
	return sendresp.Success, err
}

//CancelPaymentLink Cancel given link id 
func (r *RazorPay) CancelPaymentLink(id string) (string, error) {
	paylinkresp := new(PaymentLinkResponse) 
	resp, err := r.call("CancelPaymentLink", nil, id + "/cancel", nil)
	json.Unmarshal(resp, paylinkresp)
	return paylinkresp.Status, err

}