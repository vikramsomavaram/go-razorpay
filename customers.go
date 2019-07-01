package razorpay

import "encoding/json"

//GetCustomerByID Fetch Customer by ID
func (r *RazorPay) GetCustomerByID(id string) (*Customer, error) {
	customerresp := new(Customer)
	resp, err := r.call("GetCustomerByID", nil, id, nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, customerresp)
	if err != nil{
		return nil, err
	}
	return customerresp, err
}