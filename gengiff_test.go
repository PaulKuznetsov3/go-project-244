package code

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name string
		filepath1 string
		filepath2  string
		expected string
	}{
		{
			name: "compare json",
			filepath1: "./testdata/file1.json",
			filepath2: "./testdata/file2.json",
			expected: "./testdata/stylish.txt",
		},
			{
			name: "compare yml",
			filepath1: "./testdata/file1.yml",
			filepath2: "./testdata/file2.yml",
			expected: "./testdata/stylish.txt",
		},
	}

	for _, tc := range tests {
		data, _ := os.ReadFile(tc.expected)
		want := string(data)
		got, _ := GenDiff(tc.filepath1, tc.filepath2, "stylish")
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, want, got)
		})
	}
}
