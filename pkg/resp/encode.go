package resp

import (
	"errors"
	"fmt"
)

// EncodeStringFormat encodes the string in the input format
func EncodeStringFormat(s, format string) (string, error) {
	if s == "" {
		return "", errors.New("Input string cannot be null or empty")
	}
	return fmt.Sprintf(format, s), nil
}

// EncodeSimpleString encodes a string in RESP protocol.
//
// Example: OK => "+OK\r\n"
func EncodeSimpleString(simpleString string) (string, error) {
	return EncodeStringFormat(simpleString, "+%s\r\n")
}

// EncodeErrorMessage encodes an error message in RESP protocol.
//
// Example: Error message => "-Error message\r\n"
func EncodeErrorMessage(errorString string) (string, error) {
	return EncodeStringFormat(errorString, "-%s\r\n")
}
