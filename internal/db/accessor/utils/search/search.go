package search

import (
	"errors"
	"regexp"
	"strings"
)

// TO DO: it might make sense to rename the input rows[] to table[]

// GetRowByUniqueId search for a row with the index searched for among the rows with unique ids.
func GetRowByUniqueId(rows []string, id string) (string, error) {
	err := verifyId(id)
	if err != nil {
		return "", err
	}

	idIndexInRow := 0
	for _, row := range rows {
		rowID := GetValueByFieldIndexInRow(row, idIndexInRow)
		if rowID == id {
			return row, nil
		}
	}

	return "", errors.New("Row with the required Index are absent")
}

// GetRowByField search for the first row that contains the same field value as the one searched for.
func GetRowByField(rows []string, fieldTitle, fieldValue string) (string, error) {
	tableHeader := rows[0] // the first line of any table contains the header with the field names
	fieldIndex := GetColumnIndexInRow(tableHeader, fieldTitle)

	for _, row := range rows {
		rowValue := GetValueByFieldIndexInRow(row, fieldIndex)
		if rowValue == fieldValue {
			return row, nil
		}
	}

	return "", errors.New("Row with the required field are absent")
}

// GetRowsById search for the rows with the index searched for among the rows that may use the same ids.
func GetRowsById(rows []string, id string) ([]string, error) {
	idPrositionIndex := 0 // means that the Id field comes first in the line
	flag := false
	var foundRows []string
	// TO DO: Need to make it more readable
	for _, row := range rows {
		field := GetValueByFieldIndexInRow(row, idPrositionIndex)
		if field == id {
			foundRows = append(foundRows, row)
			flag = true
		} else if flag {
			return foundRows, nil
		}
	}

	return nil, errors.New("rows with the required id are absent")
}

// GetColumIndexInRow search for the index of the column in the row (fields are separated by tabulation).
// Returns -1 if there is no column to search for.
func GetColumnIndexInRow(row, title string) int {
	stubs := strings.Split(row, "\t")
	for index, word := range stubs {
		if word == title {
			return index
		}
	}
	return -1
}

// GetValueByIndexInRow search for a field value by a column index in the row (fields are separated by tabulation).
// Returns "" if the index is out of range.
func GetValueByFieldIndexInRow(row string, index int) string {
	values := strings.Split(row, "\t")
	for i, value := range values {
		if i == index {
			return value
		}
	}
	return ""
}

// ValidateId checks that Id matches the format, starting with nm or tt and ending with numbers, example: nm000003, tt12345.
func verifyId(id string) error {
	// ^tt means that the string must exactly start with tt
	// \d+$ means that the string must exactly continue and end with numbers
	// '|' means logical or
	// ^nm\d+$ has the same meaning as the left part of the expression, but for the prefix nm.
	rgx := regexp.MustCompile(`^tt\d+$|^nm\d+$`)
	if !rgx.Match([]byte(id)) {
		return errors.New("Incorrect id")
	}
	return nil
}
