// Package accumulate provides the implementation of exercism's "accumulate" exercise.
package accumulate

// consts
const testVersion = 1

// Accumulate returns a new collection containing the results of applying a
// given converter function to each element of an input collection
func Accumulate(collection []string, converter func(string) string) []string {
	accumulatedCollection := make([]string, len(collection))
	for i, element := range collection {
		accumulatedCollection[i] = converter(element)
	}
	return accumulatedCollection
}
