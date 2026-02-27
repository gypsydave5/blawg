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
tags:
    - category1
    - category2
published: true
description: a description of a post
---
This is the body of the post...

## A sub header
`

var badMetadata = `---
layout: post
title: "example post"
date: 2016-10-15T23:24:01
tags:
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
	assert.StringsEqual(post.Description, "a description of a post")

	expectedDate, err := time.Parse(dateFormat, "2016-10-15 23:24:01")

	if post.Date != expectedDate {
		t.Error("Did not get the expected date", post.Date)
	}

	if !reflect.DeepEqual(post.Categories, []string{"category1", "category2"}) {
		t.Error("Did not get the expected categories", post.Categories)
	}

	assert.True(post.Published, "expected post to be published")

	var expectedHTML = `<p>This is the body of the post&hellip;</p>

<h2 id="a-sub-header">A sub header</h2>
`
	assert.StringsEqual(string(post.Body), expectedHTML)

	assert.StringsEqual(post.Description, "a description of a post")
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

func TestParseDraft(t *testing.T) {
	assert := NewAssertions(t)

	t.Run("full frontmatter", func(t *testing.T) {
		assert := NewAssertions(t)
		raw := `---
title: "my draft"
date: 2024-06-15 10:00:00
published: true
---
Draft body here.
`
		draft, err := parseDraft(strings.NewReader(raw))
		assert.NotError(err)
		assert.StringsEqual(string(draft.Title), "my draft")
		assert.StringsEqual(draft.TitleText, "my draft")

		expectedDate, _ := time.Parse(dateFormat, "2024-06-15 10:00:00")
		if draft.Date != expectedDate {
			t.Errorf("expected date %v, got %v", expectedDate, draft.Date)
		}

		assert.StringContains(string(draft.Body), "Draft body here")
	})

	t.Run("missing date", func(t *testing.T) {
		assert := NewAssertions(t)
		raw := `---
title: "no date draft"
published: true
---
Body without a date.
`
		draft, err := parseDraft(strings.NewReader(raw))
		assert.NotError(err)
		assert.StringsEqual(draft.TitleText, "no date draft")

		var zeroTime time.Time
		if draft.Date != zeroTime {
			t.Errorf("expected zero time, got %v", draft.Date)
		}
	})

	t.Run("no frontmatter is skipped by GetDrafts", func(t *testing.T) {
		assert := NewAssertions(t)
		// parseDraft itself returns an error for no-frontmatter
		_, err := parseDraft(strings.NewReader("no frontmatter here"))
		if err == nil {
			t.Error("expected an error for input without frontmatter")
		}
		assert.StringsEqual(err.Error(), "no metadata block")
	})

	_ = assert
}

func TestParsePage(t *testing.T) {
	assert := NewAssertions(t)
	var rawPage = `---
title: "_example page_"
---
This is the body of the page...

## A sub header
`

	page, err := parsePage(strings.NewReader(rawPage))
	assert.NotError(err)
	assert.StringsEqual(string(page.Title), "<em>example page</em>")
	assert.StringsEqual(string(page.TitleText), "example page")

	var expectedHTML = `<p>This is the body of the page&hellip;</p>

<h2 id="a-sub-header">A sub header</h2>
`
	assert.StringsEqual(string(page.Body), expectedHTML)
}
