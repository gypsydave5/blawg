package blawg

import (
	"io/ioutil"
	"strings"
	"testing"
)

var testSiteDirectory = "test-site-directory"
var testTemplateDirectory = "example/templates"
var testPostDirectory = "example/_posts"

func TestMakeBlog(t *testing.T) {
	assert := NewAssertions(t)
	err := MakeBlog(testPostDirectory, testTemplateDirectory, testSiteDirectory)

	assert.NotError(err)
	assert.DirectoryExists(testSiteDirectory)
	assert.DirectoryExists(testSiteDirectory + "/posts")
	assert.FileExists(testSiteDirectory + "/index.html")
	assert.FileExists(testSiteDirectory + "/posts/2016/3/28/post-one/index.html")
	assert.FileExists(testSiteDirectory + "/posts/2017/10/21/post-two/index.html")
	assert.FileDoesNotExist(testSiteDirectory + "/posts/1901/1/1/not-to-be-published/index.html")

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
