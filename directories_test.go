package blawg

import (
	"time"
	"html/template"
)

func testPost(title, body string, year, month, day int) Post {
	publishTime := time.Date(year, time.Month(month), day, 7, 8, 9, 1, time.Local)
	return Post{
		Body: template.HTML(body),
		Date: publishTime,
		Metadata: Metadata{
			Title: title,
		},
	}
}

func stubTemplate() (mainTemplate *template.Template) {
	mainTemplate, _ = template.New("main").Parse("")
	return
}

func paths(posts []Post) []string {
	var paths []string

	for _, post := range posts {
		paths = append(paths, post.Path())
	}

	return paths
}
