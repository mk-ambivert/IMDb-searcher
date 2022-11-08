package filesystem

import (
	"io/ioutil"
	"os"
)

func RemoveFile(filePath string) error {
	return os.Remove(filePath)
}

func RemoveFilesInDir(dirPath string) error {
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := RemoveFile(dirPath + file.Name())
		if err != nil {
			return err
		}
	}

	return nil
}
