package main

import (
	"fmt"
	"log"
	"os"
)

// LOGFILE est un fichier de test
var LOGFILE = "C:\\temp\\log_go2.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	Log := log.New(f, "customLogAvecdeslignes", log.LstdFlags)

	Log.SetFlags(log.LstdFlags | log.Lshortfile)
	Log.Println("Salut")
	Log.Println("Autre ligne")
}

//customLogAvecdeslignes2020/04/17 09:33:52 customLogLineNumber.go:25: Salut
//customLogAvecdeslignes2020/04/17 09:33:52 customLogLineNumber.go:26: Autre ligne
