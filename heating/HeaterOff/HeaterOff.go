package HeaterOff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jake-mok-nelson/kasa-go"
)

// Turns the heater off

func HeaterOff(w http.ResponseWriter, r *http.Request) {

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Open our jsonFile
	jsonFile, err := os.Open("creds.json")
	// if we os.Open returns an error then handle it
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

	println("I was told to turn the heater off, so I guess I'll do that")
	err = heater.TurnOff()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	fmt.Print("Heater turned off successfully!")
	// all good. write our message.
	w.WriteHeader(http.StatusOK)
}
