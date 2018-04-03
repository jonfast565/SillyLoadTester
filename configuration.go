package main

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	RequestUrl      string
	RequestVerb     string
	PostBody        string
	PostContentType string
	NumberOfTasks   int
	CallsPerTask    int
}

func ReadConfig(path string, logger *ProgramLogger) *Configuration {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := new(Configuration)
	err := decoder.Decode(&configuration)
	if err != nil {
		logger.LogError(err.Error())
		panic(err)
	}
	logger.LogSuccess("Configuration read successfully.")
	return configuration
}
