package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/jake-mok-nelson/kasa-go"
)

// Turns the heater off

// Some config
var tempThreshold float64 = 15.0 // If the weather is equal to or below this, we will turn on the heater
var city string = "Melbourne"    // Where we are

func main() {

	type Credentials struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		WeatherKey string `json:"weatherkey"` // This is gotten from the openweatherapi, it's free as long as you don't query too often
	}

	jsonFile, err := os.Open("creds.json")
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

	wh, err := owm.NewCurrent("C", "en", creds.WeatherKey) // Celsius, english, weatherapi key
	if err != nil {
		log.Println(err.Error())
	}
	err = wh.CurrentByName(city)
	if err != nil {
		log.Println(err.Error())
	}
	currentTemp := wh.Main.Temp //Let's make this shit more readable
	if wh.Main.Temp <= tempThreshold {
		fmt.Printf("\nIt's pretty chilly in %s at %f, turning on the heater.\n", city, currentTemp)

		err = heater.TurnOn()
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Print("Heater turned on successfully!")
	} else {
		println("Heat is above threshold so I won't turn the heater on.")
	}
}
