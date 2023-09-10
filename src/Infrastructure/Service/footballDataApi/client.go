package footballdataapi

import (
	"io"
	"net/http"
	"os"
)

const baseUrl = "https://api.football-data.org/v4/"

// const baseUrl = "http://demo6280197.mockable.io"

type httpClient struct {
	client  http.Client
	baseUrl string
}

// Create new api client.
func NewClient() *httpClient {
	return &httpClient{
		client:  http.Client{},
		baseUrl: "https://api.football-data.org/v4/",
	}
}

// Create new api client.
func NewClientMock() *httpClient {
	return &httpClient{
		client:  http.Client{},
		baseUrl: "http://demo6280197.mockable.io/",
	}
}

func (c *httpClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", c.baseUrl+url, nil)
	if err != nil {
		return nil, err

	}

	return c.Do(req)
}

func (c *httpClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", c.baseUrl+url, body)
	if err != nil {
		return nil, err

	}

	req.Header.Set("Content-Type", contentType)

	return c.Do(req)
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Auth-Token", os.Getenv("API_AUTH_TOKEN"))

	return c.client.Do(req)
}
