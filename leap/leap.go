package leap

// Function IsLeapYear returns boolean indicating if the given year is leap year.
func IsLeapYear(n int) bool {
	return n%400 == 0 || (n%4 == 0 && n%100 != 0)
}
