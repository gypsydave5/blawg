package main

import (
	"github.com/gypsydave5/blawg"
	"log"
)

var siteDirectory = "site"
var templateDirectory = "templates"
var postDirectory = "_posts"
var extrasDirectory = "extras"

func main() {
	err := blawg.MakeBlog(postDirectory, templateDirectory, extrasDirectory, siteDirectory)
	if err != nil {
		log.Fatalf("%s", err)
	}
}
