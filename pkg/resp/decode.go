package resp

import (
	"errors"
	"strings"
)

// DecodeSimpleString takes a RESP string and decodes it human-readable language.
//
// "+OK\r\n" => OK
func DecodeSimpleString(respString string) (decoded string, err error) {
	decoded = strings.TrimPrefix(respString, "+")
	if decoded == respString {
		err = errors.New("Invalid RESP string")
	}

	decoded = strings.TrimSuffix(decoded, "\r\n")
	if decoded == respString {
		err = errors.New("Invalid RESP string")
	}
	return
}

// DecodeErrorMessage takes a RESP error message and decodes it human-readable language.
//
// "-Error message\r\n" => Error message
func DecodeErrorMessage(respString string) (decoded string, err error) {
	decoded = strings.TrimPrefix(respString, "-")
	if decoded == respString {
		err = errors.New("Invalid RESP string")
	}

	decoded = strings.TrimSuffix(decoded, "\r\n")
	if decoded == respString {
		err = errors.New("Invalid RESP string")
	}
	return
}
