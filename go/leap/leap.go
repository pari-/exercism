// Package leap provides a function which checks whether a supplied integer
// (year) qualifies as a Gregorian calendar "leap year"
package leap

const testVersion = 3

// IsLeapYear returns true or false depending  if the supplied year is a
// leap year
func IsLeapYear(year int) bool {
	if (year%4 != 0) || ((year%400 != 0) && (year%100 == 0)) {
		return false
	}
	return true
}
