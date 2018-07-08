package blawg

import (
	"fmt"
	"html/template"
	"sort"
	"strings"
	"time"
)

// Posts represents a slice of []Post. It is used for default sorting by date
// and adds some methods used in the templates.
type Posts []Post

func (ps Posts) Len() int {
	return len(ps)
}

func (ps Posts) Less(i, j int) bool {
	return ps[i].Date.After(ps[j].Date)
}

func (ps Posts) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Reverse returns a copy of the Posts, sorted into reverse date order (earliest
// first).
func (ps Posts) Reverse() Posts {
	reversedList := make(Posts, len(ps))
	copy(reversedList, ps)
	sort.Sort(sort.Reverse(reversedList))
	return reversedList
}

// Take returns a slice of the first n Posts.
func (ps Posts) Take(n int) Posts {
	return ps[:n]
}

// Drop returns a slice of Posts, without the first n.
func (ps Posts) Drop(n int) Posts {
	return ps[n:]
}

// SortPostsByDate sorts a list of Posts in place by date order.
func SortPostsByDate(ps *Posts) {
	sort.Sort(ps)
}

// Post is a representation of a single blog post.
type Post struct {
	Body      template.HTML
	Title     template.HTML
	Date      time.Time
	TitleText string
	Metadata
}

// Path is a unique file path for a blog post.
func (p *Post) Path() string {
	postPathTitle := lowerKebabCase(p.TitleText)
	postPath := fmt.Sprintf(
		"%d/%d/%d/%s/",
		p.Date.Year(),
		p.Date.Month(),
		p.Date.Day(),
		postPathTitle,
	)
	return postPath
}

func (ps *Posts) sortByDate() {
	sort.Sort(ps)
}

// Metadata represents the metadata for a blog post
type Metadata struct {
	Title      string   `yaml:"title"`
	Layout     string   `yaml:"layout"`
	Date       string   `yaml:"date"`
	Categories []string `yaml:"categories"`
	Published  bool     `yaml:"published"`
}

// PostPage represents the page for a blog post
type PostPage struct {
	Post     *Post
	PostList *Posts
}

type Page struct {
	Body template.HTML
}
type Pages []Page

func lowerKebabCase(s string) string {
	return toKebabCase(strings.ToLower(s))
}

func toKebabCase(s string) string {
	return strings.Replace(s, " ", "-", -1)
}
