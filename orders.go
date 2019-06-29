package razorpay

import (
	"encoding/json"
	"log"
)

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