package day3

import (
	"testing"
)

func TestA(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := A(tt.args.input); got != tt.want {
				t.Errorf("A() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestB(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := B(tt.args.input); got != tt.want {
				t.Errorf("B() = %v, want %v", got, tt.want)
			}
		})
	}
}
