package unpacker

import (
	"errors"
	"os"

	"github.com/IMDb-searcher/internal/config"
	"github.com/IMDb-searcher/internal/db/utils/filesystem"
	"github.com/IMDb-searcher/internal/logger"
)

type IUnpacker interface {
	UnGzipFiles() error
}

type unpacker struct {
	log    logger.ILogger
	config config.IConfig
}

func (u *unpacker) removeExtractedFiles() {
	pathToUnpackedDBFiles := u.config.GetDBPathToUnpackedFiles()

	err := filesystem.RemoveFilesInDir(pathToUnpackedDBFiles)
	if err != nil {
		u.log.Error(err)
		return
	}
}

func (u *unpacker) extractFiles() error {
	pathToPackedDBFiles := u.config.GetDBPathToPackedFiles()
	pathToUnpackedDBFiles := u.config.GetDBPathToUnpackedFiles()
	dbFileNames := u.config.GetDBFileNames()

	for i := 0; i < len(dbFileNames); i++ {
		from := pathToPackedDBFiles + dbFileNames[i]
		to := pathToUnpackedDBFiles + dbFileNames[i]

		err := filesystem.UnGzip(from, to)
		if err != nil {
			u.log.Error(err)
			return err
		}
	}

	return nil
}

func (u *unpacker) verifyFiles() error {
	pathToUnpackedDBFiles := u.config.GetDBPathToUnpackedFiles()
	dbFileNames := u.config.GetDBFileNames()

	for i := 0; i < len(dbFileNames); i++ {
		filePath := pathToUnpackedDBFiles + dbFileNames[i]
		if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
			u.log.Error("verifying error:", err)
			return err
		}
		info, err := os.Stat(filePath)
		if err != nil {
			u.log.Error("verifiying error:", err)
		}
		if info.Size() < 1000 { // TO DO: make better
			err = errors.New("verifiying error: file is suspiciously small")
			u.log.Error(err)
			return err
		}
	}
	return nil
}

func (u *unpacker) UnGzipFiles() error {
	err := u.extractFiles()
	if err != nil {
		u.log.Error(err)

		u.removeExtractedFiles()
		return err
	}

	err = u.verifyFiles()
	if err != nil {
		u.log.Error(err)

		u.removeExtractedFiles()
		return err
	}

	return nil
}

func New(logger logger.ILogger) (IUnpacker, error) {
	cfg := config.GetConfig(logger)
	return &unpacker{
		log:    logger,
		config: cfg,
	}, nil
}
