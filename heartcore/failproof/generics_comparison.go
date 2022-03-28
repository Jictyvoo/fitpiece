package failproof

import (
	"errors"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, received, expected T) {
	if received != expected {
		t.Errorf("Error! Fail in assertion:\nReceived: `%v`\nExpected: `%v`", received, expected)
	}
}

func AssertEqualSlice[T comparable](t *testing.T, received, expected []T) {
	if len(received) != len(expected) {
		t.Errorf("Size of slices are not equal %d - %d", len(received), len(expected))
		return
	}
	for index, _ := range received {
		AssertEqual(t, received[index], expected[index])
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Error(generateError(err.Error(), "nil Error"))
	}
}

func AssertError(t *testing.T, received error, expected error) {
	if errors.Is(received, expected) {
		t.Error(generateError(received.Error(), expected.Error()))
	}
}
