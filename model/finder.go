package model

import (
	"errors"
	"fmt"
	"strings"
)

type Finder interface {
	Find(name string, b Bookmarks) ([]Bookmark, error)
}

// FirstExactMatchFinder finds the first exact match in the bookmarks.
// If no exact match could be made, the Finder returns an error.
type FirstExactMatchFinder struct {
}

func (f FirstExactMatchFinder) Find(name string, b Bookmarks) ([]Bookmark, error) {
	for _, bm := range b.Bookmarks {
		if strings.ToLower(bm.Name) == strings.ToLower(name) {
			return []Bookmark{bm}, nil
		}
	}
	return []Bookmark{}, errors.New(fmt.Sprintf("FirstExactMatchFinder: bookmark with name '%s' does not exist", name))
}
