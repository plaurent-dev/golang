package main

import (
	"fmt"
)

func main() {
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}
	for i, a := range threeD {
		fmt.Println(i)
		fmt.Println(a)
		fmt.Println("__________________")
	}
}
