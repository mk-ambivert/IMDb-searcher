package formatter

import (
	"github.com/IMDb-searcher/internal/db/accessor"
	"github.com/IMDb-searcher/internal/db/models"
	"github.com/IMDb-searcher/internal/logger"

	e "github.com/IMDb-searcher/internal/errors"
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

func backendErrorWrapper(err error) error {
	switch err.(type) {
	case *e.ErrDataBaseUnpacking:
		return err
	case *e.ErrDataBaseVerifying:
		return err
	case *e.ErrDataBaseLoading:
		return err
	case *e.ErrNotFound:
		return err
	case *e.ErrBadYearFormat:
		return err
	default:
		return &e.ErrDefaultRequestProcessing{}
	}
}

func (f *formatter) requestWrapper(request string, requestHandler func(string) (models.IFormat, error)) (string, error) {
	responce, err := requestHandler(request)
	if err != nil {
		f.log.Error(err)
		return "", backendErrorWrapper(err)
	}
	yaml, err := responce.YAML()
	if err != nil {
		f.log.Error(err)
		return "", backendErrorWrapper(err)
	}

	return yaml, nil
}

func (f *formatter) FindInfoByPersonName(name string) (string, error) {
	return f.requestWrapper(name, f.accessor.FindInfoByPersonName)
}

func (f *formatter) FindTitleAndCastInfoByTitleName(title string) (string, error) {
	return f.requestWrapper(title, f.accessor.FindTitleAndCastInfoByTitleName)
}

func (f *formatter) FindTitlesByPersonName(name string) (string, error) {
	return f.requestWrapper(name, f.accessor.FindTitlesByPersonName)
}

func (f *formatter) FindAllTitlesBySpecificYear(year string) (string, error) {
	return f.requestWrapper(year, f.accessor.FindAllTitlesBySpecificYear)
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
