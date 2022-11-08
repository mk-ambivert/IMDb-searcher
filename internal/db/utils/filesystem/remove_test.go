package filesystem

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFile(t *testing.T) {
	fileName := "test_files/removingdir/test.txt"
	file, err := os.Create(fileName)
	if err != nil {
		assert.NoError(t, err)
	} else {
		file.Close()
	}

	err = RemoveFile(fileName)
	assert.NoError(t, err)

	// Making sure the file has been removed
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	assert.Error(t, os.ErrNotExist, err)
	f.Close()
}

func TestRemoveFilesInDir(t *testing.T) {
	dirPath := "test_files/removingdir/"
	fileNames := []string{
		"1.txt",
		"2.txt",
		"3.txt",
	}

	for i := 0; i < 3; i++ {
		filePath := dirPath + fileNames[i]
		file, err := os.Create(filePath)
		if err != nil {
			assert.NoError(t, err)
		} else {
			file.Close()
		}
	}

	err := RemoveFilesInDir(dirPath)
	assert.NoError(t, err)

	for i := 1; i < 3; i++ { // Making sure the all files has been removed
		filePath := dirPath + fileNames[i]
		f, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
		assert.Error(t, os.ErrNotExist, err)
		f.Close()
	}
}
