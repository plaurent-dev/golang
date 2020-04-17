package main

import (
	"fmt"
	"log"
	"os"
)

// LOGFILE est un fichier de test
var LOGFILE = "C:\\temp\\log_go.log"

func main() {
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	Log := log.New(f, "customLogAToto", log.LstdFlags)

	Log.SetFlags(log.LstdFlags)
	Log.Println("Salut")
	Log.Println("Autre ligne")
}

//customLogAToto2020/04/17 09:30:10 Salut
//customLogAToto2020/04/17 09:30:10 Autre ligne
//customLogAToto2020/04/17 09:31:47 Salut
//customLogAToto2020/04/17 09:31:47 Autre ligne
