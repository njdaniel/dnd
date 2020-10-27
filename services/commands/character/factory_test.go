package character

import "testing"

func Test_createHeritage(t *testing.T) {
	type args struct {
		nc Character
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Elf",args{Character{Race:Race(2).String()}}, "High" },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createHeritage(tt.args.nc); got != tt.want {
				t.Errorf("createHeritage() = %v, want %v", got, tt.want)
			}
		})
	}
}
