//leap is a package that determines whether a given year
// is a leap year in the Gregorian calendar.
//Ben Morrison
//ben@gbmor.dev
package leap

// IsLeapYear implements the actual test to determine leap-year-ness
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	} else {
		return false
	}
}
