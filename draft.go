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

	// SeriesPosts is always nil for drafts; it exists so the shared
	// "article" template can be used for both Posts and Drafts.
	SeriesPosts Posts
}

// Path of a draft post. URL safe.
func (d *Draft) Path() string {
	return "drafts/" + urlSafeFileName(d.TitleText) + "/"
}

// SeriesPath is the path of this draft's series index page. Only
// meaningful when the draft belongs to a series.
func (d *Draft) SeriesPath() string {
	return "series/" + urlSafeFileName(d.Series) + "/"
}

// DraftPage is the template data for rendering a draft post.
type DraftPage struct {
	Draft *Draft
}
