package dice

import (
	"reflect"
	"testing"
)

func TestParseRollString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want DiceInfo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseRollString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseRollString() = %v, want %v", got, tt.want)
			}
		})
	}
}

//{name: "d6", args:args{"d6"}, wantNumberOfDice: 1, wantTypeOfDice: 6},
//{name: "2d6", args:args{"2d6"}, wantNumberOfDice: 2, wantTypeOfDice: 6},
//{name: "10d6", args:args{"10d6"}, wantNumberOfDice: 10, wantTypeOfDice: 6},
