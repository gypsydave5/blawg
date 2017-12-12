package parse

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
	"github.com/gypsydave5/blawg"
	"io"
	"time"

	"gopkg.in/yaml.v2"
)

const dateFormat = "2006-01-02 15:04:05"

func Parse(rawPage io.Reader) (*blawg.Post, error) {
	page := new(blawg.Post)
	rawMeta, body, err := split(rawPage)
	if err != nil {
		return page, err
	}
	addMeta(rawMeta, page)
	page.Body = blackfriday.Run(body)
	return page, nil
}

func split(page io.Reader) (meta, body []byte, err error) {
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

	return m.Bytes(), b.Bytes(), nil
}

func addMeta(rawMeta []byte, page *blawg.Post) (err error) {
	meta := blawg.Metadata{}

	err = yaml.Unmarshal(rawMeta, &meta)
	if err != nil {
		return err
	}

	date, err := time.Parse(dateFormat, meta.Date)
	if err != nil {
		return err
	}

	page.Date = date
	page.Metadata = meta

	return
}
