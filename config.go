package main

import (
	"encoding/json"
)



type JsonData struct {
	Conf Config `json:"config"`
}

type Config struct {
	Devices []string `json:"devices"`
}

var devices []string


func initConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var jsondata JsonData
	err = json.Unmarshal(file, &jsondata)
	if err != nil {
		panic(err)
	}

	for _, device := range jsondata.Conf.Devices {
		fmt.Printf("devices are :%s\n", device)
	}
	devices = jsondata.Conf.Devices
}

