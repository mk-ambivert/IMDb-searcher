package config

import (
	"os"
	"path"
	"path/filepath"
	"sync"

	"github.com/IMDb-searcher/internal/logger"
	"github.com/ilyakaznacheev/cleanenv"
)

type IConfig interface {
	GetDBPathToPackedFiles() string
	GetDBPathToUnpackedFiles() string
	GetDBFileNames() []string
}

type Config struct {
	DBInfo dataBaseInfo `yaml:"DataBaseInfo"`
}

type dataBaseInfo struct {
	DBPaths     databasePaths `yaml:"DatabasePaths"`
	DBFileNames []string      `yaml:"DatabaseFileNames"`
}

type databasePaths struct {
	PathToPackedDBFiles   string `yaml:"pathToPackedDBFiles" env:"PROJECT_DIR"`
	PathToUnpackedDBFiles string `yaml:"pathToUnpackedDBFiles" env:"PROJECT_DIR"`
}

func (c *Config) GetDBPathToPackedFiles() string {
	return c.DBInfo.DBPaths.PathToPackedDBFiles
}

func (c *Config) GetDBPathToUnpackedFiles() string {
	return c.DBInfo.DBPaths.PathToUnpackedDBFiles
}

func (c *Config) GetDBFileNames() []string {
	return c.DBInfo.DBFileNames
}

var once sync.Once
var config *Config

func GetConfig(log logger.ILogger) IConfig {
	once.Do(func() {
		projectDir := os.Getenv("PROJECT_DIR")
		configName := "config.yml"
		pathConfig := filepath.Join(projectDir, configName)

		config = &Config{}
		err := cleanenv.ReadConfig(pathConfig, config)
		if err != nil {
			log.Error(err)
			help, _ := cleanenv.GetDescription(config, nil)
			log.Panic(help)
		}
		config.DBInfo.DBPaths.PathToPackedDBFiles = path.Join(config.DBInfo.DBPaths.PathToPackedDBFiles, `database`)
		config.DBInfo.DBPaths.PathToUnpackedDBFiles = path.Join(config.DBInfo.DBPaths.PathToUnpackedDBFiles, `database\unpacked`)
	})

	return config
}
