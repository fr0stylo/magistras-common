package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
)

type client struct {
	httpClient *http.Client
	baseUrl    string
}

func NewClient(url string, transport *http.Transport) *client {
	return &client{
		httpClient: &http.Client{
			Transport: transport,
		},
		baseUrl: url,
	}
}

func (c *client) Get(ctx context.Context, url string, response interface{}) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseUrl+url, nil)
	if err != nil {
		return err
	}

	return c.DoRequest(req, response)
}

func (c *client) DoRequest(req *http.Request, response interface{}) error {
	cookie, ok := req.Context().Value("cookie").(*http.Cookie)
	if ok {
		req.AddCookie(cookie)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if response != nil {
		if err := json.NewDecoder(resp.Body).Decode(response); err != nil {
			return err
		}
	}

	return nil
}
