package blawg

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// renderMermaidSVG shells out to the mermaid-cli (mmdc) to render a mermaid
// diagram definition into an SVG document.
//
// Requires mmdc (npm install -g @mermaid-js/mermaid-cli) and a
// Chrome/Chromium binary on PATH or in Puppeteer's cache; see docs/blawg.1.
func renderMermaidSVG(diagram string) (string, error) {
	dir, err := ioutil.TempDir("", "blawg-mermaid")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(dir)

	inputPath := filepath.Join(dir, "diagram.mmd")
	outputPath := filepath.Join(dir, "diagram.svg")

	if err := ioutil.WriteFile(inputPath, []byte(diagram), 0644); err != nil {
		return "", err
	}

	cmd := exec.Command("mmdc", "-i", inputPath, "-o", outputPath, "-b", "transparent")
	if err := cmd.Run(); err != nil {
		return "", err
	}

	svg, err := ioutil.ReadFile(outputPath)
	if err != nil {
		return "", err
	}

	return string(svg), nil
}
