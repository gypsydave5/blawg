package blawg

import (
	"testing"
	"time"
	"os"
)

var testSiteDirectory = "test-site-directory"

func TestBuildPostPath(t *testing.T) {
	post := testPost()

	expectedPath := "1066/5/22/the-post-title/"
	calculatedPath := post.Path()

	if calculatedPath != expectedPath {
		t.Errorf("Expected '%s' to equal '%s'", expectedPath, calculatedPath)
	}
}


func TestSavePost(t *testing.T) {
	post := testPost()

	err := os.MkdirAll(testSiteDirectory, os.FileMode(0777))
	if err != nil {
		t.Errorf("Could not create the test directory %s", err)
	}

	err = Export(testSiteDirectory, post)
	if err != nil {
		t.Errorf("Could not create the post %s", err)
	}

	expectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
	testFileExists(t, expectedFile)

	teardown(t)
}

func TestExportPosts(t *testing.T) {
	postOne := testPost()

	publishTimeTwo := time.Date(1979, 12, 5, 7, 8, 9, 1, time.Local)
	postTwo := Post{
		Date: publishTimeTwo,
		Metadata: Metadata{
			Title: "Post Number One",
		},
	}

	posts := []Post{postOne, postTwo}

	err := ExportAll(testSiteDirectory, posts)
	if err != nil {
		t.Errorf("Could not create the post %s", err)
	}

	expectedFileOne := testSiteDirectory + "/posts/" + postOne.Path() + "index.html"
	testFileExists(t, expectedFileOne)

	expectedFileTwo := testSiteDirectory + "/posts/" + postTwo.Path() + "index.html"
	testFileExists(t, expectedFileTwo)

	teardown(t)
}

func testFileExists(t *testing.T, pathToFile string) {
	_, err := os.Stat(pathToFile)
	if err != nil {
		t.Errorf("Could not find file: %s", err)
	}
}

func teardown(t *testing.T) {
	err := os.RemoveAll(testSiteDirectory)
	if err != nil {
		t.Errorf("Could not delete test directory: %s", err)
	}
}

func testPost() Post {
	publishTime := time.Date(1066, 5, 22, 7, 8, 9, 1, time.Local)
	post := Post{
		Body: nil,
		Date: publishTime,
		Metadata: Metadata{
			Title: "The Post Title",
		},
	}
	return post
}
