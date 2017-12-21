package blawg

import (
	"html/template"
	"testing"
	"bytes"
	"strings"
	"io/ioutil"
	"os"
)

var testSiteDirectory = "test-site-directory"
var testPostOne = testPost("Post Number One", "", 1066, 5, 22)
var testPostTwo = testPost("Post Number Two", "", 1979, 12, 5)

func TestTemplate(t *testing.T) {
	assert := Assertions{t}
	mainTemplate, _ := template.ParseGlob("testTemplates/*")

	mainPost := testPost("Main Post Here", "<p>main post body</p>", 1984, 6, 6)

	posts := []Post{
		testPost("First Post", "First Post Body", 1979, 12, 5),
		testPost("Second Post", "Second Post Body", 1989, 12, 5),
		mainPost,
	}

	var b bytes.Buffer
	err := WritePost(&b, &mainPost, &posts, mainTemplate)
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
	assert := Assertions{t}
	postOne := testPost("Abba", "First Post Body", 1979, 12, 5)
	postTwo := testPost("Second Post", "Second Post Body", 1989, 12, 5)
	posts := []Post{
		postOne,
		postTwo,
	}

	mainTemplate, err := template.New("main").Parse(`<p>{{.Post.Title}}</p>"`)
	assert.NotError(err)

	MakePosts(testSiteDirectory, &posts, mainTemplate)

	for _, post := range posts {
		expectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
		assert.FileExists(expectedFile)
		contents, _ := ioutil.ReadFile(expectedFile)

		assert.StringContains(string(contents), post.Title)
	}

	tearDown(t)
}

func TestBuildPostPath(t *testing.T) {
	assert := Assertions{t}
	post := testPostOne

	expectedPath := "1066/5/22/post-number-one/"
	calculatedPath := post.Path()
	assert.StringsEqual(calculatedPath, expectedPath)
}

func TestMultiplePostPaths(t *testing.T) {
	assert := Assertions{t}
	posts := []Post{testPostOne, testPostTwo}

	paths := paths(posts)

	expectedPathOne := "1066/5/22/post-number-one/"
	expectedPathTwo := "1979/12/5/post-number-two/"

	assert.StringsEqual(paths[0], expectedPathOne)
	assert.StringsEqual(paths[1], expectedPathTwo)
}

func TestSavePost(t *testing.T) {
	assert := Assertions{t}
	post := testPostOne

	err := os.MkdirAll(testSiteDirectory, os.FileMode(0777))
	assert.NotError(err)

	err = Export(testSiteDirectory, &post, nil, stubTemplate())
	assert.NotError(err)

	expectedFile := testSiteDirectory + "/posts/" + post.Path() + "index.html"
	assert.FileExists(expectedFile)

	tearDown(t)
}

func tearDown(t *testing.T) {
	err := os.RemoveAll(testSiteDirectory)
	if err != nil {
		t.Errorf("Could not delete test directory: %s", err)
	}
}

type Assertions struct {
	test *testing.T
}

func (a Assertions) StringContains(str, substr string) {
	if !strings.Contains(str, substr) {
		a.test.Errorf(`Expected to find
%s
in
%s
`, substr, str)
	}
}

func (a Assertions) StringsEqual(str1, str2 string) {
	if str1 != str2 {
		a.test.Errorf("Expected '%s' to equal '%s'", str1, str2)
	}
}

func (a Assertions) NotError(err error) {
	if err != nil {
		a.test.Errorf("unexpected error: %s", err)
	}
}

func (a Assertions) FileExists(pathToFile string) {
	_, err := os.Stat(pathToFile)
	if err != nil {
		a.test.Errorf("Could not find file: %s", err)
	}
}
