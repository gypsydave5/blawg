package blawg

import (
	"fmt"
	"html/template"
	"io"
	"os"
)

func WritePost(w io.Writer, post *Post, posts *[]Post, template *template.Template) error {
	page := Page{
		post,
		posts,
	}
	err := template.ExecuteTemplate(w, "main", &page)
	return err
}

func MakePosts(siteDirectory string, posts *[]Post, tmplt *template.Template) (err error) {
	for _, post := range *posts {
		err = MakePost(siteDirectory, &post, posts, tmplt)
		if err != nil {
			return
		}
	}
	return
}

func MakePost(siteDirectory string, post *Post, posts *[]Post, tmplt *template.Template) error {
	if !post.Published {
		return nil
	}

	path := fmt.Sprintf("%s/posts/%s", siteDirectory, post.Path())
	os.MkdirAll(path, os.FileMode(0777))
	fileName := fmt.Sprintf("%sindex.html", path)
	file, err := os.Create(fileName)
	defer file.Close()

	if err != nil {
		return err
	}

	err = WritePost(file, post, posts, tmplt)

	if err != nil {
		return err
	}

	return err
}

func MakeHomepage(siteDirectory string, posts *[]Post, t *template.Template) error {
	os.MkdirAll(siteDirectory, os.FileMode(0777))

	f, err := os.Create(siteDirectory + "/index.html")
	defer f.Close()

	if err != nil {
		return err
	}

	recentPost := (*posts)[len(*posts) - 1]
	println(recentPost.Title)

	page := Page{
		&recentPost,
		posts,
	}

	err = t.ExecuteTemplate(f, "main", page)

	return err
}
