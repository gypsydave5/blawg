package main

import (
	"flag"
	"log"

	"git.sr.ht/~dew/blawg"
)

var (
	siteDirectory      string
	templatesDirectory string
	postsDirectory     string
	extrasDirectory    string
	pagesDirectory     string
	draftsDirectory    string
)

func init() {
	flag.StringVar(&siteDirectory, "site", "site", "directory to write the website to")
	flag.StringVar(&templatesDirectory, "templates", "templates", "directory containing the templates")
	flag.StringVar(&postsDirectory, "posts", "posts", "directory containing the blog posts")
	flag.StringVar(&extrasDirectory, "extras", "extras", "directory containing the templates")
	flag.StringVar(&pagesDirectory, "pages", "pages", "directory containing the templates")
	flag.StringVar(&draftsDirectory, "drafts", "drafts", "directory containing draft posts")
}

func main() {
	flag.Parse()
	err := blawg.MakeBlawg(postsDirectory, pagesDirectory, templatesDirectory, extrasDirectory, siteDirectory)
	if err != nil {
		log.Fatalf("%s", err)
	}
	err = blawg.MakeDrafts(draftsDirectory, templatesDirectory, siteDirectory)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
