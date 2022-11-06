package models

type NameBasics struct {
	nconst            string
	PrimaryName       string
	BirthYear         string
	DeathYear         string
	PrimaryProfession string
	KnownForTitles    []TitleBasics
}

type NameBasicsMainInfo struct {
	nconst            string
	PrimaryName       string
	BirthYear         string
	DeathYear         string
	PrimaryProfession string
}
