package blawg

import (
	"testing"
	"strings"
	"fmt"
)

func TestSidebar(t *testing.T) {
	postOne := testPostOne()
	postTwo := testPostTwo()
	posts := []Post{postOne, postTwo}

	var sidebar string = newSidebar(posts)

	htmlLinkFormatString := `<a href="%s">%s</a>`
	postOneLink := fmt.Sprintf(htmlLinkFormatString, "posts/" + postOne.Path(), postOne.Title)
	postTwoLink := fmt.Sprintf(htmlLinkFormatString, "posts/" + postTwo.Path(), postTwo.Title)

	if !strings.Contains(sidebar, postOneLink) {
		t.Errorf(`Expected to find "%s" in "%s"`, postOneLink, sidebar)
	}

	if !strings.Contains(sidebar, postTwoLink) {
		t.Errorf(`Expected to find "%s" in "%s"`, postOneLink, sidebar)
	}
}

func newSidebar(posts []Post) string {
	return ""
}
