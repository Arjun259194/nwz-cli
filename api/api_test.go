package api

import (
	"testing"
)

func TestGetTopCountryNews(t *testing.T) {
	apiKey := "a481486cec3745ddae53821bb5132799"
	api := &Api{ApiKey: apiKey}

	country := "us"
	articles := api.GetTopCountryNews(country)

	if len(articles) == 0 {
		t.Errorf("Expected non-empty list of articles, but got empty list")
	}

	// You can add more assertions to check for specific properties of the articles,
	// or to test error cases.
}

