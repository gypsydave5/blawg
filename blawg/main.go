package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gypsydave5/blawg"
)

var (
	siteDirectory      string
	templatesDirectory string
	postsDirectory     string
	extrasDirectory    string
	pagesDirectory     string
)

func init() {
	flag.StringVar(&siteDirectory, "site", "site", "directory to write the website to")
	flag.StringVar(&templatesDirectory, "templates", "templates", "directory containing the templates")
	flag.StringVar(&postsDirectory, "posts", "posts", "directory containing the blog posts")
	flag.StringVar(&extrasDirectory, "extras", "extras", "directory containing the templates")
	flag.StringVar(&pagesDirectory, "pages", "pages", "directory containing the templates")
}

func main() {
	flag.Parse()
	err := blawg.MakeBlawg(postsDirectory, pagesDirectory, templatesDirectory, extrasDirectory, siteDirectory)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
