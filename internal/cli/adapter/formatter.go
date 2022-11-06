package adapter

import (
	"github.com/IMDb-searcher/internal/db/accessor"
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
}

func New() (IFormatter, error) {
	accessor, err := accessor.New()
	if err != nil {
		return nil, err
	}
	return &formatter{
		accessor: accessor,
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
