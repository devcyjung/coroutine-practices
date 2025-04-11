package readline

import (
	"io"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestFindDuplicateLines(t *testing.T) {
	for _, tc := range testcases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := FindDuplicateLines(strings.NewReader(tc.input))
			slices.Sort(tc.expected)
			slices.Sort(got)
			if err == nil && slices.Equal(got, tc.expected) {
				t.Logf("%s passed", tc.description)
			} else {
				t.Errorf(`%s failed
input:
%s
error:
%+v
got:
%s
expected:
%s`, tc.description, tc.input, err, got, tc.expected)
			}
		})
	}

	for _, tc := range testfiles {
		t.Run(tc.description, func(t *testing.T) {
			f, err := os.Open(tc.path)
			defer func() {
				closeErr := f.Close()
				if closeErr != nil {
					t.Errorf("error closing file: %v", closeErr.Error())
				}
			}()
			if err != nil {
				t.Errorf("could not open test file: %v", err)
			}
			got, err := FindDuplicateLines(f)
			slices.Sort(got)
			slices.Sort(tc.expected)
			if err == nil && slices.Equal(got, tc.expected) {
				t.Logf("%s passed", tc.description)
			} else {
				var buf strings.Builder
				_, copyErr := io.Copy(&buf, f)
				if copyErr != nil {
					t.Errorf("could not copy test file: %v", copyErr)
				}
				t.Errorf(`%s failed
input:
%s
error:
%+v
got:
%s
expected:
%s`, tc.description, buf.String(), err, got, tc.expected)
			}
		})
	}
}

var testcases = []struct {
	description string
	input       string
	expected    []string
}{
	{
		description: "empty string",
		input:       "",
		expected:    []string{},
	},
	{
		description: "single line",
		input:       "a",
		expected:    []string{},
	},
	{
		description: "multiple empty lines",
		input:       "\n\n\n\n\n\n\n",
		expected:    []string{""},
	},
	{
		description: "multiple lines",
		input:       "a\n a\n  a\na",
		expected:    []string{"a"},
	},
}

var testfiles = []struct {
	description string
	path        string
	expected    []string
}{
	{
		description: "ascii T vs unicode greek T, space vs no space",
		path:        "_testdata/test1.txt",
		expected:    []string{"This is duplicate."},
	},
}
