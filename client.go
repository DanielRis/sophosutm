package sophosutm

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	cleanhttp "github.com/DanielRis/go-cleanhttp"
)

type Client struct {
	ServerURL  string // Location of SophosUTM to use
	APIUser    string // REST API User Name
	APIKey     string // REST API authentication key
	APIVersion string // API version
	HTTP       *http.Client
}

// NewClient returns a new SophosUTM Client
func NewClient(serverURL string, apiUser string, apiKey string) (*Client, error) {
	client := Client{
		ServerURL: serverURL,
		APIUser:   apiUser,
		APIKey:    apiKey,
		HTTP:      cleanhttp.DefaultClient(),
	}
	return &client, nil
}

func (c *Client) newRequest(method string, endpoint string, body []byte) (*http.Request, error) {
	var urlStr string
	urlStr = c.ServerURL + "/api" + endpoint
	url, err := url.Parse(urlStr)

	if err != nil {
		return nil, fmt.Errorf("Error during parsing request URL: %s", err)
	}
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req, err := http.NewRequest(method, url.String(), bodyReader)
	if err != nil {
		return nil, fmt.Errorf("Error during creation of request: %s", err)
	}
	req.SetBasicAuth(c.APIUser, c.APIKey)
	req.Header.Add("Accept", "application/json")

	if method != "GET" {
		req.Header.Add("Content-Type", "application/json")
	}
	if method == "DELETE" {
		req.Header.Add("X-Restd-Err-Ack", "all")
	}

	return req, nil
}

// GetObject gets a object as a json string
func (c *Client) GetObject(endpoint string) (string, error) {
	req, err := c.newRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	/*
		// Debug output to display HTTP response
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return "", err
		}
		log.Print("[Info] Response: \n", string(b))
	*/
	if resp.StatusCode >= 400 {
		log.Printf("[Error] %s", resp.Body)
		return "", fmt.Errorf("%s", resp.Body)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()
	return s, nil
}

// CreateObject creates a object and returns a json string
func (c *Client) CreateObject(endpoint string, objectAttributes []byte) (string, error) {

	req, err := c.newRequest("POST", endpoint, objectAttributes)
	if err != nil {
		return "", err
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	/*
		// Debug output to display HTTP response
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return "", err
		}
		log.Print("[Info] Response: \n", string(b))
	*/

	if resp.StatusCode >= 400 {
		log.Printf("[Error] %s", resp.Body)
		return "", fmt.Errorf("%s", resp.Body)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()
	return s, nil
}

// DeleteObject deletes an object and retuturns a json string
func (c *Client) DeleteObject(endpoint string) (string, error) {
	req, err := c.newRequest("DELETE", endpoint, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	/*
		// Debug output to display HTTP response
		b, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return "", err
		}
		log.Print("[Info] Response: \n", string(b))
	*/

	if resp.StatusCode >= 400 {
		log.Printf("[Error] %s", resp.Body)
		return "", fmt.Errorf("%s", resp.Body)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	s := buf.String()
	return s, nil
}
