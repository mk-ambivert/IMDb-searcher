package unpacker

import "github.com/IMDb-searcher/internal/logger"

type IUnpacker interface {
	UnGzipFiles() error
}

type unpacker struct {
}

func (u *unpacker) removeExtractedFiles() {

}

func (u *unpacker) extractFiles() error {
	return nil
}

func (u *unpacker) verifyFiles() error {
	return nil
}

func (u *unpacker) UnGzipFiles() error {
	return nil
}

func New(logger logger.ILogger) (IUnpacker, error) {
	return &unpacker{}, nil
}
