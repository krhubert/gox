package xstrings

import (
	"math/rand"
	"strings"
)

// Ptr returns string reference.
func Ptr(s string) *string {
	return &s
}

// PtrToLower returns a copy of the string s with all Unicode letters
// mapped to their lower case or nil if the string was nil
func PtrToLower(s *string) *string {
	if s == nil {
		return nil
	}
	ns := strings.ToLower(*s)
	return &ns
}

// CompareSlices compares if a and b slice have the same elements, returns
// true if slice are the same, false otherwise.
func CompareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[string]int, len(a))

	for i := range a {
		m[a[i]]++
	}

	for i := range b {
		m[b[i]]--
		if m[b[i]] != 0 {
			return false
		}
	}

	return true
}

// FindCommonPrefix take prefix and try to find longest common prefix in w slice.
// It returns longest common prefix and all possibles complements.
// NOTE: w slice must be sorted.
func FindCommonPrefix(w []string, prefix string) (string, []string) {
	start := 0

	for i, s := range w {
		if strings.HasPrefix(s, prefix) {
			start = i
			break
		}
	}

	stop := start
	for _, s := range w[start:] {
		if !strings.HasPrefix(s, prefix) {
			break
		}
		stop++
	}

	newp := prefix
	oldp := prefix

loop:
	for i := len(prefix); i < len(w[start]); i++ {
		newp = newp + string(w[start][i])
		for _, s := range w[start:stop] {
			if !strings.HasPrefix(s, prefix) {
				break loop
			}
		}
		oldp = newp
	}
	return oldp, w[start:stop]
}

// SliceIndex returns the index e in a, return -1 if not found.
func SliceIndex(a []string, e string) int {
	for i, s := range a {
		if s == e {
			return i
		}
	}
	return -1
}

// SliceContains returns true if slice a contains e element, false otherwise.
func SliceContains(a []string, e string) bool {
	return SliceIndex(a, e) >= 0
}

// Pop retrives first element from string slice, returns first elemetn and the rest of slice.
func Pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}
	return s[0], s[1:]
}

// FindLongest finds the length of longest string in the slice.
// It returns -1 if the slice len equals 0.
func FindLongest(a []string) int {
	l := -1
	for _, s := range a {
		if i := len(s); i > l {
			l = i
		}
	}
	return l
}

// Alphabet ASCII
var Alphabet = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-.")

// Random generates random string of length n. It takes alpahabet as second argument
func Random(n int, a []rune) string {
	l := len(a)
	b := make([]rune, n)
	for i := range b {
		b[i] = a[rand.Intn(l)]
	}
	return string(b)
}

// RandomASCII generates random string of length n.
// It is equivalent of Random(n, ASCII).
func RandomASCII(n int) string {
	return Random(n, Alphabet)
}

// Accept test if s string consist of rs runes. Return true if it does, false otherwies.
func Accept(s string, rs []rune) bool {
	m := make(map[rune]bool, len(rs))
	for _, r := range rs {
		m[r] = true
	}

	return AcceptFunc(s, func(r rune) bool { return m[r] })
}

// AcceptFunc test if s string runes are accepted by fn function. Return true if it does, false otherwies.
func AcceptFunc(s string, fn func(rune) bool) bool {
	for _, r := range s {
		if !fn(r) {
			return false
		}
	}
	return true
}

// SplitFunc slices s into all substrings separated by rune, returns a slice of the substrings
// between those runes and slice of runes used for split s.
// fn func(rune) bool should return true if the slice s should be separeted on this rune, false otherwise.
// Special case:
// If some runes left in slices s, all of them is append to output string slice and character 0 is append
// to output rune slice.
func SplitFunc(s string, fn func(rune) bool) ([]string, []rune) {
	var (
		a     = make([]string, 0)
		ra    = make([]rune, 0)
		start = 0
	)

	for i, r := range s {
		if !fn(r) {
			continue
		}

		a = append(a, s[start:i])
		ra = append(ra, r)
		start = i + 1
	}

	if start < len(s) {
		a = append(a, s[start:])
		ra = append(ra, 0)
	}

	return a, ra
}

// CharsToString convert c string (char*) to string.
func CharsToString(c []int8) string {
	var (
		n int

		l = len(c)
		s = make([]byte, l)
	)

	for ; n < l && c[n] != 0; n++ {
		s[n] = uint8(c[n])
	}

	return string(s[:n])
}

// Unique returs slice without duplicated elements.
func Unique(s []string) []string {
	found := make(map[string]bool, len(s))
	j := 0
	for _, v := range s {
		if found[v] {
			continue
		}
		found[v] = true
		s[j] = v
		j++

	}
	return s[:j]
}
