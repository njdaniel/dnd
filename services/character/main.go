package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	fmt.Println("starting...")
	rs := MultiRolls(6, 6)
	fmt.Println("rs: ", rs)
	hrs := KeepHighestRolls(3, rs)
	fmt.Println("hrs: ", hrs)
	fmt.Println("done.")
}

type Attributes struct {
	Strength     int
	Dexterity    int
	Constitution int
	Perception   int
	Intelligence int
	Willpower    int
	Charisma     int
	Movement     int
}

func CreateCharacter() {
	//1) Add attributes
	//roll 4d6 take sum of highest 3
	fmt.Println("starting...")
	fmt.Println(Roll(6))
	fmt.Println("done.")

	//2)

}

func Roll(d int) int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := d
	return rand.Intn(max-min+1) + min
}

func MultiRolls(n, d int) []int {
	rs := []int{}
	for i := 0; i < n; i++ {
		rs = append(rs, Roll(d))
	}
	return rs
}

func MultiRollSum(n, d int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += Roll(d)
	}
	return sum
}

func SumRolls(r []int) int {
	sum := 0
	for _, v := range r {
		sum += v
	}
	return sum
}

func KeepHighestRolls(h int, rs []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(rs)))
	rs = rs[:h]
	return rs
}

func KeepLowestRolls(l int, rs []int) []int {
	sort.Ints(rs)
	rs = rs[:l]
	return rs
}
