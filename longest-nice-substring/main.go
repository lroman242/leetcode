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
		result = longestNiceSubstring(inputString)
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
