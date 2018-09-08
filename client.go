package cha

import (
	"log"
	"net/http"
	"net/url"
	"os"
)

const CHATWORK_API = "https://api.chatwork.com/v2/"

type Client struct {
	APIKey     string
	BaseURL    *url.URL
	Debug      bool
	HTTPClient *http.Client
	Logger     *log.Logger
}

func NewClient(apiKey string) (*Client, error) {
	baseUrl, err := url.Parse(CHATWORK_API)
	if err != nil {
		return nil, err
	}

	return &Client{
		APIKey:     apiKey,
		Debug:      false,
		BaseURL:    baseUrl,
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "", log.LstdFlags),
	}, nil
}

func (c *Client) SetBaseURL(baseURL string) (err error) {
	if c.BaseURL, err = url.Parse(baseURL); err != nil {
		return
	}
	return nil
}
