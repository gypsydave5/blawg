package blawg

import (
	"fmt"
	"html/template"
	"io"
	"os"
)

func writePost(w io.Writer, post *Post, posts *Posts, template *template.Template) error {
	page := PostPage{
		Post:     post,
		PostList: posts,
	}
	err := template.ExecuteTemplate(w, "post", &page)
	return err
}

func makePosts(siteDirectory string, posts *Posts, tmplt *template.Template) (err error) {
	for _, post := range *posts {
		err = makePost(siteDirectory, &post, posts, tmplt)
		if err != nil {
			return
		}
	}
	return
}

func makePages(siteDirectory string) (err error) {
	path := fmt.Sprintf("%s/pages", siteDirectory)
	os.MkdirAll(path, os.FileMode(0777))
	return
}

func makePost(siteDirectory string, post *Post, posts *Posts, tmplt *template.Template) error {
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

	err = writePost(file, post, posts, tmplt)

	if err != nil {
		return err
	}

	return err
}

func makeHomepage(siteDirectory string, posts *Posts, t *template.Template) error {
	os.MkdirAll(siteDirectory, os.FileMode(0777))

	f, err := os.Create(siteDirectory + "/index.html")
	defer f.Close()

	if err != nil {
		return err
	}

	recentPost := (*posts)[len(*posts)-1]

	page := PostPage{
		&recentPost,
		posts,
	}

	err = t.ExecuteTemplate(f, "post", page)

	return err
}

func makePostIndex(siteDirectory string, posts *Posts, t *template.Template) error {
	if t.Lookup("index") == nil {
		return nil
	}

	os.MkdirAll(siteDirectory+"/posts", os.FileMode(0777))

	f, err := os.Create(siteDirectory + "/posts/index.html")
	defer f.Close()

	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(f, "index", posts)
	return err
}
