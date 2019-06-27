package razorpay

//CapturePayment Capture a payment
func (r *RazorPay) CapturePayment(id string) (*Payment, error) {
	paymentresp := new(Payment)
	resp, err := r.call("CreatePayment", nil, id + "/capture", nil)
	json.Unmarshal(resp, paymentresp)
	return paymentresp, nil
}

//GetPaymentByID Fetch a Payment by ID
func (r *RazorPay) GetPaymentByID(id string) (*Payment, error) {
	paymentresp := new(Payment)
	resp, err := r.call("GetPaymentByID", nil, id, nil)
	json.Unmarshal(resp, paymentresp)
	return paymentresp, err
}

//GetPayments Fetch multiple payments
func (r *RazorPay) GetPayments(from time.Time, to time.Time, count int, skip int) []Payment {
	paymentresp := new(Payment)
	resp, err := r.call("GetPayments", nil, urlparams)
	json.Unmarshal(resp, paymentresp)
	return paymentresp, nil
}





