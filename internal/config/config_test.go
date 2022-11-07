package config

import (
	"testing"

	mock "github.com/IMDb-searcher/internal/logger/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	expected := &Config{
		DBInfo: dataBaseInfo{
			DBPaths: databasePaths{
				PathToPackedDBFiles:   `database`,
				PathToUnpackedDBFiles: `database\unpacked`,
			},
			DBFileNames: []string{
				"name.basics.tsv",
				"title.basics.tsv",
				"title.akas.tsv",
				"title.crew.tsv",
				"title.principals.tsv",
				"title.ratings.tsv",
				"title.episode.tsv",
			},
		},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	log := mock.NewMockILogger(ctrl)

	cfg := GetConfig(log)
	assert.Equal(t, expected, cfg)
}
