package com

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Session struct {
	Jwt string `json:"access_token,omitempty"`
}

func (c *ComClient) GetAuthHeaders() map[string]string {

	return map[string]string{
		"Authorization": c.Jwt,
	}
}

// func (c *ComClient) ComLogin() (Session, error) {
func (c *ComClient) ComLogin() (err error) {
	var (
		uri     = "https://sso.common.cloud.hpe.com/as/token.oauth2"
		rawBody = "grant_type=client_credentials&client_id=" + c.Id + "&client_secret=" + c.Secret
		session Session
	)

	// fmt.Println(uri)
	body := bytes.NewBuffer([]byte(rawBody))
	req, _ := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()
	rsBody, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal([]byte(rsBody), &session); err != nil {
		return err
	}

	c.Jwt = "Bearer " + session.Jwt
	// Update APIKey
	return err
}
