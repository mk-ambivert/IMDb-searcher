package models

import (
	"strings"

	"gopkg.in/yaml.v2"
)

type TitleCrew struct {
	tconst    string
	directors []*NameBasicsMain
	writers   []*NameBasicsMain
}

func (t *TitleCrew) YAML() (string, error) {
	model, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(model), nil
}

func CreateTitleCrew(tableRow string, directors, wirters []*NameBasicsMain) (*TitleCrew, error) {
	fields := strings.Split(tableRow, "\t")
	if fields == nil {
		return nil, errModelCreatinoError
	}

	return &TitleCrew{
		tconst:    fields[0],
		directors: directors,
		writers:   wirters,
	}, nil
}
