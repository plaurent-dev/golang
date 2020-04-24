package main

import (
	"fmt"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"Northe", "East", "South", "West"}[d]
}
func main() {
	var d Direction = East
	fmt.Print(d)
	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
