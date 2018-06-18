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
			number: "4444333322221111",
			cvv2:   "123",
			expiry: "06/2019",
		},
		&CreditCard{
			"Leonardo Faoro",
			"4444333322221111",
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
	card := New("Leonardo Faoro", "4444333322221111", "123", "06/2019")
	fmt.Println(card.Number)
	// Output: 4444333322221111
}

func ExampleCreditCard_First6() {
	card := New("Leonardo Faoro", "4444333322221111", "123", "06/2019")
	fmt.Println(card.First6())
	// Output: 444433
}

func ExampleCreditCard_Last4() {
	card := New("Leonardo Faoro", "4444333322221111", "123", "06/2019")
	fmt.Println(card.Last4())
	// Output: 1111
}

func ExampleCreditCard_Encrypt() {
	card := New("Leonardo Faoro", "4444333322221111", "123", "06/2019")
	fmt.Println(card.Encrypt("myAwesomeSalt"))
	// Output: [89 210 45 55 151 142 210 59 208 84 108 251 121 255 44 227 92 245 235 112 38 44 112 212 86 246 207 25 19 117 101 160 91 78 132 95 204 6 138 75 122 66 224 206 140 178 17 119 183 153 113 204 119 53 105 224 109 186 40 76 83 198 120 9]
}

func ExampleCreditCard_Issuer() {
	card := New("Leonardo Faoro", "4444333322221111", "123", "06/2019")
	fmt.Println(card.Issuer())
	// Output: visa
}

func ExampleCreditCard_ToJSON() {
	card := New("Leonardo Faoro", "4444333322221111", "123", "06/2019")
	fmt.Println(card.ToJSON())
	// Output: {"name":"Leonardo Faoro","number":"4444333322221111","cvv_2":"123","expiry":"06/2019"}
}
