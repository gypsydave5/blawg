package post

import (
	"reflect"
	"strings"
	"testing"

	"time"
)

var page = `---
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

func TestSplit(t *testing.T) {
	meta, body, _ := split(strings.NewReader(page))

	expectedMeta := `layout: post
title: "example post"
date: 2016-10-15 23:24:01
categories:
    - category1
    - category2
published: true
`

	expectedBody := `This is the body of the post

## A sub header
`
	if expectedMeta != string(meta) {
		t.Error("Did not get meta data")
	}

	if expectedBody != string(body) {
		t.Error("Did not get the body", body)
	}
}

func TestSplitNoMeta(t *testing.T) {
	_, _, err := split(strings.NewReader(`no meta block here!`))

	if err.Error() != "No metadata block" {
		t.Error("did not get the expected error")
	}
}

func TestSplitNoBody(t *testing.T) {
	noBody := `---
no body here
or here`

	_, _, err := split(strings.NewReader(noBody))

	if err.Error() != "No end to the metadata block" {
		t.Error("did not get the expected error", err.Error())
	}
}

func TestParse(t *testing.T) {
	newPage, err := Parse(strings.NewReader(page))
	if err != nil {
		t.Error(err)
	}

	if newPage.Layout != "post" {
		t.Error("Did not get expected layout", newPage.Layout)
	}

	var expectedHTML = `<p>This is the body of the post</p>

<h2>A sub header</h2>
`
	if string(newPage.Body) != expectedHTML {
		t.Error("Did not get expected body", newPage.Body)
	}
}

func TestAddMeta(t *testing.T) {
	rawMeta := `layout: post
title: "example post"
date: 2016-10-15 23:24:01
categories:
    - category1
    - category2
published: true
`
	page := new(Page)

	err := addMeta([]byte(rawMeta), page)

	if err != nil {
		t.Error(err)
	}

	if page.Title != "example post" {
		t.Error("Did not get the expected title", page.Title)
	}

	var timeFormat = "2006-01-02 15:04:05"
	expectedDate, err := time.Parse(timeFormat, "2016-10-15 23:24:01")
	if err != nil {
		t.Error(err)
	}

	if page.Date != expectedDate {
		t.Error("Did not get the expected date", page.Date)
	}

	if !reflect.DeepEqual(page.Categories, []string{"category1", "category2"}) {
		t.Error("Did not get the expected categories", page.Categories)
	}
}
