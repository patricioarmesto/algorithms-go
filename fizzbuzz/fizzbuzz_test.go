package fizzBuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want string
	}{{name: "15 should be FizzBuzz", args: args{val: 15}, want: "FizzBuzz"},
		{name: "5 should be Buzz", args: args{val: 5}, want: "Buzz"},
		{name: "3 should be Fizz", args: args{val: 3}, want: "Fizz"},
		{name: "2 should be 2", args: args{val: 2}, want: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.val); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
