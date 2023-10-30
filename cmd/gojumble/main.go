/*
go jumble: golang implementation of simple de-jumbler program
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// Function to hash words by alphabetizing their letters
func hash(s string) string {
	a := strings.Split(s, "")
	sort.Strings(a)
	return strings.Join(a, "")
}

// Pull up dictionary and parse words
func readWords() map[string][]string {
	var wordHash = make(map[string][]string)
	wordList, err := os.Open("../../assets/words")
	if err != nil {
		log.Fatal(err)
	}
	defer wordList.Close()
	scanner := bufio.NewScanner(wordList)
	for scanner.Scan() {
		myWord := strings.Trim(scanner.Text(), " ")
		myHash := hash(myWord)
		wordHash[myHash] = append(wordHash[myHash], myWord)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wordHash
}

func usage() {
	fmt.Printf("Usage: gojumble word\n")
}

func main() {
	var help = flag.Bool("h", false, "Print usage")
	flag.Parse()
	if *help {
		usage()
		os.Exit(0)
	}
	var guess string
	if len(flag.Args()) > 0 {
		guess = flag.Arg(0)
	} else {
		usage()
		os.Exit(1)
	}

	wordHash := readWords()
	fmt.Printf("wordHash of %q %v\n", guess, wordHash[hash(guess)])

	// Find some stats
	var biggest []string
	var biggest_len = 1
	for k := range wordHash {
		if len(wordHash[k]) > biggest_len {
			biggest_len = len(wordHash[k])
			biggest = wordHash[k]
		}
	}
	fmt.Printf("Word with the most jumbles has %v variations and the values are %v\n",
		biggest_len, biggest)
	fmt.Printf("Length of word list %v\n", len(wordHash))
}
