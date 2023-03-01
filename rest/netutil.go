package rest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"

	"example.local/hpecom/utils"
)

var (
	codes = map[int]bool{
		http.StatusOK:                   true,
		http.StatusBadRequest:           false,
		http.StatusUnauthorized:         false,
		http.StatusNotFound:             false,
		http.StatusNotAcceptable:        false,
		http.StatusConflict:             false,
		http.StatusUnsupportedMediaType: false,
		http.StatusInternalServerError:  false,
	}

	// TODO: this should have a real cert
	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// get a client
	client = &http.Client{Transport: tr}
	// client = &http.Client{}
)

// Options for REST call
type Options struct {
	Headers map[string]string
	Query   map[string]interface{}
}

// Client - generic REST api client
type Client struct {
	Method
	Id       string
	Secret   string
	Jwt      string
	Endpoint string
	Option   Options
}

// NewClient - get a new network client
func (c *Client) NewClient(id, jwt, secret, endpoint string) *Client {
	return &Client{
		Id:       id,
		Secret:   secret,
		Jwt:      jwt,
		Endpoint: endpoint,
		Option:   Options{},
	}
}

// isOkStatus - check the return status of the response
func (c *Client) isOkStatus(code int) bool {
	return codes[code]
}

// SetQueryString - set the query strings to use
func (c *Client) SetQueryString(query map[string]interface{}) {
	// TODO: uuencode the query String
	c.Option.Query = query
}

// GetQueryString - get a query string for url
func (c *Client) GetQueryStrings(u *url.URL, query map[string]interface{}) {
	if len(query) == 0 {
		return
	}
	parameters := url.Values{}
	for k, v := range query {
		if val, ok := v.([]string); ok {
			for _, va := range val {
				parameters.Add(k, va)
			}
		} else {
			parameters.Add(k, v.(string))
		}
		u.RawQuery = parameters.Encode()
	}
	return
}

// GetQueryString - get a query string for url through the Client Struct
func (c *Client) GetQueryString(u *url.URL) {
	if len(c.Option.Query) == 0 {
		return
	}
	parameters := url.Values{}
	for k, v := range c.Option.Query {
		if val, ok := v.([]string); ok {
			for _, va := range val {
				parameters.Add(k, va)
			}
		} else {
			parameters.Add(k, v.(string))
		}
		u.RawQuery = parameters.Encode()
	}
	return
}

// SetAuthHeaderOptions - set the Headers Options
func (c *Client) SetAuthHeaderOptions(headers map[string]string) {
	c.Option.Headers = headers
}

// RestAPICall - general rest method caller
// query is an variadic arg. It receives a slice of map[string]interface{}
func (c *Client) RestAPICall(method Method, path string, options interface{}, query ...map[string]interface{}) ([]byte, error) {
	// log.Debugf("RestAPICall %s - %s%s", method, utils.Sanatize(c.Endpoint), path)

	var (
		Url *url.URL
		err error
		req *http.Request
	)

	Url, err = url.Parse(utils.Sanatize(c.Endpoint))
	// fmt.Println(Url)
	if err != nil {
		return nil, err
	}

	Url.Path += path
	// fmt.Println(c.Endpoint)

	// Manage the query string
	if len(query) != 0 {
		// since query is received as slice, accessing 0th element to get filters
		c.GetQueryStrings(Url, query[0])
	} else {
		c.GetQueryString(Url)
	}

	// log.Debugf("*** url => %s", Url.String())
	// log.Debugf("*** method => %s", method.String())

	// parse url
	reqUrl, err := url.Parse(Url.String())
	// fmt.Println(reqUrl)
	if err != nil {
		return nil, fmt.Errorf("error with request: %v - %q", Url, err)
	}

	// handle options

	if options != nil {
		fmt.Println(reflect.TypeOf(options))
		OptionsJSON, err := json.Marshal(options)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequest(method.String(), reqUrl.String(), bytes.NewBuffer(OptionsJSON))

		// req, err = http.NewRequest(method.String(), reqUrl.String(), bytes.NewBuffer([]byte(fmt.Sprint(options))))
		// fmt.Println(reqUrl.String())
	} else {
		req, err = http.NewRequest(method.String(), reqUrl.String(), nil)
	}

	if err != nil {
		return nil, fmt.Errorf("Error with request: %v - %q", Url, err)
	}

	// setup proxy
	proxyUrl, err := http.ProxyFromEnvironment(req)
	if err != nil {
		return nil, fmt.Errorf("Error with proxy: %v - %q", proxyUrl, err)
	}
	if proxyUrl != nil {
		tr.Proxy = http.ProxyURL(proxyUrl)
		// log.Debugf("*** proxy => %+v", tr.Proxy)
	}

	// build the auth headerU
	// fmt.Println(c.Option.Headers)
	for k, v := range c.Option.Headers {
		// log.Debugf("Headers -> %s -> %+v\n", k, v)
		req.Header.Add(k, v)
	}

	// req.SetBasicAuth(c.User, c.APIKey)
	// req.Method = fmt.Sprintf("%s", method.String())
	// fmt.Println(req.Method)
	// fmt.Println(req.Body)
	// fmt.Println(req.Header)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// TODO: CLeanup Later
	// DEBUGGING WHILE WE WORK
	// DEBUGGING WHILE WE WORK
	// fmt.Printf("METHOD --> %+v\n", method)
	// fmt.Printf("REQ    --> %+v\n", req)
	// fmt.Printf("RESP   --> %+v\n", resp)
	// fmt.Printf("ERROR  --> %+v\n", err)
	// DEBUGGING WHILE WE WORK

	// RESET QUERY PARAMETERS AFTER EVERY CALL
	c.SetQueryString(nil)

	data, err := ioutil.ReadAll(resp.Body)
	if !c.isOkStatus(resp.StatusCode) {
		type apiErr struct {
			Message string `json:"message"`
			Details string `json:"details"`
		}
		var outErr apiErr
		json.Unmarshal(data, &outErr)
		return nil, fmt.Errorf("Error in response: %s\n Response Status: %s\n Response Details: %s", outErr.Message, resp.Status, outErr.Details)
	}

	if err != nil {
		return nil, err
	}

	// Added the condition to accomodate the response where only the response header is returned.
	if len(data) == 0 {
		if resp.Header["Location"] != nil {
			data = []byte(`{"URI":"` + resp.Header["Location"][0] + `"}`)
		}
	}
	// fmt.Println(data)

	return data, nil

}
