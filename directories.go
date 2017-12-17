package blawg

import (
	"fmt"
	"strings"
)

func(post Post) Path() string {
	postPathTitle := strings.Replace(strings.ToLower(post.Title)," ", "-", -1)
	postPath := fmt.Sprintf("%d/%d/%d/%s/", post.Date.Year(), post.Date.Month(), post.Date.Day(), postPathTitle)
	return postPath
}