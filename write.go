package blawg

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"os"
)

func writePost(w io.Writer, post *Post, posts *Posts, template *template.Template) error {
	postPage := PostPage{
		Post:     post,
		PostList: posts,
	}
	err := template.ExecuteTemplate(w, "post", &postPage)
	return err
}

func writePage(w io.Writer, page *Page, template *template.Template) error {
	err := template.ExecuteTemplate(w, "page", &page)
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

func makePages(siteDirectory string, pages []Page, tmplt *template.Template) error {
	path := fmt.Sprintf("%s/pages", siteDirectory)
	err := os.MkdirAll(path, os.FileMode(0777))
	if err != nil {
		return err
	}

	for _, page := range pages {
		err = makePage(siteDirectory, &page, tmplt)
		if err != nil {
			return err
		}
	}
	return err
}

func makePage(siteDirectory string, page *Page, tmplt *template.Template) error {
	path := fmt.Sprintf("%s/pages/%s/", siteDirectory, page.Path())
	err := os.MkdirAll(path, os.FileMode(0777))
	if err != nil {
		return err
	}
	fileName := fmt.Sprintf("%sindex.html", path)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	err = writePage(file, page, tmplt)

	return err
}

func makePost(siteDirectory string, post *Post, posts *Posts, tmplt *template.Template) error {
	path := fmt.Sprintf("%s/posts/%s", siteDirectory, post.Path())
	err := os.MkdirAll(path, os.FileMode(0777))

	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%sindex.html", path)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	err = writePost(file, post, posts, tmplt)

	if err != nil {
		return err
	}

	return err
}

func makeHomepage(siteDirectory string, posts *Posts, t *template.Template) error {
	err := os.MkdirAll(siteDirectory, os.FileMode(0777))
	if err != nil {
		return err
	}

	f, err := os.Create(siteDirectory + "/index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	if len(*posts) == 0 {
		return errors.New("no posts")
	}

	recentPost := (posts.Published())[0]

	page := PostPage{
		&recentPost,
		posts,
	}

	err = t.ExecuteTemplate(f, "post", page)

	return err
}

func makeDraft(siteDir string, draft *Draft, t *template.Template) error {
	path := fmt.Sprintf("%s/%s", siteDir, draft.Path())
	err := os.MkdirAll(path, os.FileMode(0777))
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%sindex.html", path)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if t.Lookup("draft") != nil {
		return t.ExecuteTemplate(file, "draft", &DraftPage{Draft: draft})
	}

	// Fall back to post template
	post := &Post{
		Body:      draft.Body,
		Title:     draft.Title,
		TitleText: draft.TitleText,
		Date:      draft.Date,
		Metadata:  draft.Metadata,
	}
	return t.ExecuteTemplate(file, "post", &PostPage{Post: post})
}

func makeDraftPosts(siteDir string, drafts []Draft, t *template.Template) error {
	for _, draft := range drafts {
		err := makeDraft(siteDir, &draft, t)
		if err != nil {
			return err
		}
	}
	return nil
}

func makePostIndex(siteDirectory string, posts *Posts, t *template.Template) error {
	if t.Lookup("index") == nil {
		return nil
	}

	err := os.MkdirAll(siteDirectory+"/posts", os.FileMode(0777))
	if err != nil {
		return err
	}

	f, err := os.Create(siteDirectory + "/posts/index.html")
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.ExecuteTemplate(f, "index", posts.Published())
	return err
}
