package unpacker

import (
	"errors"
	"os"
	"path"

	"github.com/IMDb-searcher/internal/config"
	"github.com/IMDb-searcher/internal/db/utils/filesystem"
	"github.com/IMDb-searcher/internal/logger"

	e "github.com/IMDb-searcher/internal/errors"
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
	err := os.MkdirAll(pathToUnpackedDBFiles, os.ModeDir)
	if err != nil {
		return err
	}

	for i := 0; i < len(dbFileNames); i++ {
		packageType := ".gz"
		from := path.Join(pathToPackedDBFiles, dbFileNames[i]+packageType)
		to := path.Join(pathToUnpackedDBFiles, dbFileNames[i])

		err := filesystem.UnGzip(from, to)
		if err != nil {
			u.log.Error(err)
			return err
		}
	}

	return nil
}

func (u *unpacker) verifyUnpackedFiles() error {
	pathToUnpackedDBFiles := u.config.GetDBPathToUnpackedFiles()
	dbFileNames := u.config.GetDBFileNames()

	for i := 0; i < len(dbFileNames); i++ {
		filePath := path.Join(pathToUnpackedDBFiles, dbFileNames[i])
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
	err := u.verifyUnpackedFiles()
	if err == nil {
		return nil // All is well, the base is already unpacked
	}

	err = u.extractFiles()
	if err != nil {
		u.log.Error(err)

		u.removeExtractedFiles()
		return &e.ErrDataBaseUnpacking{}
	}

	err = u.verifyUnpackedFiles()
	if err != nil {
		u.log.Error(err)

		u.removeExtractedFiles()
		return &e.ErrDataBaseUnpacking{}
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
