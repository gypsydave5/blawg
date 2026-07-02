package blawg

import (
	"strings"
	"testing"
)

func TestRenderMermaidSVGErrorsWhenMmdcMissing(t *testing.T) {
	t.Setenv("PATH", t.TempDir())

	_, err := renderMermaidSVG("graph TD;\nA-->B;\n")
	if err == nil {
		t.Fatal("expected an error when mmdc is not on PATH, got nil")
	}
}

func TestMermaidFenceFallsBackToCodeBlockWhenMmdcMissing(t *testing.T) {
	assert := NewAssertions(t)
	t.Setenv("PATH", t.TempDir())

	var rawPage = "---\ntitle: \"diagram page\"\n---\n" +
		"```mermaid\ngraph TD;\nA-->B;\n```\n"

	page, err := parsePage(strings.NewReader(rawPage))
	assert.NotError(err)
	assert.StringContains(string(page.Body), "<code")
	assert.StringDoesNotContain(string(page.Body), "<svg")
}
