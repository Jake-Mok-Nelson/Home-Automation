package HeaterOn

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	owm "github.com/briandowns/openweathermap"
	"github.com/jake-mok-nelson/kasa-go"
)

// Turns the heater off

// Some config
var tempThreshold float64 = 15.0 // If the weather is equal to or below this, we will turn on the heater
var city string = "Melbourne"    // Where we are

func HeaterOn(w http.ResponseWriter, r *http.Request) {

	type Credentials struct {
		Username   string `json:"username"`
		Password   string `json:"password"`
		WeatherKey string `json:"weatherkey"` // This is gotten from the openweatherapi, it's free as long as you don't query too often
	}

	jsonFile, err := os.Open("creds.json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	var creds Credentials

	err = json.Unmarshal(byteValue, &creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	api, err := kasa.Connect(creds.Username, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	heater, err := api.GetHS100("Heater")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}

	wh, err := owm.NewCurrent("C", "en", creds.WeatherKey) // Celsius, english, weatherapi key
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	err = wh.CurrentByName(city)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	currentTemp := wh.Main.Temp //Let's make this shit more readable
	if wh.Main.Temp <= tempThreshold {
		fmt.Printf("\nIt's pretty chilly in %s at %f, turning on the heater.\n", city, currentTemp)
		err = heater.TurnOn()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			log.Println(err.Error())
			return
		}
		fmt.Print("Heater turned on successfully!")
	} else {
		println("Heat is above threshold so I won't turn the heater on.")
	}

	// all good. write our message.
	w.WriteHeader(http.StatusOK)

}
