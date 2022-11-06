package accessor

import "github.com/IMDb-searcher/internal/logger"

type IFormat interface {
	YAML() (string, error)
}

type IDBAccessor interface {
	FindInfoByPersonName(name string) (IFormat, error)
	FindTitleAndCastInfoByTitleName(name string) (IFormat, error)
	FindTitlesByPersonName(name string) (IFormat, error)
	FindAllTitlesBySpecificYear(year string) (IFormat, error)
}

type dBAccessor struct {
	dbTables map[string][]string

	IDBAccessor
	log logger.ILogger
}

func New(logger logger.ILogger) (IDBAccessor, error) {
	accessor := &dBAccessor{
		dbTables: make(map[string][]string),
		log:      logger,
	}
	err := accessor.loadDB()
	if err != nil {
		return nil, err
	}

	return accessor, nil
}

func (d *dBAccessor) loadDB() error {
	return nil
}

func (d *dBAccessor) FindInfoByPersonName(name string) (IFormat, error) {
	return nil, nil
}

func (d *dBAccessor) FindTitleAndCastInfoByTitleName(titleName string) (IFormat, error) {
	return nil, nil
}

func (d *dBAccessor) FindTitlesByPersonName(name string) (IFormat, error) {
	return nil, nil
}

func (d *dBAccessor) FindAllTitlesBySpecificYear(year string) (IFormat, error) {
	return nil, nil
}
