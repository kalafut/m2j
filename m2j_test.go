package m2j

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMDToJira(t *testing.T) {
	tests := []string{
		"headings",
		"text_styles",
		"code_blocks",
		"links",
		"lists",
		"crlf",
		"comments",
	}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			input, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.md", test))
			if err != nil {
				t.Fatal(err)
			}
			expected, err := ioutil.ReadFile(fmt.Sprintf("testdata/%s.jira", test))
			if err != nil {
				t.Fatal(err)
			}
			output, err := MDToJira(string(input))
			require.NoError(t, err)
			assert.Equal(t, string(expected), output, "mismatched output for %q", test)
		})
	}

}
