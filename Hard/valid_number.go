package main

import (
	"unicode"
)

func isNumber(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}

	// Pointers for traversing the string
	i, eSeen, dotSeen, digitSeen := 0, false, false, false

	// Optional leading sign
	if s[i] == '+' || s[i] == '-' {
		i++
	}

	// Traverse the string
	for i < n {
		ch := s[i]

		if unicode.IsDigit(rune(ch)) {
			digitSeen = true
		} else if ch == '.' {
			// Dot can only appear once and not after 'e'
			if dotSeen || eSeen {
				return false
			}
			dotSeen = true
		} else if ch == 'e' || ch == 'E' {
			// 'e' can only appear once and must follow a digit
			if eSeen || !digitSeen {
				return false
			}
			eSeen = true
			digitSeen = false // Expect a digit after 'e'
		} else if ch == '+' || ch == '-' {
			// Sign after 'e' must be followed by a digit
			if s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			// Any other character is invalid
			return false
		}
		i++
	}

	// Must end with a digit
	return digitSeen
}
