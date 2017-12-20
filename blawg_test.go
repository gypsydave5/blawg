package blawg

import (
	"html/template"
	"testing"
	"bytes"
	"strings"
	"io/ioutil"
)

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
	assert.stringContains(s, "<h1>Main Post Here</h1>")
	assert.stringContains(s, `<li><a href="/post/1979/12/5/first-post/">First Post</a></li>`)
	assert.stringContains(s, `<li><a href="/post/1989/12/5/second-post/">Second Post</a></li>`)
	assert.stringContains(s, `<p>main post body</p>`)
	assert.stringContains(s, `<time datetime="1984-06-06T07:08" >Jun 06, 1984</time>`)
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
	assert.notError(err)

	MakePosts(testSiteDirectory, &posts, mainTemplate)

	expectedFileOne := testSiteDirectory + "/posts/" + postOne.Path() + "index.html"
	testFileExists(t, expectedFileOne)

	contents, _ := ioutil.ReadFile(expectedFileOne)

	assert.stringContains(string(contents), postOne.Title)

	tearDown(t)
}

func MakePosts(siteDirectory string, posts *[]Post, tmplt *template.Template) (err error) {
	for _, post := range *posts {
		err = Export(siteDirectory, &post, posts, tmplt)
		if err != nil {
			return
		}
	}
	return
}

type Assertions struct{
	test *testing.T
}

func (a Assertions) stringContains(str, substr string) {
	if !strings.Contains(str, substr) {
		a.test.Errorf(`Expected to find
%s
in
%s
`, substr, str)
	}
}

func (a Assertions) notError(err error) {
	if err != nil {
		a.test.Errorf("unexpected error: %s", err)
	}
}