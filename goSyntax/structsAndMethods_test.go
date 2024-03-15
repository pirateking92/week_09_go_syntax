package main

import "testing"

func Test_drink_calculatePriceAtCheckout(t *testing.T) {
	type fields struct {
		name   string
		volume float32
		price  int
	}
	type args struct {
		quantity int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myDrink := drink{
				name:   tt.fields.name,
				volume: tt.fields.volume,
				price:  tt.fields.price,
			}
			if got := myDrink.calculatePriceAtCheckout(tt.args.quantity); got != tt.want {
				t.Errorf("drink.calculatePriceAtCheckout() = %v, want %v", got, tt.want)
			}
		})
	}
}
