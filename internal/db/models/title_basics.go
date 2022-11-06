package models

type TitleBasics struct {
	tconst         string
	TitleType      string
	PrimaryTitle   string
	OriginalTitle  string
	IsAdult        string
	StartYear      string
	EndYear        string
	RuntimeMinutes string
	Genres         []string
}

type TitlesBasics struct {
	Titles []TitleBasics
}
