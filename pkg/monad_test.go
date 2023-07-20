package pkg

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Test struct {
	*testing.T
}

func (t *Test) AssertEqual(expected any) Bind {
	return func(actual any) Monad {
		assert.Equal(t, actual, expected)
		return Empty()
	}
}

func (t *Test) IsNotCalled() Bind {
	return func(actual any) Monad {
		assert.True(t, false)
		return Empty()
	}
}

func (t *Test) IsError(actual error) func(err error) {
	return func(expected error) {
		assert.EqualError(t, expected, actual.Error())
	}
}

func Tester(t *testing.T) *Test {
	return &Test{t}
}

func TestJust(t *testing.T) {
	Just(2).Bind(Tester(t).AssertEqual(2))
}

func TestEmpty(t *testing.T) {
	Empty().Bind(Tester(t).IsNotCalled())
}

func TestFail(t *testing.T) {
	err := errors.New("new error")
	Fail(err).Catch(Tester(t).IsError(err))
}
