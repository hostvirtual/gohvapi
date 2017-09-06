package govrapi

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

type Client struct {
	client    *http.Client
	userAgent string
	endPoint  *url.URL
	apiKey    string
}

func GetKeyFromEnv() string {
	return os.Getenv("VR_API_KEY")
}

func NewClient(apikey string) *Client {
	useragent := "govrapi/" + Version
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

func apiPath(path string) string {
	if strings.HasPrefix(path, "/") {
		return fmt.Sprintf("%s", path)
	} else {
		return fmt.Sprintf("/%s", path)
	}
}

func apiKeyPath(path, apiKey string) string {
	if strings.Contains(path, "?") {
		return path + "&key=" + apiKey
	}
	return path + "?key=" + apiKey
}

func (c *Client) get(path string, data interface{}) error {
	req, err := c.newRequest("GET", apiPath(path), nil)
	if err != nil {
		return err
	}
	return c.do(req, data)
}

func (c *Client) post(path string, values []byte, data interface{}) error {

	fmt.Println(string(values))

	req, err := c.newRequest("POST", apiPath(path), bytes.NewBuffer(values))

	if err != nil {
		return err
	}

	return c.do(req, data)
}

func (c *Client) put(path string, values []byte, data interface{}) error {

	fmt.Println(string(values))

	req, err := c.newRequest("PUT", apiPath(path), bytes.NewBuffer(values))

	if err != nil {
		return err
	}
	return c.do(req, data)
}

func (c *Client) patch(path string, values url.Values, data interface{}) error {
	req, err := c.newRequest("PATCH", apiPath(path), strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}
	return c.do(req, data)
}

func (c *Client) delete(path string, values url.Values, data interface{}) error {
	req, err := c.newRequest("DELETE", apiPath(path), nil)
	if err != nil {
		return err
	}
	return c.do(req, data)
}

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
