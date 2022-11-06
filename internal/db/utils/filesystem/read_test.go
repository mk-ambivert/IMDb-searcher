package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFileToSlice_NonexistingFile(t *testing.T) {
	filePath := "test_files/foo.txt"

	rows, err := ReadFileToSlice(filePath)
	assert.Error(t, err)
	assert.Nil(t, rows)
}

func TestReadFileToSlice_EmptyFile(t *testing.T) {
	filePath := "test_files/empty.txt"

	rows, err := ReadFileToSlice(filePath)
	assert.Error(t, err)
	assert.Nil(t, rows)
}

func TestReadFileToSlice_OnelineFile(t *testing.T) {
	expectedRows := []string{"row"}
	filePath := "test_files/row.txt"

	rows, err := ReadFileToSlice(filePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedRows, rows)
}

func TestReadFileToSlice_MultilineFile(t *testing.T) {
	expectedRows := []string{
		"row",
		"row foo",
		"row",
	}
	filePath := "test_files/rows.txt"

	rows, err := ReadFileToSlice(filePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedRows, rows)
}
