package accessor

import (
	"path"
	"strings"

	"github.com/IMDb-searcher/internal/config"
	"github.com/IMDb-searcher/internal/errors"
	"github.com/IMDb-searcher/internal/logger"

	"github.com/IMDb-searcher/internal/db/accessor/utils"
	"github.com/IMDb-searcher/internal/db/models"
	"github.com/IMDb-searcher/internal/db/unpacker"
	"github.com/IMDb-searcher/internal/db/utils/filesystem"
)

type IDBAccessor interface {
	FindInfoByPersonName(name string) (models.IFormat, error)
	FindTitleAndCastInfoByTitleName(name string) (models.IFormat, error)
	FindTitlesByPersonName(name string) (models.IFormat, error)
	FindAllTitlesBySpecificYear(year string) (models.IFormat, error)
}

type dBAccessor struct {
	dbTables map[string][]string

	unpacker unpacker.IUnpacker
	log      logger.ILogger
}

// loadDB uses IUnpacker to unpack and verify the database, and then reads the database into dbTables
func (d *dBAccessor) loadDB() error {
	err := d.unpacker.UnGzipFiles()
	if err != nil {
		d.log.Error(err)
		return err
	}

	cfg := config.GetConfig(d.log)

	pathToUnpackedFiles := cfg.GetDBPathToUnpackedFiles()

	for i := 0; i < len(cfg.GetDBFileNames()); i++ {
		fileName := cfg.GetDBFileNames()[i]

		table, err := filesystem.ReadFileToSlice(path.Join(pathToUnpackedFiles, fileName))
		if err != nil {
			d.log.Error(err)
			return &errors.ErrDataBaseLoading{}
		}
		d.dbTables[fileName] = table
	}
	return nil
}

// findRelatedTitlesByNameBasicsRow finds all references to titles in table "title.basics.tsv"
func (d *dBAccessor) findRelatedTitlesByNameBasicsRow(nameBasicsRow string) ([]*models.TitleBasics, error) {
	titleBasicsTable := d.dbTables["title.basics.tsv"]

	// nconst	primaryName	birthYear	deathYear	primaryProfession	knownForTitles
	knownForTitlesFieldId := 5
	knownForTitlesField := utils.GetValueByFieldIndexInRow(nameBasicsRow, knownForTitlesFieldId) //titles ids string, like: tt000001,tt000002

	if !utils.IsReferenceExists(knownForTitlesField) {
		d.log.Info("row does not contain references to other fields")
		return nil, nil
	}

	titlesIds := strings.Split(knownForTitlesField, ",") // slice of titles indexes

	var titles []*models.TitleBasics
	for _, titleId := range titlesIds {
		titleRow, err := utils.GetRowByUniqueId(titleBasicsTable, titleId)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}

		title, err := models.CreateTitleBasics(titleRow)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}

		titles = append(titles, title)
	}

	return titles, nil
}

// findActorsByTitle finds all the actors involved in the title
func (d *dBAccessor) findActorsByTitle(titleId string) ([]*models.NameBasicsMain, error) {
	nameBasicsTable := d.dbTables["name.basics.tsv"]
	titlePrincipalsTable := d.dbTables["title.principals.tsv"]

	principalsRows, err := utils.GetRowsById(titlePrincipalsTable, titleId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	//"title.principals.tsv": tconst	ordering	nconst	category	job	characters
	nconstFieldId := 2
	categoryFieldId := 3
	desiredCategoryMale := "actor"
	desiredCategoryFemale := "actress"

	var actorsModels []*models.NameBasicsMain
	for _, row := range principalsRows {
		category := utils.GetValueByFieldIndexInRow(row, categoryFieldId)
		if category != desiredCategoryMale && category != desiredCategoryFemale {
			continue
		}
		id := utils.GetValueByFieldIndexInRow(row, nconstFieldId)
		actorRow, err := utils.GetRowByUniqueId(nameBasicsTable, id)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}
		model, err := models.CreateNameBasicsMain(actorRow)
		if err != nil {
			d.log.Error(err)
			return nil, err
		}
		actorsModels = append(actorsModels, model)
	}

	return actorsModels, nil
}

// findTitleRatings finds title ratings by titleId in the table title.ratings.tsv
func (d *dBAccessor) findTitleRatings(titleId string) (*models.TitleRatings, error) {
	titleRatingsTable := d.dbTables["title.ratings.tsv"]

	ratingsRow, err := utils.GetRowByUniqueId(titleRatingsTable, titleId)
	if err != nil {
		if _, ok := err.(*errors.ErrNotFound); ok {
			d.log.Info(err)
			return nil, nil
		}
	}
	ratingsModel, err := models.CreateTitleRatings(ratingsRow)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}
	return ratingsModel, nil
}

// FindInfoByPersonName finds all the information about a person's name
// (including the names of titles in which he or she has participated)
func (d *dBAccessor) FindInfoByPersonName(name string) (models.IFormat, error) {
	nameBasicsTable := d.dbTables["name.basics.tsv"]

	personNameField := "primaryName"
	nameBasicsRow, err := utils.GetRowByField(nameBasicsTable, personNameField, name) // row contains the person information
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

// FindTitleAndCastInfoByTitleName finds all the information about a title, including its rating and actors.
func (d *dBAccessor) FindTitleAndCastInfoByTitleName(titleName string) (models.IFormat, error) {
	titleBasicsTable := d.dbTables["title.basics.tsv"]
	titleAkasTable := d.dbTables["title.akas.tsv"]

	// titleId	ordering	title	region	language	types	attributes	isOriginalTitle
	titleNameField := "title"
	titleIdFieldId := 0

	titleAkasRow, err := utils.GetRowByField(titleAkasTable, titleNameField, titleName) // row contains the actor information
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	titleId := utils.GetValueByFieldIndexInRow(titleAkasRow, titleIdFieldId)

	actorsModels, err := d.findActorsByTitle(titleId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	if actorsModels == nil {
		d.log.Info("the title has no actors")
	}

	ratingsModel, err := d.findTitleRatings(titleId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	if ratingsModel == nil {
		d.log.Info("the title has no ratings")
	}

	titleRow, err := utils.GetRowByUniqueId(titleBasicsTable, titleId)
	if err != nil {
		d.log.Error(err)
		return nil, err
	}

	return models.CreateTitleInfoWithActors(titleRow, ratingsModel, actorsModels)
}

// FindTitlesByPersonName finds basic information about all titles by person name.
func (d *dBAccessor) FindTitlesByPersonName(name string) (models.IFormat, error) {
	nameBasicsTable := d.dbTables["name.basics.tsv"]

	personNameField := "primaryName"
	nameBasicsRow, err := utils.GetRowByField(nameBasicsTable, personNameField, name) // row contains the actor information
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

// FindAllTitlesBySpecificYear finds basic information about all titles shot in a particular year.
func (d *dBAccessor) FindAllTitlesBySpecificYear(year string) (models.IFormat, error) {
	titleBasicsTable := d.dbTables["title.basics.tsv"]

	err := utils.VerifyYear(year)
	if err != nil {
		d.log.Error(err)
		return nil, &errors.ErrBadYearFormat{}
	}

	// tconst	titleType	primaryTitle	originalTitle	isAdult	startYear	endYear	runtimeMinutes	genres
	startYearId := 5
	endYearId := 6

	var titles []*models.TitleBasics
	for i := 1; i < len(titleBasicsTable); i++ {
		row := titleBasicsTable[i]
		startyear := utils.GetValueByFieldIndexInRow(row, startYearId)
		if startyear == year {
			endyear := utils.GetValueByFieldIndexInRow(row, endYearId)
			if endyear == `\N` {
				title, err := models.CreateTitleBasics(row)
				if err != nil {
					d.log.Error(err)
					return nil, err
				}
				titles = append(titles, title)
			}
		}
	}

	if titles == nil {
		d.log.Error(&errors.ErrNotFound{})
		return nil, &errors.ErrNotFound{}
	}

	model := models.CreateTitlesBasics(titles)
	return model, nil
}

func New(logger logger.ILogger) (IDBAccessor, error) {
	unpacker, err := unpacker.New(logger)
	if err != nil {
		logger.Error("unpacker creation error:", err)
		return nil, err
	}
	accessor := &dBAccessor{
		dbTables: make(map[string][]string),
		unpacker: unpacker,
		log:      logger,
	}
	err = accessor.loadDB()
	if err != nil {
		accessor.log.Error("database loading error", err)
		return nil, err
	}

	return accessor, nil
}
