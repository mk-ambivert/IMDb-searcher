package models

import (
	"strings"

	"gopkg.in/yaml.v2"
)

type TitleBasics struct {
	tconst         string
	TitleType      string   `yaml:"titleType"`
	PrimaryTitle   string   `yaml:"primaryTitle"`
	OriginalTitle  string   `yaml:"originalTitle"`
	IsAdult        string   `yaml:"isAdult"`
	StartYear      string   `yaml:"startYear"`
	EndYear        string   `yaml:"endYear"`
	RuntimeMinutes string   `yaml:"runtimeMinutes"`
	Genres         []string `yaml:"genres"`
}

func CreateTitleBasics(tableRow string) (*TitleBasics, error) {
	fields := strings.Split(tableRow, "\t")
	if fields == nil {
		return nil, errModelCreatinoError
	}
	if len(fields) != 9 {
		return nil, errModelCreatinoError
	}

	var genres []string
	if fields[8] != `\N` {
		genres = strings.Split(fields[8], ",")
		if genres == nil {
			return nil, errModelCreatinoError
		}
	}

	return &TitleBasics{
		tconst:         fields[0],
		TitleType:      fields[1],
		PrimaryTitle:   fields[2],
		OriginalTitle:  fields[3],
		IsAdult:        fields[4],
		StartYear:      fields[5],
		EndYear:        fields[6],
		RuntimeMinutes: fields[7],
		Genres:         genres,
	}, nil
}

type TitleInfoWithActors struct {
	Title   *TitleBasics      `yaml:"title"`
	Ratings *TitleRatings     `yaml:"ratings"`
	Actors  []*NameBasicsMain `yaml:"actors"`
}

func (t *TitleInfoWithActors) YAML() (string, error) {
	model, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(model), nil
}

func CreateTitleInfoWithActors(tableRow string,
	ratings *TitleRatings,
	actors []*NameBasicsMain) (*TitleInfoWithActors, error) {
	titleModel, err := CreateTitleBasics(tableRow)
	if err != nil {
		return nil, err
	}

	return &TitleInfoWithActors{
		Title:   titleModel,
		Ratings: ratings,
		Actors:  actors,
	}, nil
}

type TitlesBasics struct {
	Titles []*TitleBasics `yaml:"titles"`
}

func (t *TitlesBasics) YAML() (string, error) {
	model, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}
	return string(model), nil
}

func CreateTitlesBasics(titles []*TitleBasics) *TitlesBasics {
	return &TitlesBasics{
		Titles: titles,
	}
}
