package failproof

import (
	"errors"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, received, expected T) bool {
	if received != expected {
		t.Errorf("Error! Fail in assertion:\nReceived: `%v`\nExpected: `%v`", received, expected)
		return false
	}
	return true
}

func AssertEqualCompare[T any](t *testing.T, compareCallback func(a, b T) int, received, expected T) bool {
	if compareCallback(received, expected) != 0 {
		t.Errorf("Error! Fail in assertion:\nReceived: `%v`\nExpected: `%v`", received, expected)
		return false
	}
	return true
}

func AssertEqualSlice[T comparable](t *testing.T, received, expected []T) bool {
	if len(received) != len(expected) {
		t.Errorf("Size of slices are not equal %d - %d", len(received), len(expected))
		return false
	}
	valid := true
	for index, _ := range received {
		valid = valid && AssertEqual(t, received[index], expected[index])
	}
	return valid
}

func NoError(t *testing.T, err error) bool {
	if err != nil {
		t.Error(generateError(err.Error(), "nil Error"))
		return false
	}
	return true
}

func AssertError(t *testing.T, received error, expected error) {
	if errors.Is(received, expected) {
		t.Error(generateError(received.Error(), expected.Error()))
	}
}
