package models

import (
	"strings"

	"gopkg.in/yaml.v2"
)

type NameBasics struct {
	nconst            string
	PrimaryName       string         `yaml:"primaryName"`
	BirthYear         string         `yaml:"birthYear"`
	DeathYear         string         `yaml:"deathYear"`
	PrimaryProfession string         `yaml:"primaryProffession"`
	KnownForTitles    []*TitleBasics `yaml:"knownForTitles"`
}

func (r *NameBasics) YAML() (string, error) {
	model, err := yaml.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(model), nil
}

func CreateNameBasics(tableRow string, knownForTitles []*TitleBasics) (*NameBasics, error) {
	fields := strings.Split(tableRow, "\t")
	if fields == nil {
		return nil, errModelCreatinoError
	}

	return &NameBasics{
		nconst:            fields[0],
		PrimaryName:       fields[1],
		BirthYear:         fields[2],
		DeathYear:         fields[3],
		PrimaryProfession: fields[4],
		KnownForTitles:    knownForTitles,
	}, nil
}

type NameBasicsMain struct {
	nconst            string
	PrimaryName       string `yaml:"primaryName"`
	BirthYear         string `yaml:"birthYear"`
	DeathYear         string `yaml:"deathYear"`
	PrimaryProfession string `yaml:"primaryProffession"`
}

func (r *NameBasicsMain) YAML() (string, error) {
	model, err := yaml.Marshal(r)
	if err != nil {
		return "", err
	}
	return string(model), nil
}

func CreateNameBasicsMain(tableRow string) (*NameBasicsMain, error) {
	fields := strings.Split(tableRow, "\t")
	if fields == nil {
		return nil, errModelCreatinoError
	}

	return &NameBasicsMain{
		nconst:            fields[0],
		PrimaryName:       fields[1],
		BirthYear:         fields[2],
		DeathYear:         fields[3],
		PrimaryProfession: fields[4],
	}, nil
}
