package errors

type ErrNotFound struct {
}

func (e *ErrNotFound) Error() string {
	return "required data is missing"
}

type ErrBadYearFormat struct {
}

func (e *ErrBadYearFormat) Error() string {
	return "bad year format, expect years in the range of [1800:2099]"
}

type ErrDataBaseVerifying struct {
}

func (e *ErrDataBaseVerifying) Error() string {
	return "database verification error"
}

type ErrDataBaseUnpacking struct {
}

func (e *ErrDataBaseUnpacking) Error() string {
	return "database unpacking error"
}

type ErrDataBaseLoading struct {
}

func (e *ErrDataBaseLoading) Error() string {
	return "database loading error"
}

type ErrDefaultRequestProcessing struct {
}

func (e *ErrDefaultRequestProcessing) Error() string {
	return "An unexpected error occurred while processing the request, please see the logs for details."
}
