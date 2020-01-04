package coursea

import (
	"fmt"
	"testing"
)

func Test_isCyclicRotation(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		args args
		want bool
	}{
		{args: args{a: "winterbreak", b: "breakwinter"}, want: true},
		{args: args{a: "abcdef", b: "efabcd"}, want: true},
		{args: args{a: "tttttttttt", b: "tttttttttt"}, want: true},
		{args: args{a: "123456", b: "7890"}, want: false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := isCyclicRotation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isCyclicRotation() = %v, want %v", got, tt.want)
			}
		})
	}

}
