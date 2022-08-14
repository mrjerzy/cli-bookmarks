package main

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {

	tt := []struct {
		name      string
		input     string
		bookmarks Bookmarks
		err       error
	}{
		{
			name:      "empty input",
			input:     "",
			bookmarks: Bookmarks{},
			err:       nil,
		},
		{
			name:      "empty json input",
			input:     "{}",
			bookmarks: Bookmarks{},
			err:       nil,
		},
		{
			name:  "read bookmarks",
			input: `{"bookmarks":[{"name": "a","path": "b","invocations": 3},{"name": "c","path": "d","invocations": 4}]}`,
			bookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "a",
						Path:        "b",
						Invocations: 3,
					},
					{
						Name:        "c",
						Path:        "d",
						Invocations: 4,
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			reader := strings.NewReader(tc.input)
			b, err := Load(reader)
			if err != tc.err {
				t.Fatalf("Did not expect to receive error: %s", err)
			}
			if !reflect.DeepEqual(b, tc.bookmarks) {
				t.Fatalf("Did expect to receive:\n%+v\nbut received:\n%+v", tc.bookmarks, b)
			}
		})
	}
}

func TestAdd(t *testing.T) {

	tt := []struct {
		name            string
		beforeBookmarks Bookmarks
		afterBookmarks  Bookmarks
		newBookmark     Bookmark
		err             bool
	}{
		{
			name:            "add empty bookmark lead to error",
			beforeBookmarks: Bookmarks{},
			newBookmark:     Bookmark{},
			afterBookmarks:  Bookmarks{},
			err:             true,
		},
		{
			name:            "add new element should add element",
			beforeBookmarks: Bookmarks{},
			newBookmark: Bookmark{
				Name:        "a",
				Path:        "b",
				Invocations: 0,
			},
			afterBookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "a",
						Path:        "b",
						Invocations: 0,
					},
				},
			},
			err: false,
		},
		{
			name: "add element with the same name should lead to error",
			beforeBookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "a",
						Path:        "b",
						Invocations: 0,
					},
				},
			},
			newBookmark: Bookmark{
				Name:        "a",
				Path:        "b",
				Invocations: 0,
			},
			afterBookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "a",
						Path:        "b",
						Invocations: 0,
					},
				},
			},
			err: true,
		},
		{
			name: "add element with the same name should lead to error - case insensitive",
			beforeBookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "A",
						Path:        "b",
						Invocations: 0,
					},
				},
			},
			newBookmark: Bookmark{
				Name:        "a",
				Path:        "b",
				Invocations: 0,
			},
			afterBookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "A",
						Path:        "b",
						Invocations: 0,
					},
				},
			},
			err: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			err := tc.beforeBookmarks.Add(tc.newBookmark)
			if (!tc.err && err != nil) || (tc.err && err == nil) {
				t.Fatalf("Expected error to be %t but received %+v", tc.err, err)
			}
			if !reflect.DeepEqual(tc.beforeBookmarks, tc.afterBookmarks) {
				t.Fatalf("%s, Expected %+v, but received %+v", tc.name, tc.afterBookmarks, tc.beforeBookmarks)
			}

		})
	}
}

func TestSave(t *testing.T) {

	tt := []struct {
		name      string
		bookmarks Bookmarks
		output    string
		err       error
	}{
		{
			name:      "empty json input",
			bookmarks: Bookmarks{},
			output:    `{"bookmarks": null}`,
			err:       nil,
		},
		{
			name: "one element input",
			bookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "a",
						Path:        "b",
						Invocations: 3,
					},
				},
			},
			output: `{"bookmarks": [{"name":"a", "path":"b", "invocations":3}]}`,
			err:    nil,
		},
		{
			name: "multiple elements input",
			bookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "a",
						Path:        "b",
						Invocations: 3,
					},
					{
						Name:        "b",
						Path:        "c",
						Invocations: 4,
					},
				},
			},
			output: `{"bookmarks": [{"name":"a", "path":"b", "invocations":3},{"name":"b", "path":"c", "invocations":4}]}`,
			err:    nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			isBuffer := new(strings.Builder)
			err := tc.bookmarks.Save(isBuffer)
			if err != nil {
				t.Errorf("Did not expect to receive error: %s", err)
			}

			should := replaceWhitespace(tc.output)
			is := replaceWhitespace(isBuffer.String())

			if should != is {
				t.Errorf("Expected output to be: %s, but received: %s", should, is)
			}
		})
	}
}

func TestUpdate(t *testing.T) {

	tt := []struct {
		name            string
		beforeBookmarks Bookmarks
		afterBookmarks  Bookmarks
		updateBookmark  Bookmark
		err             bool
	}{
		{
			name: "update existing element with new values",
			beforeBookmarks: Bookmarks{
				[]Bookmark{{Name: "a", Path: "b", Invocations: 3}},
			},
			updateBookmark: Bookmark{Name: "a", Path: "c", Invocations: 9},
			afterBookmarks: Bookmarks{
				[]Bookmark{{Name: "a", Path: "c", Invocations: 9}},
			},
			err: false,
		},
		{
			name: "update existing element with new values, case insensitive",
			beforeBookmarks: Bookmarks{
				[]Bookmark{
					{Name: "a", Path: "b", Invocations: 3},
				},
			},
			updateBookmark: Bookmark{
				Name: "A", Path: "c", Invocations: 9},
			afterBookmarks: Bookmarks{
				[]Bookmark{
					{Name: "a", Path: "c", Invocations: 9},
				},
			},
			err: false,
		},
		{
			name: "error on non-existing name",
			beforeBookmarks: Bookmarks{
				[]Bookmark{
					{Name: "a", Path: "b", Invocations: 3},
				},
			},
			updateBookmark: Bookmark{Name: "f", Path: "f", Invocations: 12},
			afterBookmarks: Bookmarks{
				[]Bookmark{
					{Name: "a", Path: "b", Invocations: 3},
				},
			},
			err: true,
		},
		{
			name:            "error on empty input",
			beforeBookmarks: Bookmarks{},
			updateBookmark:  Bookmark{},
			afterBookmarks:  Bookmarks{},
			err:             true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.beforeBookmarks.Update(tc.updateBookmark)
			if (!tc.err && err != nil) || (tc.err && err == nil) {
				t.Fatalf("Expected error to be %t but received %+v", tc.err, err)
			}
			if !reflect.DeepEqual(tc.beforeBookmarks, tc.afterBookmarks) {
				t.Fatalf("%s, Expected %+v, but received %+v", tc.name, tc.afterBookmarks, tc.beforeBookmarks)
			}
		})
	}
}

func TestRemove(t *testing.T) {

	tt := []struct {
		name            string
		beforeBookmarks Bookmarks
		deleteBookmark  Bookmark
		afterBookmarks  Bookmarks
		err             bool
	}{
		{
			name:            "error is input for deletion is empty",
			beforeBookmarks: Bookmarks{},
			afterBookmarks:  Bookmarks{},
			deleteBookmark:  Bookmark{},
			err:             true,
		},
		{
			name: "deleting non-existing name leads to error",
			beforeBookmarks: Bookmarks{
				[]Bookmark{{Name: "a", Path: "b", Invocations: 5}},
			},
			afterBookmarks: Bookmarks{
				[]Bookmark{{Name: "a", Path: "b", Invocations: 5}},
			},
			deleteBookmark: Bookmark{Name: "f"},
			err:            true,
		},
		{
			name:            "deletion of empty set leads to error",
			beforeBookmarks: Bookmarks{},
			afterBookmarks:  Bookmarks{},
			deleteBookmark:  Bookmark{Name: "a", Path: "b", Invocations: 3},
			err:             true,
		},
		{
			name: "delete element",
			beforeBookmarks: Bookmarks{
				[]Bookmark{
					{Name: "a", Path: "b", Invocations: 5},
					{Name: "f", Path: "f", Invocations: 12},
				},
			},
			afterBookmarks: Bookmarks{
				[]Bookmark{
					{Name: "a", Path: "b", Invocations: 5},
				},
			},
			deleteBookmark: Bookmark{Name: "f"},
			err:            false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.beforeBookmarks.Remove(tc.deleteBookmark)
			if (!tc.err && err != nil) || (tc.err && err == nil) {
				t.Fatalf("Expected error to be %t but received %+v", tc.err, err)
			}
			if !reflect.DeepEqual(tc.beforeBookmarks, tc.afterBookmarks) {
				t.Fatalf("%s, Expected %+v, but received %+v", tc.name, tc.afterBookmarks, tc.beforeBookmarks)
			}
		})
	}
}

// helper struct for TestSearch
type staticElementReturnFinder struct{}

var staticElement Bookmark = Bookmark{Name: "static", Path: "element", Invocations: 4}

// helper funct for TestSearch
func (f staticElementReturnFinder) Find(name string, b Bookmarks) ([]Bookmark, error) {
	return []Bookmark{staticElement}, nil
}

type errorReturnFinder struct{}

func (e errorReturnFinder) Find(name string, b Bookmarks) ([]Bookmark, error) {
	return []Bookmark{}, errors.New("some error")
}

func TestGet(t *testing.T) {
	tt := []struct {
		name             string
		searchStr        string
		inputBookmarks   Bookmarks
		expectedBookmark Bookmark
		finder           Finder
		err              bool
	}{
		{
			name:             "return the finder element",
			inputBookmarks:   Bookmarks{[]Bookmark{}},
			searchStr:        "any",
			expectedBookmark: staticElement,
			finder:           staticElementReturnFinder{},
			err:              false,
		},
		{
			name:             "propagate error from finder",
			inputBookmarks:   Bookmarks{[]Bookmark{}},
			searchStr:        "any",
			expectedBookmark: Bookmark{},
			finder:           errorReturnFinder{},
			err:              true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			bms, err := tc.inputBookmarks.Get(tc.searchStr, tc.finder)
			if (!tc.err && err != nil) || (tc.err && err == nil) {
				t.Fatalf("Expected error to be %t but received %+v", tc.err, err)
			}
			if !reflect.DeepEqual(bms, tc.expectedBookmark) {
				t.Fatalf("%s, Expected %+v, but received %+v", tc.name, tc.expectedBookmark, bms)
			}
		})
	}
}

func replaceWhitespace(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", "")
}
