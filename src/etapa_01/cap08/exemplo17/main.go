package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name  string
	Idade int
}

type ByAge []Person

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Idade < b[j].Idade
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {

	ps := []Person{
		{"Teo", 31},
		{"Lara", 30},
		{"Nah", 33},
		{"Maria", 1},
	}

	sort.Sort(ByAge(ps))
	fmt.Println(ps)

}
