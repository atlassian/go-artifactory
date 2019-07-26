package client

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

const (
	userAgent = "go-artifactory v3"

	MediaTypePlain = "text/plain"
	MediaTypeXml   = "application/xml"

	MediaTypeJson = "application/json"
	MediaTypeForm = "application/x-www-form-urlencoded"
)

type Client struct {
	// HTTP Client used to communicate with the API.
	Client *resty.Client

	// Base URL for API requests. BaseURL should always be specified with a trailing slash.
	BaseURL *url.URL
}

// NewClient creates a Client from a provided base url for an artifactory instance and a service Client
func NewClient(baseUrl string, httpClient *http.Client) *resty.Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if !strings.HasSuffix(baseUrl, "/") {
		baseUrl += "/"
	}

	return resty.NewWithClient(httpClient).
		SetHostURL(baseUrl).
		SetHeader("User-Agent", userAgent)
}
