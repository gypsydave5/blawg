package parse

import (
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

func TestSplitNoMeta(t *testing.T) {
	_, err := Parse(strings.NewReader(`no meta block here!`))

	if err.Error() != "No metadata block" {
		t.Error("did not get the expected error: ", err)
	}
}

func TestParseNoBody(t *testing.T) {
	noBody := `---
no body here
or here`

	_, err := Parse(strings.NewReader(noBody))

	if err.Error() != "No end to the metadata block" {
		t.Error("did not get the expected error", err.Error())
	}
}

func TestParse(t *testing.T) {
	post, err := Parse(strings.NewReader(rawPost))
	if err != nil {
		t.Error(err)
	}

	if post.Layout != "post" {
		t.Error("Did not get expected layout", post.Layout)
	}

	if post.Title != "example post" {
		t.Error("Did not get the expected title", post.Title)
	}

	var timeFormat = DateFormat
	expectedDate, err := time.Parse(timeFormat, "2016-10-15 23:24:01")

	if post.Date != expectedDate {
		t.Error("Did not get the expected date", post.Date)
	}

	if !reflect.DeepEqual(post.Categories, []string{"category1", "category2"}) {
		t.Error("Did not get the expected categories", post.Categories)
	}

	var expectedHTML = `<p>This is the body of the post</p>

<h2>A sub header</h2>
`
	if string(post.Body) != expectedHTML {
		t.Error("Did not get expected body", post.Body)
	}
}
