// Package acronym provides the implementation of exercism's "acronym" exercise.
package acronym

// import path for needed packages
import (
	"bytes"
	"regexp"
	"strings"
)

// consts
const testVersion = 3

// Abbreviate takes a space or hyphen delimited string, generating an acronym
// composed of the capitalized initial letter of each word
func Abbreviate(phrase string) string {
	var acronym bytes.Buffer
	for _, element := range regexp.MustCompile(" |-").Split(strings.ToUpper(phrase), -1) {
		acronym.Write([]byte(element[0:1]))
	}
	return acronym.String()
}
