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

	// SeriesPosts holds every post in this post's series, ordered by
	// Part. It is nil for posts that aren't part of a series, and is
	// populated by AttachSeries after all posts have been loaded.
	SeriesPosts Posts
}

// The Metadata of a Post
type Metadata struct {
	Title       string   `yaml:"title"`
	Layout      string   `yaml:"layout"`
	Date        string   `yaml:"date"`
	Categories  []string `yaml:"tags"`
	Published   bool     `yaml:"published"`
	Description string   `yaml:"description"`
	Series      string   `yaml:"series"`
	Part        int      `yaml:"part"`
}

// SeriesPath is the path of this post's series index page. Only
// meaningful when the post belongs to a series.
func (p *Post) SeriesPath() string {
	return "series/" + urlSafeFileName(p.Series) + "/"
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
