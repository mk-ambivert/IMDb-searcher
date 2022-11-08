package unpacker

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/IMDb-searcher/internal/db/utils/filesystem"

	mconfig "github.com/IMDb-searcher/internal/config/mock"
	mlogger "github.com/IMDb-searcher/internal/logger/mock"
)

func TestUnGzipFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	config := mconfig.NewMockIConfig(ctrl)
	config.EXPECT().GetDBPathToPackedFiles().Return("test_files/compressed").AnyTimes()
	config.EXPECT().GetDBPathToUnpackedFiles().Return("test_files/unpacked").AnyTimes()
	config.EXPECT().GetDBFileNames().Return([]string{"title.ratings.tsv"}).AnyTimes()

	log := mlogger.NewMockILogger(ctrl)
	log.EXPECT().Error(gomock.Any()).Return().AnyTimes()
	log.EXPECT().Panic(gomock.Any()).Return().AnyTimes()

	unpacker := &unpacker{
		config: config,
		log:    log,
	}

	err := unpacker.UnGzipFiles()
	assert.NoError(t, err)

	err = filesystem.RemoveFile("test_files/unpacked/title.ratings.tsv")
	assert.NoError(t, err)
}
