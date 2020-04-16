package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	myString2 := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "Donne un arg"
	} else {
		myString = arguments[1]
		myString2 = arguments[2]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, myString2)
	io.WriteString(os.Stdout, "\n")
}
