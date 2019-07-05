# Heater-On

## Description

Turns my bedroom heater off using a HS100 TP Link (Kasa) smart plug. They expose functionality through a webapi.

Using [Kasa-go](https://github.com/ivanbeldad/kasa-go)

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
  url: https://us-central1-<PROJECT_ID>.cloudfunctions.net/HeaterOff
labels:
  deployment-tool: cli-gcloud
...
```
