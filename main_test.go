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
		{name: "WhenInputIs6_ShouldReturnFizz", args: args{6}, want: "Fizz"},
		{name: "WhenInputIs5_ShouldReturnBuzz", args: args{5}, want: "Buzz"},
		{name: "WhenInputIs10_ShouldReturnBuzz", args: args{10}, want: "Buzz"},
		{name: "WhenInputIs15_ShouldReturnFizzBuzz", args: args{15}, want: "FizzBuzz"},
		{name: "WhenInputIs30_ShouldReturnFizzBuzz", args: args{30}, want: "FizzBuzz"},
		{name: "WhenInputIs100_ShouldReturnBuzz", args: args{100}, want: "Buzz"},
		{name: "WhenInputIs90_ShouldReturnFizzBuzz", args: args{90}, want: "FizzBuzz"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.input); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}