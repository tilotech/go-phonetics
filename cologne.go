package phonetics

import (
	"strings"

	"github.com/tilotech/go-phonetics/diacrit"
)

// EncodeCologne is a function to encode string with the Cologne Phonetics algorithm.
//
// See https://en.wikipedia.org/wiki/Cologne_phonetics.
func EncodeCologne(word string) string {
	if word == "" {
		return ""
	}
	word = cologneNormalize(word)
	if word == "" {
		return ""
	}

	l := len(word)
	i := 0

	var result strings.Builder
	if word[0] == 'c' {
		if l > 1 {
			switch word[1] {
			case 'a', 'h', 'k', 'l', 'o', 'q', 'r', 'u', 'x':
				result.WriteByte('4')
			}
		}
		if result.Len() == 0 {
			result.WriteByte('8')
		}
		i = 1
	} else {
		i = 0
	}

	for ; i < l; i++ {
		switch word[i] {
		case 'a', 'e', 'i', 'o', 'u':
			if result.Len() == 0 {
				result.WriteByte('0')
			}
		case 'b', 'p':
			result.WriteByte('1')
		case 'd', 't':
			if i+1 < l {
				switch word[i+1] {
				case 'c', 's', 'z':
					result.WriteByte('8')
				default:
					result.WriteByte('2')
				}
			} else {
				result.WriteByte('2')
			}
		case 'f':
			result.WriteByte('3')
		case 'g', 'k', 'q':
			result.WriteByte('4')
		case 'c':
			if i+1 < l {
				switch word[i+1] {
				case 'a', 'h', 'k', 'o', 'q', 'u', 'x':
					switch word[i-1] {
					case 's', 'z':
						result.WriteByte('8')
					default:
						result.WriteByte('4')
					}
				default:
					result.WriteByte('8')
				}
			} else {
				result.WriteByte('8')
			}
		case 'x':
			if i > 0 {
				switch word[i-1] {
				case 'c', 'k', 'q':
					result.WriteByte('8')
				default:
					result.WriteString("48")
				}
			} else {
				result.WriteString("48")
			}
		case 'l':
			result.WriteByte('5')
		case 'm', 'n':
			result.WriteByte('6')
		case 'r':
			result.WriteByte('7')
		case 's', 'z':
			result.WriteByte('8')
		}
	}

	if result.Len() == 0 {
		return ""
	}

	return removeDuplicatesCologne(result.String())
}

func cologneNormalize(word string) string {
	word = diacrit.Normalize(strings.ToLower(word))

	runes := []rune(word)
	l := len(runes)
	var cleaned strings.Builder
	for i := 0; i < l; i++ {
		rune := runes[i]
		if rune >= 'a' && rune <= 'z' {
			switch rune {
			case 'j', 'y':
				rune = 'i'
			case 'v', 'w':
				rune = 'f'
			case 'p':
				if l > i+1 && runes[i+1] == 'h' {
					rune = 'f'
					i++
				}
			}
			cleaned.WriteRune(rune)
		}
	}

	return cleaned.String()
}

func removeDuplicatesCologne(word string) string {
	previousChar := []rune(word)[0]
	var result strings.Builder
	result.WriteRune(previousChar)
	for _, rune := range word[1:] {
		if rune != previousChar {
			result.WriteRune(rune)
		}
		previousChar = rune
	}
	return result.String()
}
