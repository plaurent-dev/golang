package main

import (
	"aPackage"
	"fmt"
	"huaweiPkg"
)

func main() {
	fmt.Println("Using aPackage!")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.MyConstant)
	huaweiPkg.A()
}
