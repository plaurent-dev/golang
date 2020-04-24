package huaweiPkg

import (
	"fmt"
)

func A() {
	fmt.Println("HUAWEI 1!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}

const MyConstant = 123
const privateConstant = 21
