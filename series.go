package blawg

import (
	"fmt"
	"sort"
)

// A Series is a named, ordered group of posts (chapters).
type Series struct {
	Name  string
	Posts Posts
}

// Path of a Series index page. URL safe.
func (s Series) Path() string {
	return "series/" + urlSafeFileName(s.Name) + "/"
}

// AttachSeries groups posts by their Metadata.Series name, orders each
// group by Part, and sets SeriesPosts on every post that belongs to a
// series. It returns the series, keyed by name, for building series index
// pages.
//
// Posts with an empty Series are left untouched. A post with a non-empty
// Series but no Part (Part <= 0), or two posts in the same series sharing
// a Part, is a build error.
func AttachSeries(posts *Posts) (map[string]Series, error) {
	grouped := map[string][]int{} // series name -> indexes into *posts

	for i, post := range *posts {
		if post.Series == "" {
			continue
		}
		grouped[post.Series] = append(grouped[post.Series], i)
	}

	series := map[string]Series{}

	for name, indexes := range grouped {
		seenParts := map[int]string{}

		for _, i := range indexes {
			post := (*posts)[i]
			if post.Part <= 0 {
				return nil, fmt.Errorf("series %q: post %q has no part number", name, post.TitleText)
			}
			if existing, ok := seenParts[post.Part]; ok {
				return nil, fmt.Errorf("series %q: posts %q and %q both have part %d", name, existing, post.TitleText, post.Part)
			}
			seenParts[post.Part] = post.TitleText
		}

		chapters := make(Posts, len(indexes))
		for j, i := range indexes {
			chapters[j] = (*posts)[i]
		}
		sort.Slice(chapters, func(a, b int) bool {
			return chapters[a].Part < chapters[b].Part
		})

		for _, i := range indexes {
			(*posts)[i].SeriesPosts = chapters
		}

		series[name] = Series{Name: name, Posts: chapters}
	}

	return series, nil
}
