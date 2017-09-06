//Package gohvapi provides a simple golang interface to the HostVirtual
//Rest API at https://bapi.vr.org/
package gohvapi

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const (
	Version      = "0.0.1"
	BaseEndpoint = "https://bapi.vr.org/"
	ContentType  = "application/json"
)

//Client is the main object (struct) to which we attach most methods/functions.
//It has the following fields: (client, userAgent, endPoint, apiKey)
type Client struct {
	client    *http.Client
	userAgent string
	endPoint  *url.URL
	apiKey    string
}

//GetKeyFromEnv is a simple function to try to yank the value for "VR_API_KEY" from
//the environment
func GetKeyFromEnv() string {
	return os.Getenv("VR_API_KEY")
}

//NewClient is the main entrypoint for instantiating a Client struct. It takes
//your API Key as it's sole argument and returns the Client struct ready to talk to the API
func NewClient(apikey string) *Client {
	useragent := "gohvapi/" + Version
	transport := &http.Transport{
		TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
	}
	client := http.DefaultClient
	client.Transport = transport
	endpoint, _ := url.Parse(BaseEndpoint)

	return &Client{
		userAgent: useragent,
		client:    client,
		endPoint:  endpoint,
		apiKey:    apikey,
	}
}

//apiPath is just a short internal function for forcing the prepending of / to the url
func apiPath(path string) string {
	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s", path)
	} else {
		return fmt.Sprintf("/%s", path)
	}
}

//apiKeyPath is just a short internal function for appending the key to the url
func apiKeyPath(path, apiKey string) string {
	if strings.Contains(path, "?") {
		return path + "&key=" + apiKey
	}
	return path + "?key=" + apiKey
}

//get method on Client struct for providing the HTTP GET call
func (c *Client) get(path string, data interface{}) error {
	req, err := c.newRequest("GET", apiPath(path), nil)
	if err != nil {
		return err
	}
	return c.do(req, data)
}

//post method on Client struct for providing the HTTP POST call
func (c *Client) post(path string, values []byte, data interface{}) error {

	fmt.Println(string(values))

	req, err := c.newRequest("POST", apiPath(path), bytes.NewBuffer(values))

	if err != nil {
		return err
	}

	return c.do(req, data)
}

//put method on Client struct for providing the HTTP PUT call
func (c *Client) put(path string, values []byte, data interface{}) error {

	fmt.Println(string(values))

	req, err := c.newRequest("PUT", apiPath(path), bytes.NewBuffer(values))

	if err != nil {
		return err
	}
	return c.do(req, data)
}

//patch method on Client struct for providing the HTTP PATCH call
func (c *Client) patch(path string, values url.Values, data interface{}) error {
	req, err := c.newRequest("PATCH", apiPath(path), strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	return c.do(req, data)
}

//delete method on Client struct for providing the HTTP DELETE call
func (c *Client) delete(path string, values url.Values, data interface{}) error {
	req, err := c.newRequest("DELETE", apiPath(path), nil)
	if err != nil {
		return err
	}
	return c.do(req, data)
}

//newRequest method on Client struct for wrapping the get/post/put/patch/delete methods
func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {

	relPath, err := url.Parse(apiKeyPath(path, c.apiKey))

	if err != nil {
		return nil, err

	}

	url := c.endPoint.ResolveReference(relPath)

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err

	}

	req.Header.Add("User-Agent", c.userAgent)
	req.Header.Add("Accept", ContentType)

	return req, nil

}

//do method on Client struct for making the HTTP calls
func (c *Client) do(req *http.Request, data interface{}) error {

	var apiError error

	resp, err := c.client.Do(req)

	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusOK {

		//fmt.Println(string(body))
		if data != nil {
			if err := json.Unmarshal(body, data); err != nil {
				return err
			}
		}
		return nil
	}

	errorCodes := map[string]bool{
		"401": true,
		"500": true,
	}

	if errorCodes[strconv.Itoa(resp.StatusCode)] {

		type Err struct {
			Error struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			} `json:"error"`
		}

		data := &Err{}

		if err := json.Unmarshal(body, data); err != nil {
			return err
		}

		fmt.Println(data.Error.Message)

		apiError = errors.New(string(data.Error.Message))

		return apiError

	}

	apiError = errors.New(string(body))

	return apiError
}
