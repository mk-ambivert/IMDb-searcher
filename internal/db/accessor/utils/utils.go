package utils

import (
	"errors"
	"regexp"
)

var errBadYearFormat error = errors.New("bad year format")

// VerifyYear checks that year matches the format,
// starting with 18,19 or 20, and end with any two digits, so a valid year must be between [1800:2099].
func VerifyYear(year string) error {
	// ^(18|19|20) means that the string must start exactly with the values: 18,19,20
	// \\d{2}$ means that the string must end exactly with any two digits
	rxp, err := regexp.Compile(`^(18|19|20)\d{2}$`)
	if err != nil {
		return err
	}
	ok := rxp.Match([]byte(year))
	if !ok {
		return errBadYearFormat
	}
	return nil
}
