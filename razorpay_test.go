package razorpay_test

import (
	"reflect"
	"testing"
	"time"
)

func TestRazorPay_GetPayments(t *testing.T) {
	type fields struct {
		apikey    string
		apisecret string
	}
	type args struct {
		from  time.Time
		to    time.Time
		count int
		skip  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Payment
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RazorPay{
				apikey:    tt.fields.apikey,
				apisecret: tt.fields.apisecret,
			}
			if got := r.GetPayments(tt.args.from, tt.args.to, tt.args.count, tt.args.skip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RazorPay.GetPayments() = %v, want %v", got, tt.want)
			}
		})
	}
}
