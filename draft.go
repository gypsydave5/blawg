package blawg

import (
	"html/template"
	"time"
)

// A Draft post - publicly accessible by URL but excluded from homepage, post index, and RSS feed.
type Draft struct {
	Body      template.HTML
	Title     template.HTML
	TitleText string
	Date      time.Time
	Metadata
}

// Path of a draft post. URL safe.
func (d *Draft) Path() string {
	return "drafts/" + urlSafeFileName(d.TitleText) + "/"
}

// DraftPage is the template data for rendering a draft post.
type DraftPage struct {
	Draft *Draft
}
