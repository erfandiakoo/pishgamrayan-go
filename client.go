package pishgamrayan

import (
	"bytes"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const (
	apiBaseURL = "https://smsapi.pishgamrayan.com/"
)

type Return struct {
	StatusCode     int    `json:"statusCode"`
	MessageId      string `json:"messageId"`
	BlackListCount int    `json:"blackListCount"`
}

type ReturnError struct {
	*Return `json:"return"`
}

type Client struct {
	BaseClient *http.Client
	apikey     string
	BaseURL    *url.URL
}

func NewClient(apikey string) *Client {
	baseURL, _ := url.Parse(apiBaseURL)
	c := &Client{
		BaseClient: http.DefaultClient,
		BaseURL:    baseURL,
		apikey:     apikey,
	}
	return c
}

func (c *Client) EndPoint(parts ...string) *url.URL {
	u, _ := url.Parse(strings.Join(parts, "/"))

	log.Println(u)

	return u
}

func (c *Client) Execute(urlStr string, b Message, v interface{}) error {
	jsonBytes, err := json.Marshal(b)
	if err != nil {
		return err
	}

	body := bytes.NewReader(jsonBytes)

	ul, _ := url.Parse(urlStr)
	u := c.BaseURL.ResolveReference(ul)
	req, _ := http.NewRequest("POST", u.String(), body)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Charset", "utf-8")
	req.Header.Add("Authorization", c.apikey)

	resp, err := c.BaseClient.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok {
			return netErr
		}
		if resp == nil {
			return &HTTPError{
				Status:  http.StatusInternalServerError,
				Message: "nil api response",
				Err:     err,
			}
		}
		return &HTTPError{
			Status:  resp.StatusCode,
			Message: resp.Status,
			Err:     err,
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		re := new(ReturnError)
		err = json.NewDecoder(resp.Body).Decode(&re)
		if err != nil {
			return &HTTPError{
				Status:  resp.StatusCode,
				Message: resp.Status,
			}
		}
		return &APIError{
			Status: re.Return.StatusCode,
		}
	}

	return json.NewDecoder(resp.Body).Decode(&v)
}
