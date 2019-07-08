package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/jake-mok-nelson/kasa-go"
)

// Turns the heater off

func main() {

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Open our jsonFile
	jsonFile, err := os.Open("creds.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err.Error())
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err.Error())
	}
	var creds Credentials

	err = json.Unmarshal(byteValue, &creds)
	if err != nil {
		log.Println(err.Error())
	}
	api, err := kasa.Connect(creds.Username, creds.Password)
	if err != nil {
		log.Println(err.Error())
	}

	heater, err := api.GetHS100("Heater")
	if err != nil {
		log.Println(err.Error())
	}
	println("I was told to turn the heater off, so I guess I'll do that")
	err = heater.TurnOff()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Print("Heater turned off successfully!")
}
