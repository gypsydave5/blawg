package main

import (
	"bytes"
	"html/template"
	"log"
	"os"
)

func main() {
	homepage := blackfriday.Run([]byte("# Title\n\nparagraph\n\nanother paragraph"))

	buildHomepage()

	page.Parse()
}

func buildHomepage() {
	os.Mkdir("site", os.FileMode(0777))

	f, err := os.Create("site/index.html")
	check(err)
	defer f.Close()

	_, err = f.Write(homepage())
	check(err)
}

func homepage() []byte {
	t, err := template.New("page").ParseFiles("template.html")
	check(err)

	var b bytes.Buffer

	err = t.Execute(&b, "")
	check(err)

	return b.Bytes()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
