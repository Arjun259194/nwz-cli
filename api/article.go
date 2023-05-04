package api

import "fmt"

type Article struct {
	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

func (a *Article) Display() {
	fmt.Printf("\n\nTitle - %v", a.Title)
	fmt.Printf("\nAuthor - %v", a.Author)
	fmt.Printf("\nDescription - %v", a.Description)
	fmt.Printf("\nURL - %v", a.URL)
}
