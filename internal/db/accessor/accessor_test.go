package accessor

import (
	"path"
	"testing"

	"github.com/IMDb-searcher/internal/db/models"
	"github.com/IMDb-searcher/internal/db/utils/filesystem"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	munpacker "github.com/IMDb-searcher/internal/db/unpacker/mock"
	mlogger "github.com/IMDb-searcher/internal/logger/mock"
)

func loadDB(t *testing.T) map[string][]string {
	files := []string{
		"name.basics.tsv",
		"title.basics.tsv",
		"title.akas.tsv",
		"title.crew.tsv",
		"title.principals.tsv",
		"title.ratings.tsv",
		"title.episode.tsv",
	}
	dbPath := "db_test_files"

	db := make(map[string][]string)
	for _, file := range files {
		tablePath := path.Join(dbPath, file)
		rows, err := filesystem.ReadFileToSlice(tablePath)
		assert.NoError(t, err)
		db[file] = rows
	}
	return db
}

func TestFindInfoByPersonName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mlogger.NewMockILogger(ctrl)
	log.EXPECT().Error(gomock.Any()).Return().AnyTimes()

	unpacker := munpacker.NewMockIUnpacker(ctrl)

	accessor := &dBAccessor{
		dbTables: loadDB(t),
		unpacker: unpacker,
		log:      log,
	}

	title1, _ := models.CreateTitleBasics(
		`tt0000001	short	Carmencita	Carmencita	0	1894	\N	1	Documentary,Short`)
	title2, _ := models.CreateTitleBasics(
		`tt0000002	short	Le clown et ses chiens	Le clown et ses chiens	0	1892	\N	5	Animation,Short`)
	title3, _ := models.CreateTitleBasics(
		`tt0000003	short	Pauvre Pierrot	Pauvre Pierrot	0	1892	\N	4	Animation,Comedy,Romance`)
	title5, _ := models.CreateTitleBasics(
		`tt0000005	tvEpisode	Spécial Pétain - Laval	Spécial Pétain - Laval	0	1993	\N	\N	Documentary,Talk-Show`)
	titles := []*models.TitleBasics{title1, title2, title3, title5}

	titlesBasics := models.CreateTitlesBasics(titles)
	nameBasicsRow := `nm0000006	Ingmar Bergman	1984	2022	writer,director,actor	tt0000001,tt0000002,tt0000003,tt0000005`
	expected, _ := models.CreateNameBasics(nameBasicsRow, titlesBasics.Titles)

	personName := "Ingmar Bergman"

	model, err := accessor.FindInfoByPersonName(personName)
	assert.NoError(t, err)
	assert.Equal(t, expected, model)
}

func TestFindTitleAndCastInfoByTitleName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mlogger.NewMockILogger(ctrl)
	log.EXPECT().Error(gomock.Any()).Return().AnyTimes()

	unpacker := munpacker.NewMockIUnpacker(ctrl)

	accessor := &dBAccessor{
		dbTables: loadDB(t),
		unpacker: unpacker,
		log:      log,
	}

	actor1, _ := models.CreateNameBasicsMain(
		`nm0000001	Fred Astaire	1905	1985	soundtrack,actor,miscellaneous	tt0000001,tt0000002,tt0000003,tt0000004`)
	actor2, _ := models.CreateNameBasicsMain(
		`nm0000004	John Belushi	1999	1982	actor,editor,writer tt0000003,tt0000002`)
	ratings, _ := models.CreateTitleRatings(
		`tt0000005	5.6	20`)
	expected, _ := models.CreateTitleInfoWithActors(
		`tt0000005	tvEpisode	Spécial Pétain - Laval	Spécial Pétain - Laval	0	1993	\N	\N	Documentary,Talk-Show`,
		ratings,
		[]*models.NameBasicsMain{actor1, actor2})

	titleName := "एपिसोड #1.1733"
	model, err := accessor.FindTitleAndCastInfoByTitleName(titleName)
	assert.NoError(t, err)
	assert.Equal(t, expected, model)
}

func TestFindTitlesByPersonName(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mlogger.NewMockILogger(ctrl)
	log.EXPECT().Error(gomock.Any()).Return().AnyTimes()

	unpacker := munpacker.NewMockIUnpacker(ctrl)

	accessor := &dBAccessor{
		dbTables: loadDB(t),
		unpacker: unpacker,
		log:      log,
	}
	title1, _ := models.CreateTitleBasics(
		`tt0000001	short	Carmencita	Carmencita	0	1894	\N	1	Documentary,Short`)
	title2, _ := models.CreateTitleBasics(
		`tt0000002	short	Le clown et ses chiens	Le clown et ses chiens	0	1892	\N	5	Animation,Short`)
	title3, _ := models.CreateTitleBasics(
		`tt0000003	short	Pauvre Pierrot	Pauvre Pierrot	0	1892	\N	4	Animation,Comedy,Romance`)
	title5, _ := models.CreateTitleBasics(
		`tt0000005	tvEpisode	Spécial Pétain - Laval	Spécial Pétain - Laval	0	1993	\N	\N	Documentary,Talk-Show`)
	titles := []*models.TitleBasics{title1, title2, title3, title5}
	expected := models.CreateTitlesBasics(titles)

	personName := "Ingmar Bergman"
	model, err := accessor.FindTitlesByPersonName(personName)
	assert.NoError(t, err)
	assert.Equal(t, expected, model)
}

func TestFindAllTitlesBySpecificYear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mlogger.NewMockILogger(ctrl)
	log.EXPECT().Error(gomock.Any()).Return().AnyTimes()

	unpacker := munpacker.NewMockIUnpacker(ctrl)

	accessor := &dBAccessor{
		dbTables: loadDB(t),
		unpacker: unpacker,
		log:      log,
	}

	titleRow := `tt0000001	short	Carmencita	Carmencita	0	1894	\N	1	Documentary,Short`
	title, _ := models.CreateTitleBasics(titleRow)
	expected := models.CreateTitlesBasics([]*models.TitleBasics{title})

	year := "1894"
	model, err := accessor.FindAllTitlesBySpecificYear(year)
	assert.NoError(t, err)
	assert.Equal(t, expected, model)
}
