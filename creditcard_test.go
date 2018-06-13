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
			if got := ValidateLuhn(tt.args.number); got != tt.want {
				t.Errorf("ValidateLuhn() = %v, want %v", got, tt.want)
			}
		})
	}
}
