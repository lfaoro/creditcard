// Copyright 2018 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package creditcard

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		name   string
		number string
		cvv2   string
		expiry string
	}
	tests := []struct {
		name string
		args args
		want *CreditCard
	}{{
		"ciao",
		args{
			name:   "Leonardo Faoro",
			number: "1234567891234567",
			cvv2:   "123",
			expiry: "06/2019",
		},
		&CreditCard{
			"Leonardo Faoro",
			"1234567891234567",
			"123",
			"06/2019",
		},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.name, tt.args.number, tt.args.cvv2, tt.args.expiry); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleNew() {
	card := New("Leonardo Faoro", "1234567891234567", "123", "06/2019")
	fmt.Println(card.Number)
	// Output: 1234567891234567
}

func ExampleCreditCard_First6() {
	card := New("Leonardo Faoro", "1234567891234567", "123", "06/2019")
	fmt.Println(card.First6())
	// Output: 123456
}

func ExampleCreditCard_Last4() {
	card := New("Leonardo Faoro", "1234567891234567", "123", "06/2019")
	fmt.Println(card.Last4())
	// Output: 4567
}

func ExampleCreditCard_Encrypt() {
	card := New("Leonardo Faoro", "1234567891234567", "123", "06/2019")
	fmt.Println(card.Encrypt("myAwesomeSalt"))
	// Output: [165 47 214 73 219 178 174 233 237 64 123 55 75 223 124 220 218 2 174 35 159 132 30 43 57 182 199 175 207 217 113 68 216 15 245 32 55 221 118 95 142 147 23 93 70 222 106 80 68 82 109 112 234 46 226 4 22 56 94 113 9 41 225 202]
}
