package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var uniqueIdTable = []string{
	`tconst	titleType	primaryTitle	originalTitle	isAdult	startYear	endYear	runtimeMinutes	genres`,
	`tt0000001	short	Carmencita	Carmencita	0	1894	\N	1	Documentary,Short`,
	`tt0000002	short	Le clown et ses chiens	Le clown et ses chiens	0	1892	\N	5	Animation,Short`,
	`tt0000004	short	Un bon bock	Un bon bock	0	1892	\N	12	Animation,Short`,
	`tt0000005	short	Blacksmith Scene	Blacksmith Scene	0	1893	\N	1	Comedy,Short`,
	`tt0000006	short	Chinese Opium Den	Chinese Opium Den	0	1894	\N	1	Short`,
	`tt0000007	short	Corbett and Courtney Before the Kinetograph	Corbett and Courtney Before the Kinetograph	0	1894	\N	1	Short,Sport`,
	`tt0000008	short	Edison Kinetoscopic Record of a Sneeze	Edison Kinetoscopic Record of a Sneeze	0	1894	\N	1	Documentary,Short`}

var multyIdTable = []string{
	`tconst	ordering	nconst	category	job	characters`,
	`tt0000001	1	nm1588970	self	\N	["Self"]`,
	`tt0000001	2	nm0005690	director	\N	\N`,
	`tt0000001	3	nm0374658	cinematographer	director of photography	\N`,
	`tt0000002	1	nm0721526	director	\N	\N`,
	`tt0000002	2	nm1335271	composer	\N	\N`,
	`tt0000003	1	nm0721526	director	\N	\N`,
	`tt0000003	2	nm1770680	producer	producer	\N`,
	`tt0000003	3	nm1335271	composer	\N	\N`,
}

func TestGetRowByUniqueId(t *testing.T) {
	expectedRow := `tt0000004	short	Un bon bock	Un bon bock	0	1892	\N	12	Animation,Short`
	existingRowId := "tt0000004"
	nonexistentRowId := "tt0000009"

	row, err := GetRowByUniqueId(uniqueIdTable, existingRowId)
	assert.NoError(t, err)
	assert.Equal(t, row, expectedRow)

	row, err = GetRowByUniqueId(uniqueIdTable, nonexistentRowId)
	assert.Error(t, err)
	assert.Equal(t, "", row)
}

func TestGetRowByField(t *testing.T) {
	value := `Chinese Opium Den`
	title := `primaryTitle`
	expectedRow := `tt0000006	short	Chinese Opium Den	Chinese Opium Den	0	1894	\N	1	Short`

	row, err := GetRowByField(uniqueIdTable, title, value)
	assert.NoError(t, err)
	assert.NotEqual(t, "", row)
	assert.Equal(t, expectedRow, row)
}

func TestGetRowsById(t *testing.T) {
	expectedRows := [3]string{
		`tt0000001	1	nm1588970	self	\N	["Self"]`,
		`tt0000001	2	nm0005690	director	\N	\N`,
		`tt0000001	3	nm0374658	cinematographer	director of photography	\N`}
	existingRowId := "tt0000001"
	nonexistentRowId := "tt0000004"

	rows, err := GetRowsById(multyIdTable, existingRowId)
	assert.NotNil(t, rows)
	assert.NoError(t, err)

	for i := 1; i < 3; i++ {
		assert.Equal(t, expectedRows[i], rows[i])
	}

	rows, err = GetRowsById(multyIdTable, nonexistentRowId)
	assert.Error(t, err)
	assert.Nil(t, rows)
}

func TestGetColumnIndexInRow(t *testing.T) {
	row := `tconst	ordering	nconst	category	job	characters`
	title := "nconst"
	nconstIndex := 2

	index := GetColumnIndexInRow(row, title)
	assert.Equal(t, nconstIndex, index)

	absentIndex := -1
	title = "foo"
	index = GetColumnIndexInRow(row, title)
	assert.Equal(t, absentIndex, index)
}

func TestGetValueByFieldIndexInRow(t *testing.T) {
	row := `tconst	ordering	nconst	category	job	characters`
	expected := "nconst"
	index := 2

	value := GetValueByFieldIndexInRow(row, index)
	assert.NotEqual(t, "", value)
	assert.Equal(t, expected, value)
}

func TestVerifyId(t *testing.T) {
	assert.NoError(t, verifyId("nm1234000"))
	assert.NoError(t, verifyId("tt0000005"))

	assert.Error(t, verifyId("nm12g34"))
	assert.Error(t, verifyId("nmnm1234"))
	assert.Error(t, verifyId("34nmnm1234"))
	assert.Error(t, verifyId("nm12g34"))
	assert.Error(t, verifyId("nm1234g"))

	assert.Error(t, verifyId("ttt1234"))
	assert.Error(t, verifyId("tt12tt4"))
	assert.Error(t, verifyId("tt1234tt"))
	assert.Error(t, verifyId("12tt34"))
	assert.Error(t, verifyId("12tt34tt"))
}
