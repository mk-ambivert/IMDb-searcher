package cli

import (
	"github.com/IMDb-searcher/internal/cli/adapter"
	"github.com/IMDb-searcher/internal/logger"
)

type IUserInterface interface {
	Run()
}

type UserInterface struct {
	IUserInterface

	formatter adapter.IFormatter
	log       logger.ILogger
}

func (u *UserInterface) Run() {

}

func New(logger logger.ILogger) (IUserInterface, error) {
	formatter, err := adapter.New(logger)
	if err != nil {
		return nil, err
	}
	return &UserInterface{
		formatter: formatter,
		log:       logger,
	}, nil
}
