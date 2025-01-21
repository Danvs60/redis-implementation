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
	decoded = strings.TrimSuffix(decoded, "\r\n")
	if decoded == respString {
		err = errors.New("Invalid RESP string")
	}
	return
}
