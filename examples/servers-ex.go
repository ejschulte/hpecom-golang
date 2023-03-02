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

}
