package character

import (
	"testing"
)

func Test_createHeritage(t *testing.T) {
	type args struct {
		nc Character
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Elf",args{Character{Race:Race(2).String()}}, []string{"High", "Drow", "Wood"} },
		{"Dwarf", args{Character{Race:Race(3).String()}}, []string{"IronHills", "Mountain", "Deep"}},
		{"Human", args{Character{Race:Race(1).String()}}, []string{"Imperial", "Nord", "Vardisan", "Lumdrani", "Nimalese", "Minskite"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createHeritage(tt.args.nc)
			match := false
			for _, v := range tt.want {
				if got == v {
					match = true
					break
				}
			}
			if !match {
				t.Errorf("createHeritage() = %v, want %v", got, tt.want)
			}
		})
	}
}
