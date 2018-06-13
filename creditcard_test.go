// Copyright 2018 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package creditcard

import "testing"

func TestValidateLuhn(t *testing.T) {
	type args struct {
		number string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{
		"Numbers",
		args{
			"7992739871",
		},
		true,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Luhn(tt.args.number); got != tt.want {
				t.Errorf("Luhn() = %v, want %v", got, tt.want)
			}
		})
	}
}
