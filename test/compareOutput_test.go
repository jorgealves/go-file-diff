package test

import (
	"github.com/jorgealves/go-file-diff/internal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleComparison(t *testing.T) {
	source := internal.Data{
		ID:   1,
		Name: "Jon",
	}
	target := internal.Data{
		Name: "Jon",
		ID:   1,
	}

	assert.Equal(t, source, target)
}

func TestCompareOutputReturnsTrue(t *testing.T) {
	outfileSourceOutput := []internal.Data{
		{ID: 1, Name: "Jon"},
		{ID: 2, Name: "Doe"},
	}
	outfileTargetFileParsed := []internal.Data{
		{ID: 2, Name: "Doe"},
		{ID: 1, Name: "Jon"},
	}

	result := internal.CompareOutput(outfileSourceOutput, outfileTargetFileParsed)
	assert.Truef(t, result, "")
}


func TestCompareDifferentLenOutputReturnsFalse(t *testing.T) {
	outfileSourceOutput := []internal.Data{
		{ID: 1, Name: "Jon"},
		{ID: 2, Name: "Doe"},
	}
	outfileTargetFileParsed := []internal.Data{
		{ID: 2, Name: "Doe"},
		{ID: 1, Name: "Jon"},
		{ID: 3, Name: "Mary"},
	}

	result := internal.CompareOutput(outfileSourceOutput, outfileTargetFileParsed)
	assert.Falsef(t, result, "")
}

func TestCompareDiffferentOutput1ReturnsFalse(t *testing.T) {
	outfileSourceOutput := []internal.Data{
		{ID: 1, Name: "Jon"},
		{ID: 3, Name: "Mary"},
	}
	outfileTargetFileParsed := []internal.Data{
		{ID: 2, Name: "Doe"},
		{ID: 1, Name: "Jon"},
	}

	result := internal.CompareOutput(outfileSourceOutput, outfileTargetFileParsed)
	assert.Falsef(t, result, "")
}
