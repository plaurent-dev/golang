package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Donnez un ou plusieurs floats")
		os.Exit(1)
	}

	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 32)
	max, _ := strconv.ParseFloat(arguments[2], 32)

	fmt.Println(min * max)
}
