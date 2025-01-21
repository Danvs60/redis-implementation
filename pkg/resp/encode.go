package resp

import (
	"errors"
	"fmt"
)

// EncodeSimpleString encodes a string in RESP protocol.
//
// Example: OK => "+OK\r\n"
func EncodeSimpleString(simpleString string) (string, error) {
	if simpleString == "" {
		return "", errors.New("Input string cannot be null or empty")
	}
	return fmt.Sprintf("+%s\r\n", simpleString), nil
}
