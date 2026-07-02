package blawg

import (
	"io"

	"gopkg.in/russross/blackfriday.v2"
)

const mermaidLanguage = "mermaid"

// mermaidRenderer wraps blackfriday's HTML renderer, replacing ```mermaid
// code fences with an inline SVG rendered by mermaid-cli. If rendering
// fails, the code block falls back to normal HTML rendering.
type mermaidRenderer struct {
	*blackfriday.HTMLRenderer
}

func newMermaidRenderer() *mermaidRenderer {
	return &mermaidRenderer{
		HTMLRenderer: blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
			Flags: blackfriday.CommonHTMLFlags,
		}),
	}
}

func (r *mermaidRenderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if node.Type == blackfriday.CodeBlock && string(node.CodeBlockData.Info) == mermaidLanguage {
		svg, err := renderMermaidSVG(string(node.Literal))
		if err == nil {
			w.Write([]byte(svg))
			return blackfriday.GoToNext
		}
	}

	return r.HTMLRenderer.RenderNode(w, node, entering)
}

func markdownRenderer() blackfriday.Option {
	return blackfriday.WithRenderer(newMermaidRenderer())
}
