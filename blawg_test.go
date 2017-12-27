package blawg

import (
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
	assert.FileExists(testSiteDirectory + "/posts/2016/3/28/(even-more)-memoization-in-javascript/index.html")
	assert.FileExists(testSiteDirectory + "/posts/2017/10/21/lambda-calculus-3---logic-with-church-booleans/index.html")

	tearDownTestSite(t)
}
