package cli

import (
	"github.com/IMDb-searcher/internal/cli/adapter"
)

type IUserInterface interface {
	Run()
}

type UserInterface struct {
	IUserInterface

	formatter adapter.IFormatter
}

func (u *UserInterface) Run() {

}

func New() (IUserInterface, error) {
	formatter, err := adapter.New()
	if err != nil {
		return nil, err
	}
	return &UserInterface{
		formatter: formatter,
	}, nil
}
