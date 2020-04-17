package main

import (
	"io"
	"os"
)

func main() {
	maChaine := ""
	arguments := os.Args
	if len(arguments) == 1 {
		maChaine = "File un argument"
	} else {
		maChaine = arguments[1]
	}

	io.WriteString(os.Stdout, "Sortie standard \n")
	io.WriteString(os.Stderr, maChaine)
	io.WriteString(os.Stderr, "\n")
}
