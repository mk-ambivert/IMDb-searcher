package models

import (
	"strings"
)

type TitleRatings struct {
	tconst        string
	averageRating string
	numVotes      string
}

func CreateTitleRatings(tableRow string) (*TitleRatings, error) {
	fields := strings.Split(tableRow, "\t")
	if fields == nil {
		return nil, errModelCreatinoError
	}

	return &TitleRatings{
		tconst:        fields[0],
		averageRating: fields[1],
		numVotes:      fields[2],
	}, nil
}
