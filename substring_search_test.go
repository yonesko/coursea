package coursea

import (
	"fmt"
	"github.com/thanhpk/randstr"
	"math"
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
		{args: args{a: "aaaaaaaaaaa", b: "aaaaaaaaaaa"}, want: true},
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
			if got := isCyclicRotation(b+a, a+b); got != true {
				t.Errorf("isCyclicRotation() = %v, want %v", got, true)
			}
		})
	}
}

func Benchmark_isCyclicRotation(bench *testing.B) {
	for p := 1; p <= 15; p++ {
		bench.Run(fmt.Sprintf("2^%d", p), func(bench *testing.B) {
			a := strings.Repeat("123", int(math.Pow(2, float64(p))))
			b := strings.Repeat("124", int(math.Pow(2, float64(p))))
			for i := 0; i < bench.N; i++ {
				isCyclicRotation(a, b)
			}
		})
	}
}
func Benchmark_tandemRepeat(bench *testing.B) {
	for p := 1; p <= 6; p++ {
		bench.Run(fmt.Sprintf("pow-%d", p), func(bench *testing.B) {
			a := randstr.String(int(math.Pow10(p)))
			for i := 0; i < bench.N; i++ {
				tandemRepeat(a)
			}
		})
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
		{args: args{a: "abcabcababcaba"}, want: "abcab"},
		{args: args{a: "123456565678"}, want: "56"},
		{args: args{a: "1234567890"}, want: "1234567890"},
		{args: args{a: "abcabc"}, want: "abc"},
		{args: args{a: "aaaaaaaaaaaa"}, want: "a"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.args), func(t *testing.T) {
			if got := tandemRepeat(tt.args.a); got != tt.want {
				t.Errorf("tandemRepeat() = %v, want %v", got, tt.want)
			}
		})
	}

	b := randstr.String(3)
	t.Run(strings.Repeat(b, 5), func(t *testing.T) {
		if got := tandemRepeat("12" + strings.Repeat(b, 5) + "44"); got != b {
			t.Errorf("tandemRepeat() = %v, want %v", got, strings.Repeat(b, 5))
		}
	})
}
