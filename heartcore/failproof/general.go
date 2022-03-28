package failproof

import (
	"errors"
	"fmt"
)

// Printable TODO: Use in Go 1.18 generics
type Printable interface {
	String() string
}

func generateError[T any](result, expected T) error {
	errorMessage := fmt.Sprintf("Error! Fail in assertion:\nReceived: `%v`\nExpected: `%v`", result, expected)
	return errors.New(errorMessage)
}
