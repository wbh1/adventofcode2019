package helpers

import (
	"strings"
)

// SplitOnNewLines splits a multi-line file based on new-line characters
func SplitOnNewLines(content []byte) []string {
	return strings.Split(string(content), "\n")

}
