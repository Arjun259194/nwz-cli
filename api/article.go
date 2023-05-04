package api

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

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
	// Format the output with colors and indentation
	fmt.Printf("%s\n\n", color.GreenString(a.Title))
	fmt.Printf("%s %s\n\n", color.BlueString("Author:"), a.Author)
	fmt.Printf("%s %s\n\n", color.BlueString("Source:"), a.Source.Name)
	fmt.Printf("%s %s\n\n", color.RedString("URL:"), a.URL)
	fmt.Printf("%s\n%s\n\n", color.WhiteString("Content:"), indent(a.Content, 2))
}

// Indent a string by n spaces
func indent(s string, n int) string {
	indentation := strings.Repeat(" ", n)
	return indentation + strings.ReplaceAll(s, "\n", "\n"+indentation)
}
