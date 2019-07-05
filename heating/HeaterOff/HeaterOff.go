package HeaterOff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/ivandelabeldad/kasa-go"
	"github.com/jake-mok-nelson/home-automation/common/writers"
)

// Turns the heater off
func NewMessageWriter(message string) *JSONMessageWriter {
	return &JSONMessageWriter{
		Message: message,
	}
}

func (jw *JSONMessageWriter) JSONString() (string, error) {
	messageResponse := map[string]interface{}{
		"data": map[string]string{
			"message": jw.Message,
		},
	}
	bytesValue, err := json.Marshal(messageResponse)
	if err != nil {
		return "", err
	}
	return string(bytesValue), nil
}

func HeaterOff(w http.ResponseWriter, r *http.Request) {

	type Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Open our jsonFile
	jsonFile, err := os.Open(".creds.json")
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

	jw := writers.NewMessageWriter(message)
	jsonString, err := jw.JSONString()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Println(err.Error())
		return
	}
	// all good. write our message.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonString))
}
