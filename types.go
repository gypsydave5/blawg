package blawg

import "time"

type Post struct {
	Body []byte
	Date time.Time
	Metadata
}

type Metadata struct {
	Title      string   `yaml:"title"`
	Layout     string   `yaml:"layout"`
	Date       string   `yaml:"date"`
	Categories []string `yaml:"categories"`
}

type Page struct {
	Post Post
	PostList *[]Post
}
