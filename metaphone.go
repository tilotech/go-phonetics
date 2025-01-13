// Copyright 2013 Vitaly Domnikov. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package phonetics

import (
	"strings"
)

// EncodeMetaphone is a function to encode string with Metaphone algorithm.
// Metaphone is a phonetic algorithm, published by Lawrence Philips in 1990, for indexing words by their English pronunciation.
// With Michael Kuhn modification (mkuhn@rhlab.UUCP)
func EncodeMetaphone(word string) string {
	if word == "" {
		return ""
	}
	word = strings.ToUpper(word)
	word = removeDuplicatesMetaphone(word)
	wordLen := len(word)
	if wordLen > 1 {
		switch word[0:2] {
		case "PN", "AE", "KN", "GN", "WR":
			word = word[1:]
		case "WH":
			word = "W" + word[2:]
		}
		if word[0:1] == "X" {
			word = "W" + word[1:]
		}
	}

	var result strings.Builder
	for i, rune := range word {
		switch rune {
		case 'B':
			{
				if i != wordLen-1 || safeSubString(word, i-1, 2) != "MB" {
					result.WriteByte('B')
				}
			}
		case 'C':
			{
				if safeSubString(word, i, 3) == "CIA" || safeSubString(word, i, 2) == "CH" {
					result.WriteByte('X')
				} else if safeSubString(word, i, 2) == "CI" || safeSubString(word, i, 2) == "CE" || safeSubString(word, i, 2) == "CY" {
					result.WriteByte('S')
				} else if safeSubString(word, i-1, 3) != "SCI" || safeSubString(word, i-1, 3) != "SCE" || safeSubString(word, i-1, 3) != "SCY" {
					result.WriteByte('K')
				}
			}
		case 'D':
			{
				if safeSubString(word, i, 3) == "DGE" || safeSubString(word, i, 3) == "DGY" || safeSubString(word, i, 3) == "DGI" {
					result.WriteByte('J')
				} else {
					result.WriteByte('T')
				}
			}
		case 'F':
			result.WriteByte('F')
		case 'G':
			{
				prev := safeSubString(word, i+1, 1)
				if (safeSubString(word, i, 2) == "GH" && !isVowel(safeSubString(word, i+2, 1))) ||
					safeSubString(word, i, 2) == "GN" ||
					safeSubString(word, i, 4) == "GNED" ||
					safeSubString(word, i, 3) == "GDE" ||
					safeSubString(word, i, 3) == "GDY" ||
					safeSubString(word, i, 3) == "GDI" {
				} else if prev == "I" || prev == "E" || prev == "Y" {
					result.WriteByte('J')
				} else {
					result.WriteByte('K')
				}
			}
		case 'H':
			{
				if !isVowel(safeSubString(word, i+1, 1)) &&
					safeSubString(word, i-2, 2) != "CH" &&
					safeSubString(word, i-2, 2) != "SH" &&
					safeSubString(word, i-2, 2) != "PH" &&
					safeSubString(word, i-2, 2) != "TH" &&
					safeSubString(word, i-2, 2) != "GH" {
					result.WriteByte('H')
				}
			}
		case 'J':
			result.WriteByte('J')
		case 'K':
			{
				if safeSubString(word, i-1, 1) != "C" {
					result.WriteByte('K')
				}
			}
		case 'L':
			result.WriteByte('L')
		case 'M':
			result.WriteByte('M')
		case 'N':
			result.WriteByte('N')
		case 'P':
			{
				if safeSubString(word, i+1, 1) == "H" {
					result.WriteByte('F')
				} else {
					result.WriteByte('P')
				}
			}
		case 'Q':
			result.WriteByte('K')
		case 'R':
			result.WriteByte('R')
		case 'S':
			{
				if safeSubString(word, i+1, 1) == "H" || safeSubString(word, i, 3) == "SIO" || safeSubString(word, i, 3) == "SIA" {
					result.WriteByte('X')
				} else {
					result.WriteByte('S')
				}
			}
		case 'T':
			{
				if safeSubString(word, i, 3) == "TIO" || safeSubString(word, i, 3) == "TIA" {
					result.WriteByte('X')
				} else if safeSubString(word, i+1, 1) == "H" {
					result.WriteByte('0')
				} else if safeSubString(word, i, 3) != "TCH" {
					result.WriteByte('T')
				}
			}
		case 'V':
			result.WriteByte('F')
		case 'W':
			{
				if isVowel(safeSubString(word, i+1, 1)) {
					result.WriteByte('W')
				}
			}
		case 'X':
			result.WriteString("KS")
		case 'Y':
			{
				if isVowel(safeSubString(word, i+1, 1)) {
					result.WriteByte('Y')
				}
			}
		case 'Z':
			result.WriteByte('S')
		}
	}
	return result.String()
}

func safeSubString(word string, start, count int) string {
	wordLen := len(word)
	if start < 0 {
		start = 0
		count = count + start
	}
	if start+count > wordLen {
		count = wordLen - start
	}
	return word[start : start+count]
}

func isVowel(char string) bool {
	return strings.Contains("AEIOU", char)
}

func removeDuplicatesMetaphone(word string) string {
	previousChar := []rune(word)[0]
	var result strings.Builder
	result.WriteRune(previousChar)
	for _, rune := range word[1:] {
		if rune != previousChar || rune == 'C' {
			result.WriteRune(rune)
		}
		previousChar = rune
	}
	return result.String()
}
