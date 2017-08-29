// Package pangram provides the implementation of exercism's "pangram" exercise.
package pangram

// import path for needed packages.
import (
	"strings"
)

// consts
const testVersion = 2

// IsPangram takes an input string and checks if it's a pangram.
func IsPangram(pangram string) bool {
	m := make(map[rune]bool)
	for _, c := range strings.ToLower(pangram) {
		if c >= rune('a') && c <= rune('z') {
			m[c] = true
		}
	}
	if len(m) == 26 {
		return true
	}
	return false
}
