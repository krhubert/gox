package xstrconv

import "strconv"

// IsBool is wrapper arout strconv.ParseBool and it returns
// true if s could be parse to bool, false otherwise.
func IsBool(s string) bool {
	_, err := strconv.ParseBool(s)
	return err == nil
}

// IsFloat is wrapper arout strconv.ParseFloat and it returns
// true if s could be parse to float, false otherwise.
func IsFloat(s string, bitSize int) bool {
	_, err := strconv.ParseFloat(s, bitSize)
	return err == nil
}

// IsInt is wrapper arout strconv.ParseInt and it returns
// true if s could be parse to int, false otherwise.
func IsInt(s string, base int, bitSize int) bool {
	_, err := strconv.ParseInt(s, base, bitSize)
	return err == nil
}

// IsUint is wrapper arout strconv.ParseUint and it returns
// true if s could be parse to uint, false otherwise.
func IsUint(s string, base int, bitSize int) bool {
	_, err := strconv.ParseUint(s, base, bitSize)
	return err == nil
}

// CompareFloat returns an integer comparing two strings as float.
// The result will be 0 if a==b, negative if a < b, and positive if a > b.
func CompareFloat(a, b string, bitSize int) int {
	n, _ := strconv.ParseFloat(a, bitSize)
	m, _ := strconv.ParseFloat(b, bitSize)
	if n == m {
		return 0
	}
	if n < m {
		return -1
	}
	return +1
}

// CompareInt returns an integer comparing two strings as int.
// The result will be 0 if a==b, negative if a < b, and positive if a > b.
func CompareInt(a, b string, base, bitSize int) int {
	n, _ := strconv.ParseInt(a, base, bitSize)
	m, _ := strconv.ParseInt(b, base, bitSize)
	if n == m {
		return 0
	}
	if n < m {
		return -1
	}
	return +1
}

// CompareUint returns an integer comparing two strings as uint.
// The result will be 0 if a==b, negative if a < b, and positive if a > b.
func CompareUint(a, b string, base, bitSize int) int {
	n, _ := strconv.ParseUint(a, base, bitSize)
	m, _ := strconv.ParseUint(b, base, bitSize)
	if n == m {
		return 0
	}
	if n < m {
		return -1
	}
	return +1
}
