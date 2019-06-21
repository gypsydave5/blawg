package blawg

import "sort"

// Posts represents a slice of []Post. It is used for default sorting by date
// and adds some methods used in the templates.
type Posts []Post

func (ps Posts) Len() int {
	return len(ps)
}

func (ps Posts) Less(i, j int) bool {
	return ps[i].Date.After(ps[j].Date)
}

func (ps Posts) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Reverse returns a copy of the Posts, sorted into reverse date order (earliest
// first).
func (ps Posts) Reverse() Posts {
	reversedList := make(Posts, len(ps))
	copy(reversedList, ps)
	sort.Sort(sort.Reverse(reversedList))
	return reversedList
}

// Take returns a slice of the first n Posts.
func (ps Posts) Take(n int) Posts {
	return ps[:n]
}

// Drop returns a slice of Posts, without the first n.
func (ps Posts) Drop(n int) Posts {
	return ps[n:]
}

// SortPostsByDate sorts a list of Posts in place by date order.
func SortPostsByDate(ps *Posts) {
	sort.Sort(ps)
}

func (ps *Posts) sortByDate() {
	sort.Sort(ps)
}
