// Package twofer implements Exercism's 'two-fer'-exercise best described here:
// https://github.com/exercism/go/blob/master/exercises/two-fer/README.md
package twofer

import (
	"fmt"
)

// ShareWith returns the string "One for %s, one for me".
// The value of '%s' depends on the supplied input parameter (subject). If it
// is an empty string, it'll evaluate to 'you', otherwise to the the value of
// subject
func ShareWith(subject string) string {
	if len(subject) == 0 {
		subject = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", subject)
}
