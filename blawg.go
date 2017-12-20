package blawg

import (
	"html/template"
	"io"
)

func WritePost(w io.Writer, post *Post, posts *[]Post, template *template.Template) error {
	page := Page {
		post,
		posts,
	}
	err := template.ExecuteTemplate(w, "main", &page)
	return err
}
