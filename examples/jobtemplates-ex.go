package main

import (
	"errors"
	"fmt"

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
	//  Get all jobTemplates
	jobTemplates, err := cClient.GetJobTemplates(nil, "", "")
	if err == nil {
		for i := 0; i < jobTemplates.Count; i++ {
			fmt.Println(jobTemplates.Items[i].Name)
		}
	} else {
		fmt.Println(err)
	}
}
