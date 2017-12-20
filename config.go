package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Configuration struct {
	ServerPort string
}

var configFileName = "config.json"
var err error
var config Configuration

func ReadConfig() (Configuration, error) {
	configFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Print("Unable to read " + configFileName + ", switching flag mode")
		return Configuration{}, err
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil || config.ServerPort == "" {
		log.Print("Invalid JSON, expecting port from command line flag")
		return Configuration{}, err
	}

	return config, nil
}
