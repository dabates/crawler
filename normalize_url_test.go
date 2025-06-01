package main

import (
	"reflect"
	"testing"
)

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

func TestGetURLsFromHTML(t *testing.T) {
	testCases := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
		wantErr   bool
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
			wantErr:  false,
		},
		{
			name:      "empty HTML body",
			inputURL:  "https://blog.boot.dev",
			inputBody: "",
			expected:  []string{},
			wantErr:   false,
		},
		{
			name:     "no links in HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<p>No links here</p>
	</body>
</html>
`,
			expected: []string{},
			wantErr:  false,
		},
		{
			name:     "multiple relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">Link 1</a>
		<a href="/path/two">Link 2</a>
		<a href="/path/three">Link 3</a>
	</body>
</html>
`,
			expected: []string{
				"https://blog.boot.dev/path/one",
				"https://blog.boot.dev/path/two",
				"https://blog.boot.dev/path/three",
			},
			wantErr: false,
		},
		{
			name:     "mixed relative and absolute URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/relative/path">Relative</a>
		<a href="https://external.com/path">External</a>
		<a href="/">Home</a>
		<a href="https://blog.boot.dev/direct">Direct</a>
	</body>
</html>
`,
			expected: []string{
				"https://blog.boot.dev/relative/path",
				"https://external.com/path",
				"https://blog.boot.dev/",
				"https://blog.boot.dev/direct",
			},
			wantErr: false,
		},
		{
			name:     "invalid HTML",
			inputURL: "https://blog.boot.dev",
			inputBody: `
   <html>
       <body>
           <a href=">Invalid HTML
       </body>
   </html>
   `,
			expected: []string{},
			wantErr:  false, // Change to false since html.Parse won't error
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := getURLsFromHTML(tc.inputBody, tc.inputURL)

			// Check error cases
			if (err != nil) != tc.wantErr {
				t.Errorf("getURLsFromHTML() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			// If we expect an error, don't check the results
			if tc.wantErr {
				return
			}

			// Compare results using reflect.DeepEqual
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("getURLsFromHTML() = %v, want %v", got, tc.expected)
			}
		})
	}
}
