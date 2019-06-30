package razorpay

import (
	"encoding/json"
	"strconv"
	"time"
)

//Refund issue refund
func (r *RazorPay) CreateRefund(id string) (*Refund, error) {
	refundresp := new(Refund)
	resp, err := r.call("CreateRefund", nil, id+ "/refund", nil)
	err = json.Unmarshal(resp, refundresp)
	if err != nil {
		return nil, err
	}
	return refundresp, err
}

//GetRefundsByPaymentID Fetch Multiple Refunds for a Payment
func (r *RazorPay) GetRefundsByPaymentID(id string, from time.Time, to time.Time, count int, skip int) (*[]Refunds, error) {
	refundresp := new([]Refunds)
	queryparams := make(map[string]string)
	queryparams["id"]= id
	queryparams["from"]= strconv.Itoa(int(from.Unix()))
	queryparams["to"]= strconv.Itoa(int(to.Unix()))
	queryparams["count"]= strconv.Itoa(count)
	queryparams["skip"]= strconv.Itoa(skip)
	resp, err := r.call("GetRefundsByPaymentID", nil, id+ "/refunds", queryparams)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, refundresp)
	if err != nil{
		return nil, err
	}
	return refundresp, nil
}

//func (r *RazorPay) GetRefundByPaymentID(id string) (*Refund, error) {
//	refundresp := new(Refund)
//	refundreq := refund
//	createreqjson, err := json.Marshal(&refundreq)
//	resp, err := r.call("CreateRefund", createreqjson, nil)
//	json.Unmarshal(resp, refundresp)
//	return refundresp, nil
//}

//GetRefunds Fetch Multiple Refunds
func (r *RazorPay) GetRefunds(from time.Time, to time.Time, count int, skip int) (*[]Refunds, error) {
	refundresp := new([]Refunds)
	queryparams := make(map[string]string)
	queryparams["from"]= strconv.Itoa(int(from.Unix()))
	queryparams["to"]= strconv.Itoa(int(to.Unix()))
	queryparams["count"]= strconv.Itoa(count)
	queryparams["skip"]= strconv.Itoa(skip)
	resp, err := r.call("GetRefunds", nil, "", queryparams)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, refundresp)
	if err != nil{
		return nil, err
	}
	return refundresp, err
}

//GetRefundByID Fetch Refund by ID
func (r *RazorPay) GetRefundByID(id string) (*Refund, error) {
	refundresp := new(Refund)
	resp, err := r.call("GetRefundByID", nil, id, nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, refundresp)
	if err != nil{
		return nil, err
	}
	return refundresp, err
}