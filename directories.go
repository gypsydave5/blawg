package blawg

import (
	"fmt"
	"strings"
	"os"
)

func(post Post) Path() string {
	postPathTitle := strings.Replace(strings.ToLower(post.Title)," ", "-", -1)
	postPath := fmt.Sprintf("%d/%d/%d/%s/", post.Date.Year(), post.Date.Month(), post.Date.Day(), postPathTitle)
	return postPath
}

func Export(siteDirectory string, post Post) error {
	path := fmt.Sprintf("%s/posts/%s", siteDirectory, post.Path())
	os.MkdirAll(path, os.FileMode(0777))
	fileName := fmt.Sprintf("%sindex.html", path)
	file, err := os.Create(fileName)
	defer file.Close()
	return err
}