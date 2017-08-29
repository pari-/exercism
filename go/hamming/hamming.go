// Package hamming provides the implementation of exercism's "hamming" exercise.
package hamming

// import path for needed packages
import "errors"

// consts
const testVersion = 6

// Distance computes the hamming distance of two strings and returns an error
// in case of unequal length
func Distance(a, b string) (int, error) {
	var error error
	v := 0
	if len(a) == len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				v++
			}
		}
	} else {
		error = errors.New("unequal length :(")
	}
	return v, error
}
