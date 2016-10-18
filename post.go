package page

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"gopkg.in/yaml.v2"
)

const dateFormat = "2006-01-02 15:04:05"

type Page struct {
	Body string
	Date time.Time
	metadata
}

type metadata struct {
	Title      string   `yaml:"title"`
	Layout     string   `yaml:"layout"`
	Date       string   `yaml:"date"`
	Categories []string `yaml:"categories"`
}

func parse(rawPage io.Reader) (*Page, error) {
	page := new(Page)
	rawMeta, body, err := split(rawPage)
	if err != nil {
		return page, err
	}
	addMeta(rawMeta, page)
	page.Body = body
	return page, nil
}

func split(page io.Reader) (meta, body string, err error) {
	var m bytes.Buffer
	var b bytes.Buffer
	scanner := bufio.NewScanner(page)

	scanner.Scan()
	if line := scanner.Text(); line != "---" {
		return meta, body, errors.New("No metadata block")
	}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" {
			break
		}
		m.WriteString(fmt.Sprintln(line))
	}

	if scanner.Text() != "---" {
		return meta, body, errors.New("No end to the metadata block")
	}

	for scanner.Scan() {
		line := scanner.Text()
		b.WriteString(fmt.Sprintln(line))
	}

	return m.String(), b.String(), nil
}

func addMeta(rawMeta string, page *Page) (err error) {
	meta := metadata{}

	err = yaml.Unmarshal([]byte(rawMeta), &meta)
	if err != nil {
		return err
	}

	date, err := time.Parse(dateFormat, meta.Date)
	if err != nil {
		return err
	}

	page.Date = date
	page.metadata = meta

	return
}
