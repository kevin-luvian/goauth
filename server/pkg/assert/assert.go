package assert

import (
	"errors"

	tassert "github.com/stretchr/testify/assert"
)

var (
	ErrMock = errors.New("an error")
)

func WantError(t tassert.TestingT, want bool, err error, msgAndArgs ...interface{}) bool {
	if want {
		return tassert.Error(t, err, msgAndArgs...)
	} else {
		return tassert.NoError(t, err, msgAndArgs...)
	}
}

func Equal(t tassert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return tassert.Equal(t, expected, actual, msgAndArgs...)
}
