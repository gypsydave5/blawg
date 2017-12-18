package blawg

import (
	"testing"
	"strings"
	"fmt"
	"html/template"
	"bytes"
)

func TestPostListHTML(t *testing.T) {
	postOne := testPostOne()
	postTwo := testPostTwo()
	posts := []Post{postOne, postTwo}

	var postList string = postListHTML(posts)

	htmlLinkFormatString := `<li><a href="%s">%s</a></li>`
	postOneLink := fmt.Sprintf(htmlLinkFormatString, "posts/" + postOne.Path(), postOne.Title)
	postTwoLink := fmt.Sprintf(htmlLinkFormatString, "posts/" + postTwo.Path(), postTwo.Title)

	if !strings.Contains(postList, postOneLink) {
		t.Errorf(`Expected to find "%s" in "%s"`, postOneLink, postList)
	}

	if !strings.Contains(postList, postTwoLink) {
		t.Errorf(`Expected to find "%s" in "%s"`, postOneLink, postList)
	}

	if !strings.HasPrefix(postList, "<ul>") {
		t.Errorf(`Expected to find "%s" to start with "<ul>"`, postList)
	}

	if !strings.HasSuffix(postList, "</ul>") {
		t.Errorf(`Expected to find "%s" to end with "</ul>"`, postList)
	}
}

func postListHTML(posts []Post) string {
	const listTemplate = `<ul>
		{{range .}}<li><a href="posts/{{.Path}}">{{.Title}}</a></li>{{end}}
</ul>`

	t, _ := template.New("list").Parse(listTemplate)

	var b bytes.Buffer

	t.Execute(&b, posts)

	return b.String()
}
