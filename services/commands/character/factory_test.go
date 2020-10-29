package character

import (
	"regexp"
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
		{"Elf", args{Race(2).String()}, []string{"High", "Drow", "Wood"}},
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

func Test_createName(t *testing.T) {
	type args struct {
		race   string
		gender string
	}
	tests := []struct {
		name string
		args args
		//want string
	}{
		{"human male", args{Race(1).String(), Gender(1).String()} },
	}
	re := regexp.MustCompile(`(\w+)`)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createName(tt.args.race, tt.args.gender); re.MatchString(got) {
				t.Errorf("createName() = %v, want %v", got, re)
			}
		})
	}
}
