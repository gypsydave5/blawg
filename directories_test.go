package blawg

import (
	"testing"
	"time"
	"os"
	"html/template"
)

var testSiteDirectory = "test-site-directory"

func TestBuildPostPath(t *testing.T) {
	post := testPostOne()

	expectedPath := "1066/5/22/post-number-one/"
	calculatedPath := post.Path()

	if calculatedPath != expectedPath {
		t.Errorf("Expected '%s' to equal '%s'", expectedPath, calculatedPath)
	}
}

func TestMultiplePostPaths(t *testing.T) {
	posts := []Post{testPostOne(), testPostTwo()}

	paths := paths(posts)

	expectedPathOne := "1066/5/22/post-number-one/"
	expectedPathTwo := "1979/12/5/post-number-two/"

	if paths[0] != expectedPathOne {
		t.Errorf("Expected %s, got %s", expectedPathOne, paths[0])
	}

	if paths[1] != expectedPathTwo {
		t.Errorf("Expected %s, got %s", expectedPathTwo, paths[1])
	}
}

func paths(posts []Post) []string {
	var paths []string

	for _, post := range posts {
		paths = append(paths, post.Path())
	}

	return paths
}

func TestSavePost(t *testing.T) {
	post := testPostOne()

	err := os.MkdirAll(testSiteDirectory, os.FileMode(0777))
	if err != nil {
		t.Errorf("Could not create the test directory %s", err)
	}

	err = Export(testSiteDirectory, &post, nil, stubTemplate())
	if err != nil {
		t.Errorf("Could not create the post %s", err)
	}

	expectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
	testFileExists(t, expectedFile)

	tearDown(t)
}

func TestExportPosts(t *testing.T) {
	postOne := testPostOne()

	postTwo := testPostTwo()

	posts := []Post{postOne, postTwo}

	err := ExportAll(testSiteDirectory, &posts, stubTemplate())
	if err != nil {
		t.Errorf("Could not create the post %s", err)
	}

	expectedFileOne := testSiteDirectory + "/posts/" + postOne.Path() + "index.html"
	testFileExists(t, expectedFileOne)

	expectedFileTwo := testSiteDirectory + "/posts/" + postTwo.Path() + "index.html"
	testFileExists(t, expectedFileTwo)

	tearDown(t)
}

func testFileExists(t *testing.T, pathToFile string) {
	_, err := os.Stat(pathToFile)
	if err != nil {
		t.Errorf("Could not find file: %s", err)
	}
}

func tearDown(t *testing.T) {
	err := os.RemoveAll(testSiteDirectory)
	if err != nil {
		t.Errorf("Could not delete test directory: %s", err)
	}
}

func testPost(title, body string, year, month, day int) Post {
	publishTime := time.Date(year, time.Month(month), day, 7, 8, 9, 1, time.Local)
	return Post{
		Body: template.HTML(body),
		Date: publishTime,
		Metadata: Metadata{
			Title: title,
		},
	}
}

func testPostOne() Post {
	publishTime := time.Date(1066, 5, 22, 7, 8, 9, 1, time.Local)
	post := Post{
		Date: publishTime,
		Metadata: Metadata{
			Title: "Post Number One",
		},
	}
	return post
}

func testPostTwo() Post {
	publishTimeTwo := time.Date(1979, 12, 5, 7, 8, 9, 1, time.Local)
	postTwo := Post{
		Date: publishTimeTwo,
		Metadata: Metadata{
			Title: "Post Number Two",
		},
	}
	return postTwo
}

func stubTemplate() (mainTemplate *template.Template) {
	mainTemplate, _ = template.New("main").Parse("")
	return
}

