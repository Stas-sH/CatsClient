package client

import (
	"Stas-sH/clietntForCts/internal/responses"
	"Stas-sH/clietntForCts/pkg/cat"
	"encoding/json"

	"io/ioutil"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	return &Client{
		client: http.DefaultClient,
	}
}

func (c *Client) GetCats() ([]cat.Cat, error) {
	resp, err := c.client.Get("https://catfact.ninja/breeds")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respData responses.BreedsResponse
	if err = json.Unmarshal(body, &respData); err != nil {
		return nil, err
	}

	return respData.Data, nil
}
