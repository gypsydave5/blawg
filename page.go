package blawg

import (
	"html/template"
	"strings"
	"unicode"
)

// A Page on the blog
type Page struct {
	Body      template.HTML
	Title     template.HTML
	TitleText string
}

// Path of a blog Page. URL safe.
func (p *Page) Path() string {
	return urlSafeFileName(p.TitleText)
}

func urlSafeFileName(s string) string {
	removeUnsafeRunes := func(r rune) rune {
		switch {
		case r >= '0' && r <= '9':
			return r
		case r >= 'a' && r <= 'z':
			return r
		case r >= 'A' && r <= 'Z':
			return unicode.ToLower(r)
		case r == ' ':
			return '-'
		case strings.ContainsRune(".~_-", r):
			return r
		}
		return rune(-1)
	}

	return strings.Map(removeUnsafeRunes, s)
}
