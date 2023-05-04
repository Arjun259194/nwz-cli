package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Api struct {
	ApiKey string
}

func (a *Api) GetTopCountryNews(country string) []Article {
	url := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%v&apiKey=%v", country, a.ApiKey)
	responseByte := getResponseData(url)
	articles := getArticlesFromResponseBytes(responseByte)
	return articles
}

func (a *Api) GetNewsFromKeyWord(keyword string) []Article {
	url := fmt.Sprintf("https://newsapi.org/v2/everything?q=%v&apiKey=%v", keyword, a.ApiKey)
	responseByte := getResponseData(url)
	articles := getArticlesFromResponseBytes(responseByte)
	return articles
}

func getResponseData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Printf("could not create new request - %v", err)
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Not able to read reasponse.Body - %v", err)
	}
	return responseByte
}

func getArticlesFromResponseBytes(byteData []byte) []Article {
	apiRes := &APIResponse{}
	if err := json.Unmarshal(byteData, apiRes); err != nil {
		log.Printf("could not unmarch api response - %v", err)
	}
	if apiRes.Status != "ok" {
		log.Printf("Can't fetch these data")
	}
	articles := apiRes.getArticles()
	return articles
}
