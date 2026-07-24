package blawg

import (
	"fmt"
	"io"

	"gopkg.in/russross/blackfriday.v2"
)

const mermaidLanguage = "mermaid"

// footnoteReturnLink is the contents of the link at the end of each footnote
// that takes the reader back to the reference in the body of the text.
const footnoteReturnLink = `<span class="footnote-return-icon">&#8617;&#xFE0E;</span>`

// mermaidRenderer wraps blackfriday's HTML renderer, replacing ```mermaid
// code fences with an inline SVG rendered by mermaid-cli. If rendering
// fails, the code block falls back to normal HTML rendering.
type mermaidRenderer struct {
	*blackfriday.HTMLRenderer
}

func newMermaidRenderer() *mermaidRenderer {
	return &mermaidRenderer{
		HTMLRenderer: blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			Flags:                      blackfriday.CommonHTMLFlags | blackfriday.FootnoteReturnLinks,
			FootnoteReturnLinkContents: footnoteReturnLink,
		}),
	}
}

func (r *mermaidRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if node.Type == blackfriday.CodeBlock && string(node.CodeBlockData.Info) == mermaidLanguage {
		svg, err := renderMermaidSVG(string(node.Literal))
		if err == nil {
			fmt.Fprintf(w, `<div class="mermaid-diagram">%s</div>`, svg)
			return blackfriday.GoToNext
		}
	}

	return r.HTMLRenderer.RenderNode(w, node, entering)
}

func markdownRenderer() blackfriday.Option {
	return blackfriday.WithRenderer(newMermaidRenderer())
}
