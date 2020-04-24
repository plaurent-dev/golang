package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Record struct {
	Nom    string
	Prenom string
	Tel    []Telephone
}

type Telephone struct {
	Mobile bool
	Nombre string
}

func loadFromJSON(filename string, key interface{}) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}

	decodeJSON := json.NewDecoder(in)
	err = decodeJSON.Decode(key)
	if err != nil {
		return err
	}
	in.Close()
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Merci de fournir un nom de fichier")
		return
	}

	filename := arguments[1]

	var myRecord Record
	err := loadFromJSON(filename, &myRecord)

	if err == nil {
		fmt.Println(myRecord.Tel[0].Nombre)
	} else {
		fmt.Println(err)
	}

	fmt.Println("___________________________")

	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result)
	fmt.Println("---------")

	for k, v := range result {
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}
}
