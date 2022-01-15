package db

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var (
	allWordTypes []string = []string{"adjective", "adverb", "animal", "bodyPart", "gerund", "noun", "pluralNoun", "verb"}
)

// GetAllWordTypes returns all word types
func GetWordTypes() []string {
	return allWordTypes
}

func GetWords() []string {
	var words []string
	for _, wordType := range allWordTypes {
		words = append(words, getWordFile(wordType)...)
	}
	return words
}

func GetWordsByType(wordType string) []string {
	return getWordFile(wordType)
}

func GetWord() string {
	var word string
	wordType := getRandomWord(allWordTypes)
	words := getWordFile(wordType)
	word = getRandomWord(words)
	return word
}
func GetWordByType(types string) string {
	var word string
	words := getWordFile(types)
	if words == nil {
		return types
	}
	word = getRandomWord(words)
	return word
}

func getWordFile(wordtype string) []string {
	var words []string
	// read file content from data folder
	file, err := os.Open("db/data/" + wordtype + "s.txt")
	if err != nil {
		return nil
	}
	defer file.Close()
	// read file line by line and store to words
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words
}

func getRandomWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(words))
	return words[random]
}
