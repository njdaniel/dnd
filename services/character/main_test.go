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
		min  int
		max  int
	}{
		{name: "d6", args: args{6}, min: 1, max: 6},
		{name: "d10", args: args{10}, min: 1, max: 10},
		{name: "d20", args: args{20}, min: 1, max: 20},
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

func TestCalcAttrBonus(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"20", args{20}, 5},
		{"19", args{19}, 4},
		{"18", args{18}, 4},
		{"17", args{17}, 3},
		{"16", args{16}, 3},
		{"15", args{15}, 2},
		{"14", args{14}, 2},
		{"13", args{13}, 1},
		{"12", args{12}, 1},
		{"11", args{11}, 0},
		{"10", args{10}, 0},
		{"9", args{9}, -1},
		{"8", args{8}, -1},
		{"7", args{7}, -2},
		{"6", args{6}, -2},
		{"5", args{5}, -3},
		{"4", args{4}, -3},
		{"3", args{3}, -4},
		{"2", args{2}, -4},
		{"1", args{1}, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcAttrBonus(tt.args.a); got != tt.want {
				t.Errorf("CalcAttrBonus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadCSV(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCSV(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
