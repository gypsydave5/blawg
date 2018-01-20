package blawg

import (
	"io/ioutil"
	"strings"
	"testing"
)

var testSiteDirectory = "test-site-directory"
var testTemplateDirectory = "example/templates"
var testExtrasDirectory = "example/extras"
var testPostDirectory = "example/posts"

func TestMakeBlog(t *testing.T) {
	assert := NewAssertions(t)
	err := MakeBlawg(testPostDirectory, testTemplateDirectory, testExtrasDirectory, testSiteDirectory)

	assert.NotError(err)
	directoryExists(testSiteDirectory)
	directoryExists(testSiteDirectory + "/posts")
	directoryExists(testSiteDirectory + "/css")

	fileExists(testSiteDirectory + "/index.html")

	if fileDoesNotExist(testSiteDirectory + "/about.html") {
		t.Errorf("Expected about.html to exist")
	}
	if fileDoesNotExist(testSiteDirectory + "/public.txt") {
		t.Error("Expected public.txt to exist")
	}
	if fileDoesNotExist(testSiteDirectory + "/css/styles.css") {
		t.Error("Expected /css/styles.css to exist")
	}
	if fileDoesNotExist(testSiteDirectory + "/posts/index.html") {
		t.Errorf("expected post index to exist")
	}
	if fileDoesNotExist(testSiteDirectory + "/posts/2016/3/28/post-one/index.html") {
		t.Errorf("expected post one to exist")
	}
	if fileDoesNotExist(testSiteDirectory + "/posts/2017/10/21/post-two/index.html") {
		t.Errorf("expected post two to exist")
	}
	if fileExists(testSiteDirectory + "/posts/1901/1/1/not-to-be-published/index.html") {
		t.Errorf("expected not-to-be-published not to exist")
	}

	file, err := ioutil.ReadFile(testSiteDirectory + "/posts/2017/10/21/post-two/index.html")
	assert.NotError(err)
	post := string(file)

	split := strings.Split(post, "should be first")
	if !strings.Contains(split[0], "post one") {
		t.Errorf("post one appears too soon")
	}
	if strings.Contains(split[1], "post one") {
		t.Errorf("post one not in second half of list")
	}

	tearDownTestSite(t)
}
