package main

import (
	"reflect"
	"testing"
)

func TestRoll(t *testing.T) {
	type args struct {
		n int
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
			if got := Roll(tt.args.n); got != tt.want {
				t.Errorf("Roll() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestKeepHighestRolls(t *testing.T) {
	type args struct {
		h  int
		rs []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeepHighestRolls(tt.args.h, tt.args.rs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeepHighestRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}
