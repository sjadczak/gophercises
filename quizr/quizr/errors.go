package quizr

import "errors"

var (
	ErrCSVPath   = errors.New("quizr: csv path invalid")
	ErrNotCSV    = errors.New("quizr: couldn't parse as csv file")
	ErrCSVFormat = errors.New("quizr: csv format incorrect")
	As           = errors.As
	Is           = errors.Is
)

func PubError(err error, msg string) error {
	return &PublicError{err, msg}
}

type PublicError struct {
	err error
	msg string
}

func (pe PublicError) Error() string {
	return pe.err.Error()
}

func (pe PublicError) Public() string {
	return pe.msg
}

func (pe PublicError) Unwrap() error {
	return pe.err
}
