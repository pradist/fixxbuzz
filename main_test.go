package main

import (
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "WhenInputIs1_ShouldReturn1", args: args{1}, want: "1"},
		{name: "WhenInputIs2_ShouldReturn2", args: args{2}, want: "2"},
		{name: "WhenInputIs3_ShouldReturnFizz", args: args{3}, want: "Fizz"},
		{name: "WhenInputIs4_ShouldReturn4", args: args{4}, want: "4"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.input); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}