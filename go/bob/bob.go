// Package bob provides the implementation of exercism's "bob" exercise.
package bob

// import path
import (
	"strings"
)

// consts
const testVersion = 3

// Hey implements Bob's lackadaisical teenager responses depending on input
func Hey(input string) string {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return "Fine. Be that way!"
	}
	endsWith := input[len(input)-1]
	if strings.ToUpper(input) == input && strings.ContainsAny(input, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return "Whoa, chill out!"
	} else if endsWith == '?' {
		return "Sure."
	}
	return "Whatever."
}
