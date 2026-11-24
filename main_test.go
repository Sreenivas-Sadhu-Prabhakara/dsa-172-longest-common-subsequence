package main

import "testing"

func TestLCS(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"abcde", "ace", 3},
		{"abc", "abc", 3},
		{"abc", "def", 0},
		{"", "abc", 0},
		{"", "", 0},
		{"abcbdab", "bdcaba", 4}, // classic: bcba / bdab
		{"aaaa", "aa", 2},
	}
	for _, c := range cases {
		if got := LCS(c.a, c.b); got != c.want {
			t.Errorf("LCS(%q,%q) = %d, want %d", c.a, c.b, got, c.want)
		}
		if got := LCS(c.b, c.a); got != c.want {
			t.Errorf("LCS symmetric check failed for %q,%q", c.a, c.b)
		}
	}
}
