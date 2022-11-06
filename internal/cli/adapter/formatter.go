package adapter

import (
	"github.com/IMDb-searcher/internal/db/accessor"
	"github.com/IMDb-searcher/internal/logger"
)

type IFormatter interface {
	FindInfoByPersonName(string) (string, error)
	FindTitleAndCastInfoByTitleName(string) (string, error)
	FindTitlesByPersonName(string) (string, error)
	FindAllTitlesBySpecificYear(string) (string, error)
}

type formatter struct {
	IFormatter

	accessor accessor.IDBAccessor
	log      logger.ILogger
}

func New(logger logger.ILogger) (IFormatter, error) {
	accessor, err := accessor.New(logger)
	if err != nil {
		return nil, err
	}
	return &formatter{
		accessor: accessor,
		log:      logger,
	}, nil
}

func (f *formatter) FindInfoByPersonName(name string) (string, error) {
	return "", nil
}

func (f *formatter) FindTitleAndCastInfoByTitleName(title string) (string, error) {
	return "", nil
}

func (f *formatter) FindTitlesByPersonName(name string) (string, error) {
	return "", nil
}

func (f *formatter) FindAllTitlesBySpecificYear(year string) (string, error) {
	return "", nil
}
