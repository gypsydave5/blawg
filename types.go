package blawg

import (
	"fmt"
	"html/template"
	"sort"
	"strings"
	"time"
)

type Posts []Post

func (ps Posts) Len() int {
	return len(ps)
}

func (ps Posts) Less(i, j int) bool {
	return ps[i].Date.Before(ps[j].Date)
}

func (ps Posts) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type Post struct {
	Body template.HTML
	Date time.Time
	TitleText string
	Metadata
}

func (p *Post) Path() string {
	postPathTitle := lowerKebabCase(p.Title)
	postPath := fmt.Sprintf(
		"%d/%d/%d/%s/",
		p.Date.Year(),
		p.Date.Month(),
		p.Date.Day(),
		postPathTitle,
	)
	return postPath
}

func (ps *Posts) SortByDate() {
	sort.Sort(ps)
}

type Metadata struct {
	Title      string   `yaml:"title"`
	Layout     string   `yaml:"layout"`
	Date       string   `yaml:"date"`
	Categories []string `yaml:"categories"`
	Published  bool     `yaml:"published"`
}

type Page struct {
	Post     *Post
	PostList *[]Post
}

func lowerKebabCase(s string) string {
	return toKebabCase(strings.ToLower(s))
}

func toKebabCase(s string) string {
	return strings.Replace(s, " ", "-", -1)
}
