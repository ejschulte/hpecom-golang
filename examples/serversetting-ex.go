package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ejschulte/hpecom-golang/com"
)

func main() {

	var (
		ClientCom com.ComClient
	)
	// fmt.Println("Calling NewComClient")
	cClient := ClientCom.NewComClient()
	// fmt.Println(cClient.Endpoint)
	err := cClient.ComLogin()
	if err != nil {
		panic(errors.New(fmt.Sprintln("failed to establish session")))
	}
	// Set appropriate API endpoint
	cClient.Endpoint = "https://us-west2-api.compute.cloud.hpe.com"

	//  Get all ServerSettings
	serverSettings, err := cClient.GetServerSettings(nil, "", "")
	if err == nil {
		for i := 0; i < serverSettings.Count; i++ {
			fmt.Println(serverSettings.Items[i].Name)
		}
	} else {
		fmt.Println(err)
	}

	// Get Bios ServerSettings
	biosSetting, err := cClient.GetBiosServerSettings()
	if err == nil {
		for i := 0; i < biosSetting.Count; i++ {
			fmt.Printf("BIOS setting - %s\n", biosSetting.Items[i].Name)
			//
			val := reflect.ValueOf(biosSetting.Items[i].Settings)
			field := val.Type()

			for j := 0; j < field.NumField(); j++ {
				fmt.Printf("\t%s ---> %v\n", field.Field(j).Name, val.Field(j).Interface())
			}
		}
	} else {
		fmt.Println(err)
	}

	// Get Firmware ServerSettings
	firmwareSetting, err := cClient.GetFirmwareServerSettings()
	if err == nil {
		for i := 0; i < firmwareSetting.Count; i++ {
			fmt.Printf("Firmware setting - %s\n", firmwareSetting.Items[i].Name)

			val := reflect.ValueOf(firmwareSetting.Items[i].Settings)
			field := val.Type()

			for j := 0; j < field.NumField(); j++ {
				fmt.Printf("\t%s firmware bundle id ---> %v\n", field.Field(j).Name, val.Field(j).Interface())
			}
		}
	} else {
		fmt.Println(err)
	}

	// Get Storage ServerSettings
	storageSetting, err := cClient.GetStorageServerSettings()
	if err == nil {
		for i := 0; i < storageSetting.Count; i++ {
			fmt.Printf("Storage setting - %s\n", storageSetting.Items[i].Name)

			val := reflect.ValueOf(storageSetting.Items[i].Settings)
			field := val.Type()

			for j := 0; j < field.NumField(); j++ {
				fmt.Printf("\t%s ---> %v\n", field.Field(j).Name, val.Field(j).Interface())
			}
		}
	} else {
		fmt.Println(err)
	}

}
