package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "C:\\temp\\log_go_ch02.log"

func one(aLog *log.Logger) {
	aLog.Println("__ FUNCTION one __")
	defer aLog.Println("__ FUNCTION one __")

	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func two(aLog *log.Logger) {
	aLog.Println("__ FUNCTION two __")
	defer aLog.Println("__ FUNCTION two __")

	for i := 10; i > 0; i-- {
		aLog.Println(i)
	}
}

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	iLog := log.New(f, "logDefer", log.LstdFlags)
	iLog.Println("salut")
	iLog.Println("Une autre entrée")

	one(iLog)
	two(iLog)
}
