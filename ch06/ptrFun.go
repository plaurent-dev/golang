package main

import (
	"fmt"
)

func getPtr(v *float64) float64 {
	return *v * *v
}

func getPtrSansPtr(v float64) float64 {
	return v * v
}

func main() {
	x := 12.2
	fmt.Println(getPtr(&x))
	x = 12
	fmt.Println(getPtr(&x))
	fmt.Println("________________")
	x = 12.2
	fmt.Println(getPtrSansPtr(x))
	x = 12
	fmt.Println(getPtrSansPtr(x))
}
