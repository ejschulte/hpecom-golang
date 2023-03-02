package main

import (
	"errors"
	"fmt"

	"github.com/ejschulte/hpecom-golang/com"
)

func main() {

	var (
		ClientCom com.ComClient
		// u *url.URL
	)

	// fmt.Println("Calling NewComClient")
	cClient := ClientCom.NewComClient()
	// fmt.Println(cClient.Endpoint)
	err := cClient.ComLogin()
	if err != nil {
		panic(errors.New(fmt.Sprintln("failed to establish session")))
	}
	// fmt.Println(cClient.Jwt)

	// Set appropriate API endpoint
	cClient.Endpoint = "https://us-west2-api.compute.cloud.hpe.com"

	// Get servers with query
	query := []string{}
	query = append(query, "hardware/model eq 'ProLiant DL325 Gen10 Plus'")
	serverList, err := cClient.GetServers(query, "", "")
	if err == nil {
		for i := 0; i < serverList.Count; i++ {
			fmt.Println(serverList.Items[i].Name)
			fmt.Println(serverList.Items[i].Hardware.Model)
		}
	}

	//  Get server by name
	server, err := cClient.GetServerByName("HPE-HOL07")
	if err == nil {
		s := fmt.Sprintf("Found server %s", server.Name)
		fmt.Println(s)
	} else {
		fmt.Println(err)
	}

	// Get groups
	groupList, err := cClient.GetGroups(nil, "", "")
	if err == nil {
		for i := 0; i < groupList.Count; i++ {
			fmt.Println(groupList.Items[i].Name)
		}
	} else {
		fmt.Println(err)
	}

	// Get All ServerSettings
	allSettings, err := cClient.GetServerSettings(nil, "", "")
	if err == nil {
		for i := 0; i < allSettings.Count; i++ {
			fmt.Println(allSettings.Items[i].Settings)
		}
	} else {
		fmt.Println(err)
	}

	// Get Bios ServerSettings
	biosSetting, err := cClient.GetBiosServerSettings()
	if err == nil {
		for i := 0; i < biosSetting.Count; i++ {
			fmt.Println(biosSetting.Items[i].Settings)
		}
	} else {
		fmt.Println(err)
	}

	// Get Firmware ServerSettings
	firmwareSetting, err := cClient.GetFirmwareServerSettings()
	if err == nil {
		for i := 0; i < firmwareSetting.Count; i++ {
			fmt.Println(firmwareSetting.Items[i].Settings)
		}
	} else {
		fmt.Println(err)
	}

	// Get Storage ServerSettings
	storageSetting, err := cClient.GetStorageServerSettings()
	if err == nil {
		for i := 0; i < storageSetting.Count; i++ {
			fmt.Println(storageSetting.Items[i].Settings)
		}
	} else {
		fmt.Println(err)
	}

	// Get job template by name
	id := com.PowerOffNew.String()
	fmt.Println(id)

	jobTemplate, err := cClient.GetJobTemplateByName(com.PowerOffNew)
	if err == nil {
		fmt.Println(jobTemplate.Name)
	} else {
		fmt.Println(err)
	}
}
