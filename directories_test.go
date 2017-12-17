package blawg

import (
	"testing"
	"time"
	"os"
	"fmt"
)

func TestBuildPostPath(t *testing.T) {
	publishTime := time.Date(1066, 5, 22, 7, 8, 9, 1, time.Local)

	post := Post{
		Body: nil,
		Date: publishTime,
		Metadata: Metadata{
			Title: "The Post Title",
		},
	}

	expectedPath := "1066/5/22/the-post-title/"
	calculatedPath := post.Path()

	if calculatedPath != expectedPath {
		t.Errorf("Expected '%s' to equal '%s'", expectedPath, calculatedPath)
	}
}

func TestSavePost(t *testing.T) {
	publishTime := time.Date(1066, 5, 22, 7, 8, 9, 1, time.Local)

	post := Post{
		Body: nil,
		Date: publishTime,
		Metadata: Metadata{
			Title: "The Post Title",
		},
	}

	testSiteDirectory := "test-site-directory"

	err := os.MkdirAll(testSiteDirectory, os.FileMode(0777))
	if err != nil {
		t.Errorf("Could not create the test directory %s", err)
	}

	err = Export(testSiteDirectory, post)
	if err != nil {
		t.Errorf("Could not create the post %s", err)
	}

	expectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
	fmt.Println(expectedFile)

	_, err = os.Stat(expectedFile)
	if err != nil {
		t.Errorf("Could not find generated post file: %s", err)
	}

	err = os.RemoveAll(testSiteDirectory)
	if err != nil {
		t.Errorf("Could not delete test directory", err)
	}
}
