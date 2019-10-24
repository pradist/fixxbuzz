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
		{name: "WhenInputIs15_ShouldReturnBuzz", args: args{15}, want: "Buzz"},
		{name: "WhenInputIs30_ShouldReturnFizz", args: args{30}, want: "Fizz"},
		{name: "WhenInputIs100_ShouldReturnBuzz", args: args{100}, want: "Buzz"},
		{name: "WhenInputIs90_ShouldReturnFizzBuzz", args: args{90}, want: "FizzBuzz"},
		{name: "WhenInputIs13_ShouldReturnFizz", args: args{13}, want: "Fizz"},
		{name: "WhenInputIs31_ShouldReturnFizz", args: args{31}, want: "Fizz"},
		{name: "WhenInputIs23_ShouldReturnFizz", args: args{23}, want: "Fizz"},
		{name: "WhenInputIs51_ShouldReturnBuzz", args: args{51}, want: "Buzz"},
		{name: "WhenInputIs52_ShouldReturnBuzz", args: args{52}, want: "Buzz"},
		{name: "WhenInputIs56_ShouldReturnBuzz", args: args{56}, want: "Buzz"},
		{name: "WhenInputIs35_ShouldReturnFizzBuzz", args: args{35}, want: "FizzBuzz"},
		{name: "WhenInputIs53_ShouldReturnFizzBuzz", args: args{53}, want: "FizzBuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FizzBuzz(tt.args.input); got != tt.want {
				t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
			}
		})
	}
}
