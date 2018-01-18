package blawg

import (
	"fmt"
	"html/template"
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
	_, err := parse(strings.NewReader(`no meta block here!`))

	if err.Error() != "no metadata block" {
		t.Error("did not get the expected error: ", err)
	}
}

func TestParseNoBody(t *testing.T) {
	noBody := `---
no body here
or here`

	_, err := parse(strings.NewReader(noBody))

	if err.Error() != "no end to the metadata block" {
		t.Error("did not get the expected error", err.Error())
	}
}

func TestParse(t *testing.T) {
	assert := NewAssertions(t)
	post, err := parse(strings.NewReader(rawPost))
	assert.NotError(err)
	assert.StringsEqual(post.Layout, "post")
	assert.StringsEqual(string(post.Title), "example post")

	expectedDate, err := time.Parse(dateFormat, "2016-10-15 23:24:01")

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
	titleWithHTML := "the _title_<h1>is BIG</h1>"
	rawPost := fmt.Sprintf(`---
title: %s
date: 2016-10-15 23:24:01
---`, titleWithHTML)

	post, err := parse(strings.NewReader(rawPost))
	assert.NotError(err)
	if post.Title != template.HTML("the <em>title</em><h1>is BIG</h1>") {
		t.Errorf("did not get the expected title HTML, %s", post.Title)
	}

	assert.StringsEqual(post.TitleText, "the title is BIG")
}

func TestMetadataParseError(t *testing.T) {
	assert := NewAssertions(t)
	_, err := parse(strings.NewReader(badMetadata))
	assert.ErrorMessage(err, `parsing time "2016-10-15T23:24:01" as "2006-01-02 15:04:05": cannot parse "T23:24:01" as " "`)
}
