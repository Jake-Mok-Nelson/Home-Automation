# Heater-On

## Description

Turns my bedroom heater on using a HS100 TP Link (Kasa) smart plug. They expose functionality through a webapi.

I set variables for the temperature threshold and my location (city).

If the temperature in my city is lower than the threshold
at the time of running the function, it will turn on the heater.

Using [Kasa-go](https://github.com/ivanbeldad/kasa-go) with the [openweathermap](https://github.com/briandowns/openweathermap) Go library.

Requires a free account on [Openweathermap](https://openweathermap.org/)

## How to use

Create creds.json file and populate the fields.

> go get

> go run HeaterOff.go
