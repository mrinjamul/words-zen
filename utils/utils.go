package utils

import (
	"strings"

	"github.com/mrinjamul/words-zen/db"
)

func isVowel(word string) bool {
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
	if strings.Contains(query, ",") {
		words = strings.Split(query, ",")
	}
	// Trim word for trailing character
	for id, word := range words {
		word = strings.Trim(word, "$")
		word = strings.TrimSpace(word)
		words[id] = word
	}
	phrase := phraseGenerator(words)
	if words[0] == "a" {
		if isVowel(phrase) {
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
		if id > 0 {
			if word == "is" && words[id-1] == "pluralNoun" {
				word = "are"
			}
			if word == "was" && words[id-1] == "pluralNoun" {
				word = "were"
			}
		}
		// get randow word here
		word = db.GetWordByType(word)
		phrase += word + " "
	}
	return phrase
}
