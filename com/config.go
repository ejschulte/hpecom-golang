package com

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	ClientId     string `json:"id"`     //Com application API ID
	ClientSecret string `json:"secret"` //Com application API Secret
	Jwt          string `json:"jwt"`
	Endpoint     string `json:"endpoint"` //Com region API endpoint (ex. us-west2-api.compute.cloud.hpe.com)
}

func ConfigFromFile() (Config, error) {
	var config Config
	_, filename, _, _ := runtime.Caller(1)
	configFilePath := filepath.Join(filepath.Dir(filename), "ComConfig.json")
	configF, err := os.Open(configFilePath)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error opening json file")
		return config, err
	}
	defer configF.Close()

	configFile := json.NewDecoder(configF)
	err_unmarshal := configFile.Decode(&config)
	if err_unmarshal != nil {
		fmt.Println("error unmarshaling json file")
		fmt.Println(err_unmarshal)

	}
	return config, nil
}

func ConfigFromOs() (Config, error) {
	var config Config
	config.ClientId = os.Getenv("ComClientID")
	config.ClientSecret = os.Getenv("ComClientSecret")

	// jsonData := map[string]string{
	// // jsonData := map[string]string{
	// 	"id":       os.Getenv("ComClientID"),
	// 	"secret":   os.Getenv("ComClientSecret"),
	// 	"endpoint": os.Getenv("ComEndpoint"),
	// }

	if config.ClientId == "" || config.ClientSecret == "" {
		return config, errors.New("variable load error")
	}

	// if jsonData["id"] == "" || jsonData["secret"] == "" {
	// 	return config, errors.New("variable load error")
	// }

	// marshalled, err := json.Marshal(jsonData)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// jsonErr := json.Unmarshal(marshalled, &config)
	// if jsonErr != nil {
	// 	fmt.Println("error parsing json")
	// }
	return config, nil

}

func LoadConfig() (Config, error) {
	var config Config

	// Try OS variables first
	config, err := ConfigFromOs()
	if err != nil {

		fmt.Println("os variables not found or incomplete")
		fmt.Println("attempting config load from file")

		// No OS vars, try from file
		config, err := ConfigFromFile()

		if err != nil {
			return config, err
		}
	}
	// fmt.Println(config.ClientId)
	return config, nil
}
