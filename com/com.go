package com

import (
	"errors"
	"fmt"

	"example.local/hpecom/rest"
)

type ComClient struct {
	rest.Client
}

func (c *ComClient) NewComClient() *ComClient {

	config, err := LoadConfig()
	// fmt.Println(config)
	if err != nil {
		panic(errors.New(fmt.Sprintln("unable to load configuration parameters")))
		// os.Exit(2)
	}
	// fmt.Println("out of loadConfig")
	// fmt.Println(config.ClientId)

	// session, err := c.ComLogin(config)
	// fmt.Println(session)

	c = &ComClient{
		rest.Client{
			Id:       config.ClientId,
			Secret:   config.ClientSecret,
			Jwt:      config.Jwt,
			Endpoint: config.Endpoint,
		},
	}

	return c
}

func (c *ComClient) SetQueryParams(filters []string, offset, limit string) map[string]interface{} {

	q := make(map[string]interface{})

	if len(filters) > 0 {
		q["filter"] = filters
	}

	if offset != "" {
		q["offset"] = offset
	}

	if limit != "" {
		q["limit"] = limit
	}

	return q
}
