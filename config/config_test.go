package config

import (
	"os"
	"reflect"
	"testing"

	"github.com/mrjerz/bookmarks/model"
)

func TestLoad(t *testing.T) {

	tmpFile, err := os.CreateTemp("", "tmpfile")
	if err != nil {
		t.Errorf("Could not prepare testdata: %s", err)
	}

	tt := []struct {
		name        string
		content     string
		filePath    string
		expected    model.Bookmarks
		shouldError bool
	}{
		{
			name:        "empty existing file",
			content:     "",
			filePath:    tmpFile.Name(),
			expected:    model.Bookmarks{},
			shouldError: false,
		},
		{
			name:        "non existing file",
			content:     "",
			filePath:    "/tmp/i-do-not-exist",
			expected:    model.Bookmarks{},
			shouldError: true,
		},
		{
			name:     "proper file",
			content:  `{"bookmarks": [{"name":"test", "path": "/test"}]}`,
			filePath: tmpFile.Name(),
			expected: model.Bookmarks{
				Bookmarks: []model.Bookmark{
					{
						Name: "test",
						Path: "/test",
					},
				},
			},
			shouldError: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.content != "" {
				tmpFile.WriteString(tc.content)
			}
			bm, err := Read(tc.filePath)
			if (!tc.shouldError && err != nil) || (tc.shouldError && err == nil) {
				t.Errorf("Error behavior not correct. Expected: %t, Received: %s", tc.shouldError, err)
			}

			if !reflect.DeepEqual(bm, tc.expected) {
				t.Errorf("Expected content: %+v, Received: %+v", tc.expected, bm)
			}
		})
	}
}

func TestWrite(t *testing.T) {

	tt := []struct {
		name        string
		bm          model.Bookmarks
		shouldError bool
	}{
		{
			name: "should write 2 entires from temporary file",
			bm: model.Bookmarks{
				Bookmarks: []model.Bookmark{
					{
						Name: "test1",
						Path: "/test1",
					},
					{
						Name: "test2",
						Path: "/test2",
					},
				},
			},
			shouldError: false,
		},
		{
			name: "should write empty if no bookmarks present",
			bm: model.Bookmarks{
				Bookmarks: []model.Bookmark{},
			},
			shouldError: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			tmpFile, err := os.CreateTemp("", "file")
			if err != nil {
				t.Errorf("Could not prepare testdata. %s", err)
			}

			err = Write(tmpFile.Name(), tc.bm)
			if (!tc.shouldError && err != nil) || (tc.shouldError && err == nil) {
				t.Errorf("Error behavior not correct. Expected: %t, Received: %s", tc.shouldError, err)
			}
			tmpFile.Sync()

			readBookmarks, err := Read(tmpFile.Name())
			if err != nil {
				t.Errorf("Did not expect error here: %s", err)
			}

			if !reflect.DeepEqual(readBookmarks, tc.bm) {
				t.Errorf("Did not write back, what it should. Expected %+v, Received: %+v. File: %s", tc.bm, readBookmarks, tmpFile.Name())
			}
		})
	}
}
