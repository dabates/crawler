package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "URL with HTTPS and trailing slash",
			input:    "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with HTTPS without trailing slash",
			input:    "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with HTTP and trailing slash",
			input:    "http://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with HTTP without trailing slash",
			input:    "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "URL with additional paths",
			input:    "https://blog.boot.dev/path/to/something/",
			expected: "blog.boot.dev/path/to/something",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := normalizeURL(tc.input)
			if got != tc.expected {
				t.Errorf("normalizeURL(%q) = %q; want %q",
					tc.input, got, tc.expected)
			}
		})
	}
}
