package accessor

import (
	"errors"
	"strings"

	"github.com/IMDb-searcher/internal/config"
	"github.com/IMDb-searcher/internal/db/accessor/utils"
	"github.com/IMDb-searcher/internal/db/models"
	"github.com/IMDb-searcher/internal/db/unpacker"
	"github.com/IMDb-searcher/internal/db/utils/filesystem"
	"github.com/IMDb-searcher/internal/db/utils/search"
	"github.com/IMDb-searcher/internal/logger"
)

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

	unpacker unpacker.IUnpacker
	log      logger.ILogger
}

func New(logger logger.ILogger) (IDBAccessor, error) {
	unpacker, err := unpacker.New(logger)
	if err != nil {
		return nil, err
	}
	accessor := &dBAccessor{
		dbTables: make(map[string][]string),
		unpacker: unpacker,
		log:      logger,
	}
	err = accessor.loadDB()
	if err != nil {
		return nil, err
	}

	return accessor, nil
}

var errNotFound error = errors.New("required values were not found")

func (d *dBAccessor) loadDB() error {
	cfg := config.GetConfig(d.log)

	pathToUnpackedFiles := cfg.GetDBPathToUnpackedFiles()

	for i := 0; i < len(cfg.GetDBFileNames()); i++ {
		fileName := cfg.GetDBFileNames()[i]

		table, err := filesystem.ReadFileToSlice(pathToUnpackedFiles + fileName) // TO DO: possible need for better error wrapping
		if err != nil {
			d.log.Error(err)
			return err
		}
		d.dbTables[fileName] = table
	}
	return nil
}

func (d *dBAccessor) findRelatedTitlesByNameBasicsRow(nameBasicsRow string) ([]models.TitleBasics, error) {
	titleBasicsTable := d.dbTables["title.basics.tsv"]

	// nconst	primaryName	birthYear	deathYear	primaryProfession	knownForTitles
	knownForTitlesFieldId := 5
	knownForTitlesField := search.GetValueByFieldIndexInRow(nameBasicsRow, knownForTitlesFieldId) //titles ids string, like: tt000001,tt000002

	titlesIds := strings.Split(knownForTitlesField, ",") // slice of titles indexes

	var titles []models.TitleBasics
	for _, titleId := range titlesIds {
		titleRow, err := search.GetRowByUniqueId(titleBasicsTable, titleId)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}

		title, err := models.CreateTitleBasics(titleRow)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}

		titles = append(titles, *title)
	}

	return titles, nil
}

func (d *dBAccessor) findPersonsByTheirIdsField(idsField string) ([]models.NameBasicsMain, error) {
	nameBasicsTable := d.dbTables["name.basics.tsv"]
	ids := strings.Split(idsField, ",")

	var perosns []models.NameBasicsMain
	for _, id := range ids {
		if id == `\N` {
			continue
		}
		row, err := search.GetRowByUniqueId(nameBasicsTable, id)
		if err != nil {
			return nil, err
		}
		nameBasicsModel, err := models.CreateNameBasicsMain(row)
		if err != nil {
			return nil, err
		}
		perosns = append(perosns, *nameBasicsModel)
	}
	return perosns, nil
}

func (d *dBAccessor) FindInfoByPersonName(name string) (IFormat, error) {
	nameBasicsTable := d.dbTables["name.basics.tsv"]

	personNameField := "primaryName"
	nameBasicsRow, err := search.GetRowByField(nameBasicsTable, personNameField, name) // row contains the person information
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	titles, err := d.findRelatedTitlesByNameBasicsRow(nameBasicsRow)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	return models.CreateNameBasics(nameBasicsRow, titles)
}

func (d *dBAccessor) FindTitleAndCastInfoByTitleName(titleName string) (IFormat, error) {
	titleBasicsTable := d.dbTables["title.basics.tsv"]
	titleAkasTable := d.dbTables["title.akas.tsv"]
	titleCrewTable := d.dbTables["title.crew.tsv"]

	titleNameField := "title"
	row, err := search.GetRowByField(titleAkasTable, titleNameField, titleName) // row contains the actor information
	if err != nil {
		return nil, err
	}

	titleId := search.GetValueByFieldIndexInRow(row, 0)
	if titleId == "" {
		return nil, errors.New("error gettring id")
	}

	titleCrewRow, err := search.GetRowByUniqueId(titleCrewTable, titleId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}
	// tconst	directors	writers
	dirFieldId := 1
	wrFieldId := 2

	directorsId := search.GetValueByFieldIndexInRow(titleCrewRow, dirFieldId)
	directorsModels, err := d.findPersonsByTheirIdsField(directorsId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	writersId := search.GetValueByFieldIndexInRow(titleCrewRow, wrFieldId)
	writersModels, err := d.findPersonsByTheirIdsField(writersId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	crew, err := models.CreateTitleCrew(titleCrewRow, directorsModels, writersModels)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	titleBasicsRow, err := search.GetRowByUniqueId(titleBasicsTable, titleId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	return models.CreateTitleBasicsWithCrew(titleBasicsRow, *crew)
}

func (d *dBAccessor) FindTitlesByPersonName(name string) (IFormat, error) {
	nameBasicsTable := d.dbTables["name.basics.tsv"]

	personNameField := "primaryName"
	nameBasicsRow, err := search.GetRowByField(nameBasicsTable, personNameField, name) // row contains the actor information
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	titles, err := d.findRelatedTitlesByNameBasicsRow(nameBasicsRow)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	return models.CreateTitlesBasics(titles), nil
}

func (d *dBAccessor) FindAllTitlesBySpecificYear(year string) (IFormat, error) {
	titleBasicsTable := d.dbTables["title.basics.tsv"]

	err := utils.VerifyYear(year)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	startYearId := 5
	endYearId := 6

	var titles []models.TitleBasics
	for i := 1; i < len(titleBasicsTable); i++ {
		row := titleBasicsTable[i]
		startyear := search.GetValueByFieldIndexInRow(row, startYearId)
		if startyear == year {
			endyear := search.GetValueByFieldIndexInRow(row, endYearId)
			if endyear == `\N` {
				title, err := models.CreateTitleBasics(row)
				if err != nil {
					d.log.Error(err)
					return nil, err
				}
				titles = append(titles, *title)
			}
		}
	}

	if titles == nil {
		d.log.Error(errNotFound)
		return nil, errNotFound
	}

	model := models.CreateTitlesBasics(titles)
	return model, nil
}
