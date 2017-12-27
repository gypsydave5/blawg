package main

import (
	"github.com/gypsydave5/blawg"
	"log"
)

var siteDirectory = "site"
var templateDirectory = "templates"
var postDirectory = "_posts"

func main() {
	err := blawg.MakeBlog(postDirectory, templateDirectory, siteDirectory)
	if err != nil {
		log.Fatalf("Something went very wrong %s", err)
	}
}
