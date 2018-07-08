package blawg

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// MakeBlawg creates a blog in `siteDirectory` out of a directory of posts, using
// standard go html/templates from a directory.
func MakeBlawg(postDirectory, templateDirectory, extrasDirectory, siteDirectory string) error {
	posts, err := GetPosts(postDirectory)
	if err != nil {
		return err
	}

	SortPostsByDate(posts)

	t, err := GetTemplates(templateDirectory)
	if err != nil {
		return err
	}

	err = makePosts(siteDirectory, posts, t)
	if err != nil {
		return err
	}

	err = makePostIndex(siteDirectory, posts, t)
	if err != nil {
		return err
	}

	err = makeHomepage(siteDirectory, posts, t)
	if err != nil {
		return err
	}

	err = makePages(siteDirectory)
	if err != nil {
		return err
	}

	err = copyExtrasDirectoryContents(extrasDirectory, siteDirectory)
	return err
}

func copyExtrasDirectoryContents(publicDirectory, siteDirectory string) (err error) {
	err = filepath.Walk(publicDirectory, func(sourcePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		destinationPath := siteDirectory + strings.TrimPrefix(sourcePath, publicDirectory)

		if info.IsDir() {
			err = os.MkdirAll(destinationPath, info.Mode())
			return err
		}

		source, err := os.Open(sourcePath)
		defer source.Close()
		if err != nil {
			return err
		}

		destination, err := os.Create(destinationPath)
		defer destination.Close()
		if err != nil {
			return err
		}

		_, err = io.Copy(destination, source)
		if err != nil {
			return err
		}

		err = destination.Sync()
		if err != nil {
			return err
		}

		err = destination.Chmod(info.Mode())
		return err
	})

	return err
}
