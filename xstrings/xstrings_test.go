package xstrings

import (
	"sort"
	"testing"
)

func TestPtrToLower(t *testing.T) {
	if PtrToLower(nil) != nil {
		t.Errorf("PtrToLower(nil) != nil")
	}

	if *PtrToLower(Ptr("A")) != "a" {
		t.Errorf("PtrToLower(A) != a")
	}
}

func TestCompareSlices(t *testing.T) {
	tests := []struct {
		a    []string
		b    []string
		want bool
	}{
		{nil, nil, true},
		{[]string{""}, nil, false},
		{[]string{""}, []string{""}, true},
		{[]string{"a"}, []string{"a"}, true},
		{[]string{"a"}, []string{"b"}, false},
		{[]string{"a"}, []string{"a", "b"}, false},
		{[]string{"b", "a"}, []string{"a", "b"}, true},
		{[]string{"b", "a"}, []string{"a", "a"}, false},
		{[]string{"b", "a"}, []string{"a", "b", "c"}, false},
	}
	for _, tt := range tests {
		if got := CompareSlices(tt.a, tt.b); got != tt.want {
			t.Errorf("CompareSlices(%s, %s) - got: %t, want: %t\n", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestFindCommonPrefix(t *testing.T) {
	tests := []struct {
		prefix     string
		array      []string
		wantPrefix string
		wantComp   []string
	}{
		{"a", []string{"abc", "abcd"}, "abc", []string{"abc", "abcd"}},
		{"a", []string{"a", "abcd"}, "a", []string{"a", "abcd"}},
		{"a", []string{"a", "d"}, "a", []string{"a"}},
		{"a", []string{"ab", "abc", "d"}, "ab", []string{"ab", "abc"}},
		{"a", []string{"b", "bc", "d"}, "a", []string{}},
		{"d", []string{"b", "dc", "bc", "d"}, "d", []string{"d", "dc"}},
	}
	for _, tt := range tests {
		sort.Strings(tt.array)
		gotPrefix, gotComp := FindCommonPrefix(tt.array, tt.prefix)
		if gotPrefix != tt.wantPrefix || !CompareSlices(gotComp, tt.wantComp) {
			t.Errorf("FindCommonPrefix(%s, %s) - got: (%s, %s), want: (%s, %s)\n",
				tt.prefix, tt.array, gotPrefix, gotComp, tt.wantPrefix, tt.wantComp)
		}
	}
}
