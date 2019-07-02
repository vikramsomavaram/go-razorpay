package razorpay

import (
	"encoding/json"
	"strconv"
	"time"
)

//GetSettlements Fetch Multiple Settlements
func (r *RazorPay) GetSettlements(from time.Time, to time.Time, count int, skip int) (*[]Settlements, error) {
	settlementsresp := new([]Settlements)
	queryparams := make(map[string]string)
	queryparams["from"]= strconv.Itoa(int(from.Unix()))
	queryparams["to"]= strconv.Itoa(int(to.Unix()))
	queryparams["count"]= strconv.Itoa(count)
	queryparams["skip"]= strconv.Itoa(skip)
	resp, err := r.call("GetSettlements", nil, "", queryparams)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, settlementsresp)
	if err != nil{
		return nil, err
	}
	return settlementsresp, err
}

//GetSettlementByID Fetch Settlement using ID
func (r *RazorPay) GetSettlementByID(id string) (*Settlement, error) {
	settlementsresp := new(Settlement)
	resp, err := r.call("GetSettlementByID", nil, id, nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, settlementsresp)
	if err != nil{
		return nil, err
	}
	return settlementsresp, err
}
