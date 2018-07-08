package blawg

import (
	"io/ioutil"
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
	assert.DirectoryExists(testSiteDirectory)
	assert.DirectoryExists(testSiteDirectory + "/posts")
	assert.DirectoryExists(testSiteDirectory + "/css")
	assert.FileExists(testSiteDirectory + "/index.html")

	assert.FileExists(testSiteDirectory + "/public.txt")
	assert.FileExists(testSiteDirectory + "/css/styles.css")
	assert.FileExists(testSiteDirectory + "/posts/index.html")
	assert.FileExists(testSiteDirectory + "/posts/2016/3/28/post-one/index.html")
	assert.FileExists(testSiteDirectory + "/posts/2017/10/21/post-two/index.html")
	assert.FileDoesntExist(testSiteDirectory + "/posts/1901/1/1/not-to-be-published/index.html")

	file, err := ioutil.ReadFile(testSiteDirectory + "/posts/2017/10/21/post-two/index.html")
	assert.NotError(err)

	post := string(file)
	assert.StringContainsInOrder(post, "post two", "post one")

	tearDownTestSite(t)
}
