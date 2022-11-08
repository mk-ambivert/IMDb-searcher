package filesystem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnGzip(t *testing.T) {
	pathToPackedFile := "test_files/compressed/name.basics.tsv.gz"
	pathToUnpackedFile := "test_files/name.basics.tsv"
	expected := []string{
		`nconst	primaryName	birthYear	deathYear	primaryProfession	knownForTitles`,
		`nm0000001	Fred Astaire	1905	1985	soundtrack,actor,miscellaneous	tt0053137,tt0031983,tt0050419,tt0072308`,
		`nm0000002	Lauren Bacall	1930	2020	actress,writer	tt0071877,tt0037382,tt0117057,tt0038355`,
		`nm0000003	Brigitte Bardot	2000	\N	actress,director,producer	tt0056404,tt0057345,tt0049189,tt0054452`,
		`nm0000004	John Belushi	1999	1982	actor,editor,writer	tt0077975,tt0078723,tt0072562,tt0080455`,
		`nm0000005	Ingmar Bergman	1984	2022	writer,director,actor	tt0069467,tt0050986,tt0060827,tt0050976`}

	err := UnGzip(pathToPackedFile, pathToUnpackedFile)
	assert.NoError(t, err)
	defer RemoveFile(pathToUnpackedFile)

	fileRows, err := ReadFileToSlice(pathToUnpackedFile)

	assert.NoError(t, err)
	assert.NotNil(t, fileRows)
	for i, row := range fileRows {
		assert.Equal(t, expected[i], row)
	}
}
