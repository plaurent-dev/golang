package main

import (
	"dockerServices"
	"fmt"
)

type ip struct {
	adress   string
	port     int
	username string
	password string
}

func (i ip) Docker() string {
	// call ssh(ip)
	fmt.Println(i.username)
	return "2.22"
}
func (i ip) Kernel() string {
	return "1.117.c"
}

func (i ip) Memory() float64 {
	return 10.25
}

func Get(x dockerServices.Infos) {
	fmt.Println(x.Docker())
	fmt.Println(x.Kernel())
	fmt.Println(x.Memory())
}

func main() {
	credentials := ip{adress: "10.10.10.10", port: 22, username: "pierre", password: "tenzin"}
	Get(credentials)
}
