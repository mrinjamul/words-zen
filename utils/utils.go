package utils

import (
	"strings"

	"github.com/mrinjamul/words-zen/db"
)

func vowelTester(word string) bool {
	if len(word) > 0 {
		switch word[0] {
		case 'a', 'e', 'i', 'o', 'u':
			return true
		}
	}
	return false
}

func PhraseResolver(query string) string {
	words := strings.Fields(query)
	phrase := phraseGenerator(words)
	if words[0] == "a" {
		if vowelTester(phrase) {
			phrase = "an " + phrase
		} else {
			phrase = "a " + phrase
		}
	}
	return phrase[:len(phrase)-1]
}

func phraseGenerator(words []string) string {
	phrase := ""
	for id, word := range words {
		if (word == "a" && id == 0) || word == "" {
			continue
		}
		word = strings.Replace(word, "$", "", -1)
		// get randow word here
		word = db.GetWordByType(word)
		phrase += word + " "
	}
	return phrase
}
