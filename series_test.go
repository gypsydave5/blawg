package blawg

import "testing"

func TestAttachSeriesOrdersChaptersByPart(t *testing.T) {
	posts := Posts{
		{TitleText: "chapter two", Metadata: Metadata{Series: "A Series", Part: 2}},
		{TitleText: "chapter one", Metadata: Metadata{Series: "A Series", Part: 1}},
		{TitleText: "unrelated post"},
	}

	series, err := AttachSeries(&posts)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	s, ok := series["A Series"]
	if !ok {
		t.Fatal("expected series \"A Series\" to be present")
	}
	if len(s.Posts) != 2 {
		t.Fatalf("expected 2 posts in series, got %d", len(s.Posts))
	}
	if s.Posts[0].TitleText != "chapter one" || s.Posts[1].TitleText != "chapter two" {
		t.Fatalf("expected chapters ordered by part, got %v", s.Posts)
	}

	for _, p := range posts {
		if p.Series == "A Series" && len(p.SeriesPosts) != 2 {
			t.Fatalf("expected post %q to have SeriesPosts attached", p.TitleText)
		}
	}
	if posts[2].SeriesPosts != nil {
		t.Fatalf("expected unrelated post to have no SeriesPosts, got %v", posts[2].SeriesPosts)
	}
}

func TestAttachSeriesErrorsOnMissingPart(t *testing.T) {
	posts := Posts{
		{TitleText: "chapter one", Metadata: Metadata{Series: "A Series", Part: 1}},
		{TitleText: "chapter two", Metadata: Metadata{Series: "A Series"}},
	}

	_, err := AttachSeries(&posts)
	if err == nil {
		t.Fatal("expected an error for a post with no part number, got nil")
	}
}

func TestAttachSeriesErrorsOnDuplicatePart(t *testing.T) {
	posts := Posts{
		{TitleText: "chapter one", Metadata: Metadata{Series: "A Series", Part: 1}},
		{TitleText: "also chapter one", Metadata: Metadata{Series: "A Series", Part: 1}},
	}

	_, err := AttachSeries(&posts)
	if err == nil {
		t.Fatal("expected an error for duplicate part numbers, got nil")
	}
}
