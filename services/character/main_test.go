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
		min int
		max int
	}{
		{name:"d6", args:args{6}, min: 1, max: 6},
		{name:"d10", args:args{10}, min:1, max: 10},
		{name:"d20", args:args{20}, min: 1, max: 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Roll(tt.args.n); got < tt.min || got > tt.max {
				t.Errorf("Roll() = %v, want between %v and %v", got, tt.min, tt.max)
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
