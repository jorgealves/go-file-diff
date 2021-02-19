package test

import (
	"github.com/jorgealves/go-file-diff/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonDiff(t *testing.T) {
	tableCases := []struct {
		source   string
		target   string
		expected bool
	}{
		{
			source:   "../examples/first.json",
			target:   "../examples/first.json",
			expected: true,
		},
		{
			source:   "../examples/first.json",
			target:   "../examples/second.json",
			expected: false,
		},
		{
			source:   "../examples/first.json",
			target:   "../examples/third.json",
			expected: false,
		},
		{
			source:   "../examples/first.json",
			target:   "../examples/fourth.json",
			expected: false,
		},
	}

	for _, test := range tableCases {
		result := cmd.CompareJsonFiles(test.source, test.target)

		assert.Equal(t, test.expected, result)
	}
}
