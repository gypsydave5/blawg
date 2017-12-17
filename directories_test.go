package blawg

import (
	"testing"
	"time"
)

func TestBuildPostPath(t *testing.T) {
	publishTime := time.Date(1066, 5, 22, 7, 8, 9, 1, time.Local)

	post := Post{
		Body: nil,
		Date: publishTime,
		Metadata: Metadata{
			Title:      "The Post Title",
		},
	}

	expectedPath := "1066/5/22/the-post-title/"
	calculatedPath := post.Path()

	if calculatedPath != expectedPath {
		t.Errorf("Expected '%s' to equal '%s'", expectedPath, calculatedPath)
	}
}
