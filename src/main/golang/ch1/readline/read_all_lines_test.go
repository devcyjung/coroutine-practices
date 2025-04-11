package readline

import (
	"errors"
	"fmt"
	"strings"
	"testing"
	"testing/iotest"
)

func TestReadAllLines(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := ReadAllLines(strings.NewReader(tc.input))
			if got == tc.expected && err == nil {
				t.Logf("%s passed", tc.description)
			} else {
				t.Errorf(`
%s failed,
input
%s
expected
'%s'
got
'%s'`,
					tc.description, tc.input, tc.expected, got)
			}
		})
	}
	t.Run("error case", func(t *testing.T) {
		given := fmt.Errorf("some error")
		got, err := ReadAllLines(iotest.ErrReader(given))
		if !(errors.Is(err, given) && got == "") {
			t.Errorf("got, err = %+v, %s", got, err)
		}
	})
}

var testCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "single newline",
		input:       "a\nb c\nde f g\n",
		expected:    "ab cde f g",
	},
	{
		description: "double newline",
		input:       "a\nb c\n\nde f g\n",
		expected:    "ab cde f g",
	},
	{
		description: "ends without newline",
		input:       "a\nb\nc\nde f g",
		expected:    "abcde f g",
	},
}
