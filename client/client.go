package client

import (
	"log"
	"net/http"
	"net/url"
)

const CHATWORK_API = "https://api.chatwork.com/v2/"

type HTTPClient struct {
	Get()
	Post()
	Put()
	Delete()
}

type Client struct {
	APIKey     string
	BaseURL    *url.URL
	Debug      bool
	HTTPClient *http.Client
	Logger     *log.Logger
	HTTPClient
}

func NewClient(apiKey string) (*Client, error) {
	baseUrl, err := url.Parse(CHATWORK_API)
	if err != nil {
		return nil, err
	}

	return &Client{
		APIKey:     "",
		Debug:      false,
		BaseURL:    baseUrl,
		HTTPClient: http.DefaultClient,
	}, nil
}

func (c *Client) SetBaseURL(baseURL string) (err error) {
	if c.BaseURL, err = url.Parse(baseURL); err != nil {
		return
	}
	return nil
}
