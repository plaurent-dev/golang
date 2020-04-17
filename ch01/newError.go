package main

import (
	"errors"
	"fmt"
)

func returnError(a, b int) error {
	if a == b {
		err := errors.New("Error dans la fonction returnError()")
		return err
	} else {
		return nil
	}
}

func main() {
	err := returnError(1, 2)

	if err == nil {
		fmt.Println("returnError() a fini normalement")
	} else {
		fmt.Println(err)
	}
	fmt.Println("______________________________")

	err = returnError(10, 10)

	if err == nil {
		fmt.Println("returnError() a fini normalement")
	} else {
		fmt.Println(err)
	}

	if err.Error() == "Error dans la fonction returnError()" {
		fmt.Println("!!!!")
	}

}
