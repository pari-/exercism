// Package raindrops provides the implementation of exercism's "raindrops" exercise.
package raindrops

// import path for needed packages.
import (
	"bytes"
	"strconv"
)

// consts
const testVersion = 3

// Convert converts a given number to the outlined 'raindrop-speak'-string.
func Convert(number int) string {
	var raindropSpeak bytes.Buffer
	if (number % 3) == 0 {
		raindropSpeak.WriteString("Pling")
	}
	if (number % 5) == 0 {
		raindropSpeak.WriteString("Plang")
	}
	if (number % 7) == 0 {
		raindropSpeak.WriteString("Plong")
	}
	if len(raindropSpeak.String()) == 0 {
		raindropSpeak.WriteString(strconv.Itoa(number))
	}

	return raindropSpeak.String()
}
