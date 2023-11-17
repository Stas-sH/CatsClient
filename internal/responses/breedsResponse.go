package responses

import (
	"Stas-sH/clietntForCts/pkg/cat"
	"Stas-sH/clietntForCts/pkg/link"
)

type BreedsResponse struct {
	Current_page   int         `json:"current_page"`
	Data           []cat.Cat   `json:"data"`
	First_page_url string      `json:"first_page_url"`
	From           int         `json:"from"`
	Last_page      int         `json:"last_page"`
	Last_page_url  string      `json:"last_page_url"`
	Links          []link.Link `json:"links"`
	Next_page_url  string      `json:"next_page_url"`
	Path           string      `json:"path"`
	Per_page       int         `json:"per_page"`
	Prev_page_url  string      `json:"prev_page_url"`
	To             int         `json:"to"`
	Total          int         `json:"total"`
}
