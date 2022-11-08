package filesystem

import (
	"compress/gzip"
	"io"
	"os"
)

func UnGzip(sourcePath, targetPath string) error {
	reader, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	writer, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	if err != nil {
		return err
	}

	return nil
}
