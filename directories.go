package blawg

import (
	"fmt"
	"strings"
	"os"
	"html/template"
)

func(post Post) Path() string {
	postPathTitle := strings.Replace(strings.ToLower(post.Title)," ", "-", -1)
	postPath := fmt.Sprintf("%d/%d/%d/%s/", post.Date.Year(), post.Date.Month(), post.Date.Day(), postPathTitle)
	return postPath
}

func Export(siteDirectory string, post *Post, posts *[]Post) error {
	path := fmt.Sprintf("%s/posts/%s", siteDirectory, post.Path())
	os.MkdirAll(path, os.FileMode(0777))
	fileName := fmt.Sprintf("%sindex.html", path)
	file, err := os.Create(fileName)
	defer file.Close()
	
	return err
}

func ExportAll(siteDirectory string, posts *[]Post) error {
	t, err := template.New("page").ParseFiles("template.html")

	for _, post := range *posts {
		err := Export(siteDirectory, &post, posts)
		if err != nil {
			return err
		}
	}
	return nil
}