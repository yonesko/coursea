package coursea

import (
	"fmt"
	"github.com/thanhpk/randstr"
	"strings"
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
		{args: args{a: "wiwnterbreak", b: "breakwiwnter"}, want: true},
		{args: args{a: "winterbrwintereak", b: "brwintereakwinter"}, want: true},
		{args: args{a: "abcdef", b: "efabcd"}, want: true},
		{args: args{a: "tttttttttt", b: "tttttttttt"}, want: true},
		{args: args{a: "123456", b: "7890"}, want: false},
		{args: args{a: "7g7Xn2qO", b: "n2qO7g7X"}, want: true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := isCyclicRotation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("isCyclicRotation() = %v, want %v", got, tt.want)
			}
		})
	}

	for i := 0; i < 1000; i++ {
		a, b := randstr.String(4), randstr.String(4)
		t.Run(a+" "+b, func(t *testing.T) {
			if got := isCyclicRotation(a+b, b+a); got != true {
				t.Errorf("isCyclicRotation() = %v, want %v", got, true)
			}
		})
	}
}

func Benchmark_isCyclicRotation(bench *testing.B) {
	a := strings.Repeat("123", 1)
	b := strings.Repeat("abc", 1)
	for i := 0; i < bench.N; i++ {
		isCyclicRotation(strings.Repeat(a, i), strings.Repeat(b, i))
	}
}

func Test_tandemRepeat(t *testing.T) {
	type args struct {
		a string
	}
	tests := []struct {
		args args
		want string
	}{
		{args: args{a: "abcabcababcaba"}, want: "abcababcab"},
		{args: args{a: "123456565678"}, want: "565656"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := tandemRepeat(tt.args.a); got != tt.want {
				t.Errorf("tandemRepeat() = %v, want %v", got, tt.want)
			}
		})
	}

	repeat := strings.Repeat(randstr.String(3), 5)
	t.Run(repeat, func(t *testing.T) {
		if got := tandemRepeat("12" + repeat + "44"); got != repeat {
			t.Errorf("tandemRepeat() = %v, want %v", got, repeat)
		}
	})
}
