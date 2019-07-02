package razorpay

import (
	"encoding/json"
)

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

//CaptureCustomer Create a customer
func (r *RazorPay) CreateCustomer(customerInput CustomerInput) (*Customer, error) {
	customerresp := new(Customer)
	createreqjson, err := json.Marshal(customerInput)
	if err != nil{
		return nil, err
	}
	resp, err := r.call("CreateCustomer", createreqjson, "", nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, customerresp)
	if err != nil{
		return nil, err
	}
	return customerresp, err
}

//EditCustomerByID Edit a customer
func (r *RazorPay) EditCustomerByID(id string, customerInput CustomerInput) (*Customer, error) {
	customerresp := new(Customer)
	createreqjson, err := json.Marshal(customerInput)
	if err != nil{
		return nil, err
	}
	resp, err := r.call("EditCustomerByID", createreqjson, id, nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, customerresp)
	if err != nil{
		return nil, err
	}
	return customerresp, err
}

//GetCustomers Fetch Multiple Customers
func (r *RazorPay) GetCustomers() (*[]Customers, error) {
	customerresp := new([]Customers)
	resp, err := r.call("GetCustomers", nil, "", nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, customerresp)
	if err != nil{
		return nil, err
	}
	return customerresp, err
}

