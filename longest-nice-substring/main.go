package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := []string{"HeOlLoplaygroundGGAZ", "YazaAay", "Bb", "c", "dDzeE", "ijIJwuUnW", "wWOExoVhvXebB", "VaiOlVBrVyoGqygbrELjHNyRAVmHDhtSsvLMCIeFStnoyTygcrMduvyYfJQ", "hiIWIllOlRbOhXZbZZlNXxIIIOBIoioBCXlzbbioIIOOiwHhIozhblolIxoiXoOziLo"}
	expected := []string{"OlLo", "aAa", "Bb", "", "dD", "ijIJ", "wW", "Ss", "ioIIOOi"}

	var result string
	for i, inputString := range input {
		//result = longestNiceSubstring(inputString)
		result = longestNiceSubstring1(inputString)
		if result != expected[i] {
			fmt.Printf("Wrong result for %s. expected %s but got %s \n", inputString, expected[i], result)
		} else {
			fmt.Printf("%s -> %s\n", inputString, result)
		}
	}
}

func longestNiceSubstring(s string) string {
	subStrings := removeNotNiceRunes(s)

	if len(subStrings) == 0 {
		return ""
	}

	maxStr := ""
	for _, str := range subStrings {
		if len(str) > len(maxStr) {
			maxStr = str
		}
	}

	return maxStr
}

func removeNotNiceRunes(s string) []string {
	runes := []rune(s)
	list := make([]string, 0, 0)

	for i, c := range runes {
		// check if rune has its nice pair
		if isUpper(c) && !strings.Contains(s, string(unicode.ToLower(c))) || !isUpper(c) && !strings.Contains(s, string(unicode.ToUpper(c))) {
			// convert "not nice" runes into '_'
			runes[i] = '_'
		}
	}

	// check if input already nice
	if s == string(runes) {
		list = append(list, s)
		return list
	}

	tmp := ""
	for _, c := range runes {
		if c == '_' && tmp == "" {
			continue
		} else if c == '_' && tmp != "" && len(tmp) > 1 {
			results := removeNotNiceRunes(tmp)
			if len(results) > 0 {
				for _, r := range results {
					list = append(list, r)
				}
			}
			tmp = ""
		} else if c == '_' && tmp != "" && len(tmp) == 1 {
			tmp = ""
		} else if c != '_' {
			tmp = tmp + string(c)
		}
	}

	if len(tmp) > 1 {
		results := removeNotNiceRunes(tmp)
		if len(results) > 0 {
			for _, r := range results {
				list = append(list, r)
				tmp = ""
			}
		}
	}

	return list
}

func isUpper(c rune) bool {
	return c < rune('a')
	//return 'A' <= c && 'Z' >= c
}

// another interesting solution
func longestNiceSubstring1(s string) string {
	var lower, upper int32  // 2 * 24 bits for letter presence flags

	if len(s) <= 1 {  // recursion stops here
		return ""
	}
	for _,c := range s {  // go over the string to set upper and lower letter flags
		if c < rune('a') { // isupper relying on c being a latin letter
			upper |= 1 << (c - rune('A'))  // set the bit for 'seen the letter in upper case'
		} else {  // islower
			lower |= 1 << (c - rune('a'))  // set the bit for 'seen the letter in lower case'
		}
	}
	for i,c := range s {
		var bit int32  // bit corresponding to the letter
		if c < rune('a') {
			bit = 1 << (c - rune('A'))
		} else {
			bit = 1 << (c - rune('a'))
		}
		if (upper & bit) != (lower & bit) {  // the letter was not seen in upper or lower case
			// use recursion to find nice substring to the left of the letter
			ns1 := longestNiceSubstring1(s[:i])
			// use recursion to find nice substring to the right of the letter
			ns2 := longestNiceSubstring1(s[i+1:])
			// return longest of the nice substrings
			if len(ns1) < len(ns2) {
				return ns2
			}
			return ns1
		}
	}
	// entire string is nice: saw every letter in both cases
	return s
}