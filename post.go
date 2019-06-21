package blawg

import (
	"fmt"
	"html/template"
	"time"
)

// A Post for the blog
type Post struct {
	Body      template.HTML
	Title     template.HTML
	Date      time.Time
	TitleText string
	Metadata
}

// The Metadata of a Post
type Metadata struct {
	Title       string   `yaml:"title"`
	Layout      string   `yaml:"layout"`
	Date        string   `yaml:"date"`
	Categories  []string `yaml:"tags"`
	Published   bool     `yaml:"published"`
	Description string   `yaml:"description"`
}

// Path of a post
func (p *Post) Path() string {
	postPathTitle := urlSafeFileName(p.TitleText)
	postPath := fmt.Sprintf(
		"%d/%d/%d/%s/",
		p.Date.Year(),
		p.Date.Month(),
		p.Date.Day(),
		postPathTitle,
	)
	return postPath
}

// A PostPage is the page with a post on
type PostPage struct {
	Post     *Post
	PostList *Posts
}
