package blawg

// MakeBlog creates a blog in `siteDirectory` out of a directory of posts, using
// standard go html/templates from a directory.
func MakeBlog(postDirectory, templateDirectory, siteDirectory string) error {
	posts, err := GetPosts(postDirectory)
	if err != nil {
		return err
	}

	SortPostsByDate(posts)

	t, err := GetTemplates(templateDirectory)
	if err != nil {
		return err
	}

	err = MakePosts(siteDirectory, &posts, t)
	if err != nil {
		return err
	}

	err = MakeHomepage(siteDirectory, &posts, t)
	return err
}
