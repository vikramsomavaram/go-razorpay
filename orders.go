package razorpay

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
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
	err = json.Unmarshal(resp, orderresp)
	if err != nil{
		return nil, err
	}
	return orderresp, err
}

//GetOrders Fetch Multiple Orders
func (r *RazorPay) GetOrders(from time.Time, to time.Time, count int, skip int, authorized string, receipt string) (*[]Order, error) {
	orderresp := new([]Order)
	queryparams := make(map[string]string)
	queryparams["from"]= strconv.Itoa(int(from.Unix()))
	queryparams["to"]= strconv.Itoa(int(to.Unix()))
	queryparams["count"]= strconv.Itoa(count)
	queryparams["skip"]= strconv.Itoa(skip)
	queryparams["authorized"]= authorized
	queryparams["receipt"]= receipt
	resp, err := r.call("GetOrders", nil, "", queryparams)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, orderresp)
	if err != nil{
		return nil, err
	}
	return orderresp, nil
}

//GetOrderByID Fetch an Order with Id
func (r *RazorPay) GetOrderByID(id string) (*Order, error) {
	orderresp := new(Order)
	resp, err := r.call("GetOrderByID", nil, id, nil)
	if err != nil{
		return nil, err
	}
	err = json.Unmarshal(resp, orderresp)
	if err != nil{
		return nil, err
	}
	return orderresp, nil
}