package main

import (
	"reflect"
	"testing"
)

func TestFirstExactMatchFinder(t *testing.T) {

	tt := []struct {
		name           string
		bookmarks      Bookmarks
		search         string
		expectedResult []Bookmark
		err            bool
	}{
		{
			name:           "find empty string leads to error",
			bookmarks:      Bookmarks{},
			expectedResult: []Bookmark{},
			search:         "",
			err:            true,
		},
		{
			name:           "search empty bookmarks leads to error",
			bookmarks:      Bookmarks{},
			expectedResult: []Bookmark{},
			search:         "something",
			err:            true,
		},
		{
			name: "search returns correct entry",
			bookmarks: Bookmarks{
				[]Bookmark{
					{
						Name:        "a",
						Path:        "b",
						Invocations: 4,
					},
					{
						Name:        "b",
						Path:        "c",
						Invocations: 5,
					},
					{
						Name:        "c",
						Path:        "d",
						Invocations: 6,
					},
				},
			},
			expectedResult: []Bookmark{
				{
					Name:        "b",
					Path:        "c",
					Invocations: 5,
				},
			},
			search: "b",
			err:    false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var f FirstExactMatchFinder
			bm, err := f.Find(tc.search, tc.bookmarks)
			if (!tc.err && err != nil) || (tc.err && err == nil) {
				t.Fatalf("Expected error to be %t but received %+v", tc.err, err)
			}
			if !reflect.DeepEqual(tc.expectedResult, bm) {
				t.Fatalf("%s, Expected %+v, but received %+v", tc.name, tc.expectedResult, bm)
			}
		})
	}
}
