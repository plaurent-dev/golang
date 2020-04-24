package main

import "fmt"

func ajoutun(mess string, s ...int) int {
	fmt.Println(mess)
	somme := 0
	for i, a := range s {
		fmt.Println(i, a)
		somme = somme + a
	}
	s[0] = -1000
	return somme
}

func main() {
	sum := ajoutun(" AJOUT ", 1, 2, 3, 4, 5, 6, 7)
	print(sum)
}
