package api

import (
	"encoding/json"
	"log"
)

type APIResponse struct {
	Status       string          `json:"status"`
	TotalResults int             `json:"totalResults"`
	Articles     json.RawMessage `json:"articles"`
}

func (apres *APIResponse) getArticles() []Article {
	articles := []Article{}
	if err := json.Unmarshal(apres.Articles, &articles); err != nil {
		log.Printf("Could not unmarsh Articles - %v", err)
	}
	return articles
}
