// Package documentme ...
// This pacakge is for showcasing the documentation capabilities of Go
// It is a naive package
package documentme

// Pie is a global vvariable
// This is a silly comment !
const Pie = 3.1415912

// S1 ...
// The S1() function finds the length of a string
// It iterates over the string using range
func S1(s string) int {
	if s == "" {
		return 0
	}

	n := 0
	for range s {
		n++
	}
	return n
}

// F1 ...
// The F1() function returns the double value of its inteeger
// A better function name would have been Double() !
func F1(n int) int {
	return 2 * n
}
