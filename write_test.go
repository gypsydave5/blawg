package blawg

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

var testPostOne = testPost("Post Number One", "", 1066, 5, 22)
var testPostTwo = testPost("Post Number Two", "", 1979, 12, 5)

func TestTemplate(t *testing.T) {
	assert := NewAssertions(t)
	mainTemplate, _ := template.ParseGlob("testTemplates/*")

	mainPost := testPost("Main Post Here", "<p>main post body</p>", 1984, 6, 6)

	posts := Posts{
		testPost("First Post", "First Post Body", 1979, 12, 5),
		testPost("Second Post", "Second Post Body", 1989, 12, 5),
		mainPost,
	}

	var b bytes.Buffer
	err := writePost(&b, &mainPost, &posts, mainTemplate)
	if err != nil {
		t.Error("unexpected error")
	}

	s := b.String()
	assert.StringContains(s, "<h1>Main Post Here</h1>")
	assert.StringContains(s, `<li><a href="/post/1979/12/5/first-post/">First Post</a></li>`)
	assert.StringContains(s, `<li><a href="/post/1989/12/5/second-post/">Second Post</a></li>`)
	assert.StringContains(s, `<p>main post body</p>`)
	assert.StringContains(s, `<time datetime="1984-06-06T07:08" >Jun 06, 1984</time>`)
}

func TestMakePosts(t *testing.T) {
	assert := NewAssertions(t)
	postOne := testPost("Abba", "First Post Body", 1979, 12, 5)
	postTwo := testPost("Second Post", "Second Post Body", 1989, 12, 5)
	posts := Posts{
		postOne,
		postTwo,
	}

	postTemplate, err := template.New("post").Parse(`<p>{{.Post.Title}}</p>"`)
	assert.NotError(err)

	makePosts(testSiteDirectory, &posts, postTemplate)

	for _, post := range posts {
		expectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
		if fileDoesNotExist(expectedFile) {
			t.Error("expected index.html to exist")
		}
		contents, _ := ioutil.ReadFile(expectedFile)
		assert.StringContains(string(contents), string(post.Title))
	}

	tearDownTestSite(t)
}

func TestPostsIndex(t *testing.T) {
	assert := NewAssertions(t)
	postOne := testPost("Abba", "First Post Body", 1979, 12, 5)
	postTwo := testPost("Second Post", "Second Post Body", 1989, 12, 5)
	posts := Posts{
		postOne,
		postTwo,
	}

	indexTemplate, err := template.New("index").Parse(`<p>{{range .}}{{.Title}}{{end}}</p>"`)

	assert.NotError(err)
	err = makePostIndex(testSiteDirectory, &posts, indexTemplate)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}

	if fileDoesNotExist(testSiteDirectory + "/posts/index.html") {
		t.Error("expected index.html to exist")
	}
}

func TestBuildPostPath(t *testing.T) {
	assert := NewAssertions(t)
	post := testPostOne

	expectedPath := "1066/5/22/post-number-one/"
	calculatedPath := post.Path()
	assert.StringsEqual(calculatedPath, expectedPath)
}

func TestMultiplePostPaths(t *testing.T) {
	assert := NewAssertions(t)
	posts := []Post{testPostOne, testPostTwo}

	paths := paths(posts)

	expectedPathOne := "1066/5/22/post-number-one/"
	expectedPathTwo := "1979/12/5/post-number-two/"

	assert.StringsEqual(paths[0], expectedPathOne)
	assert.StringsEqual(paths[1], expectedPathTwo)
}

func TestSavePost(t *testing.T) {
	assert := NewAssertions(t)
	post := testPostOne

	err := os.MkdirAll(testSiteDirectory, os.FileMode(0777))
	assert.NotError(err)

	err = makePost(testSiteDirectory, &post, nil, stubTemplate())
	assert.NotError(err)

	expectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
	if fileDoesNotExist(expectedFile) {
		t.Errorf("expected %s to exist", expectedFile)
	}

	tearDownTestSite(t)
}

func TestNotSavingUnpublishedPost(t *testing.T) {
	assert := NewAssertions(t)
	post := testPost("do not publish me", "", 1901, 1, 1)
	post.Published = false

	err := os.MkdirAll(testSiteDirectory, os.FileMode(0777))
	assert.NotError(err)

	err = makePost(testSiteDirectory, &post, nil, stubTemplate())
	assert.NotError(err)

	unexpectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
	if fileExists(unexpectedFile) {
		t.Errorf("expected %s not to exist", unexpectedFile)
	}

	tearDownTestSite(t)
}

func tearDownTestSite(t *testing.T) {
	err := os.RemoveAll(testSiteDirectory)
	if err != nil {
		t.Errorf("Could not delete test directory: %s", err)
	}
}

func paths(posts []Post) []string {
	var paths []string

	for _, post := range posts {
		paths = append(paths, post.Path())
	}

	return paths
}

func testPost(title, body string, year, month, day int) Post {
	publishTime := time.Date(year, time.Month(month), day, 7, 8, 9, 1, time.Local)
	return Post{
		Body:      template.HTML(body),
		Date:      publishTime,
		Title:     template.HTML(title),
		TitleText: title,
		Metadata: Metadata{
			Title:     title,
			Published: true,
		},
	}
}

func stubTemplate() (mainTemplate *template.Template) {
	mainTemplate, _ = template.New("post").Parse("")
	return
}
