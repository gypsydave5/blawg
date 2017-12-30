package blawg

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"gopkg.in/russross/blackfriday.v2"
	"io"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
	"html/template"
	"os"
	"path/filepath"
)

const DateFormat = "2006-01-02 15:04:05"

var markdownExtensions = blackfriday.WithExtensions(
	blackfriday.Footnotes |
		blackfriday.CommonExtensions,
)

func Parse(rawPage io.Reader) (*Post, error) {
	post := new(Post)
	rawMeta, body, err := split(rawPage)

	if err != nil {
		return post, err
	}

	err = addMeta(rawMeta, post)

	if err != nil {
		return post, err
	}

	postHTML := blackfriday.Run(body, markdownExtensions)
	post.Body = template.HTML(postHTML)
	post.TitleText, _ = textTitle(post)
	return post, nil
}

func textTitle(p *Post) (string, error) {
	n, err := html.Parse(strings.NewReader(p.Title))
	if err != nil {
		return "", err
	}

	text := ""

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			text = text + n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return text, nil
}

func split(page io.Reader) (meta, body []byte, err error) {
	var m bytes.Buffer
	var b bytes.Buffer
	scanner := bufio.NewScanner(page)

	scanner.Scan()
	if line := scanner.Text(); line != "---" {
		return meta, body, errors.New("no metadata block")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			break
		}
		m.WriteString(fmt.Sprintln(line))
	}

	if scanner.Text() != "---" {
		return meta, body, errors.New("no end to the metadata block")
	}

	for scanner.Scan() {
		line := scanner.Text()
		b.WriteString(fmt.Sprintln(line))
	}

	return m.Bytes(), b.Bytes(), nil
}

func addMeta(rawMeta []byte, post *Post) (err error) {
	meta := Metadata{}

	err = yaml.Unmarshal(rawMeta, &meta)
	if err != nil {
		return err
	}

	date, err := time.Parse(DateFormat, meta.Date)
	if err != nil {
		return err
	}

	post.Date = date
	post.Metadata = meta

	return
}

func GetTemplates(templateDirectory string) (*template.Template, error) {
	return template.ParseGlob(templateDirectory + "/**")
}

func GetPosts(postDir string) (posts []Post, err error) {
	err = filepath.Walk(postDir, func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		defer f.Close()

		post, err := Parse(f)
		if err != nil {
			return errors.New(fmt.Sprintf("error parsing post %s : \n\t%s", fileInfo.Name(), err))
		}
		posts = append(posts, *post)
		return err
	})

	return posts, err
}
