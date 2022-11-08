package accessor

import (
	"testing"

	"github.com/IMDb-searcher/internal/db/models"
	"github.com/IMDb-searcher/internal/db/utils/filesystem"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock "github.com/IMDb-searcher/internal/logger/mock"
)

func loadDB() map[string][]string {
	dbTables := make(map[string][]string)
	pathToDB := "db_test_files"

	fileNames := []string{
		"name.basics.tsv",
		"title.basics.tsv",
		"title.akas.tsv",
		"title.crew.tsv",
		"title.principals.tsv",
		"title.ratings.tsv",
		"title.episode.tsv",
	}

	for i := 0; i < len(fileNames); i++ {
		fileName := fileNames[i]

		table, _ := filesystem.ReadFileToSlice(pathToDB + fileName) // TO DO: possible need for better error wrapping
		dbTables[fileName] = table
	}

	return nil
}

func TestFindAllTitlesBySpecificYear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mock.NewMockILogger(ctrl)
	titleRow := `tt0000001	short	Carmencita	Carmencita	0	1894	\N	1	Documentary,Short`
	title, _ := models.CreateTitleBasics(titleRow)
	expected := models.CreateTitlesBasics([]*models.TitleBasics{title})
	year := "1894"

	accessor, err := New(log)
	assert.NoError(t, err)

	titles, err := accessor.FindAllTitlesBySpecificYear(year)
	assert.NoError(t, err)
	assert.Equal(t, expected, titles)
}

func TestFindTitleAndCastInfoByTitleName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mock.NewMockILogger(ctrl)

	var accessor IDBAccessor = &dBAccessor{
		log:      log,
		dbTables: loadDB(),
	}

	model, err := accessor.FindTitleAndCastInfoByTitleName("titleName")
	assert.NoError(t, err)
	assert.Equal(t, nil, model)
}

func TestFindTitlesByPersonName(t *testing.T) {

}

func TestFindInfoByPersonName(t *testing.T) {
	title1, _ := models.CreateTitleBasics(`tt0000001	short	Carmencita	Carmencita	0	1894	\N	1	Documentary,Short`)
	title2, _ := models.CreateTitleBasics(`tt0000002	short	Le clown et ses chiens	Le clown et ses chiens	0	1892	\N	5	Animation,Short`)
	title3, _ := models.CreateTitleBasics(`tt0000003	short	Pauvre Pierrot	Pauvre Pierrot	0	1892	\N	4	Animation,Comedy,Romance`)
	title5, _ := models.CreateTitleBasics(`tt0000005	tvEpisode	Spécial Pétain - Laval	Spécial Pétain - Laval	0	1993	\N	\N	Documentary,Talk-Show`)
	titles := []*models.TitleBasics{title1, title2, title3, title5}

	titlesBasics := models.CreateTitlesBasics(titles)
	nameBasicsRow := `nm0000006	Ingmar Bergman	1984	2022	writer,director,actor	tt0000001,tt0000002,tt0000003,tt0000005`
	expected, _ := models.CreateNameBasics(nameBasicsRow, titlesBasics.Titles)
	personName := "Ingmar Bergman"

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mock.NewMockILogger(ctrl)

	var accessor IDBAccessor = &dBAccessor{
		log:      log,
		dbTables: loadDB(),
	}

	model, err := accessor.FindInfoByPersonName(personName)
	assert.NoError(t, err)
	assert.Equal(t, expected, model)
}
