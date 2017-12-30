package blawg

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"time"
)

var rawPost = `---
layout: post
title: "example post"
date: 2016-10-15 23:24:01
categories:
    - category1
    - category2
published: true
---
This is the body of the post

## A sub header
`

var badMetadata = `---
layout: post
title: "example post"
date: 2016-10-15T23:24:01
categories:
    - category1
    - category2
published: true
---
This is the body of the post

## A sub header
`

func TestSplitNoMeta(t *testing.T) {
	_, err := Parse(strings.NewReader(`no meta block here!`))

	if err.Error() != "no metadata block" {
		t.Error("did not get the expected error: ", err)
	}
}

func TestParseNoBody(t *testing.T) {
	noBody := `---
no body here
or here`

	_, err := Parse(strings.NewReader(noBody))

	if err.Error() != "no end to the metadata block" {
		t.Error("did not get the expected error", err.Error())
	}
}

func TestParse(t *testing.T) {
	assert := NewAssertions(t)
	post, err := Parse(strings.NewReader(rawPost))
	assert.NotError(err)
	assert.StringsEqual(post.Layout, "post")
	assert.StringsEqual(post.Title, "example post")

	expectedDate, err := time.Parse(DateFormat, "2016-10-15 23:24:01")

	if post.Date != expectedDate {
		t.Error("Did not get the expected date", post.Date)
	}

	if !reflect.DeepEqual(post.Categories, []string{"category1", "category2"}) {
		t.Error("Did not get the expected categories", post.Categories)
	}

	assert.True(post.Published, "expected post to be published")

	var expectedHTML = `<p>This is the body of the post</p>

<h2>A sub header</h2>
`
	assert.StringsEqual(string(post.Body), expectedHTML)
}

func TestTitleTextParse(t *testing.T) {
	assert := NewAssertions(t)
	titleWithHTML := "the <em>title</em>"
	rawPost := fmt.Sprintf(`---
title: %s
date: 2016-10-15 23:24:01
---`, titleWithHTML)

	post, err := Parse(strings.NewReader(rawPost))
	assert.NotError(err)
	assert.StringsEqual(post.Title, titleWithHTML)
	assert.StringsEqual(post.TitleText, "the title")
}

func TestMetadataParseError(t *testing.T) {
	assert := NewAssertions(t)
	_, err := Parse(strings.NewReader(badMetadata))
	assert.ErrorMessage(err, `parsing time "2016-10-15T23:24:01" as "2006-01-02 15:04:05": cannot parse "T23:24:01" as " "`)
}
