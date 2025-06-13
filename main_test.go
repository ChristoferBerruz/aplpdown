package main

import "testing"

func TestIsFileLink(t *testing.T) {
	tests := []struct {
		link     string
		expected bool
	}{
		{"http://example.com/file.txt", true},
		{"http://example.com/folder/", false},
		{"http://example.com/image.png", true},
		{"http://example.com/folder", false},
	}

	for _, tt := range tests {
		result := isFileLink(tt.link)
		if result != tt.expected {
			t.Errorf("isFileLink(%q) = %v; want %v", tt.link, result, tt.expected)
		}
	}
}
