package main

import (
	"bytes"
	"fmt"
	"github.com/gypsydave5/blawg/post"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var posts []post.Page
	posts = createPosts()

	for _, post := range posts {
		fmt.Println(post.Metadata)
	}

	buildHomepage()
}

func createPosts() []post.Page {
	postDir := "_posts"
	posts := []post.Page{}

	err := filepath.Walk(postDir, func(path string, fileInfo os.FileInfo, err error) error {
		if fileInfo.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		check(err)
		defer f.Close()
		post, err := post.Parse(f)
		if err != nil {
			log.Fatalf("Failed for %s : %s", path, err)
		}
		posts = append(posts, *post)
		return err
	})

	check(err)

	return posts
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
