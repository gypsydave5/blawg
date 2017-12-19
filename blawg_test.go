package blawg

import (
	"html/template"
	"testing"
	"bytes"
	"strings"
)

func TestTemplate(t *testing.T) {
	mainTemplate, _ := template.ParseGlob("testTemplates/*")

	posts := []Post{
		testPost("First Post", "First Post Body", 1979, 12, 5),
		testPost("Second Post", "Second Post Body", 1989, 12, 5),
	}

	page := Page {
		testPost("Main Post Here", "main post body", 1984, 6, 6),
		&posts,
	}

	var b bytes.Buffer
	mainTemplate.ExecuteTemplate(&b, "main", &page)
	s := b.String()

	assertContains(t, s, "<h1>Main Post Here</h1>")
	assertContains(t, s, `<li><a href="/post/1979/12/5/first-post/">First Post</a></li>`)
	assertContains(t, s, `<li><a href="/post/1989/12/5/second-post/">Second Post</a></li>`)
}

func assertContains(t *testing.T, str, substr string) {
	if !strings.Contains(str, substr) {
		t.Errorf(`Expected to find
%s
in
%s
`, substr, str)
		}
}