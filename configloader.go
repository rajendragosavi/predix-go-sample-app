package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
)

// AppConfig ...
type AppConfig struct {
	Postgres struct {
		Hostname string `json:"hostname"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"postgres"`
	WebServer struct {
		Port string `json:"port"`
	} `json:"webserver"`
}

func loadConfig(filename string, config *AppConfig) error {
	var err error
	if filename != "" {
		log.Println("-> Loading configuration file...")
		err = loadFile(filename, config)
		log.Println("-> DONE!")
	}
	return errors.Wrap(err, "filename cannot be an empty string")
}

func loadFile(filename string, config *AppConfig) error {
	configFile, err := os.Open(filename)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)

	if err = decoder.Decode(&config); err != nil {
		return errors.Wrap(err, "failed to decode config file")
	}
	return nil
}
