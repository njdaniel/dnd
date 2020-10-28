package character

import (
	"testing"
)

func Test_createHeritage(t *testing.T) {
	type args struct {
		race string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Elf",args{Race(2).String()}, []string{"High", "Drow", "Wood"} },
		{"Dwarf", args{Race(3).String()}, []string{"IronHills", "Mountain", "Deep"}},
		{"Human", args{Race(1).String()}, []string{"Imperial", "Nord", "Vardisan", "Lumdrani", "Nimalese", "Minskite"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createHeritage(tt.args.race)
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
