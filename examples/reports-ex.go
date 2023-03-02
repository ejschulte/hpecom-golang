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

	//  Get all reports
	reports, err := cClient.GetReports(nil, "", "")
	if err == nil {
		for i := 0; i < reports.Count; i++ {
			fmt.Println(reports.Items[i].Name)
		}
	} else {
		fmt.Println(err)
	}
}
