package filesystem

import (
	"bufio"
	"errors"
	"os"
)

// ReadFileToSlice reads a file located by filePath into a strings slice.
// The division into strings is based on bufio.ScanLines (`\r?\n` regexp).
func ReadFileToSlice(filePath string) ([]string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if info, err := file.Stat(); err != nil {
		return nil, err
	} else if info.Size() == 0 {
		return nil, errors.New("File is empty")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var rows []string // TO DO: need to optimize the allocation for the slice based on a priori knowledge.
	for fileScanner.Scan() {
		rows = append(rows, fileScanner.Text())
	}

	err = fileScanner.Err()

	return rows, err
}
