// Package triangle provides the implementation of exercism's "triangle" exercise.
package triangle

// import path
import (
	"math"
)

// consts
const (
	testVersion = 3
	NaT         = iota
	Equ         = iota
	Iso         = iota
	Sca         = iota
)

// Kind (see consts)
type Kind uint8

// KindFromSides returns the classification of the supplied triangle params
func KindFromSides(a, b, c float64) Kind {
	if a > 0 && b > 0 && c > 0 && a < math.Inf(1) && b < math.Inf(1) && c < math.Inf(1) {
		if a+b >= c && a+c >= b && b+c >= a {
			if a == b && b == c {
				return Equ
			} else if a == b || a == c || b == c {
				return Iso
			} else {
				return Sca
			}
		}
	}
	return NaT
}
