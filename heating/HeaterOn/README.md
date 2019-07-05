# Heater-On

## Description

Turns my bedroom heater on using a HS100 TP Link (Kasa) smart plug. They expose functionality through a webapi.

I set variables for the temperature threshold and my location (city).

If the temperature in my city is lower than the threshold
at the time of running the function, it will turn on the heater.

Using [Kasa-go](https://github.com/ivanbeldad/kasa-go) with the [openweathermap](https://github.com/briandowns/openweathermap) Go library.

Requires a free account on [Openweathermap](https://openweathermap.org/)

How to use...

Send the cloud function to ... you know, the cloud.
`gcloud functions deploy HeaterOn --runtime go111 --trigger-http`

Record the output for your IFTTT app or whatever you want to call it with.

e.g.

```
Deploying function (may take a while - up to 2 minutes)...done.
availableMemoryMb: 256
entryPoint: HelloWorld
httpsTrigger:
  url: https://us-central1-<PROJECT_ID>.cloudfunctions.net/HeaterOn
labels:
  deployment-tool: cli-gcloud
...
```
