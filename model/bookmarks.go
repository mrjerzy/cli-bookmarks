package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

type Bookmark struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Bookmarks struct {
	Bookmarks []Bookmark `json:"bookmarks"`
}

func (b *Bookmarks) Add(bookmark Bookmark) error {
	if bookmark == (Bookmark{}) {
		return errors.New("Add: input-bookmark cannot be empty")
	}

	for _, bm := range b.Bookmarks {
		if strings.ToLower(bm.Name) == strings.ToLower(bookmark.Name) {
			return errors.New(fmt.Sprintf("Add: bookmark with name %s exists already", bookmark.Name))
		}
	}

	b.Bookmarks = append(b.Bookmarks, bookmark)

	return nil
}

func (b *Bookmarks) Update(bookmark Bookmark) error {
	if bookmark == (Bookmark{}) {
		return errors.New("Update: input-bookmark cannot be empty")
	}

	var found bool
	for i, bm := range b.Bookmarks {
		if strings.ToLower(bm.Name) == strings.ToLower(bookmark.Name) {
			found = true
			bookmark.Name = strings.ToLower(bookmark.Name)
			b.Bookmarks = append(b.Bookmarks[:i], bookmark)
			b.Bookmarks = append(b.Bookmarks, b.Bookmarks[i+1:]...)
		}
	}

	if !found {
		return errors.New(fmt.Sprintf("Update: bookmark with name '%s' does not exist", bookmark.Name))
	}

	return nil
}

func (b *Bookmarks) Remove(bookmark Bookmark) error {

	if bookmark == (Bookmark{}) {
		return errors.New(fmt.Sprintf("Remove: no element with name '%s' exists", bookmark.Name))
	}

	var found bool
	for i, bm := range b.Bookmarks {
		if strings.ToLower(bm.Name) == strings.ToLower(bookmark.Name) {
			found = true
			b.Bookmarks = append(b.Bookmarks[:i], b.Bookmarks[i+1:]...)
		}
	}

	if !found {
		return errors.New(fmt.Sprintf("Remove: bookmark with name '%s' does not exist", bookmark.Name))
	}

	return nil
}

func (b Bookmarks) Get(name string, finder Finder) (Bookmark, error) {
	bms, err := finder.Find(name, b)
	if err != nil {
		return Bookmark{}, err
	}
	if len(bms) < 1 {
		return Bookmark{}, errors.New("Get: received an empty result from Finder")
	}

	return bms[0], err
}

func Load(r io.Reader) (Bookmarks, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return Bookmarks{}, err
	}

	if bytes.Equal(content, []byte("")) {
		return Bookmarks{}, nil
	}

	var b Bookmarks
	err = json.Unmarshal(content, &b)
	if err != nil {
		return Bookmarks{}, err
	}
	return b, nil
}

func (b Bookmarks) Save(w io.Writer) error {
	j, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return err
	}
	w.Write(j)
	return nil
}
